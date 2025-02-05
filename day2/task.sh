#!/bin/bash

# Define the directory structure
DIRS=(
    "url-shortener"
    "url-shortener/handlers"
    "url-shortener/db"
    "url-shortener/models"
    "url-shortener/config"
    "url-shortener/nginx"
)

FILES=(
    "url-shortener/main.go"
    "url-shortener/handlers/shorten.go"
    "url-shortener/handlers/redirect.go"
    "url-shortener/db/postgres.go"
    "url-shortener/db/redis.go"
    "url-shortener/models/url.go"
    "url-shortener/config/config.go"
    "url-shortener/nginx/nginx.conf"
    "url-shortener/docker-compose.yml"
    "url-shortener/go.mod"
    "url-shortener/go.sum"
)

# Create directories
for dir in "${DIRS[@]}"; do
    mkdir -p "$dir"
done

# Create empty files
for file in "${FILES[@]}"; do
    touch "$file"
done

echo "Project structure created successfully!"
