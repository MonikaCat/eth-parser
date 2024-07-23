package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Database represents the database connection
type Database struct {
	SQL *sqlx.DB
}

// ConnectToDatabase handles connection to the database
func ConnectToDatabase(cfg DatabaseConfig) (*Database, error) {
	db, err := sqlx.Open("sqlite3", cfg.DNS)
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
