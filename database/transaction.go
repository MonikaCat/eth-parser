package database

import (
	"fmt"
	"os"

	"github.com/MonikaCat/eth-parser/types"
)

// SaveTransaction saves a transaction to the database
func (db *Database) SaveTransaction(tx types.Transaction) error {
	txQuery, err := os.ReadFile("database/sql/insert_transaction.sql")
	if err != nil {
		return fmt.Errorf("error reading database/sql/insert_transaction.sql: %v", err)
	}

	_, err = db.SQL.NamedExec(string(txQuery), tx)
	if err != nil {
		return fmt.Errorf("error while inserting tx: %v", err)
	}

	return nil
}
