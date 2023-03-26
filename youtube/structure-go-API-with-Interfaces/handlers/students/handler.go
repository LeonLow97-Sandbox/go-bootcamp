package students

import (
	"net/http"

	repository "github.com/LeonLow97/store"
	"github.com/LeonLow97/types"
)

type st struct {
	db repository.DatabaseRepo
}

func New(db repository.DatabaseRepo) *st {
	return &st{db: db}
}

func (s st) NewStudent(rw http.ResponseWriter, r *http.Request) {
	var student types.Student
	if err := student.DecodeFromJSON(r.Body); err != nil {
		http.Error(rw, "Failed to decode", http.StatusBadRequest)
	}

	if err := student.Validate(); err != nil {
		http.Error(rw, "Failed validation", http.StatusBadRequest)
	}

	if err := s.db.InsertNewStudent(&student); err != nil {
		http.Error(rw, "Failed to insert", http.StatusBadRequest)
	}

	rw.WriteHeader(http.StatusCreated)
	student.EncodeToJSON(rw)
	return
}
