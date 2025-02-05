package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=admin password=admin dbname=url_shortener sslmode=disable host=localhost port=5432"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		short_code TEXT UNIQUE NOT NULL,
		original_url TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	fmt.Println("Database connected & table ensured.")
}
