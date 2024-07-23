package node

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/MonikaCat/eth-parser/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var USDCAddress = common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")

func (n *Node) GetTransaction(tx *ethtypes.Transaction) types.Transaction {

	transaction, isPending, err := n.client.TransactionByHash(context.Background(), tx.Hash())

	if err != nil {
		fmt.Println("error while getting transaction %s: error: %v ", tx.Hash().String(), err)
	}

	if !isPending {
		usdcTransferTx, err := n.ParseTransactionDetails(1, transaction)
		if err != nil {
			fmt.Println("error while parsing transaction details: %v", err)
		}

		return usdcTransferTx
	}

	fmt.Printf("\n TRANSACTION DETAILS %v \n", transaction)
	fmt.Printf("\n IS PENDING %v \n", isPending)

}

func (n *Node) ParseTransactionDetails(blockNumber int64, transaction *ethtypes.Transaction) (types.Transaction, error) {

	if transaction.To() == nil || strings.ToLower(transaction.To().Hex()) != USDCAddress.Hex() {
		return types.Transaction{}, nil
	}

	txReceipt, err := n.client.TransactionReceipt(n.ctx, transaction.Hash())
	if err != nil {
		fmt.Errorf("error while getting transaction receipt: %v", err)
	}

	if len(transaction.Data()) >= 10 && string(transaction.Data()[:10]) == "0xa9059cbb" {
		var transferTo common.Address

		data := transaction.Data()[10:]
		transferTo.SetBytes(data[:32])
		value := new(big.Int).SetBytes(data[32:])

		txFrom, err := ethtypes.Sender(ethtypes.NewEIP155Signer(transaction.ChainId()), transaction)
		if err != nil {
			fmt.Errorf("error while getting transaction sender: %v", err)
		}

		accessListJson, err := json.Marshal(transaction.AccessList())
		if err != nil {
			fmt.Errorf("error marshalling logsBloom: %v", err)
		}

		v, r, s := transaction.RawSignatureValues()
		var yParity int
		if v.BitLen() > 0 {
			yParity = int(v.Uint64() % 2)
		}

		txDetails := types.Transaction{
			BlockNumber:          blockNumber,
			BlockHash:            txReceipt.BlockHash.String(),
			From:                 txFrom.String(),
			To:                   transferTo.String(),
			TransactionHash:      transaction.Hash().String(),
			TransactionIndex:     txReceipt.TransactionIndex,
			Value:                value.String(),
			Type:                 transaction.Type(),
			ChainId:              transaction.ChainId().String(),
			Gas:                  transaction.Gas(),
			GasPrice:             transaction.GasPrice().Uint64(),
			MaxFeePerGas:         transaction.GasFeeCap().Uint64(),
			MaxPriorityFeePerGas: transaction.GasTipCap().Uint64(),
			InputData:            hex.EncodeToString(transaction.Data()),
			Nonce:                transaction.Nonce(),
			AccessList:           string(accessListJson),
			V:                    v.String(),
			R:                    r.String(),
			S:                    s.String(),
			YPairity:             string(yParity),
		}
		return txDetails, nil

	}

	return types.Transaction{}, nil
}
