import { ethers } from 'ethers';
import * as fs from 'fs';
import * as path from 'path';

async function generateKeys(count: number) {
    const keys = [];
    
    for (let i = 0; i < count; i++) {
        const wallet = ethers.Wallet.createRandom();
        keys.push({
            privateKey: wallet.privateKey,
            address: wallet.address
        });
        
        console.log(`Generated key ${i + 1}/${count}`);
    }
    
    const filePath = path.join(__dirname, './keys.json');
    fs.writeFileSync(filePath, JSON.stringify(keys, null, 2));
    console.log(`Keys saved to ${filePath}`);
}

generateKeys(100).catch(console.error);