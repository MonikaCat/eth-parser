package config

import (
	"github.com/MonikaCat/eth-parser/database"
)

type Config struct {
	Database database.DatabaseConfig `yaml:"database"`
}
