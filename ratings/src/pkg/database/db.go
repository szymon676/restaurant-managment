package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=ratings port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	DB = db

	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS ratings (
            id SERIAL PRIMARY KEY,
            stars FLOAT(3),
            comment TEXT,
			username TEXT,
			date DATE
		);
    `)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrated database")
}
