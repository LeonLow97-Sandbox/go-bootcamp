package repository

import "github.com/LeonLow97/types"

type DatabaseRepo interface {
	InsertNewStudent(s *types.Student) error
	GetAllStudents() []types.Student
}
