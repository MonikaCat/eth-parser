package node

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var USDCAddress = common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")

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

	return types.Transaction{}
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
		// value := new(big.Int).SetBytes(data[32:])

		// parse the sender
		txFrom, err := n.parseSender(transaction)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("error while parsing sender: %v", err)
		}

		accessListJSON, err := json.Marshal(transaction.AccessList())
		if err != nil {
			fmt.Errorf("error marshalling logsBloom: %v", err)
		}

		v, r, s := transaction.RawSignatureValues()
		var yParity int
		if v.BitLen() > 0 {
			yParity = int(v.Uint64() % 2)
		}

		txDetails := types.NewTransaction(
			string(blockNumber),
			txReceipt.BlockHash.String(),
			txFrom.String(),
			transferTo.String(),
			transaction.Hash().String(),
			Uint64ToHex(uint64(txReceipt.TransactionIndex)),
			BigIntToHex(transaction.Value()),
			Uint64ToHex(uint64(transaction.Type())),
			BigIntToHex(transaction.ChainId()),
			Uint64ToHex(transaction.Gas()),
			BigIntToHex(transaction.GasPrice()),
			BigIntToHex(transaction.GasFeeCap()),
			BigIntToHex(transaction.GasTipCap()),
			hex.EncodeToString(transaction.Data()),
			Uint64ToHex(transaction.Nonce()),
			string(accessListJSON),
			BigIntToHex(v),
			BigIntToHex(r),
			BigIntToHex(s),
			Uint64ToHex(uint64(rune(yParity))),
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
