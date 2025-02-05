package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"url-shortener/db"
)

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenResponse struct {
	ShortCode string `json:"short_code"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Generate short code using SHA-1
	hash := sha1.New()
	hash.Write([]byte(req.OriginalURL))
	shortCode := hex.EncodeToString(hash.Sum(nil))[:6]

	// Store in PostgreSQL
	_, err = db.DB.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2) ON CONFLICT (short_code) DO NOTHING", shortCode, req.OriginalURL)
	if err != nil {
		log.Println("DB Insert Error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Store in Redis cache
	db.RedisClient.Set(db.Ctx, shortCode, req.OriginalURL, 0)

	resp := ShortenResponse{ShortCode: shortCode}
	json.NewEncoder(w).Encode(resp)
}
