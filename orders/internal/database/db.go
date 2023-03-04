package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "host=localhost port=5432 user=postgres password=1234 dbname=orders sslmode=disable"
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
        CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            product TEXT,
            tableid INTEGER
		);
    `)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrated database")
}
