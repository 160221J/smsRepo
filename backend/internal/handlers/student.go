package handlers

import (
	"net/http"
	"strconv"
	"student-management-system/internal/models"
	"sync"

	"github.com/gin-gonic/gin"
)

func NewStudentHandler() *StudentHandler {
	return &StudentHandler{
		students: make(map[int]models.Student),
		nextID:   1,
	}
}

type StudentHandler struct {
	students map[int]models.Student
	mu       sync.Mutex
	nextID   int
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

	h.mu.Lock()
	defer h.mu.Unlock()
	s := models.Student{
		ID: h.nextID, FirstName: req.FirstName, LastName: req.LastName,
		Email: req.Email, Phone: req.Phone, Course: req.Course,
	}
	h.students[s.ID] = s
	h.nextID++
	c.JSON(http.StatusCreated, s)
	// assign id, store in map, c.JSON(201, saved)
}

func (h *StudentHandler) List(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()
	out := make([]models.Student, 0, len(h.students))
	for _, s := range h.students {
		out = append(out, s)
	}
	c.JSON(http.StatusOK, out)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	//  h.mu.Lock()
	//  defer h.mu.Unlock()
	s, ok := h.students[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}
