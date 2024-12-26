package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"library/internal/db"
	"log"
)

func main() {
	conect, err := db.Connect("postgres://postgres:sql091233@localhost:5430/Library?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer conect.Close()
	fmt.Println("ok")
}
