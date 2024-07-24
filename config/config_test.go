package config_test

import (
	"testing"

	"github.com/MonikaCat/eth-parser/config"
	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/node"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewConfig tests the NewConfig function
func TestNewConfig(t *testing.T) {
	rpc := "http://localhost:8545"
	dns := "test.db"
	maxOpenConnections := 20
	maxIdleConnections := 10

	cfg := config.NewConfig(*node.NewNodeConfig(rpc), database.NewDatabaseConfig(dns, maxOpenConnections, maxIdleConnections))

	require.NotNil(t, cfg)
	assert.Equal(t, rpc, cfg.Node.RPC)
	assert.Equal(t, dns, cfg.Database.DNS)
	assert.Equal(t, maxOpenConnections, cfg.Database.MaxOpenConnections)
	assert.Equal(t, maxIdleConnections, cfg.Database.MaxIdleConnections)
}

// TestDefaultConfig tests the GetDefaultConfig function
func TestDefaultConfig(t *testing.T) {
	rpc := "http://localhost:8545"
	dns := "test.db"
	maxOpenConnections := 20
	maxIdleConnections := 10

	cfg := config.GetDefaultConfig()

	require.NotNil(t, cfg)
	assert.Equal(t, rpc, cfg.Node.RPC)
	assert.Equal(t, dns, cfg.Database.DNS)
	assert.Equal(t, maxOpenConnections, cfg.Database.MaxOpenConnections)
	assert.Equal(t, maxIdleConnections, cfg.Database.MaxIdleConnections)
}
