package types

// Transaction represents an ethereum transaction
type Transaction struct {
	BlockNumber          string `json:"block_number" db:"block_number"`
	BlockHash            string `json:"block_hash" db:"block_hash"`
	From                 string `json:"tx_from" db:"tx_from"`
	To                   string `json:"tx_to" db:"tx_to"`
	TransactionHash      string `json:"transaction_hash" db:"transaction_hash"`
	TransactionIndex     string `json:"transaction_index" db:"transaction_index"`
	Value                string `json:"tx_value" db:"tx_value"`
	Type                 string `json:"tx_type" db:"tx_type"`
	ChainId              string `json:"chain_id" db:"chain_id"`
	Gas                  string `json:"gas" db:"gas"`
	GasPrice             string `json:"gas_price" db:"gas_price"`
	MaxFeePerGas         string `json:"max_fee_per_gas" db:"max_fee_per_gas"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas" db:"max_priority_fee_per_gas"`
	InputData            string `json:"input_data" db:"input_data"`
	Nonce                string `json:"nonce" db:"nonce"`
	AccessList           string `json:"access_list" db:"access_list"`
	V                    string `json:"v" db:"v"`
	R                    string `json:"r" db:"r"`
	S                    string `json:"s" db:"s"`
	YPairity             string `json:"y_parity" db:"y_parity"`
}

// NewTransaction creates a new Transaction instance
func NewTransaction(
	blockNumber, blockHash, from, to, transactionHash,
	transactionIndex, value, txType, chainId,
	gas, gasPrice, maxFeePerGas, maxPriorityFeePerGas,
	inputData, nonce, accessList,
	v, r, s, yParity string,
) Transaction {
	return Transaction{
		BlockNumber:          blockNumber,
		BlockHash:            blockHash,
		From:                 from,
		To:                   to,
		TransactionHash:      transactionHash,
		TransactionIndex:     transactionIndex,
		Value:                value,
		Type:                 txType,
		ChainId:              chainId,
		Gas:                  gas,
		GasPrice:             gasPrice,
		MaxFeePerGas:         maxFeePerGas,
		MaxPriorityFeePerGas: maxPriorityFeePerGas,
		InputData:            inputData,
		Nonce:                nonce,
		AccessList:           accessList,
		V:                    v,
		R:                    r,
		S:                    s,
		YPairity:             yParity,
	}
}
