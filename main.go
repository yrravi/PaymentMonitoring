package main

import (
	"log"
	"payment-Api/handlers"
	"payment-Api/repositories"
	"payment-Api/services"
	"payment-Api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

func main() {
	// Initialize database connection
	fmt.Println("ENTERS THE MAIN")
	dsn := "root:new_password@tcp(127.0.0.1:3306)/myDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto migrate the database
	db.AutoMigrate(&models.Payment{})

	// Initialize repository, service, and handler
	paymentRepo := repositories.NewPaymentRepository(db)
	//fmt.Println("-------",paymentRepo)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	// Set up the Gin router
	r := gin.Default()

	// Define the API routes
	fmt.Println("ENTERS THE payments")
	r.POST("/payments", paymentHandler.CreatePaymentHandler)

	// Start the server
	fmt.Println("Server running on port 8080")

if err := r.Run(":8080"); err != nil {
    log.Fatalf("Error starting server: %v", err)
}

	//fmt.Println("ENTERS THE END")
}
