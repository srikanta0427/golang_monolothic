package main

import (
	// "fmt"
	// "log"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/srikanta0427/listing_app/internal/config"
	"github.com/srikanta0427/listing_app/internal/db"
	"github.com/srikanta0427/listing_app/internal/handler"
)

func main() {
	fmt.Println("Start")
	cfg := config.MustLoadConfig()
	db, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	mux := http.NewServeMux()

	lh := handler.NewListingHandler(db)

	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}",lh.DeleteList)
	// c

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 30,
	}

	fmt.Println("Server is running on port", os.Getenv("PORT"))
	if server.ListenAndServe() != nil {
		log.Fatalf("Server Failed To Build")
	}
	
}