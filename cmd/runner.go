package cmd

import (
	"fmt"
	"os"
	"path"

	initcmd "github.com/MonikaCat/eth-parser/cmd/init"
	parsecmd "github.com/MonikaCat/eth-parser/cmd/parse"
	preparecmd "github.com/MonikaCat/eth-parser/cmd/prepare"
	versioncmd "github.com/MonikaCat/eth-parser/cmd/version"
	"github.com/spf13/cobra"
)

var (
	HomeFlag = "home"
)

// Runner builds cobra command that contains a root command for executing
// the default sub-commands implementations
func Runner(config *Config) *cobra.Command {
	// create a new root command
	rootCmd := RootCmd(config.GetName())

	rootCmd.AddCommand(
		initcmd.InitCmd(config.GetInitConfig()),
		parsecmd.ParseCmd(config.config, config.GetName()),
		preparecmd.PrepareDatabaseCmd(config.config, config.GetName()),
		versioncmd.VersionCmd(),
	)

	return rootCmd
}

// RootCmd allows to build a default root command with the given name
func RootCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("%s is a simple block and USDC txs parser for Ethereum chains", name),
	}

	// get home directory path
	home, _ := os.UserHomeDir()
	// create default config path
	defaultConfigPath := path.Join(home, fmt.Sprintf(".%s", name))
	// add persistent flag to allow user to set the home folder of the application
	cmd.PersistentFlags().String(HomeFlag, defaultConfigPath, "Set the home folder to store config.yaml and other files")

	return cmd

}
