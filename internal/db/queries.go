package db

import (
	"database/sql"
	"fmt"
	"library/internal"
)

const (
	InsertError = "Cant insert data in db, %v"
	OutputError = "Cant output data from db, %v"
	WriteError  = "Cant  write data from db, %v"
)

type Book struct {
	ID     int
	Name   string
	Author string
	Year   string
}

func InsertBook(DB *sql.DB, Book *internal.Book) error {
	query := `insert into books(bookName,author,yearrelease) values ($1,$2,$3)`
	if _, err := DB.Exec(query, Book.BookName, Book.Author, Book.YearRealease); err != nil {
		return fmt.Errorf(InsertError, err)
	}
	return nil
}
func AllBooks(DB *sql.DB) ([]Book, error) {
	query := `select * from books`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf(OutputError, err)
	}

	var data []Book

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Year); err != nil {
			return nil, fmt.Errorf(WriteError, err)
		}

		data = append(data, book)
	}

	return data, nil
}
func InfoForID(DB *sql.DB, id int) (Book, error) {

}
