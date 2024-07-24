package node

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/MonikaCat/eth-parser/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// GetBlock queries block by number
// it returns parsed block details, transactions and error
func (n *Node) GetBlock(blockNumber big.Int) (types.Block, ethtypes.Transactions, error) {

	block, err := n.client.BlockByNumber(context.Background(), &blockNumber)
	if err != nil {
		return types.Block{}, ethtypes.Transactions{}, fmt.Errorf("error while getting block %v: error: %v ", blockNumber, err)
	}

	blockDetails, txs, err := n.ParseBlockDetails(block)
	if err != nil {
		return types.Block{}, ethtypes.Transactions{}, fmt.Errorf("error while parsing block details: %v", err)
	}

	return blockDetails, txs, nil
}

// ParseBlockDetails parses block details
// it returns block details, transactions and error
func (n *Node) ParseBlockDetails(block *ethtypes.Block) (types.Block, ethtypes.Transactions, error) {

	// marshal the logsBloom
	logsBloomJSON, err := json.Marshal(block.Bloom())
	if err != nil {
		return types.Block{}, ethtypes.Transactions{}, fmt.Errorf("error marshalling logsBloom: %v", err)
	}

	// parse the total difficulty of the block
	totalDifficulty, err := n.parseTotalDifficulty(block.Number())
	if err != nil {
		return types.Block{}, ethtypes.Transactions{}, fmt.Errorf("error parsing total difficulty: %v", err)
	}

	// create a new block details
	blockDetails := types.NewBlock(
		BigIntToHex(block.Number()),
		block.Hash().String(),
		block.Coinbase().String(),
		BigIntToHex(block.BaseFee()),
		Uint64ToHex(*block.BlobGasUsed()),
		BigIntToHex(block.Difficulty()),
		Uint64ToHex(*block.ExcessBlobGas()),
		StringToHex(hex.EncodeToString(block.Extra())),
		Uint64ToHex(block.GasLimit()),
		Uint64ToHex(block.GasUsed()),
		string(logsBloomJSON),
		block.MixDigest().String(),
		Uint64ToHex(block.Nonce()),
		block.BeaconRoot().String(),
		block.ParentHash().String(),
		block.ReceiptHash().String(),
		block.UncleHash().String(),
		Uint64ToHex(block.Size()),
		block.Root().String(),
		Uint64ToHex(block.Time()),
		BigIntToHex(totalDifficulty),
	)

	return blockDetails, block.Transactions(), nil
}

// parseTotalDifficulty calls the RPC endpoint
// and parses the total difficulty of the block
func (n *Node) parseTotalDifficulty(blockNumber *big.Int) (*big.Int, error) {
	var results map[string]interface{}
	// get the total difficulty of the block
	err := n.rpc.CallContext(n.ctx, &results, "eth_getBlockByNumber", fmt.Sprintf("0x%x", blockNumber), false)
	if err != nil {
		return nil, fmt.Errorf("error getting block from rpc: error: %v", err)
	}

	// get the total difficulty from the results
	totalDiff, ok := results["totalDifficulty"].(string)
	if !ok {
		return nil, fmt.Errorf("error while getting block total difficulty")
	}

	// convert the total difficulty to a big int
	totalDifficulty := new(big.Int)
	// remove the 0x prefix
	totalDifficulty.SetString(totalDiff[2:], 16)

	return totalDifficulty, nil
}
