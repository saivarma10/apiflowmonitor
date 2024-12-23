package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// var db *sql.DB
// type DBConnection struct {
var db *sql.DB
// }

func Init() error {
	db, err := sql.Open("sqlite3", "./apimonitor.db")
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Print("Pong\n")
	return createDatabase()
	
}
func createDatabase() error {
	_, err := db.Exec(createTablesSQL)
	if err != nil {
		return err
	}
	fmt.Print("Tables Created Successfully\n")
	return nil
}

func Close() {
	db.Close()
}
