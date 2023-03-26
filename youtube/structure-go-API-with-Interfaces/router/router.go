package router

import (
	"net/http"

	"github.com/LeonLow97/handlers/students"
	repository "github.com/LeonLow97/store"
	"github.com/gorilla/mux"
)

func NewRouter(db repository.DatabaseRepo) *mux.Router {
	r := mux.NewRouter()

	studentH := students.New(db)
	r.HandleFunc("/student", studentH.NewStudent).Methods(http.MethodPost)
	return r
}
