package db

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"apimonitor/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

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


func StoreConfigInDB(config *config.ConfigFile) error {
	// Insert API configurations
	for _, api := range config.APIs {
		_, err := db.Exec(
			"INSERT INTO api_config (url, method, request_structure, response_structure) VALUES (?, ?, ?, ?)",
			api.URL, api.Method, api.RequestStructure, api.ResponseStructure,
		)
		if err != nil {
			return fmt.Errorf("failed to insert API config: %v", err)
		}
	}

	// Insert transactions
	for _, transaction := range config.Transactions {
		result, err := db.Exec("INSERT INTO transaction (name) VALUES (?)", transaction.Name)
		if err != nil {
			return fmt.Errorf("failed to insert transaction: %v", err)
		}

		transactionID, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get last insert ID for transaction: %v", err)
		}

		// Insert transaction APIs
		for _, api := range transaction.APIs {
			dependencyJSON, err := json.Marshal(api.Dependency)
			if err != nil {
				return fmt.Errorf("failed to marshal dependency: %v", err)
			}

			_, err = db.Exec(
				"INSERT INTO transaction_api (transaction_id, api_id, sequence, dependency) VALUES (?, ?, ?, ?)",
				transactionID, api.APIIndex+1, api.Sequence, string(dependencyJSON),
			)
			if err != nil {
				return fmt.Errorf("failed to insert transaction API: %v", err)
			}
		}
	}

	return nil
}