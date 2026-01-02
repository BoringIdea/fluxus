import { ethers } from 'ethers';
import * as dotenv from 'dotenv';
import * as fs from 'fs';
import * as path from 'path';

dotenv.config();

interface WalletInfo {
    privateKey: string;
    address: string;
}

class BatchTransfer {
    private provider: ethers.JsonRpcProvider;
    private fundingWallet: ethers.Wallet;
    private wallets: WalletInfo[];
    private transferAmount = ethers.parseEther("0.05");  // 0.05 ETH
    
    constructor(rpcUrl: string, fundingPrivateKey: string) {
        this.provider = new ethers.JsonRpcProvider(rpcUrl);
        this.fundingWallet = new ethers.Wallet(fundingPrivateKey, this.provider);
        
        const keysPath = path.join(__dirname, './keys.json');
        this.wallets = JSON.parse(fs.readFileSync(keysPath, 'utf8'));
    }

    async checkFundingBalance() {
        const balance = await this.provider.getBalance(this.fundingWallet.address);
        const requiredAmount = this.transferAmount * BigInt(this.wallets.length);
        
        console.log(`Funding wallet address: ${this.fundingWallet.address}`);
        console.log(`Current balance: ${ethers.formatEther(balance)} ETH`);
        console.log(`Required amount: ${ethers.formatEther(requiredAmount)} ETH`);
        
        if (balance < requiredAmount) {
            throw new Error(`Insufficient balance. Need ${ethers.formatEther(requiredAmount)} ETH but only have ${ethers.formatEther(balance)} ETH`);
        }
    }

    async transferToAll() {
        await this.checkFundingBalance();
        
        console.log(`Starting transfers to ${this.wallets.length} addresses...`);
        
        const logPath = path.join(__dirname, './transfer_logs.json');
        const transferLogs: any[] = [];

        for (let i = 0; i < this.wallets.length; i++) {
            const wallet = this.wallets[i];
            try {
                const currentBalance = await this.provider.getBalance(wallet.address);
                
                if (currentBalance >= this.transferAmount) {
                    console.log(`Address ${wallet.address} already has sufficient balance, skipping...`);
                    continue;
                }

                const tx = await this.fundingWallet.sendTransaction({
                    to: wallet.address,
                    value: this.transferAmount,
                    gasLimit: 21000 
                });

                const receipt = await tx.wait();
                
                console.log(`Transfer successful to ${wallet.address}`);
                console.log(`Transaction hash: ${receipt?.hash}`);
                
                transferLogs.push({
                    to: wallet.address,
                    amount: ethers.formatEther(this.transferAmount),
                    txHash: receipt?.hash,
                    status: 'success',
                    timestamp: new Date().toISOString()
                });
            } catch (error: any) {
                console.error(`Failed to transfer to ${wallet.address}: ${error.message}`);
                
                transferLogs.push({
                    to: wallet.address,
                    amount: ethers.formatEther(this.transferAmount),
                    error: error.message,
                    status: 'failed',
                    timestamp: new Date().toISOString()
                });
            }

            // write log
            fs.writeFileSync(logPath, JSON.stringify(transferLogs, null, 2));

            // delay 1000ms
            await new Promise(resolve => setTimeout(resolve, 1000));
        }

        console.log('All transfers completed. Check transfer_logs.json for details.');
    }
}

async function main() {
    const RPC_URL = process.env.RPC_URL!;
    const FUNDING_PRIVATE_KEY = process.env.FUNDING_PRIVATE_KEY!;
    console.log(RPC_URL, FUNDING_PRIVATE_KEY);

    if (!RPC_URL || !FUNDING_PRIVATE_KEY) {
        throw new Error('Please set RPC_URL and FUNDING_PRIVATE_KEY in .env file');
    }

    const batchTransfer = new BatchTransfer(RPC_URL, FUNDING_PRIVATE_KEY);
    await batchTransfer.transferToAll();
}

main().catch((error) => {
    console.error('Fatal error:', error);
    process.exit(1);
});