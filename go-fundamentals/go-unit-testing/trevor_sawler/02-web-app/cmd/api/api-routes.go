package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// serving static html
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))

	mux.Route("/web", func(mux chi.Router) {
		mux.Post("/auth", app.authenticate)
		mux.Get("/refresh-token", app.refreshUsingCookie)
		mux.Get("/logout", app.deleteRefreshCookie)
	})

	// authentication routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh-token", app.refresh)

	// test handler
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		var payload = struct {
			Message string `json:"message"`
		}{
			Message: "Hello Jie Wei!",
		}

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	// protected routes
	mux.Route("/users", func(mux chi.Router) {
		// use auth middleware
		mux.Use(app.authRequired)

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.getUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/", app.insertUser)
		mux.Patch("/", app.updateUser)
	})

	return mux
}
