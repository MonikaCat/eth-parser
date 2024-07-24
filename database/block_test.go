package database_test

import (
	"testing"

	"github.com/MonikaCat/eth-parser/database"
	"github.com/MonikaCat/eth-parser/types"
)

func TestSaveBlock(t *testing.T) {
	// Create a new database configuration
	cfg := database.DatabaseConfig{
		DNS:                "test.db",
		MaxOpenConnections: 20,
		MaxIdleConnections: 10,
	}

	db, err := database.ConnectToDatabase(cfg) // Create a new instance of the database
	if err != nil {
		t.Errorf("Failed to connect to db: %v", err)
	}

	// Create a sample block
	block := types.Block{
		BlockNumber:           "0x1349fa2",
		BlockHash:             "0xe9db63650e35813afebe410cdc4f0d4f01af11c5e86ba22d2ce63a9f26175d96",
		ParentHash:            "0x519789f08f6a83ad3265e44af18e9a098b73fc19f3b7c06ce99c07cee4f6ddce",
		Nonce:                 "0x0",
		Miner:                 "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		BaseFeePerGas:         "0x3bdec102f",
		BlobGasUsed:           "0x80000",
		Difficulty:            "0x0",
		ExcessBlobGas:         "0x360000",
		ExtraData:             "0x546974616e2028746974616e6275696c6465722e78797a29",
		GasLimit:              "0x1c9c380",
		GasUsed:               "0x174c26d",
		LogsBloom:             "0xf6231116318b01cc4749a278f913f9a07da82d4d695c6a865b61a0604c1ee563a0b1a5c8a029126ce3966f758be709958343941cbb077a229791e56a713eec84fd9a560959da458d789cc24a8adcf02a170098cda7c409bc6bcbee14983cd9205e7f6646a7e4a0af241cd7b44c390f17a21f1f314019aeb087200a1f78a9a65f63aaa7e9191a42c8ee48b1042af0c5aac1b63465553900a921ace946747bf2b182be15abb2bf745036afbadf89d71b88044fa318521130871cb489f7925e726db1a00673185145d2cacf8314417bd0b11a21665ed6e32957340a7b16884de70514536ee82705e58d00af1de1a5797d5e0f3cb1129b73aec0b108fd1ef415d58c",
		MixHash:               "0xcef5592ad3e55dd5db89c310840c445a709278cccfaf55488b6606c241d10d24",
		ParentBeaconBlockRoot: "0xe2852a0ec439273a5526f1997f16329b517b170c9c3f0772c59a3dd9519e47cc",
		ReceiptsRoot:          "0x4158c96fd59a0c5d96dfe31414835d6a91afc55288f49697f673fc3f57c524cc",
		Sha3Uncles:            "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		BlockSize:             "0x1a636",
		StateRoot:             "0x140f873db76ca158deb8f2e76a12b81840179b747f11398fe481a3b6aef6906c",
		Timestamp:             "0x668542a3",
		TotalDifficulty:       "0xc70d815d562d3cfa955",
	}

	// Save the block to the database
	err = db.SaveBlock(block)
	if err != nil {
		t.Errorf("Failed to save block: %v", err)
	}

	// Retrieve the saved block from the database
	savedBlock, err := db.GetBlock(block.BlockNumber)
	if err != nil {
		t.Errorf("Failed to retrieve block: %v", err)
	}

	// Compare the retrieved block with the original block
	if savedBlock.BlockNumber != block.BlockNumber ||
		savedBlock.BlockHash != block.BlockHash ||
		savedBlock.ParentHash != block.ParentHash ||
		savedBlock.Nonce != block.Nonce ||
		savedBlock.Miner != block.Miner ||
		savedBlock.BaseFeePerGas != block.BaseFeePerGas ||
		savedBlock.BlobGasUsed != block.BlobGasUsed ||
		savedBlock.Difficulty != block.Difficulty ||
		savedBlock.ExcessBlobGas != block.ExcessBlobGas ||
		savedBlock.ExtraData != block.ExtraData ||
		savedBlock.GasLimit != block.GasLimit ||
		savedBlock.GasUsed != block.GasUsed ||
		savedBlock.MixHash != block.MixHash ||
		savedBlock.ParentBeaconBlockRoot != block.ParentBeaconBlockRoot ||
		savedBlock.ReceiptsRoot != block.ReceiptsRoot ||
		savedBlock.Sha3Uncles != block.Sha3Uncles ||
		savedBlock.BlockSize != block.BlockSize ||
		savedBlock.StateRoot != block.StateRoot ||
		savedBlock.Timestamp != block.Timestamp ||
		savedBlock.TotalDifficulty != block.TotalDifficulty {
		t.Errorf("Retrieved block does not match the original block")
	}

}
