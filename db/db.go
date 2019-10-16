package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//Db - connection to database
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		panic(err)
	}
}
