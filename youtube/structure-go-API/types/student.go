package types

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	Name  string `json:"name" validate:"required,min=1,max=20"`
	Email string `json:"email" validate:"email,required"`
	Age   int64 `json:"age" validate:"required,min=1,max=30"`
	Class string `json:"class" validate:"required"`
}

func (s *Student) DecodeFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

func (s *Student) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

func (s *Student) EncodeToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

var StudentsList = []Student{
	{
		Name:  "Leon",
		Email: "leon@email.com",
		Age:   26,
		Class: "302",
	},
	{
		Name:  "Jie Wei",
		Email: "jiewei@email.com",
		Age:   27,
		Class: "402",
	},
}
