package main

import (
	"log"

	"github.com/P-H-Pancholi/Blogging-api/pkg/database"
	"github.com/P-H-Pancholi/Blogging-api/pkg/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal("Error connecting database : %w", err)
	}

	defer database.Db.DbConn.Close()
	router := gin.Default()
	router.GET("/posts", handlers.GetAllPostsHandler)
	router.POST("/posts", handlers.CreatePostHandler)
	router.PUT("/posts/:id", handlers.UpdatePostHandler)

	router.Run("localhost:8080")

}
