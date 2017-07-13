package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	time.Sleep(10000 * time.Millisecond)
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=postgres dbname=daylog sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
