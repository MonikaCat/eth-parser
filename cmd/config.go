package cmd

import (
	initcmd "github.com/MonikaCat/eth-parser/cmd/init"
	"github.com/MonikaCat/eth-parser/config"
)

// Config keeps all configuration information
type Config struct {
	name string

	// pointer to an instance of the InitConfig type
	// means that it can store the address of a InitConfig instance allowing
	// the Config struct to reference and manipulate the original InitConfig data directly
	initConfig *initcmd.InitConfig
	config     *config.Config
}

// NewConfig allows to build new Config instance
// it returns a pointer to a Config struct
func NewConfig(name string) *Config {
	// return Config address (i.e., a pointer to the new Config instance)
	return &Config{
		name: name,
	}
}

// GetName returns the name of the Config instance
// c is a receiver name, and it is a pointer to the Config struct
// that allows to modify the Config instance directly
func (c *Config) GetName() string {
	return c.name
}

// GetInitConfig returns the init config
func (c *Config) GetInitConfig() *initcmd.InitConfig {
	if c.initConfig == nil {
		return initcmd.NewInitConfig()
	}
	return c.initConfig
}

// GetConfig returns the parser config
func (c *Config) GetConfig() *config.Config {
	if c.config == nil {
		return &config.Config{}
	}
	return c.config
}

// WithConfig returns the config with provided parser config
func (c *Config) WithConfig(cfg *config.Config) *Config {
	c.config = cfg
	return c
}
