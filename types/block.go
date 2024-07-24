package types

// Block represents an ethereum block
type Block struct {
	BlockNumber           string `json:"block_number" db:"block_number"`
	BlockHash             string `json:"block_hash" db:"block_hash"`
	ParentHash            string `json:"parent_hash" db:"parent_hash"`
	Nonce                 string `json:"nonce" db:"nonce"`
	Miner                 string `json:"miner" db:"miner"`
	BaseFeePerGas         string `json:"base_fee_per_gas" db:"base_fee_per_gas"`
	BlobGasUsed           string `json:"blob_gas_used" db:"blob_gas_used"`
	Difficulty            string `json:"difficulty" db:"difficulty"`
	ExcessBlobGas         string `json:"excess_blob_gas" db:"excess_blob_gas"`
	ExtraData             string `json:"extra_data" db:"extra_data"`
	GasLimit              string `json:"gas_limit" db:"gas_limit"`
	GasUsed               string `json:"gas_used" db:"gas_used"`
	LogsBloom             string `json:"logs_bloom" db:"logs_bloom"`
	MixHash               string `json:"mix_hash" db:"mix_hash"`
	ParentBeaconBlockRoot string `json:"parent_beacon_block_root" db:"parent_beacon_block_root"`
	ReceiptsRoot          string `json:"receipts_root" db:"receipts_root"`
	Sha3Uncles            string `json:"sha3_uncles" db:"sha3_uncles"`
	BlockSize             string `json:"block_size" db:"block_size"`
	StateRoot             string `json:"state_root" db:"state_root"`
	Timestamp             string `json:"timestamp" db:"timestamp"`
	TotalDifficulty       string `json:"total_difficulty" db:"total_difficulty"`
}

// NewBlock creates a new Block instance
func NewBlock(blockNumber string, blockHash,
	miner, baseFeePerGas, blobGasUsed, difficulty,
	excessBlobGas, extraData, gasLimit, gasUsed,
	logsBloom, mixHash, nonce, parentBeaconBlockRoot,
	parentHash, receiptsRoot, sha3Uncles, blockSize,
	stateRoot string, timestamp, totalDifficulty string,
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
