package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestApplication_AllUsers(t *testing.T) {
	// create some mock rows, and all one row
	var mockedRows = mockedDB.NewRows([]string{"id", "email", "first_name", "last_name", "password", "active", "created_at", "updated_at", "has_token"})
	mockedRows.AddRow("1", "me@here.com", "Jack", "Smith", "abc123", "1", time.Now(), time.Now(), "0")

	// tell mockedDB what queries we expect
	// select\\\\* is a regular expression stating select with something after that
	mockedDB.ExpectQuery("select \\\\* ").WillReturnRows(mockedRows)

	// create a test recorder which satisfied the requirements for a ResponseRecorder
	rr := httptest.NewRecorder()
	// create request
	request, _ := http.NewRequest("POST", "/admin/users", nil)
	// call the handler
	handler := http.HandlerFunc(testApp.AllUsers)
	handler.ServeHTTP(rr, request)

	// check for expected status code
	if rr.Code != http.StatusOK {
		t.Error("AllUsers returned wrong status code of", rr.Code)
	}
}
