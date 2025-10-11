package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5433/db?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ping db: %v", err)
	}

	fmt.Println("DB OK")
}