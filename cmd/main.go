package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/P-H-Pancholi/Blogging-api/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, pass, dbHost, port, dbName)

	conn, err := database.Connect(connURL)
	if err != nil {
		log.Fatal("Error connecting database : %w", err)
	}

	defer conn.Close(context.Background())
}
