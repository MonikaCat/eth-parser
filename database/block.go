package database

import (
	"fmt"
	"os"

	"github.com/MonikaCat/eth-parser/types"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

func (db *Database) SaveBlock(block types.Block) error {
	blockQuery, err := os.ReadFile("database/sql/insert_block.sql")
	if err != nil {
		return fmt.Errorf("error reading database/sql/insert_block.sql: %v", err)
	}

	_, err = db.SQL.NamedExec(string(blockQuery), block)
	if err != nil {
		return fmt.Errorf("error while inserting block %d: %v", block.BlockNumber, err)
	}

	return nil
}
