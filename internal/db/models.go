package db

import "database/sql"

func (d *DBConnection) Close() error {
	return d.DB.Close()
}
func (d *DBConnection) GetDB() *sql.DB {
	return d.DB
}
