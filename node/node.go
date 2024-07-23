package node

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Node struct {
	client *ethclient.Client
	ctx    context.Context
}

func NewNode(cfg NodeConfig) (*Node, error) {
	client, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %s", err)
	}

	return &Node{
		ctx:    context.Background(),
		client: client,
	}, nil
}
