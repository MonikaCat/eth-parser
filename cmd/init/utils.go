package init

import (
	"fmt"

	"github.com/MonikaCat/eth-parser/config"

	"github.com/MonikaCat/eth-parser/utils"
	"github.com/spf13/cobra"
)

var (
	HomeFlag = "home"
)

// CreateCfgFile marshals data to YAML format and writes it inside config.yaml
func CreateCfgFile(cfg interface{}, path string) error {
	// create new config.yaml file with given path
	cfgFile, err := utils.CreateNewFile(path)
	if err != nil {
		return fmt.Errorf("error while creating config file: %v", err)
	}
	defer cfgFile.Close()

	// encode the data to YAML format
	// and write it to the config.yaml file
	err = utils.SerializeToYAML(cfgFile, cfg)
	if err != nil {
		return fmt.Errorf("error while serializing config file: %v", err)
	}

	return nil
}

// SetupHome sets up the home directory of the init cmd
func SetupHome(cmd *cobra.Command, _ []string) error {
	homePath, err := cmd.Flags().GetString(HomeFlag)
	if err != nil {
		return fmt.Errorf("error while getting home path: %v", err)
	}
	config.HomePath = homePath
	return nil
}
