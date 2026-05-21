package main

import (
	"log"
	"net/http"

	"github.com/SLEEPZ74889/yoink-api/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Keine .env Datei gefunden")
	}
	r := router.New()

	log.Println("Server läuft auf :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
