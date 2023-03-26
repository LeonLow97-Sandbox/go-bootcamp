package dbrepo

import (
	"fmt"

	"github.com/LeonLow97/types"
)

type DB struct {
}

func New() *DB {
	return &DB{}
}

func (d DB) InsertNewStudent(s *types.Student) error {
	fmt.Println("Entered InsertNewStudent!")
	types.StudentsList = append(types.StudentsList, *s)
	return nil
}

func (d DB) GetAllStudents() []types.Student {
	fmt.Println("Entered GetAllStudents")
	return types.StudentsList
}
