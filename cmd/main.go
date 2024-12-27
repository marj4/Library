package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"library/config"
	"library/internal"
	"library/internal/db"
	"log"
)

func main() {
	cgf := config.StartConfig()

	DB, err := db.Connect(cgf.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	if err := db.Ping(DB); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conneсt to DB is OK!")

	BookName := ""
	Author := ""
	YearRealease := ""

	fmt.Scanln(&BookName, &Author, &YearRealease)

	Book := &internal.Book{
		BookName:     BookName,
		Author:       Author,
		YearRealease: YearRealease,
	}

	if err := db.InsertBook(DB, Book); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Date is insert succsesfully")

}
