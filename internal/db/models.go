package db


// Create tables
var createTablesSQL = []string{
	`CREATE TABLE IF NOT EXISTS api_config (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL,
		method TEXT NOT NULL,
		request_structure TEXT,
		response_structure TEXT
	);`,

	`CREATE TABLE IF NOT EXISTS transaction_list (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`,

	`CREATE TABLE IF NOT EXISTS transaction_api (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		transaction_id INTEGER NOT NULL,
		api_id INTEGER NOT NULL,
		sequence INTEGER NOT NULL,
		dependency TEXT,
		FOREIGN KEY (transaction_id) REFERENCES transaction(id),
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
