INSERT INTO block (
    block_number, block_hash, miner, 
    base_fee_per_gas, blob_gas_used, 
    difficulty, excess_blob_gas, extra_data,
    gas_limit, gas_used, logs_bloom, mix_hash,
    nonce, parent_beacon_block_root, parent_hash,
    receipts_root, sha3_uncles, block_size,
    state_root, timestamp, total_difficulty
) VALUES (
    :block_number, :block_hash, :miner, 
    :base_fee_per_gas, :blob_gas_used, 
    :difficulty, :excess_blob_gas, :extra_data, 
    :gas_limit, :gas_used, :logs_bloom, :mix_hash, 
    :nonce, :parent_beacon_block_root, :parent_hash, 
    :receipts_root, :sha3_uncles, :block_size, 
    :state_root, :timestamp, :total_difficulty
) ON CONFLICT DO NOTHING;