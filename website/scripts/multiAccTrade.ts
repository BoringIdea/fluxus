import { ethers } from 'ethers';
import * as dotenv from 'dotenv';
import * as fs from 'fs';
import * as path from 'path';

dotenv.config();

const Fluxus_CONTRACT_ADDRESS = process.env.CONTRACT_ADDRESS!;
const RPC_URL = process.env.RPC_URL!;
console.log(Fluxus_CONTRACT_ADDRESS, RPC_URL);

const ABI = [
    "function mint() public payable",
    "function quickBuy() public payable",
    "function sell(uint256 tokenId) public",
    "function getBuyPriceAfterFee() public view returns (uint256)",
    "function getSellPriceAfterFee() public view returns (uint256)",
    "function balanceOf(address owner) public view returns (uint256)",
    "function tokenOfOwnerByIndex(address owner, uint256 index) public view returns (uint256)"
];

interface WalletInfo {
    privateKey: string;
    address: string;
}

class TradeManager {
    private provider: ethers.JsonRpcProvider;
    private wallets: WalletInfo[];
    private currentWalletIndex: number = 0;

    constructor(rpcUrl: string, wallets: WalletInfo[]) {
        this.provider = new ethers.JsonRpcProvider(rpcUrl);
        this.wallets = wallets;
    }

    private getNextWallet(): ethers.Wallet {
        const walletInfo = this.wallets[this.currentWalletIndex];
        this.currentWalletIndex = (this.currentWalletIndex + 1) % this.wallets.length;
        return new ethers.Wallet(walletInfo.privateKey, this.provider);
    }

    async executeTrade() {
        const wallet = this.getNextWallet();
        const contract = new ethers.Contract(Fluxus_CONTRACT_ADDRESS, ABI, wallet);

        try {
            const balance = await this.provider.getBalance(wallet.address);
            console.log(`Current wallet: ${wallet.address}, Balance: ${ethers.formatEther(balance)} ETH`);

            if (balance === BigInt(0)) {
                console.log('Skipping wallet with zero balance');
                return;
            }

            const operation = Math.floor(Math.random() * 3);
            // const operation = 1;
            
            switch (operation) {
                case 0: // Mint
                    const mintPrice = await contract.getBuyPriceAfterFee();
                    console.log(`Mint price: ${ethers.formatEther(mintPrice)} ETH`);
                    if (balance > mintPrice) {
                        await contract.mint({ value: mintPrice });
                        console.log(`Minted new token from wallet ${wallet.address}`);
                    }
                    break;

                case 1: // Quick Buy
                    const buyPrice = await contract.getBuyPriceAfterFee();
                    console.log(`Buy price: ${ethers.formatEther(buyPrice)} ETH`);
                    if (balance > buyPrice) {
                        await contract.quickBuy({ value: buyPrice });
                        console.log(`Bought token via quickBuy from wallet ${wallet.address}`);
                    }
                    break;

                case 2: // Sell
                    const nftBalance = await contract.balanceOf(wallet.address);
                    if (nftBalance > 0) {
                        const tokenId = await contract.tokenOfOwnerByIndex(wallet.address, 0);
                        await contract.sell(tokenId);
                        console.log(`Sold token ${tokenId} from wallet ${wallet.address}`);
                    }
                    break;
            }
        } catch (error) {
            console.error(`Error executing trade for wallet ${wallet.address}:`, error);
        }
    }

    async startTrading(intervalSeconds: number) {
        console.log(`Starting automated trading with ${this.wallets.length} wallets`);
        
        await this.executeTrade();
        
        const scheduleNextTrade = () => {
            const randomDelay = Math.floor(Math.random() * intervalSeconds * 1000);
            setTimeout(() => {
                this.executeTrade().then(() => scheduleNextTrade());
            }, randomDelay);
        };
        
        scheduleNextTrade();
    }
}

async function main() {
    const keysPath = path.join(__dirname, './keys.json');
    const wallets: WalletInfo[] = JSON.parse(fs.readFileSync(keysPath, 'utf8'));

    const tradeManager = new TradeManager(RPC_URL, wallets);
    
    await tradeManager.startTrading(10);
}

main().catch((error) => {
    console.error('Fatal error:', error);
    process.exit(1);
});