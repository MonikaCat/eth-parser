package preparedatabase

import (
	"fmt"

	"github.com/MonikaCat/eth-parser/cmd/parse"
	"github.com/MonikaCat/eth-parser/config"
	"github.com/spf13/cobra"
)

// PrepareDatabaseCmd returns the command that allows to prepare database tables
func PrepareDatabaseCmd(cfg *config.Config, appName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "prepare-database",
		Short:   "Prepare database and create required tables",
		PreRunE: config.GetConfigPreRunE(cfg, appName),
		RunE: func(_ *cobra.Command, _ []string) error {
			// create a new parser config
			parser, err := parse.NewParserConfig(cfg)
			if err != nil {
				return fmt.Errorf("failed to create parser config %v", err)
			}

			// prepare the database tables
			err = parser.Database.PrepareDatabaseTables()
			if err != nil {
				return fmt.Errorf("error while preparing the database tables: %v", err)
			}

			return nil
		},
	}

	return cmd
}
