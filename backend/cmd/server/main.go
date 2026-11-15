package main

import (
	"fmt"
	"student-management-system/internal/config"
	hlr "student-management-system/internal/handlers"
	"student-management-system/internal/repository"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("sms project starting")

	cfg := config.Load()

	r := gin.Default()

	r.GET("/health", hlr.NewHealthHandler().Get)

	hlo := hlr.NewHelloHandler()
	r.GET("/api/hello", hlo.Get)

	// studentHandler /api/students:= hlr.NewStudentHandler()
	var studentHandler *hlr.StudentHandler
	studentHandler = hlr.NewStudentHandler(repository.NewStudentRepository())
	r.POST("/api/students", studentHandler.Create)
	r.GET("/api/students", studentHandler.List)
	r.GET("/api/students/:id", studentHandler.GetByID)
	r.POST("/api/students/import", studentHandler.Import)
	// r.Run(":8090")

	r.Run(":" + cfg.Port)

}
