package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Create tables
var createTablesSQL = []string{
	`CREATE TABLE IF NOT EXISTS api_config (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL,
		method TEXT NOT NULL,
		request TEXT
	);`,

	`CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`,

	`CREATE TABLE IF NOT EXISTS transaction_api (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		transaction_id INTEGER NOT NULL,
		api_id INTEGER NOT NULL,
		sequence INTEGER NOT NULL,
		dependency TEXT,
		FOREIGN KEY (transaction_id) REFERENCES transactions(id),
		FOREIGN KEY (api_id) REFERENCES api_config(id)
	);`,

	`CREATE TABLE IF NOT EXISTS reponse_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		transaction_id INTEGER NOT NULL,
		api_id INTEGER NOT NULL,
		response TEXT,
		FOREIGN KEY (transaction_id) REFERENCES transactions(id),
		FOREIGN KEY (api_id) REFERENCES api_config(id)
	);`,

	`CREATE TABLE IF NOT EXISTS api_statistics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		api_id INTEGER NOT NULL,
		status_code INTEGER NOT NULL,
		response_time REAL NOT NULL,
		executed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (api_id) REFERENCES api_config(id)
	);`,
	
}

// var db *sql.DB
// type DBConnection struct {
var db *sql.DB
// }

func Init() error {
	var err error
	db, err = sql.Open("sqlite3", "./apimonitor.db")
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	return createDatabase()
	
}
func createDatabase() error {
	for i, table := range createTablesSQL{
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
		if i==1{break}
	}
	fmt.Print("Tables Created Successfully\n")
	return nil
}

func Close() {
	db.Close()
}

