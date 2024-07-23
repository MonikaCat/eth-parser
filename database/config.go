package database

type DatabaseConfig struct {
	DNS                string `yaml:"dns"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
}

func NewDatabaseConfig(dns string, maxOpen, maxIdle int) DatabaseConfig {
	return DatabaseConfig{
		DNS:                dns,
		MaxOpenConnections: maxOpen,
		MaxIdleConnections: maxIdle,
	}
}

func DefaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DNS:                "test.db",
		MaxOpenConnections: 20,
		MaxIdleConnections: 10,
	}
}
