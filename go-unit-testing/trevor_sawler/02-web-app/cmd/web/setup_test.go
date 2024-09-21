package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application

// Set up Test Environment
func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	// Setting up test environment for sessions
	app.Session = getSession()

	// Virtual postgresql connection with no connection for testing purposes
	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
