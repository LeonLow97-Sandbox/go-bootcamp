package dbrepo

import (
	"database/sql"
	"errors"
	"time"
	"webapp/pkg/data"
)

type TestDBRepo struct{}

// Not going to connect to database in this test repository
func (m *TestDBRepo) Connection() *sql.DB {
	return nil
}

// AllUsers returns all users as a slice of *data.User
func (m *TestDBRepo) AllUsers() ([]*data.User, error) {
	var users []*data.User

	return users, nil
}

// GetUser returns one user by id
func (m *TestDBRepo) GetUser(id int) (*data.User, error) {
	var user = data.User{}
	if id == 1 {
		user = data.User{
			ID:        1,
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@example.com",
		}

		return &user, nil
	}

	return nil, errors.New("user not found")
}

// GetUserByEmail returns one user by email address
func (m *TestDBRepo) GetUserByEmail(email string) (*data.User, error) {
	if email == "admin@example.com" {
		user := data.User{
			ID:        1,
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@example.com",
			Password:  "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK",
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		return &user, nil
	}
	return nil, errors.New("not found")
}

// UpdateUser updates one user in the database
func (m *TestDBRepo) UpdateUser(u data.User) error {
	if u.ID == 1 {
		return nil
	}
	return errors.New("update failed - no user found")
}

// DeleteUser deletes one user from the database, by id
func (m *TestDBRepo) DeleteUser(id int) error {
	return nil
}

// InsertUser inserts a new user into the database, and returns the ID of the newly inserted row
func (m *TestDBRepo) InsertUser(user data.User) (int, error) {
	return 2, nil
}

// ResetPassword is the method we will use to change a user's password.
func (m *TestDBRepo) ResetPassword(id int, password string) error {
	return nil
}

// InsertUserImage inserts a user profile image into the database.
func (m *TestDBRepo) InsertUserImage(i data.UserImage) (int, error) {
	return 1, nil
}
