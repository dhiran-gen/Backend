package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Initialize Redis client
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0, // Default DB
})

// Simulated database call (mock function)
func fetchDataFromDB(id string) string {
	time.Sleep(2 * time.Second) // Simulate delay
	return fmt.Sprintf("Data for ID: %s", id)
}

// API Handler with Redis Caching
func getDataHandler(c *gin.Context) {
	id := c.Param("id")

	// Check Redis cache
	cachedData, err := redisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		// Data not found in cache, fetch from DB
		data := fetchDataFromDB(id)

		// Store data in Redis with an expiration time
		err = redisClient.Set(ctx, id, data, 10*time.Second).Err()
		if err != nil {
			log.Println("Failed to set cache:", err)
		}

		c.JSON(http.StatusOK, gin.H{"source": "DB", "data": data})
	} else if err != nil {
		// Redis error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
	} else {
		// Data found in Redis
		c.JSON(http.StatusOK, gin.H{"source": "Cache", "data": cachedData})
	}
}

func main() {
	// Initialize Gin Router
	router := gin.Default()

	// Define Route
	router.GET("/data/:id", getDataHandler)

	// Start the Server
	log.Println("Server running on port 8080")
	router.Run(":8080")
}
