package main

import (
	"fmt"
	"log"
	"net/http"
	"context"

	"github.com/SLEEPZ74889/yoink-api/internal/router"
	"github.com/SLEEPZ74889/yoink-api/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Keine .env Datei gefunden")
	}

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Fehler beim Verbinden zur Datenbank: %v", err)
	}
	defer conn.Close(context.Background())
	fmt.Println("Erfolgreich mit der Datenbank verbunden")

	r := router.New()

	log.Println("Server läuft auf :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
