package cmd_test

import (
	"testing"

	"github.com/MonikaCat/eth-parser/cmd"
	"github.com/MonikaCat/eth-parser/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRunner tests the Runner function
func TestRunner(t *testing.T) {
	config := cmd.NewConfig("eth-parser").
		WithConfig(&config.Config{})

	cmd := cmd.Runner(config)

	require.NotNil(t, cmd)
	assert.Equal(t, "eth-parser", cmd.Use)
}

// TestRootCmd tests the RootCmd function
func TestRootCmd(t *testing.T) {
	cmd := cmd.RootCmd("eth-parser")

	require.NotNil(t, cmd)
	assert.Equal(t, "eth-parser", cmd.Use)
	assert.Equal(t, "eth-parser is a simple block and USDC txs parser for Ethereum chains", cmd.Short)
}
