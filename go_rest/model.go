package main

import (
	"database/sql"
)

type book struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Author   string  `json:"author"`
	Price    float64 `json:"price"`
}

func (b *book) getBook(db *sql.DB) error {
	return db.QueryRow("SELECT title, category, author, price FROM books WHERE id=$1",
		b.ID).Scan(&b.Title, &b.Category, &b.Author, &b.Price)
}

func (b *book) updateBook(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE books SET title=$1, category=$2, author=$3, price=$4 WHERE id=$5",
			b.Title, b.Category, b.Author, b.Price, b.ID)

	return err
}

func (b *book) deleteBook(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM books WHERE id=$1", b.ID)

	return err
}

func (b *book) addBook(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO books(title, category, author, price) VALUES($1, $2, $3, $4) RETURNING id",
		b.Title, b.Category, b.Author, b.Price).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func getBooks(db *sql.DB, start, count int) ([]book, error) {
	rows, err := db.Query(
		"SELECT id, title, category, author,  price FROM books LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []book{}

	for rows.Next() {
		var b book
		if err := rows.Scan(&b.ID, &b.Title, &b.Category, &b.Author, &b.Price); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}
