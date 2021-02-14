package models

import (
	"log"

	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/lib/pq"
)

// User Table
type User struct {
	gorm.Model
	Orders []Order
	Data   string `sql:"type: JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

// Order Table
type Order struct {
	gorm.Model
	User User
	Data string `sql:"type: JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// TableName for user table
func (User) TableName() string {
	return "user"
}

// TableName for order table
func (Order) TableName() string {
	return "order"
}

// InitDB tables
func InitDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "postgres://sample:sample@localhost/sample?sslmode=disable")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !db.HasTable("user") {
		db.CreateTable(&User{})
	}
	if !db.HasTable("order") {
		db.CreateTable(&Order{})
	}
	db.AutoMigrate(&User{}, &Order{})
	return db, nil
}
