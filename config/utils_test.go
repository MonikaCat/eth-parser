package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MonikaCat/eth-parser/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// TestGetCfgFilePath tests the GetCfgFilePath function
func TestGetCfgFilePath(t *testing.T) {
	expectedFilePath := filepath.Join(config.HomePath, "config.yaml")
	actualFilePath := config.GetCfgFilePath()

	assert.Equal(t, expectedFilePath, actualFilePath, "File path should match")
}

// TestGetConfigPreRunE tests the GetConfigPreRunE function
func TestGetConfigPreRunE(t *testing.T) {
	// create a new Config object
	cfg := &config.Config{}
	cmd := &cobra.Command{}

	// get the config.yaml file home path
	home, _ := os.UserHomeDir()
	config.HomePath = filepath.Join(home, ".eth-parser")

	// get the GetConfigPreRunE function
	preRunE := config.GetConfigPreRunE(cfg, "eth-parser")

	err := preRunE(cmd, []string{})
	assert.NoError(t, err)
	assert.NotEmpty(t, cfg.Node.RPC, "RPC should not be empty")
	assert.NotEmpty(t, cfg.Database.DNS, "DNS should not be empty")
	assert.NotEmpty(t, cfg.Database.MaxOpenConnections, "MaxOpenConnections should not be empty")
	assert.NotEmpty(t, cfg.Database.MaxOpenConnections, "MaxOpenConnections should not be empty")
}
