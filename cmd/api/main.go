package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/SLEEPZ74889/yoink-api/internal/db"
	"github.com/SLEEPZ74889/yoink-api/internal/router"
	"github.com/joho/godotenv"
)

//go:embed assets/banner.txt
var banner string

func main() {
	fmt.Println(banner)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())
	fmt.Println("Successfully connected to the database")

	r := router.New()

	log.Fatal(http.ListenAndServe(":8080", r))
}
