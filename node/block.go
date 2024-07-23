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

func (n *Node) GetBlock(blockNumber big.Int) (types.Block, error) {

	block, err := n.client.BlockByNumber(context.Background(), &blockNumber)
	if err != nil {
		return types.Block{}, fmt.Errorf("error while getting block %d: error: %v ", blockNumber, err)
	}

	blockDetails, err := n.ParseBlockDetails(block)
	if err != nil {
		return types.Block{}, fmt.Errorf("error while parsing block details: %v", err)
	}

	return blockDetails, nil
}

func (n *Node) ParseBlockDetails(block *ethtypes.Block) (types.Block, error) {

	logsBloomJson, err := json.Marshal(block.Bloom())
	if err != nil {
		return types.Block{}, fmt.Errorf("error marshalling logsBloom: %v", err)
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
		LogsBloom:             string(logsBloomJson),
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

	return blockDetails, nil
}
