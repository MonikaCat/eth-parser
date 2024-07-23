package types

import (
	"math/big"
	"time"
)

type Block struct {
	BlockNumber           int64     `json:"block_number"`
	BlockHash             string    `json:"block_hash"`
	Miner                 string    `json:"miner"`
	BaseFeePerGas         *big.Int  `json:"base_fee_per_gas"`
	BlobGasUsed           uint64    `json:"blob_gas_used"`
	Difficulty            uint64    `json:"difficulty"`
	ExcessBlobGas         uint64    `json:"excess_blob_gas"`
	ExtraData             string    `json:"extra_data"` // TODO: change type
	GasLimit              uint64    `json:"gas_limit"`
	GasUsed               uint64    `json:"gas_used"`
	LogsBloom             string    `json:"logs_bloom"`
	MixHash               string    `json:"mix_hash"`
	Nonce                 uint64    `json:"nonce"`
	ParentBeaconBlockRoot string    `json:"parent_beacon_block_root"`
	ParentHash            string    `json:"parent_hash"`
	ReceiptsRoot          string    `json:"receipts_root"`
	Sha3Uncles            string    `json:"sha3_uncles"`
	BlockSize             uint64    `json:"block_size"`
	StateRoot             string    `json:"state_root"`
	Timestamp             time.Time `json:"timestamp"`
	TotalDifficulty       string    `json:"total_difficulty"`
}

func NewBlock(blockNumber int64, blockHash,
	miner string, baseFeePerGas *big.Int, blobGasUsed, difficulty uint64,
	excessBlobGas uint64, extraData string, gasLimit, gasUsed uint64,
	logsBloom string, mixHash string, nonce uint64, parentBeaconBlockRoot,
	parentHash, receiptsRoot, sha3Uncles string, blockSize uint64,
	stateRoot string, timestamp time.Time, totalDifficulty string,
) Block {
	return Block{
		BlockNumber:           blockNumber,
		BlockHash:             blockHash,
		Miner:                 miner,
		BaseFeePerGas:         baseFeePerGas,
		BlobGasUsed:           blobGasUsed,
		Difficulty:            difficulty,
		ExcessBlobGas:         excessBlobGas,
		ExtraData:             extraData,
		GasLimit:              gasLimit,
		GasUsed:               gasUsed,
		LogsBloom:             logsBloom,
		MixHash:               mixHash,
		Nonce:                 nonce,
		ParentBeaconBlockRoot: parentBeaconBlockRoot,
		ParentHash:            parentHash,
		ReceiptsRoot:          receiptsRoot,
		Sha3Uncles:            sha3Uncles,
		BlockSize:             blockSize,
		StateRoot:             stateRoot,
		Timestamp:             timestamp,
		TotalDifficulty:       totalDifficulty,
	}
}
