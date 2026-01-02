import { db } from '../src/db/db';
import { readFileSync } from 'fs';
import { join } from 'path';

async function setupChatTables() {
  try {
    console.log('Setting up chat tables...');

    // Read SQL file
    const sqlPath = join(__dirname, '../src/db/chat-tables.sql');
    const sqlContent = readFileSync(sqlPath, 'utf8');

    // Split by semicolon and execute each statement
    const statements = sqlContent
      .split(';')
      .map((stmt) => stmt.trim())
      .filter((stmt) => stmt.length > 0 && !stmt.startsWith('--'));

    for (const statement of statements) {
      if (statement.trim()) {
        console.log(`Executing: ${statement.substring(0, 50)}...`);
        await db.execute(statement);
      }
    }

    console.log('Chat tables setup completed successfully!');
  } catch (error) {
    console.error('Error setting up chat tables:', error);
    process.exit(1);
  }
}

setupChatTables();
