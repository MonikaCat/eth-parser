package utils_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/MonikaCat/eth-parser/config"
	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/node"
	"github.com/MonikaCat/eth-parser/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCreateNewFile tests the CreateNewFile function
func TestCreateNewFile(t *testing.T) {
	// test valid file creation
	tmpfile, err := os.CreateTemp("", "test_file.yaml")
	require.NoError(t, err, "Should return no error for valid file creation")

	defer os.Remove(tmpfile.Name())

	file, err := utils.CreateNewFile(tmpfile.Name())
	require.NoError(t, err, "Should return no error for valid file creation")
	require.NotNil(t, file, "Should return non-nil file handle for valid file creation")
	if file != nil {
		file.Close()
	}

	// test invalid file path
	invalidFileName := "/invalid/path/to/file.yaml"
	file, err = utils.CreateNewFile(invalidFileName)
	require.Error(t, err, "Should return error for invalid file path")
	require.Nil(t, file, "Should return nil file handle for invalid file path")
}

// TestSerializeToYAML tests the SerializeToYAML function
func TestSerializeToYAML(t *testing.T) {
	// create a buffer to capture the serialized YAML output
	var buf bytes.Buffer

	// create a new Config object
	cfg := &config.Config{
		Node: node.NodeConfig{
			RPC: "http://localhost:8545",
		},
		Database: database.DatabaseConfig{
			DNS:                "test.db",
			MaxOpenConnections: 20,
			MaxIdleConnections: 10,
		},
	}

	err := utils.SerializeToYAML(&buf, cfg)
	if err != nil {
		t.Fatalf("SerializeToYAML failed: %v", err)
	}

	// define expected YAML output as a string
	expectedYAML := `node:
  rpc_url: http://localhost:8545
database:
  dns: test.db
  max_open_connections: 20
  max_idle_connections: 10
`
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, buf.String(), "Serialized YAML does not match expected")
	assert.NotEmpty(t, buf.String(), "Serialized YAML should not be empty")
}
