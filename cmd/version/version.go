package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCmd returns the command that allows to
// display the current version of the application
func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display the version of the application",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("Version v0.1.0")
		},
	}
}
