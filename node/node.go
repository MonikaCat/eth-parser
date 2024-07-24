package node

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// NodeConfig contains all configuration for the node
type Node struct {
	client *ethclient.Client
	ctx    context.Context
	rpc    *rpc.Client
}

// NewNode creates a new Node instance
func NewNode(cfg NodeConfig) (*Node, error) {
	client, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %s", err)
	}

	rpc, err := rpc.Dial(cfg.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum RPC: %s", err)
	}

	return &Node{
		ctx:    context.Background(),
		client: client,
		rpc:    rpc,
	}, nil
}
