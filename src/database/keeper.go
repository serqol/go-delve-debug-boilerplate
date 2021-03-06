package database

import (
	"database/sql"
	"fmt"
	"log"
	"serqol/go-demo/utils"
	"sync"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

var instance *Database
var once sync.Once

func Instance() *Database {
	once.Do(func() {
		host, user, password := utils.GetEnv("DB_HOST", "postgres"), utils.GetEnv("DB_USER", "postgres"), utils.GetEnv("DB_PASSWORD", "postgres")
		connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=golang sslmode=disable", host, user, password)
		connection, err := sql.Open("postgres", connStr)
		if err == nil {
			err = connection.Ping()
		}
		if err != nil {
			log.Fatal(err)
		}
		instance = &Database{connection}
	})
	return instance
}
