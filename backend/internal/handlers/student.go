package handlers

import (
	"net/http"
	"strconv"
	"student-management-system/internal/models"
	"student-management-system/internal/repository"

	"github.com/gin-gonic/gin"
)

func NewStudentHandler(repo *repository.StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

type StudentHandler struct {
	repo *repository.StudentRepository
}

func (h *StudentHandler) Create(c *gin.Context) {
	var req models.CreateStudentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	if req.FirstName == "" || req.Email == "" || req.Course == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "first_name, email, and course are required"})
		return
	}

	s := h.repo.Create(models.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Course:    req.Course,
	})

	c.JSON(http.StatusCreated, s)
	// assign id, store in map, c.JSON(201, saved)
}

// func (h *StudentHandler) List(c *gin.Context) {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	out := make([]models.Student, 0, len(h.students))
// 	for _, s := range h.students {
// 		out = append(out, s)
// 	}
// 	c.JSON(http.StatusOK, out)
// }

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	s, ok := h.repo.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}
