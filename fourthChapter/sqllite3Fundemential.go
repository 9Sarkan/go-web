package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books_db.db")
	if err != nil {
		log.Fatal("open db: ", err)
	}
	// Create Book Table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS book (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Fatal("create table: ", err)
	}
	log.Println("Book table created!")
	statement.Exec()
	// insert into book table
	statement, err = db.Prepare("insert into book(id, isbn, name, author) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal("insert book: ", err)
	}
	statement.Exec(1, 123445, "test book", "test author")
	log.Println("an item inserted to book table!")
	// read from sqlite3
	raw, _ := db.Query("select id, name, author from book")
	var tempBook book
	for raw.Next() {
		raw.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("read book>>\tid: %d,\tname: %s, author: %s.\n", tempBook.id, tempBook.name, tempBook.author)
	}
	// update
	statement, _ = db.Prepare("update book set name=? where id=?")
	statement.Exec("Go Rest", 1)
	// insert into book table
	statement, err = db.Prepare("insert into book(id, isbn, name, author) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal("insert book: ", err)
	}
	statement.Exec(2, 23445, "test book 2", "test author")
	// read a book
	row := db.QueryRow("select id, name, author from book where id=?", 2)
	var tempbook2 book
	switch err = row.Scan(&tempbook2.id, &tempbook2.name, &tempbook2.author); err {
	case sql.ErrNoRows:
		log.Println("There is not row")
	case nil:
		log.Printf("id: %d, name: %s, author: %s\n", tempbook2.id, tempbook2.name, tempbook2.author)
	default:
		log.Fatal(err)
	}
	// delete a book
	statement, _ = db.Prepare("delete from book where id=?")
	statement.Exec(2)
	db.Close()
}
