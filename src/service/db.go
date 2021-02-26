package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

var instance *Database

func DatabaseInstance() *Database {
	if instance == nil {
		user, password := GetEnv("DB_USER", "postgres"), GetEnv("DB_PASSWORD", "postgres")
		connStr := fmt.Sprintf("host=postgres user=%s password=%s dbname=golang sslmode=disable", user, password)
		connection, err := sql.Open("postgres", connStr)
		if err == nil {
			err = connection.Ping()
		}
		if err != nil {
			log.Fatal(err)
		}
		instance = &Database{connection}
	}
	return instance
}
