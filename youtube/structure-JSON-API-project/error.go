package main

import "net/http"

var ErrUserInvalid = apiError{Err: "Invalid User", Status: http.StatusUnauthorized}

type apiError struct {
	Err    string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}
