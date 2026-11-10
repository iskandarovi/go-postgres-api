package main

import (
	"context"
	"go-postgres-api/database"
	"go-postgres-api/handlers"
	"go-postgres-api/repositories"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db, err := database.ConnectDB(pingCtx)
	if err != nil {
		log.Fatalf("Database ulanishida xato: %x ❗️❗️❗️", err)
	}
	defer db.Close()

	log.Println("PostgreSQL ga ulundingiz ✔️")

	bookRepo := repositories.NewBookRepository(db)
	bookHandler := handlers.NewBookHandler(bookRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /kitoblar", bookHandler.KitoblarOlish)
	mux.HandleFunc("GET /books/{id}", bookHandler.IdBoyichaOlish)
	mux.HandleFunc("POST /books", bookHandler.YangiKitob)
	mux.HandleFunc("PUT /books/{id}", bookHandler.KitobYangilash)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.KItobOchirish)

	log.Println("Server :8080 ishlayapti ✔️")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server xatosi: %v", err)
	}
}