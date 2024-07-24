package database

import (
	_ "embed"
	"fmt"

	"github.com/MonikaCat/eth-parser/types"
)

//go:embed sql/insert_transaction.sql
var insertTxQuerySQL string

//go:embed sql/select_transactions_by_block.sql
var getTxsByBlockQuerySQL string

// SaveTransaction saves a transaction to the database
func (db *Database) SaveTransaction(tx types.Transaction) error {

	_, err := db.SQL.NamedExec(insertTxQuerySQL, tx)
	if err != nil {
		return fmt.Errorf("error while inserting tx: %v", err)
	}

	return nil
}

// GetTransactionsByBlock returns requested transactions from the database
func (db *Database) GetTransactionsByBlock(blockNumber string) (types.Transaction, error) {
	var tx types.Transaction
	err := db.SQL.Get(&tx, getTxsByBlockQuerySQL, blockNumber)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("error while getting transactions for block %s: %v", blockNumber, err)
	}

	return tx, nil
}