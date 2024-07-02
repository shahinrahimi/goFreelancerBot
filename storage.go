package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	GetUser()
}

type SqliteStore struct {
	db *sql.DB
}

func NewSqliteStore() (*SqliteStore, error) {
	if err := CreateDirecryIfNotExist("database"); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", "./database/mydb.db")
	if err != nil {
		return nil, err
	}
	log.Println("DB Connected!")

	return &SqliteStore{
		db,
	}, nil
}
