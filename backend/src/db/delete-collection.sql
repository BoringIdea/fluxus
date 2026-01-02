-- Delete all records related to a specific collection
-- Parameters: :collection_id - The ID of the collection to delete

-- 1. Delete related records from txs table
DELETE FROM txs 
WHERE collection_id = :collection_id;

-- 2. Delete related records from nft_ownership table
DELETE FROM nft_ownership 
WHERE collection_id = :collection_id;

-- 3. Delete related records from chain_logs table
-- Note: Here we assume that the tx_hash in chain_logs is related to the tx_hash in txs
DELETE FROM chain_logs 
WHERE tx_hash IN (
    SELECT tx_hash 
    FROM txs 
    WHERE collection_id = :collection_id
);

-- 4. Delete records from collections table
DELETE FROM collections 
WHERE id = :collection_id; 