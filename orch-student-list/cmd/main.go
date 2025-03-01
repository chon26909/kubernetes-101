package main

import (
	"log"
	"orch-student-list/config"
	"orch-student-list/internal/handler"
	"orch-student-list/internal/repository"
	"orch-student-list/internal/service"
	"orch-student-list/pkg/resty_client"

	"github.com/labstack/echo/v4"
)

func main() {
	// โหลด Config
	cfg := config.LoadConfig()

	// สร้าง Resty Client
	client := resty_client.NewRestyClient()

	// สร้าง Repository และ Service
	studentRepo := repository.NewStudentRepository(client, cfg.StudentAPIURL)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	// ตั้งค่า Echo
	e := echo.New()
	e.GET("/students/:id", studentHandler.GetStudentHandler)

	// รันเซิร์ฟเวอร์
	log.Println("Starting Server on Port 8080")
	e.Start(":8080")
}
