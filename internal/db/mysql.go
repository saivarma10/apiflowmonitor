package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:passwd@tcp(0.0.0.0:3306)/user")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Print("Pong\n")
	defer db.Close()
}
func createDatabase() {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS api_monitor")
	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully Created\n")
	defer db.Close()

}
