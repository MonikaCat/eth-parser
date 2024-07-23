package config

import (
	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/node"
)

// Config contains all configuration for config file
type Config struct {
	Node     node.NodeConfig         `yaml:"node"`
	Database database.DatabaseConfig `yaml:"database"`
}

// NewConfig creates a new Config instance
func NewConfig(nodeCfg node.NodeConfig, dbCfg database.DatabaseConfig) Config {
	return Config{
		Node:     nodeCfg,
		Database: dbCfg,
	}
}

// GetDefaultConfig returns the default configuration
func GetDefaultConfig() Config {
	return NewConfig(
		*node.DefaultNodeConfig(),
		*database.DefaultDatabaseConfig(),
	)
}
