package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"webapp/pkg/data"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// authenticate is the handler used to try to authenticate a user, and
// send them a JWT token if successful.
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	// read a json payload
	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// look up the user by email address
	user, err := app.DB.GetUserByEmail(creds.Username)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// generate tokens
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "leon-refresh_token",
		Path:     "/",
		Value:    tokenPairs.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   false,
	})

	// send token to user
	_ = app.writeJSON(w, http.StatusOK, tokenPairs)
}

// refresh is the handler called to request a new token pair, when
// the jwt token has expired. We expect the refresh token to come
// from a POST request. We validate it, look up the user in the db,
// and if everything is good we send back a new token pair
// as JSON. We also set an http only, secure cookie with the refresh
// token stored inside.
func (app *application) refresh(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	refreshToken := r.Form.Get("refresh_token")
	claims := &Claims{}

	_, err = jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.JWTSecret), nil
	})

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// If the calculated time remaining until the token expiration is greater
	// than 30 seconds,
	if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
		app.errorJSON(w, errors.New("refresh token does not need renewed yet"), http.StatusTooEarly)
		return
	}

	// get the user id from the claims
	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DB.GetUser(userID)
	if err != nil {
		app.errorJSON(w, errors.New("unknown user"), http.StatusBadRequest)
		return
	}

	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Setting cookie on browser for SPA
	http.SetCookie(w, &http.Cookie{
		Name:     "leon-refresh_token",
		Path:     "/",
		Value:    tokenPairs.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	})

	// Also return json with tokens in case it is an API to API call
	_ = app.writeJSON(w, http.StatusOK, tokenPairs)
}

func (app *application) refreshUsingCookie(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == "leon-refresh_token" {
			claims := &Claims{}
			refreshToken := cookie.Value

			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})

			if err != nil {
				app.errorJSON(w, err, http.StatusBadRequest)
				return
			}

			// If the calculated time remaining until the token expiration is greater
			// than 30 seconds,
			// if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
			// 	app.errorJSON(w, errors.New("refresh token does not need renewed yet"), http.StatusTooEarly)
			// 	return
			// }

			// get the user id from the claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, err, http.StatusBadRequest)
				return
			}

			user, err := app.DB.GetUser(userID)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusBadRequest)
				return
			}

			tokenPairs, err := app.generateTokenPair(user)
			if err != nil {
				app.errorJSON(w, err, http.StatusBadRequest)
				return
			}

			// Setting cookie on browser for SPA
			http.SetCookie(w, &http.Cookie{
				Name:     "leon-refresh_token",
				Path:     "/",
				Value:    tokenPairs.RefreshToken,
				Expires:  time.Now().Add(refreshTokenExpiry),
				MaxAge:   int(refreshTokenExpiry.Seconds()),
				SameSite: http.SameSiteStrictMode,
				Domain:   "localhost",
				HttpOnly: true,
				Secure:   true,
			})

			// send back JSON
			_ = app.writeJSON(w, http.StatusOK, tokenPairs)
			return
		}
	}

	app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
}

func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.DB.AllUsers()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, users)
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	// get the userID from the url
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DB.GetUser(userID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, user)
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.UpdateUser(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	// get the userID from the url
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.DeleteUser(userID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = app.DB.InsertUser(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) deleteRefreshCookie(w http.ResponseWriter, r *http.Request) {
	delCookie := http.Cookie{
		Name:     "leon-refresh_token",
		Path:     "/",
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &delCookie)
	w.WriteHeader(http.StatusAccepted)
}
