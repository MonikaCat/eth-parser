package parse

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/MonikaCat/eth-parser/config"

	"github.com/spf13/cobra"
)

// ParseCmd returns the command that allows to parse
// USDC transfer transactions for a given block number
func ParseCmd(cfg *config.Config, appName string) *cobra.Command {
	return &cobra.Command{
		Use:     "parse [block number]",
		Short:   "Parse USDC transfer transactions for a given block number",
		Args:    cobra.ExactArgs(1),
		PreRunE: config.GetConfigPreRunE(cfg, appName),
		RunE: func(_ *cobra.Command, args []string) error {

			blockHeight, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("make sure the given start height is not null")
			}

			return parse(cfg, blockHeight)
		},
	}
}

// parse represents the function that should be called
// when the parse command is executed
func parse(cfg *config.Config, blockHeight int64) error {

	parser, err := NewParserConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to create parser config")
	}

	block, txs, err := parser.Node.GetBlock(*big.NewInt(blockHeight))
	if err != nil {
		return fmt.Errorf("failed to get block: %s", err)
	}

	err = parser.Database.SaveBlock(block)
	if err != nil {
		return fmt.Errorf("error while saving block: %v", err)
	}

	for _, tx := range txs {
		tx, err := parser.Node.GetTransaction(tx)
		if err != nil {
			return fmt.Errorf("error while getting transaction: %v", err)
		}

		err = parser.Database.SaveTransaction(tx)
		if err != nil {
			return fmt.Errorf("error while saving transaction: %v", err)
		}
	}

	return nil
}
