INSERT INTO transaction (
    block_number, block_hash,
    tx_from, tx_to, transaction_hash, 
    transaction_index, tx_value, tx_type, 
    chain_id, gas, gas_price, max_fee_per_gas,
    max_priority_fee_per_gas, input_data,
    nonce, access_list, r, s, v, y_parity ) VALUES (
        :block_number, :block_hash,
        :tx_from, :tx_to, :transaction_hash,
        :transaction_index, :tx_value, :tx_type,
        :chain_id, :gas, :gas_price, :max_fee_per_gas,
        :max_priority_fee_per_gas, :input_data,
        :nonce, :access_list, :r, :s, :v, :y_parity
    ) ON CONFLICT DO NOTHING;