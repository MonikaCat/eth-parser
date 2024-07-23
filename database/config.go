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
