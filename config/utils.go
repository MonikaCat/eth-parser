package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	HomePath = ""
)

// GetCfgFilePath returns the confg.yaml file path
func GetCfgFilePath() string {
	return path.Join(HomePath, "config.yaml")
}

func GetHomePath() string {
	return HomePath
}

// GetConfigPreRunE allows to read the config file before executing the cobra command
func GetConfigPreRunE(cmdCfg *Config, name string) func(cmd *cobra.Command, args []string) error {
	return func(_ *cobra.Command, _ []string) error {
		// get the config.yaml file home path
		home, _ := os.UserHomeDir()
		HomePath = path.Join(home, fmt.Sprintf(".%s", name))
		cfgFilePath := GetCfgFilePath()

		// check if config.yaml file exists
		if _, err := os.Stat(cfgFilePath); os.IsNotExist(err) {
			return fmt.Errorf("config.yaml file not found at %s", cfgFilePath)
		}

		// read the config.yaml file
		data, err := os.ReadFile(cfgFilePath)
		if err != nil {
			return fmt.Errorf("error while reading config.yaml file: %w", err)
		}

		// unmarshal the config.yaml file into Config type
		if err := yaml.Unmarshal(data, cmdCfg); err != nil {
			return fmt.Errorf("error while unmarshaling config.yaml file: %w", err)
		}

		return nil
	}
}
