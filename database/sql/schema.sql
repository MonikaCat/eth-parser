----------------------------------------------
--                BLOCK                     --
----------------------------------------------
CREATE TABLE block(
    block_number                BIGINT PRIMARY UNIQUE,
    block_hash                  TEXT NOT NULL,
    miner                       TEXT NOT NULL,
    base_fee_per_gas            TEXT NOT NULL, -- TODO update base_fee_per_gas type
    blob_gas_used               TEXT NOT NULL, -- TODO update blob_gas_used type
    difficulty                  TEXT NOT NULL, -- TODO update difficulty type
    excess_blob_gas             TEXT NOT NULL, -- TODO update excess_blob_gas type
    extra_data                  TEXT NOT NULL, -- TODO update extra_data type
    gas_limit                   TEXT NOT NULL, -- TODO update gas_limit type
    gas_used                    TEXT NOT NULL, -- TODO update gas_used type
    logs_bloom                  TEXT NOT NULL, -- TODO update logs_bloom type
    mix_hash                    TEXT NOT NULL, -- TODO update mix_hash type
    nonce                       TEXT NOT NULL, -- TODO update nonce tyep  
    parent_beacon_block_root    TEXT NOT NULL, -- TODO update parent_beacon_block_root type
    parent_hash                 TEXT NOT NULL,
    receipts_root               TEXT NOT NULL, -- TODO update receipts_root type
    sha3_uncles                 TEXT NOT NULL, -- TODO update sha3_uncles type
    block_size                  TEXT NOT NULL, -- TODO update block_size type
    state_root                  TEXT NOT NULL, -- TODO update state_root type
    timestamp                   TEXT NOT NULL, -- TODO update timestamp type
    total_difficulty            TEXT NOT NULL, -- TODO update total_difficulty type
)
CREATE INDEX block_number_index ON block(block_number);
CREATE INDEX block_miner_index ON block(miner);

----------------------------------------------
--             TRANSACTION                  --
----------------------------------------------
CREATE TABLE transaction(
    block_number                BIGINT references block(block_number) ON DELETE CASCADE,
    block_hash                  TEXT NOT NULL references block(block_hash) ON DELETE CASCADE,
    from                        TEXT NOT NULL,
    to                          TEXT NOT NULL,
    transaction_hash            TEXT NOT NULL,
    transaction_index           TEXT NOT NULL, -- TODO update transaction_index type
    value                       TEXT NOT NULL, -- TODO update value type
    type                        TEXT NOT NULL, -- TODO update type type
    chain_id                    TEXT NOT NULL,
    gas                         BIGINT DEFAULT 0,
    gas_price                   BIGINT DEFAULT 0,
    max_fee_per_gas             BIGINT DEFAULT 0,
    max_priority_fee_per_gas    BIGINT DEFAULT 0,
    input_data                  TEXT NOT NULL, -- TODO update input type
    nonce                       TEXT NOT NULL, -- TODO update nonce type
    access_list                 []TEXT NOT NULL,
    v                           TEXT NOT NULL, -- TODO update v type
    r                           TEXT NOT NULL, -- TODO update r type
    s                           TEXT NOT NULL, -- TODO update s type
    y_parity                    TEXT NOT NULL, -- TODO update y_parity type
)
CREATE INDEX transaction_hash_index ON transaction(transaction_hash);
CREATE INDEX transaction_block_number_index ON transaction(block_number);
CREATE INDEX transaction_from_index ON transaction(from);
CREATE INDEX transaction_to_index ON transaction(to);