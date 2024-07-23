package init

import (
	"github.com/MonikaCat/eth-parser/config"
	"github.com/spf13/cobra"
)

type ConfigCreator func(cmd *cobra.Command) interface{}

// DefaultConfigCreator creates a default configuration
func DefaultConfigCreator(_ *cobra.Command) interface{} {
	return config.GetDefaultConfig()
}

// InitConfig holds  a function that can
// create a ConfigWriter based on a *cobra.Command
type InitConfig struct {
	createConfig ConfigCreator
}

// NewInitConfig builds new InitConfig instance
// initially it will be nil
func NewInitConfig() *InitConfig {
	return &InitConfig{}
}

// GetConfigCreator returns the ConfigCreator stored in the createConfig field.
// If createConfig is nil, it returns the DefaultConfigCreator function.
func (c *InitConfig) GetConfigCreator() ConfigCreator {
	if c.createConfig == nil {
		return DefaultConfigCreator
	}
	return c.createConfig
}