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

	fmt.Printf("\n THE BLOCK BlockNumber %d \n", blockDetails.BlockNumber)
	fmt.Printf("\n THE BLOCK BlockHash %s \n", blockDetails.BlockHash)
	fmt.Printf("\n THE BLOCK Miner %s \n", blockDetails.Miner)
	fmt.Printf("\n THE BLOCK BaseFeePerGas %v \n", blockDetails.BaseFeePerGas)
	fmt.Printf("\n THE BLOCK BlobGasUsed %d \n", blockDetails.BlobGasUsed)
	fmt.Printf("\n THE BLOCK Difficulty %d \n", blockDetails.Difficulty)
	fmt.Printf("\n THE BLOCK ExcessBlobGas %v \n", blockDetails.ExcessBlobGas)
	fmt.Printf("\n THE BLOCK ExtraData %s \n", blockDetails.ExtraData)
	fmt.Printf("\n THE BLOCK GasLimit %d \n", blockDetails.GasLimit)
	fmt.Printf("\n THE BLOCK GasUsed %d \n", blockDetails.GasUsed)
	fmt.Printf("\n THE BLOCK LogsBloom %s \n", blockDetails.LogsBloom)
	fmt.Printf("\n THE BLOCK MixHash %s \n", blockDetails.MixHash)
	fmt.Printf("\n THE BLOCK Nonce %d \n", blockDetails.Nonce)
	fmt.Printf("\n THE BLOCK ParentBeaconBlockRoot %s \n", blockDetails.ParentBeaconBlockRoot)
	fmt.Printf("\n THE BLOCK ParentHash %s \n", blockDetails.ParentHash)
	fmt.Printf("\n THE BLOCK ReceiptsRoot %s \n", blockDetails.ReceiptsRoot)
	fmt.Printf("\n THE BLOCK Sha3Uncles %s \n", blockDetails.Sha3Uncles)
	fmt.Printf("\n THE BLOCK BlockSize %d \n", blockDetails.BlockSize)
	fmt.Printf("\n THE BLOCK StateRoot %s \n", blockDetails.StateRoot)
	fmt.Printf("\n THE BLOCK Timestamp %v \n", blockDetails.Timestamp)

	return blockDetails, nil
}
