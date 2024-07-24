package database

import (
	_ "embed"
	"fmt"

	"github.com/MonikaCat/eth-parser/types"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

//go:embed sql/insert_block.sql
var insertBlockQuerySQL string

//go:embed sql/select_block.sql
var getBlockQuerySQL string

// SaveBlock saves a block to the database
func (db *Database) SaveBlock(block types.Block) error {
	_, err := db.SQL.NamedExec(insertBlockQuerySQL, block)
	if err != nil {
		return fmt.Errorf("error while inserting block %s: %v", block.BlockNumber, err)
	}

	return nil
}

// GetBlock returns requested block from the database
func (db *Database) GetBlock(blockNumber string) (types.Block, error) {
	var block types.Block
	err := db.SQL.Get(&block, getBlockQuerySQL, blockNumber)
	if err != nil {
		return types.Block{}, fmt.Errorf("error while getting details from db for block %s: %v", blockNumber, err)
	}

	return block, nil
}
