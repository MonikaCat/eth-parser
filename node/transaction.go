package node

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/MonikaCat/eth-parser/types"
	"github.com/MonikaCat/eth-parser/utils"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

var USDCAddress = common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")

// GetTransaction queries transaction by hash and parse it's details
func (n *Node) GetTransaction(blockNumber int64, tx *ethtypes.Transaction) (types.Transaction, error) {

	// get the transaction details by tx hash
	transaction, isPending, err := n.client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		return types.Transaction{}, fmt.Errorf("error while getting transaction %s: error: %v ", tx.Hash().String(), err)
	}

	// handle if the transaction is not pending
	if !isPending {
		usdcTransferTx, err := n.ParseTransactionDetails(blockNumber, transaction)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("error while parsing transaction details: %v", err)
		}

		return usdcTransferTx, nil
	}

	return types.Transaction{}, nil
}

// ParseTransactionDetails parses transaction details
func (n *Node) ParseTransactionDetails(blockNumber int64, transaction *ethtypes.Transaction) (types.Transaction, error) {
	// log the processing of the transaction
	n.logger.WithFields(logrus.Fields{
		"block":   blockNumber,
		"tx hash": transaction.Hash().String(),
	}).Debug(utils.ProcessingTx)

	// check if the transaction receipient is a USDC address
	if transaction.To() == nil || transaction.To().Hex() != USDCAddress.Hex() {
		return types.Transaction{}, nil
	}

	// get the transaction receipt
	txReceipt, err := n.client.TransactionReceipt(n.ctx, transaction.Hash())
	if err != nil {
		return types.Transaction{}, fmt.Errorf("error while getting transaction receipt: %v", err)
	}
	// check if the transaction is a USDC transfer
	if len(transaction.Data()) >= 10 && strings.HasPrefix(common.Bytes2Hex(transaction.Data()[:4]), "a9059cbb") {
		var transferTo common.Address

		// parse the transferTo and value
		data := transaction.Data()[10:]
		transferTo.SetBytes(data[:32])
		value := new(big.Int).SetBytes(data[32:])

		// parse the sender
		txFrom, err := n.parseSender(transaction)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("error while parsing sender: %v", err)
		}

		// marshal the access list
		accessListJSON, err := json.Marshal(transaction.AccessList())
		if err != nil {
			return types.Transaction{}, fmt.Errorf("error marshalling access list: %v", err)
		}

		// get the raw signature values
		v, r, s := transaction.RawSignatureValues()
		var yParity int
		if v.BitLen() > 0 {
			yParity = int(v.Uint64() % 2)
		}

		// create a new transaction details
		txDetails := types.NewTransaction(
			BigIntToHex(big.NewInt(blockNumber)),
			txReceipt.BlockHash.String(),
			txFrom.String(),
			transferTo.String(),
			transaction.Hash().String(),
			Uint64ToHex(uint64(txReceipt.TransactionIndex)),
			BigIntToHex(value),
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
