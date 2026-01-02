import * as fs from 'fs';
import * as path from 'path';
import { fileURLToPath } from 'url';

// Get the directory path of the current file
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

interface NFTMetadata {
    description: string;
    image: string;
    name: string;
}

interface CollectionInfo {
    description: string;
    name: string;
    baseImgUrl: string;
    symbol: string;
    collectionImage: string;
    bannerImage?: string;
    totalSupply: number;
}

interface CollectionMetadata {
    name: string;
    symbol: string;
    description: string;
    image: string;
    banner_image?: string;
}

function generateMetadata(collectionInfo: CollectionInfo) {
    // Create output directory
    const outputDir = path.join(__dirname, 'metadata');
    if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir, { recursive: true });
    }

    // Generate collection.json
    const collectionMetadata: CollectionMetadata = {
        name: collectionInfo.name,
        symbol: collectionInfo.symbol,
        description: collectionInfo.description,
        image: collectionInfo.collectionImage,
        banner_image: collectionInfo.bannerImage
    };

    // Write to collection.json
    const collectionPath = path.join(outputDir, 'collection.json');
    fs.writeFileSync(collectionPath, JSON.stringify(collectionMetadata, null, 2));
    console.log('Generated collection.json');

    // Generate metadata for each NFT
    for (let i = 0; i < collectionInfo.totalSupply; i++) {
        const metadata: NFTMetadata = {
            description: collectionInfo.description,
            image: `${collectionInfo.baseImgUrl}/${i}.png`,
            name: collectionInfo.name
        };

        // Write to file
        const filePath = path.join(outputDir, `${i}.json`);
        fs.writeFileSync(filePath, JSON.stringify(metadata, null, 2));
        
        // Print progress
        if (i % 1000 === 0) {
            console.log(`Generated ${i} metadata files`);
        }
    }

    console.log(`Completed! Generated ${collectionInfo.totalSupply} metadata files`);
}

// Example usage
const collectionInfo: CollectionInfo = {
    description: "Nakamigos is a collection of 10,000 Nakamigos NFTs.",
    name: "Nakamigos",
    symbol: "Nakamigos",
    baseImgUrl: "https://ipfs.io/ipfs/bafybeiaq27mmrwvs3alh53t3246cubfo3wnvbxvwc5va6uqc6fjm3emzf4",
    collectionImage: "https://ipfs.io/ipfs/bafybeiaq27mmrwvs3alh53t3246cubfo3wnvbxvwc5va6uqc6fjm3emzf4/0.png",
    bannerImage: "https://ipfs.io/ipfs/bafybeiaq27mmrwvs3alh53t3246cubfo3wnvbxvwc5va6uqc6fjm3emzf4/0.png",
    totalSupply: 10000
};

generateMetadata(collectionInfo); 