-- collection table
CREATE TABLE collections (
    id SERIAL PRIMARY KEY,
    chain INTEGER, -- 20143: Monad Devnet
    address TEXT NOT NULL UNIQUE, -- Collection address
    price_contract TEXT NOT NULL, -- Price contract address
    creator TEXT NOT NULL, -- Creator address
    name TEXT, -- Collection name
    symbol TEXT, -- Collection symbol
    max_supply INTEGER, -- Max supply of the collection
    initial_price NUMERIC(78, 0), -- Initial price of the collection
    creator_fee NUMERIC(78, 0), -- Creator fee of the collection
    base_uri TEXT, -- Base URI of the collection
    owners INTEGER, -- Number of owners of the collection
    current_supply INTEGER, -- Current supply of the collection
    total_supply INTEGER, -- Total supply of the collection
    total_volume NUMERIC(78, 0), -- Total volume of the collection
    is_registered BOOLEAN DEFAULT FALSE, -- Collection is registered
    meta_data JSONB, -- Collection meta data
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create index on address
CREATE INDEX idx_collections_address ON collections (address);

-- txs table
CREATE TABLE txs (
    tx_hash TEXT PRIMARY KEY,
    chain INTEGER, -- 20143: Monad Devnet
    collection_id INTEGER REFERENCES collections(id),
    tx_type INTEGER, -- 1: mint, 2: buy, 3: sell, 4: blukBuy, 5: bulkSell, 6: bulkMint
    sender TEXT,
    price NUMERIC(78, 0),
    token_id TEXT,
    token_ids TEXT[],
    status INTEGER, -- 0: pending, 1: processing, 2: processed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- ALTER TABLE txs ADD COLUMN status INTEGER DEFAULT 0 NOT NULL;
-- ALTER TABLE txs ADD COLUMN token_ids INTEGER[] DEFAULT '{}';

-- nft ownership table
CREATE TABLE nft_ownership (
    id SERIAL PRIMARY KEY,
    collection_id INTEGER REFERENCES collections(id),
    token_id TEXT,
    owner TEXT,
    acquired_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create index on token_id and collection_id
CREATE INDEX idx_nft_ownership_token_id_collection_id ON nft_ownership (token_id, collection_id);
ALTER TABLE nft_ownership ADD CONSTRAINT unique_ownership UNIQUE (collection_id, token_id, owner);

-- block number table
CREATE TABLE block_number (
    chain INTEGER PRIMARY KEY, -- 20143: Monad Devnet
    block_number NUMERIC(78, 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- user rewards table
CREATE TABLE user_rewards (
    user_address TEXT PRIMARY KEY,
    reward_amount NUMERIC(78, 0)
);

-- chain logs
CREATE TABLE chain_logs (
    id SERIAL PRIMARY KEY,
    chain INTEGER, -- 20143: Monad Devnet
    block_number NUMERIC(78, 0), -- Block number
    tx_hash TEXT, -- Transaction hash
    log_index INTEGER, -- Log index
    first_topic TEXT, -- Topic
    log JSONB, -- Log data
    status INTEGER, -- 0: pending, 1: processing, 2: processed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE UNIQUE INDEX IF NOT EXISTS idx_chain_logs_chain_block_tx ON chain_logs (chain, block_number, tx_hash);

CREATE TABLE IF NOT EXISTS node_locks (
  id SERIAL PRIMARY KEY,
  chain INTEGER NOT NULL,
  node_id TEXT NOT NULL,
  locked_at TIMESTAMP NOT NULL,
  locked_until TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS "node_locks_chain_idx" ON "node_locks" ("chain");
CREATE INDEX IF NOT EXISTS "node_locks_locked_until_idx" ON "node_locks" ("locked_until");

-- Delete all tables
DROP TABLE IF EXISTS user_rewards;
DROP TABLE IF EXISTS nft_ownership;
DROP TABLE IF EXISTS txs;
DROP TABLE IF EXISTS block_number;
DROP TABLE IF EXISTS collections;
DROP TABLE IF EXISTS chain_logs;
DROP TABLE IF EXISTS node_locks;