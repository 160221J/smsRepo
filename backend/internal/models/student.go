package models

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Course    string `json:"course"`
}

type CreateStudentRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Course    string `json:"course" binding:"required"`
	// FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	// Email     string `json:"email"`
	Phone string `json:"phone"`
	// Course    string `json:"course"`
}
