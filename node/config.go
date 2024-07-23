package node

// NodeConfig contains the node configuration
type NodeConfig struct {
	RPC string `yaml:"rpc_url"`
}

// NewNodeConfig creates a new NodeConfig instance
func NewNodeConfig(rpc string) *NodeConfig {
	return &NodeConfig{
		RPC: rpc,
	}
}

// DefaultNodeConfig returns the default node configuration
func DefaultNodeConfig() *NodeConfig {
	return NewNodeConfig(
		"http://localhost:8545",
	)
}

// GetRPC returns the RPC URL
func (n *NodeConfig) GetRPC() string {
	return n.RPC
}
