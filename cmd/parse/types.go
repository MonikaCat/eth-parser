package parse

import (
	"fmt"

	"github.com/MonikaCat/eth-parser/config"
	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/node"
)

type ParserConfig struct {
	Node     *node.Node
	Database *database.Database
}

func NewParserConfig(cfg *config.Config) (*ParserConfig, error) {
	node, err := node.NewNode(cfg.Node)
	if err != nil {
		return nil, fmt.Errorf("error while creating new node instance: %s", err)
	}

	db, err := database.ConnectToDatabase(cfg.Database)
	if err != nil {
		fmt.Printf("error while connecting to database: error: %s", err)
	}

	return &ParserConfig{
		Node:     node,
		Database: db,
	}, nil
}
