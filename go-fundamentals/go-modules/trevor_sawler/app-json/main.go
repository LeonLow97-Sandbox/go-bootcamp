package main

import (
	"log"
	"net/http"

	"github.com/LeonLow97/toolkit"
)

type RequestPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ResponsePayload struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
}

func main() {
	// create a default serve mux
	mux := http.NewServeMux()

	// register routes
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/receive-post", receivePost)
	mux.HandleFunc("/remote-service", remoteService)
	mux.HandleFunc("/simulated-service", simulatedService)

	// print a message
	log.Println("Starting Service...")

	// start server
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func receivePost(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message:    "hit the handler ok, and sending response...",
		StatusCode: http.StatusOK,
	}

	err = t.WriteJSON(w, http.StatusAccepted, responsePayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}
}

func remoteService(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	// call a remote service
	_, statusCode, err := t.PushJSONToRemote("http://localhost:8081/simulated-service", requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message:    "hit the handler ok, and sending response...",
		StatusCode: statusCode,
	}

	err = t.WriteJSON(w, http.StatusAccepted, responsePayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}
}

func simulatedService(w http.ResponseWriter, r *http.Request) {
	payload := ResponsePayload{
		Message: "ok",
	}

	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, payload)
}
