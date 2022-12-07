package main

import (
	"go-book-management/controllers"
	"go-book-management/database"
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	engine := gin.Default()

	health := new(controllers.HealthController)

	book := new(controllers.BookController)

	engine.GET("/", health.Check)

	engine.GET("/books", book.GetBooks)

	engine.POST("/books", book.AddBook)

	return engine
}

func init() {
	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.InitDatabase()
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
