package types

type Transaction struct {
	BlockNumber          int64  `json:"block_number"`
	BlockHash            string `json:"block_hash"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	TransactionHash      string `json:"transaction_hash"`
	TransactionIndex     uint   `json:"transaction_index"`
	Value                string `json:"value"`
	Type                 uint8  `json:"type"`
	ChainId              string `json:"chain_id"`
	Gas                  uint64 `json:"gas"`
	GasPrice             uint64 `json:"gas_price"`
	MaxFeePerGas         uint64 `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas uint64 `json:"max_priority_fee_per_gas"`
	InputData            string `json:"input_data"`
	Nonce                uint64 `json:"nonce"`
	AccessList           string `json:"access_list"`
	V                    string `json:"v"`
	R                    string `json:"r"`
	S                    string `json:"s"`
	YPairity             string `json:"y_parity"`
}

func NewTransaction(
	blockNumber int64, blockHash, from, to, transactionHash string,
	transactionIndex uint, value string, txType uint8, chainId string,
	gas, gasPrice, maxFeePerGas, maxPriorityFeePerGas uint64,
	inputData string, nonce uint64, accessList string,
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
