package node

type NodeConfig struct {
	RPC string `yaml:"rpc_url"`
}

func NewNodeConfig(rpc string) *NodeConfig {
	return &NodeConfig{
		RPC: rpc,
	}
}

func DefaultNodeConfig() *NodeConfig {
	return NewNodeConfig(
		"http://localhost:8545",
	)
}

func (n *NodeConfig) GetRPC() string {
	return n.RPC
}
