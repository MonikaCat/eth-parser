package node

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

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

	logsBloomJSON, err := json.Marshal(block.Bloom())
	if err != nil {
		return types.Block{}, ethtypes.Transactions{}, fmt.Errorf("error marshalling logsBloom: %v", err)
	}

	blockDetails := types.Block{
		BlockNumber:           block.Number().Int64(),
		BlockHash:             block.Hash().String(),
		Miner:                 block.Coinbase().String(),
		BaseFeePerGas:         block.BaseFee(),
		BlobGasUsed:           block.GasUsed(),
		Difficulty:            block.Difficulty().Uint64(),
		ExcessBlobGas:         *block.ExcessBlobGas(),
		ExtraData:             hex.EncodeToString(block.Extra()),
		GasLimit:              block.GasLimit(),
		GasUsed:               block.GasUsed(),
		LogsBloom:             string(logsBloomJSON),
		MixHash:               block.MixDigest().String(),
		Nonce:                 block.Nonce(),
		ParentBeaconBlockRoot: block.BeaconRoot().String(),
		ParentHash:            block.ParentHash().String(),
		ReceiptsRoot:          block.ReceiptHash().String(),
		Sha3Uncles:            block.UncleHash().String(),
		BlockSize:             block.Size(),
		StateRoot:             block.Root().String(),
		Timestamp:             time.Unix(int64(block.Time()), 0),
		TotalDifficulty:       "0",
	}

	return blockDetails, block.Transactions(), nil
}
