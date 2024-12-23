package db

import (
	"database/sql"
	"fmt"
)

// var db *sql.DB
type DBConnection struct {
	DB *sql.DB
}

func NewConnection(host, user, password, dbname string, port int) (*DBConnection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return &DBConnection{DB: db}, nil
}
