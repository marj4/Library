package main

import (
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"library/config"
	"library/internal"
	"library/internal/db"
	"log"
	"os"
	"strings"
)

func main() {
	cfg := config.StartConfig()

	DB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	if err := db.Ping(DB); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conneсt to DB is OK!")

	scanOBJ := bufio.NewReader(os.Stdin)

	bookName, _ := scanOBJ.ReadString('\n')
	bookName = strings.TrimSpace(bookName)

	author, _ := scanOBJ.ReadString('\n')
	author = strings.TrimSpace(author)

	yearRealease, _ := scanOBJ.ReadString('\n')
	yearRealease = strings.TrimSpace(yearRealease)

	Book := &internal.Book{
		BookName:     bookName,
		Author:       author,
		YearRealease: yearRealease,
	}

	if err := db.InsertBook(DB, Book); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Date is insert succsesfully")

	selectData, err := db.AllBooks(DB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(selectData)

	info, err := db.InfoForID(DB, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)

}
