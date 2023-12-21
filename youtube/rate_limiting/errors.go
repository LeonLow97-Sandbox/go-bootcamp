package main

import "net/http"

func (app *application) rateLimitExceededResponse(w http.ResponseWriter) {
	message := "rate limit exceeded"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
	w.WriteHeader(http.StatusTooManyRequests)
}
