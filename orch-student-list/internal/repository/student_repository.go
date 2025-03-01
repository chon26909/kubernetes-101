package repository

import (
	"log"

	"github.com/go-resty/resty/v2"
)

// StudentRepository ติดต่อ API ภายนอก
type StudentRepository struct {
	client *resty.Client
	apiURL string
}

// NewStudentRepository สร้าง Repository
func NewStudentRepository(client *resty.Client, apiURL string) *StudentRepository {
	return &StudentRepository{
		client: client,
		apiURL: apiURL,
	}
}

// GetStudentByID ดึงข้อมูลนักเรียนจาก API ภายนอก
func (r *StudentRepository) GetStudentById(id string) (string, error) {
	resp, err := r.client.R().
		SetQueryParam("id", id).
		Get(r.apiURL + "/students")

	if err != nil {
		log.Println("Error fetching student:", err)
		return "", err
	}

	return resp.String(), nil
}
