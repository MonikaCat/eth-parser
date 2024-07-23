package config

import (
	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/node"
)

type Config struct {
	Node     node.NodeConfig         `yaml:"node"`
	Database database.DatabaseConfig `yaml:"database"`
}

func NewConfig(nodeCfg node.NodeConfig, dbCfg database.DatabaseConfig) Config {
	return Config{
		Node:     nodeCfg,
		Database: dbCfg,
	}
}

func GetDefaultConfig() Config {
	return NewConfig(
		*node.DefaultNodeConfig(),
		*database.DefaultDatabaseConfig(),
	)
}
