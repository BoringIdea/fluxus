import 'dotenv/config';

import { drizzle } from 'drizzle-orm/postgres-js';
import * as postgres from 'postgres';
import * as schema from './schema';

const connectionString = process.env.DATABASE_URL;
if (!connectionString) {
  throw new Error('DATABASE_URL is not set');
}

// Disable prefetch as it is not supported for "Transaction" pool mode
export const queryClient = postgres(connectionString, {
  prepare: false,
  max: 20, // Maximum number of connections
  idle_timeout: 20, // Close idle connections after 20 seconds
  max_lifetime: 60 * 30, // Close connections after 30 minutes
});
export const db = drizzle(queryClient, { schema });
