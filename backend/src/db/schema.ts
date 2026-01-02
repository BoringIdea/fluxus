import {
  pgTable,
  text,
  integer,
  numeric,
  timestamp,
  serial,
  boolean,
  jsonb,
} from 'drizzle-orm/pg-core';
import { sql } from 'drizzle-orm';

export const collections = pgTable('collections', {
  id: serial('id').primaryKey(),
  chain: integer('chain').notNull(), // 20143: Monad Devnet
  address: text('address').notNull().unique(), // Collection address
  price_contract: text('price_contract').notNull().unique(), // Price contract address
  creator: text('creator').notNull(), // Creator address
  name: text('name').notNull(), // Collection name
  symbol: text('symbol').notNull(), // Collection symbol
  max_supply: integer('max_supply').default(0), // Max supply of the collection
  initial_price: integer('initial_price').default(0), // Initial price of the collection
  creator_fee: integer('creator_fee').default(0), // Creator fee of the collection
  base_uri: text('base_uri').default(''), // Base URI of the collection
  owners: integer('owners').default(0), // Number of owners of the collection
  current_supply: integer('current_supply').default(0), // Current supply of the collection
  total_supply: integer('total_supply').default(0), // Total supply of the collection
  total_volume: integer('total_volume').default(0), // Total volume of the collection
  is_registered: boolean('is_registered').default(true), // Collection is registered
  meta_data: jsonb('meta_data'), // Collection meta data
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
});

export const txs = pgTable('txs', {
  tx_hash: text('tx_hash').primaryKey(),
  chain: integer('chain').notNull(), // 20143: Monad Devnet
  collection_id: integer('collection_id').references(() => collections.id),
  tx_type: integer('tx_type').notNull(), // 1: mint, 2: buy, 3: sell, 4: blukBuy, 5: bulkSell, 6: bulkMint
  sender: text('sender').notNull(),
  price: numeric('price').notNull(),
  token_id: text('token_id').notNull(),
  token_ids: text('token_ids').array(),
  status: integer('status').notNull(), // 0: pending, 1: processing, 2: processed, 3: failed
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
});

export const nftOwnership = pgTable('nft_ownership', {
  id: serial('id').primaryKey(),
  collection_id: integer('collection_id').references(() => collections.id),
  token_id: text('token_id').notNull(),
  owner: text('owner').notNull(),
  acquired_at: timestamp('acquired_at').default(sql`CURRENT_TIMESTAMP`),
});

export const blockNumber = pgTable('block_number', {
  chain: integer('chain').primaryKey(), // 20143: Monad Devnet
  block_number: integer('block_number').notNull(),
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
  updated_at: timestamp('updated_at').default(sql`CURRENT_TIMESTAMP`),
});

export const userRewards = pgTable('user_rewards', {
  user_address: text('user_address').primaryKey(),
  reward_amount: numeric('reward_amount'),
});

export const chainLogs = pgTable('chain_logs', {
  id: serial('id').primaryKey(),
  chain: integer('chain').notNull(), // 20143: Monad Devnet
  block_number: integer('block_number').notNull(), // Block number
  tx_hash: text('tx_hash').notNull(), // Transaction hash
  log_index: integer('log_index').notNull(), // Log index
  first_topic: text('first_topic').notNull(), // Topic
  log: jsonb('log').notNull(), // Log data
  status: integer('status').notNull(), // 0: pending, 1: processing, 2: processed, 3: failed
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
  updated_at: timestamp('updated_at').default(sql`CURRENT_TIMESTAMP`),
});

export const nodeLocks = pgTable('node_locks', {
  id: serial('id').primaryKey(),
  chain: integer('chain').notNull(),
  node_id: text('node_id').notNull(),
  locked_at: timestamp('locked_at').notNull(),
  locked_until: timestamp('locked_until').notNull(),
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
});

// Chat system tables
export const users = pgTable('fluxus_users', {
  id: serial('id').primaryKey(),
  wallet_address: text('wallet_address').notNull().unique(),
  nonce: text('nonce'),
  last_active_at: timestamp('last_active_at'),
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
  updated_at: timestamp('updated_at').default(sql`CURRENT_TIMESTAMP`),
});

export const chatRooms = pgTable('chat_rooms', {
  id: serial('id').primaryKey(),
  collection_address: text('collection_address').notNull().unique(),
  room_name: text('room_name').notNull(),
  description: text('description'),
  is_active: boolean('is_active').default(true),
  max_members: integer('max_members').default(1000),
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
  updated_at: timestamp('updated_at').default(sql`CURRENT_TIMESTAMP`),
});

export const chatRoomMembers = pgTable('chat_room_members', {
  id: serial('id').primaryKey(),
  room_id: integer('room_id').references(() => chatRooms.id, {
    onDelete: 'cascade',
  }),
  user_id: integer('user_id').references(() => users.id, {
    onDelete: 'cascade',
  }),
  joined_at: timestamp('joined_at').default(sql`CURRENT_TIMESTAMP`),
  last_read_at: timestamp('last_read_at'),
  is_active: boolean('is_active').default(true),
});

export const chatMessages = pgTable('chat_messages', {
  id: serial('id').primaryKey(),
  room_id: integer('room_id').references(() => chatRooms.id, {
    onDelete: 'cascade',
  }),
  user_id: integer('user_id').references(() => users.id, {
    onDelete: 'cascade',
  }),
  message_type: text('message_type').default('text'),
  content: text('content').notNull(),
  reply_to_message_id: integer('reply_to_message_id').references(
    () => chatMessages.id,
  ),
  is_edited: boolean('is_edited').default(false),
  edited_at: timestamp('edited_at'),
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
});

export const userTokens = pgTable('user_tokens', {
  id: serial('id').primaryKey(),
  user_id: integer('user_id').references(() => users.id, {
    onDelete: 'cascade',
  }),
  token_hash: text('token_hash').notNull(),
  expires_at: timestamp('expires_at').notNull(),
  is_active: boolean('is_active').default(true),
  created_at: timestamp('created_at').default(sql`CURRENT_TIMESTAMP`),
});

export const onlineUsers = pgTable('online_users', {
  id: serial('id').primaryKey(),
  user_id: integer('user_id').references(() => users.id, {
    onDelete: 'cascade',
  }),
  room_id: integer('room_id').references(() => chatRooms.id, {
    onDelete: 'cascade',
  }),
  socket_id: text('socket_id').notNull(),
  connected_at: timestamp('connected_at').default(sql`CURRENT_TIMESTAMP`),
  last_ping_at: timestamp('last_ping_at').default(sql`CURRENT_TIMESTAMP`),
});
