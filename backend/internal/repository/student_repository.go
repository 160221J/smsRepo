package repository

import (
	"student-management-system/internal/models"
	"sync"
)

type StudentRepository struct {
	students map[int]models.Student
	mu       sync.Mutex
	nextID   int
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		students: map[int]models.Student{},
		nextID:   1,
	}
}

func (r *StudentRepository) GetByID(id int) (models.Student, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s, ok := r.students[id]
	return s, ok
}

func (r *StudentRepository) Create(studentModel models.Student) models.Student {
	r.mu.Lock()
	defer r.mu.Unlock()
	studentModel.ID = r.nextID
	r.students[studentModel.ID] = studentModel
	r.nextID++
	return studentModel
}
