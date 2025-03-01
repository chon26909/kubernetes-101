package service

import (
	"log"

	"orch-student-list/internal/repository"
)

// StudentService จัดการ Business Logic
type StudentService struct {
	studentRepo *repository.StudentRepository
}

// NewStudentService สร้าง Service
func NewStudentService(studentRepo *repository.StudentRepository) *StudentService {
	return &StudentService{studentRepo: studentRepo}
}

// GetStudent ดึงข้อมูลนักเรียน
func (s *StudentService) GetStudent(id string) (string, error) {
	student, err := s.studentRepo.GetStudentById(id)
	if err != nil {
		log.Println("Error in service:", err)
		return "", err
	}
	return student, nil
}
