package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/db"
	"url-shortener/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	db.InitRedis()

	r := mux.NewRouter()
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortCode}", handlers.RedirectURL).Methods("GET")

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
