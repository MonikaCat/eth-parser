package node

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)


func (n *Node) GetTransaction(tx *ethtypes.Transaction, db database.Database) types.Transaction {

	transaction, isPending, err := n.client.TransactionByHash(context.Background(), tx.Hash())

	if err != nil {
		fmt.Println("error while getting transaction %s: error: %v ", tx.Hash().String(), err)
	}

	if !isPending {
		usdcTransferTx, err := n.ParseTransactionDetails(1, transaction, db)
		if err != nil {
			fmt.Println("error while parsing transaction details: %v", err)
		}

		return usdcTransferTx
	}

	fmt.Printf("\n TRANSACTION DETAILS %v \n", transaction)
	fmt.Printf("\n IS PENDING %v \n", isPending)

}

func (n *Node) ParseTransactionDetails(blockNumber int64, transaction *ethtypes.Transaction, db database.Database) (types.Transaction, error) {

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

		// parse the sender
		txFrom, err := n.parseSender(transaction)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("error while parsing sender: %v", err)
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

		txDetails := types.NewTransaction(
			blockNumber,
			txReceipt.BlockHash.String(),
			txFrom.String(),
			transferTo.String(),
			transaction.Hash().String(),
			txReceipt.TransactionIndex,
			value.String(),
			transaction.Type(),
			transaction.ChainId().String(),
			transaction.Gas(),
			transaction.GasPrice().Uint64(),
			transaction.GasFeeCap().Uint64(),
			transaction.GasTipCap().Uint64(),
			hex.EncodeToString(transaction.Data()),
			transaction.Nonce(),
			string(accessListJson),
			v.String(),
			r.String(),
			s.String(),
			string(yParity),
		)

		err = db.SaveTransaction(txDetails)
		if err != nil {
			fmt.Errorf("error while saving transaction: %v", err)
		}

		return txDetails, nil

	}

	return types.Transaction{}, nil
}

// parseSender parses the sender of a transaction, depending on the tx type
func (n *Node) parseSender(transaction *ethtypes.Transaction) (common.Address, error) {
	chainID, err := n.client.NetworkID(n.ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("error while getting chain id: %v", err)
	}

	var sender common.Address
	switch transaction.Type() {
	case ethtypes.LegacyTxType:
		signer := ethtypes.NewEIP155Signer(chainID)
		sender, err = ethtypes.Sender(signer, transaction)
		if err != nil {
			return common.Address{}, fmt.Errorf("error while getting tx sender: %v", err)
		}
	case ethtypes.AccessListTxType:
		signer := ethtypes.NewEIP2930Signer(chainID)
		sender, err = ethtypes.Sender(signer, transaction)
		if err != nil {
			return common.Address{}, fmt.Errorf("error while getting tx sender: %v", err)
		}
	case ethtypes.DynamicFeeTxType:
		signer := ethtypes.NewLondonSigner(chainID)
		sender, err = ethtypes.Sender(signer, transaction)
		if err != nil {
			return common.Address{}, fmt.Errorf("error while getting tx sender: %v", err)
		}
	case ethtypes.BlobTxType:
		signer := ethtypes.NewCancunSigner(chainID)
		sender, err = ethtypes.Sender(signer, transaction)
		if err != nil {
			return common.Address{}, fmt.Errorf("error while getting tx sender: %v", err)
		}
	default:
		fmt.Printf("\nerror: unsupported transaction type %v", transaction.Type())
	}

	return sender, nil
}
