package main

import (
	"context"
	"log"
	"logitrack"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := "postgres://postgres:MySecretPassword123@localhost:5432/logitrack?sslmode=disable"

	// 1. Connecting the DB
	pool, err := logitrack.ConnectDB(ctx, connStr)
	if err != nil {
		log.Fatalf("Database Connection Failed.\nError:%v", err)
	}
	defer pool.Close()

	// 2. Setup tables and seed data
	err = devServerSetup(connStr, pool)

	// 3. Initialize the DB Storage Layer
	repo := logitrack.NewPostgresRepository(pool)

	// 4. Initialize the server wrapper
	server := logitrack.NewServer(repo)

	log.Println("🚀 LogiTrack HTTP server booting up on http://localhost:8080")

	err = http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatalf("Server crahsed.\nError:%v", err)
	}
}
