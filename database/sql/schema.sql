PRAGMA foreign_keys = ON;

----------------------------------------------
--                BLOCK                     --
----------------------------------------------
CREATE TABLE IF NOT EXISTS block(
    block_number                TEXT PRIMARY KEY UNIQUE,
    block_hash                  TEXT NOT NULL UNIQUE,
    parent_hash                 TEXT NOT NULL,
    nonce                       TEXT NOT NULL,
    miner                       TEXT NOT NULL,
    base_fee_per_gas            TEXT NOT NULL,
    blob_gas_used               TEXT NOT NULL,
    difficulty                  TEXT NOT NULL,
    excess_blob_gas             TEXT NOT NULL,
    extra_data                  TEXT NOT NULL, 
    gas_limit                   TEXT NOT NULL,
    gas_used                    TEXT NOT NULL,
    logs_bloom                  TEXT NOT NULL,
    mix_hash                    TEXT NOT NULL,
    parent_beacon_block_root    TEXT NOT NULL,
    receipts_root               TEXT NOT NULL,
    sha3_uncles                 TEXT NOT NULL,
    block_size                  TEXT NOT NULL,
    state_root                  TEXT NOT NULL,
    timestamp                   TEXT NOT NULL,
    total_difficulty            TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS block_number_index ON block(block_number);
CREATE INDEX IF NOT EXISTS block_miner_index ON block(miner);

----------------------------------------------
--             TRANSACTIONS                 --
----------------------------------------------
CREATE TABLE IF NOT EXISTS transactions(
    block_number                TEXT NOT NULL,
    block_hash                  TEXT NOT NULL,
    tx_from                     TEXT NOT NULL,
    tx_to                       TEXT NOT NULL,
    transaction_hash            TEXT NOT NULL UNIQUE,
    transaction_index           TEXT NOT NULL,
    tx_value                    TEXT NOT NULL,
    tx_type                     TEXT NOT NULL,
    chain_id                    TEXT NOT NULL,
    gas                         TEXT NOT NULL,
    gas_price                   TEXT NOT NULL,
    max_fee_per_gas             TEXT NOT NULL,
    max_priority_fee_per_gas    TEXT NOT NULL,
    input_data                  TEXT NOT NULL,
    nonce                       TEXT NOT NULL,
    access_list                 TEXT NOT NULL,
    v                           TEXT NOT NULL,
    r                           TEXT NOT NULL,
    s                           TEXT NOT NULL,
    y_parity                    TEXT NOT NULL,

    CONSTRAINT unique_transaction UNIQUE (transaction_hash, transaction_index),
    FOREIGN KEY (block_number) REFERENCES block(block_number),
    FOREIGN KEY (block_hash) REFERENCES block(block_hash)
);
CREATE INDEX IF NOT EXISTS transaction_hash_index ON transactions(transaction_hash);
CREATE INDEX IF NOT EXISTS transaction_block_number_index ON transactions(block_number);
CREATE INDEX IF NOT EXISTS transaction_from_index ON transactions(tx_from);
CREATE INDEX IF NOT EXISTS transaction_to_index ON transactions(tx_to);