package database

import (
	_ "embed"
	"fmt"

	"github.com/MonikaCat/eth-parser/types"
)

//go:embed sql/insert_transaction.sql
var insertTxQuerySQL string

// SaveTransaction saves a transaction to the database
func (db *Database) SaveTransaction(tx types.Transaction) error {

	_, err := db.SQL.NamedExec(insertTxQuerySQL, tx)
	if err != nil {
		return fmt.Errorf("error while inserting tx: %v", err)
	}

	return nil
}
