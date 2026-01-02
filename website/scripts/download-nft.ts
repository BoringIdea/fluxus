import axios from 'axios';
import fs from 'fs';
import path from 'path';
import { promisify } from 'util';
import { fileURLToPath } from 'url';

const writeFileAsync = promisify(fs.writeFile);
const mkdirAsync = promisify(fs.mkdir);

interface NFTMetadata {
    name: string;
    image: string;
    attributes: Array<{
        trait_type: string;
        value: string;
    }>;
}

const BASE_URI = 'ipfs://bafybeibc5sgo2plmjkq2tzmhrn54bk3crhnc23zd2msg4ea7a4pxrkgfna';
const START_ID = 1;
const END_ID = 8888;
const CONCURRENT_DOWNLOADS = 20; // Concurrent download number
const IPFS_GATEWAYS = [
    'https://ipfs.io/ipfs/',
    'https://cloudflare-ipfs.com/ipfs/',
    'https://gateway.pinata.cloud/ipfs/',
    'https://dweb.link/ipfs/'
];

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const OUTPUT_DIR = path.join(__dirname, 'nfts');
const RETRY_ATTEMPTS = 3;
const RETRY_DELAY = 1000; // 1 second

async function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function ipfsToHttp(ipfsUrl: string, gatewayIndex = 0): string {
    if (!ipfsUrl.startsWith('ipfs://')) return ipfsUrl;
    const cidPath = ipfsUrl.replace('ipfs://', '');
    return IPFS_GATEWAYS[gatewayIndex] + cidPath;
}

async function downloadWithGateways<T>(urls: string[], outputPath: string, isJson = false): Promise<T | null> {
    for (let i = 0; i < urls.length; i++) {
        try {
            const response = await axios.get(urls[i], {
                responseType: isJson ? 'json' : 'arraybuffer',
                timeout: 10000,
                headers: {
                    'User-Agent': 'Mozilla/5.0 (compatible; NFT-Downloader/1.0)'
                }
            });
            if (!isJson) {
                await writeFileAsync(outputPath, response.data);
            }
            return response.data as T;
        } catch (error: any) {
            if (i === urls.length - 1) {
                return null;
            }
            await sleep(RETRY_DELAY);
        }
    }
    return null;
}

async function downloadNFT(id: number): Promise<{ success: boolean; id: number }> {
    const metaIpfsUrl = `${BASE_URI}/${id}`;
    const metaUrls = IPFS_GATEWAYS.map(gw => gw + metaIpfsUrl.replace('ipfs://', ''));
    
    try {
        const metaJson = await downloadWithGateways<NFTMetadata>(metaUrls, '', true);
        if (!metaJson || !metaJson.image) {
            throw new Error('Metadata or image field missing');
        }

        const imageIpfsUrl = metaJson.image;
        const imageUrls = IPFS_GATEWAYS.map(gw => gw + imageIpfsUrl.replace('ipfs://', ''));
        const outputPath = path.join(OUTPUT_DIR, `${id}${path.extname(imageIpfsUrl) || '.png'}`);
        const imageBuffer = await downloadWithGateways<Buffer>(imageUrls, outputPath, false);
        
        return { success: !!imageBuffer, id };
    } catch (error: any) {
        console.error(`Download failed ${id}: ${error?.message || 'Unknown error'}`);
        return { success: false, id };
    }
}

async function processBatch(startId: number, endId: number): Promise<{ success: number; fail: number }> {
    const tasks = [];
    for (let i = startId; i <= endId; i++) {
        tasks.push(downloadNFT(i));
    }
    
    const results = await Promise.all(tasks);
    return {
        success: results.filter(r => r.success).length,
        fail: results.filter(r => !r.success).length
    };
}

async function main() {
    if (!fs.existsSync(OUTPUT_DIR)) {
        await mkdirAsync(OUTPUT_DIR, { recursive: true });
    }

    console.log('Starting to download NFT images...');
    console.log(`Concurrent number: ${CONCURRENT_DOWNLOADS}`);
    
    let totalSuccess = 0;
    let totalFail = 0;
    
    // Process in batches
    for (let i = START_ID; i <= END_ID; i += CONCURRENT_DOWNLOADS) {
        const endId = Math.min(i + CONCURRENT_DOWNLOADS - 1, END_ID);
        console.log(`Processing batch ${i}-${endId}...`);
        
        const { success, fail } = await processBatch(i, endId);
        totalSuccess += success;
        totalFail += fail;
        
        console.log(`Current progress: ${endId}/${END_ID} (Success: ${totalSuccess}, Fail: ${totalFail})`);
        
        // Add small delay to avoid request too fast
        await sleep(100);
    }

    console.log('\nDownload completed!');
    console.log(`Total success: ${totalSuccess}`);
    console.log(`Total fail: ${totalFail}`);
}

main().catch(console.error); 