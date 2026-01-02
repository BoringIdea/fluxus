-- Chat system database tables

-- Users table for chat authentication
CREATE TABLE IF NOT EXISTS fluxus_users (
    id SERIAL PRIMARY KEY,
    wallet_address TEXT NOT NULL UNIQUE,
    nonce TEXT, -- For signature verification
    last_active_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_fluxus_users_wallet_address ON fluxus_users (wallet_address);

-- Chat rooms table
CREATE TABLE IF NOT EXISTS chat_rooms (
    id SERIAL PRIMARY KEY,
    collection_address TEXT NOT NULL UNIQUE,
    room_name TEXT NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    max_members INTEGER DEFAULT 1000,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_chat_rooms_collection_address ON chat_rooms (collection_address);

-- Chat room members table
CREATE TABLE IF NOT EXISTS chat_room_members (
    id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES chat_rooms(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES fluxus_users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_read_at TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(room_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_chat_room_members_room_id ON chat_room_members (room_id);
CREATE INDEX IF NOT EXISTS idx_chat_room_members_user_id ON chat_room_members (user_id);

-- Chat messages table
CREATE TABLE IF NOT EXISTS chat_messages (
    id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES chat_rooms(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES fluxus_users(id) ON DELETE CASCADE,
    message_type TEXT DEFAULT 'text', -- Only support text type
    content TEXT NOT NULL,
    reply_to_message_id INTEGER REFERENCES chat_messages(id),
    is_edited BOOLEAN DEFAULT FALSE,
    edited_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_chat_messages_room_id_created_at ON chat_messages (room_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_chat_messages_user_id ON chat_messages (user_id);

-- User tokens table for JWT authentication
CREATE TABLE IF NOT EXISTS user_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES fluxus_users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_tokens_user_id ON user_tokens (user_id);
CREATE INDEX IF NOT EXISTS idx_user_tokens_expires_at ON user_tokens (expires_at);

-- Online users table for tracking active connections
CREATE TABLE IF NOT EXISTS online_users (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES fluxus_users(id) ON DELETE CASCADE,
    room_id INTEGER REFERENCES chat_rooms(id) ON DELETE CASCADE,
    socket_id TEXT NOT NULL,
    connected_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_ping_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_online_users_room_id ON online_users (room_id);
CREATE INDEX IF NOT EXISTS idx_online_users_user_id ON online_users (user_id);

-- Drop all tables (in reverse order of dependencies)
DROP TABLE IF EXISTS online_users CASCADE;
DROP TABLE IF EXISTS user_tokens CASCADE;
DROP TABLE IF EXISTS chat_messages CASCADE;
DROP TABLE IF EXISTS chat_room_members CASCADE;
DROP TABLE IF EXISTS chat_rooms CASCADE;
DROP TABLE IF EXISTS fluxus_users CASCADE;
