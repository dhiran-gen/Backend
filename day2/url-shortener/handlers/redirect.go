package handlers

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/db"
	"github.com/gorilla/mux"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Check Redis cache first
	originalURL, err := db.RedisClient.Get(db.Ctx, shortCode).Result()
	if err == nil {
		http.Redirect(w, r, originalURL, http.StatusFound)
		return
	}

	// If not in cache, check PostgreSQL
	err = db.DB.QueryRow("SELECT original_url FROM urls WHERE short_code=$1", shortCode).Scan(&originalURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Store in cache for future requests
	db.RedisClient.Set(db.Ctx, shortCode, originalURL, 0)

	http.Redirect(w, r, originalURL, http.StatusFound)
}
