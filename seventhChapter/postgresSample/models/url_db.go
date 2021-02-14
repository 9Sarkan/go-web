package models

import (
	"database/sql"
	"log"

	// use for db
	_ "github.com/lib/pq"
)

// InitDB will be initial db for link shortner
func InitDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "postgres://sample:sample@localhost/sample?sslmode=disable")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS url_shorter(id serial primary key, url text not null);")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	response, err := statement.Exec()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(response)
	return db, nil
}
