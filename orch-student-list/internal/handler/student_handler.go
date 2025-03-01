package handler

import (
	"net/http"
	"orch-student-list/internal/service"

	"github.com/labstack/echo/v4"
)

// StudentHandler จัดการ HTTP Request
type StudentHandler struct {
	studentService *service.StudentService
}

// NewStudentHandler สร้าง Handler
func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

// GetStudentHandler รับ Request และเรียก Service
func (h *StudentHandler) GetStudentHandler(c echo.Context) error {
	id := c.Param("id")
	student, err := h.studentService.GetStudent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch student"})
	}
	return c.JSON(http.StatusOK, student)
}
