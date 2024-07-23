package database

// DatabaseConfig contains configuration for database
type DatabaseConfig struct {
	DNS                string `yaml:"dns"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
}

// NewDatabaseConfig creates a new DatabaseConfig instance
func NewDatabaseConfig(dns string, maxOpen, maxIdle int) DatabaseConfig {
	return DatabaseConfig{
		DNS:                dns,
		MaxOpenConnections: maxOpen,
		MaxIdleConnections: maxIdle,
	}
}

// DefaultDatabaseConfig returns the default database configuration
func DefaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DNS:                "test.db",
		MaxOpenConnections: 20,
		MaxIdleConnections: 10,
	}
}
