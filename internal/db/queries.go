package db

import (
	"database/sql"
	"fmt"
	"library/internal"
)

const (
	ErrInsertData = "Can't insert data into the database: %v"
	ErrOutputData = "Can't retrieve data from the database: %v"
	ErrWriteData  = "Can't write data to the database: %v"
	ErrQueryByID  = "Can't execute query by ID in the database: %v"
	ErrDeleteByID = "Can't delete book by ID from the database: %v"
)

func InsertBook(DB *sql.DB, Book internal.Book) error {
	query := `insert into books(bookName,author,yearrelease) values ($1,$2,$3)`
	if _, err := DB.Exec(query, Book.BookName, Book.Author, Book.YearRealease); err != nil {
		return fmt.Errorf(ErrInsertData, err)
	}
	return nil
}
func AllBooks(DB *sql.DB) ([]internal.Book, error) {
	query := `select * from books`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf(ErrOutputData, err)
	}

	var data []internal.Book

	for rows.Next() {
		var book internal.Book
		if err := rows.Scan(&book.ID, &book.BookName, &book.Author, &book.YearRealease); err != nil {
			return nil, fmt.Errorf(ErrWriteData, err)
		}

		data = append(data, book)
	}

	return data, nil
}
func InfoForID(DB *sql.DB, id int) (internal.Book, error) {
	query := `select bookname,author,yearrelease from books where id=$1`
	row := DB.QueryRow(query, id)
	var book internal.Book
	if err := row.Scan(&book.BookName, &book.Author, &book.YearRealease); err != nil {
		return internal.Book{}, fmt.Errorf(ErrQueryByID, err)
	}
	return book, nil
}
func DeleteFromID(DB *sql.DB, id int) error {
	query := `delete from books where id=$1`
	if _, err := DB.Exec(query, id); err != nil {
		return fmt.Errorf(ErrDeleteByID, err)
	}
	return nil
}
