package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"student-management-system/internal/apperrors"
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

	slog.Info("student created", "student_id", s.ID, "first_name", s.FirstName)

	c.JSON(http.StatusCreated, s)
	// assign id, store in map, c.JSON(201, saved)
}

func (h *StudentHandler) List(c *gin.Context) {
	out := h.repo.List()
	c.JSON(http.StatusOK, out)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	s, err2 := h.repo.GetByID(id)
	if err2 != nil {
		if errors.Is(err2, apperrors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		slog.Error("get student failed", "err", err2)
		return
	}
	c.JSON(http.StatusOK, s)
}
