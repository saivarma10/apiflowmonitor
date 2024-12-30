package db

import (
	"apimonitor/pkg/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

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
	for _, table := range createTablesSQL{
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}
	fmt.Print("Tables Created Successfully\n")
	return nil
}

func Close() {
	db.Close()
}

func StoreTransaction(transaction *utils.Transactions, response []map[string]interface{})error{
	
	// Insert transactions
	result, err := db.Exec("INSERT INTO transactions (name) VALUES (?)", transaction.Name)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID for transaction: %v", err)
	}

	// Insert transaction APIs
	for sequence, api := range transaction.APIs {
		dependencyJSON, err := json.Marshal(api.Dependency)
		if err != nil {
			return fmt.Errorf("failed to marshal dependency: %v", err)
		}

		result, err := db.Exec("INSERT INTO api_config (url, method, request) VALUES (?, ?, ?)", api.URL, api.Method, api.Request)

		if err != nil {
			return fmt.Errorf("failed to insert API config: %v", err)
		}
		apiID, err := result.LastInsertId()
		_, err = db.Exec(
			"INSERT INTO transaction_api (transaction_id, api_id, sequence, dependency) VALUES (?, ?, ?, ?)",
			transactionID, apiID, sequence, string(dependencyJSON),
		)
		if err != nil {
			return fmt.Errorf("failed to insert transaction API: %v", err)
		}

		// Insert response data
		responseJSON, err := json.Marshal(response[sequence])
		if err != nil {
			return fmt.Errorf("failed to marshal response data: %v", err)
		}

		_, err = db.Exec("INSERT INTO reponse_data (transaction_id, api_id, response) VALUES (?, ?, ?)", transactionID, apiID, string(responseJSON))
		if err != nil {
			return fmt.Errorf("failed to insert response data: %v", err)
		}
	}
	
	return nil
}


func GetTransaction(transactionID int) (*utils.TransactionResponse, error) {
	transactionResp := &utils.TransactionResponse{}

	transactionResp.TransactionID = strconv.Itoa(transactionID)

	// Get transaction name
	err := db.QueryRow("SELECT name FROM transactions WHERE id = ?", transactionID).Scan(&transactionResp.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction name: %v", err)
	}

	// get apis
	rows, err := db.Query("SELECT url, method, request FROM api_config WHERE id IN (SELECT api_id FROM transaction_api WHERE transaction_id = ?)", transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction APIs: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		api := utils.TransactionAPIResp{}
		err := rows.Scan(&api.URL, &api.Method, &api.Request)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction API: %v", err)
		}
		transactionResp.APIs = append(transactionResp.APIs, api)
	}

	// Get transaction APIs
	rows, err = db.Query("SELECT api_id, sequence, dependency FROM transaction_api WHERE transaction_id = ?", transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction APIs: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dependency map[string]utils.Dependency
		var dependencyJSON string
		var apiID, sequence int

		err := rows.Scan(&apiID, &sequence, &dependencyJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction API: %v", err)
		}

		err = json.Unmarshal([]byte(dependencyJSON), &dependency)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal dependency: %v", err)
		}

		transactionResp.APIs[sequence].Dependency = dependency
	}

	return transactionResp, nil
}
