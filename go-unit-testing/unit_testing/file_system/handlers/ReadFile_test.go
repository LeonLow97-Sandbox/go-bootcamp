package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonLow97/handlers"
)

func Test_ReadFile(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/view", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(handlers.ReadFile).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, but got %v", http.StatusOK, rr.Code)
	}
}
