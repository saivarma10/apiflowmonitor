package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() {
	db, err := sql.Open("sqlite3", "./apimonitor.db")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Print("Pong\n")
	createDatabase()
}
func createDatabase() {
	_, err := db.Exec(createTablesSQL)
	if err != nil {
		panic(err)
	}
	fmt.Print("Tables Created Successfully\n")
	defer db.Close()

}

func Close() {
	db.Close()
}
