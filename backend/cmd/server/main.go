package main

import (
	"fmt"
	hlr "student-management-system/internal/handlers"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("sms project starting")

	r := gin.Default()

	r.GET("/health", hlr.NewHealthHandler().Get)

	hlo := hlr.NewHelloHandler()
	r.GET("/api/hello", hlo.Get)

	studentHandler := hlr.NewStudentHandler()
	r.POST("/api/students", studentHandler.Create)
	r.GET("/api/students", studentHandler.List)
	r.GET("/api/students/:id", studentHandler.GetByID)
	r.Run(":8090")

}

// func FullName(first string, last string) (int, string) {
// 	defer fmt.Println("DEFER 1")
// 	defer fmt.Println("DEFER 2")
// 	defer fmt.Println("DEFER 3")
// 	fmt.Println("Calculating full name1...")
// 	fmt.Println("Calculating full name2...")
// 	return len(first + " " + last), first + " " + last
// }

// type HealthResponse struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Count   int    `json:"count,omitempty"`
// }
