package init

import (
	"fmt"
	"os"

	"github.com/MonikaCat/eth-parser/config"
	"github.com/spf13/cobra"
)

// InitCmd returns the command that allows to properly setup and initialise
// working directory and config.yaml file
func InitCmd(cfg *InitConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialise config.yaml file",
		RunE: func(cmd *cobra.Command, args []string) error {
			// setup home path
			if err := SetupHome(cmd, args); err != nil {
				return err
			}

			// create config path if it doesn't exist
			// it creates new directory if it doesnt exist yet
			if _, err := os.Stat(config.HomePath); os.IsNotExist(err) {
				if err = os.MkdirAll(config.HomePath, os.ModePerm); err != nil {
					return err
				}
			}

			// get the config.yaml file path
			cfgFilePath := config.GetCfgFilePath()
			// return an error if config.yaml file already exists
			if _, err := os.Stat(cfgFilePath); err == nil {
				return fmt.Errorf("config.yaml file already exists at %s", cfgFilePath)
			} else if !os.IsNotExist(err) { // also check !os.IsNotExist(err) to make sure it's not a different error
				return err
			}

			// create the default configuration
			// when we finally have directory and config file created
			yamlCfg := cfg.GetConfigCreator()(cmd)
			// write the configuration into config.yaml file
			return CreateCfgFile(yamlCfg, cfgFilePath)
		},
	}

	return cmd
}
