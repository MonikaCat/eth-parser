package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	SQL *sql.DB
}

func ConnectToDatabase(cfg DatabaseConfig) (*Database, error) {
	db, err := sql.Open("sqlite3", cfg.DNS)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to the database: %v", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error while pinging the database: %v", err)
	}

	return &Database{SQL: db}, nil
}
