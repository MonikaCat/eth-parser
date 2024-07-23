package node

import "github.com/ethereum/go-ethereum/ethclient"

type Node struct {
	client *ethclient.Client
}

func NewNode(client *ethclient.Client) Node {
	return Node{
		client: client,
	}
}