package database

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

func All(table string) *sql.Rows {
	db := Instance()
	closure := func() *sql.Rows {
		rows, _ := squirrel.Select("*").From(table).RunWith(db.Connection).Query()
		defer rows.Close()
		return rows
	}
	channel := make(chan *sql.Rows)
	go execute(channel, closure)
	entries := <-channel
	return entries
}

func execute(channel chan *sql.Rows, closure func() *sql.Rows) {
	channel <- closure()
}
