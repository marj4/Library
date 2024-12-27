package db

import (
	"database/sql"
	"fmt"
	"library/internal"
)

const (
	InsertError = "Cant insert data in db, %v"
)

func InsertBook(DB *sql.DB, Book *internal.Book) error {
	query := "insert into books(bookName,author,YearRelease values ($1,$2,$3)"
	if err, _ := DB.Exec(query, Book.BookName, Book.Author, Book.YearRealease); err != nil {
		return fmt.Errorf(InsertError, err)
	}
	return nil

}
