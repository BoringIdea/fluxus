import { ethers } from 'ethers';
import { FACTORY_CONTRACT_ABI } from '@/src/contract';

// Factory contract addresses for different chains
const FACTORY_ADDRESSES: Record<number, string> = {
  84532: process.env.NEXT_PUBLIC_BASE_FACTORY_CONTRACT_ADDRESS|| "", // Base Sepolia
  97: process.env.NEXT_PUBLIC_BSC_FACTORY_CONTRACT_ADDRESS || "",   // BSC Testnet
};

// RPC URLs for different chains
const RPC_URLS: Record<number, string> = {
  84532: "https://sepolia.base.org", // Base Sepolia
  97: "https://data-seed-prebsc-1-s1.binance.org:8545", // BSC Testnet
};

export interface ContractParams {
  name: string;
  symbol: string;
  initialPrice: bigint; // BigInt like in the script
  maxSupply: bigint;
  maxPrice: bigint; // BigInt like in the script
  creatorFeePercent: bigint; // BigInt like in the script
  imageUrl: string;
  gatewayAddress: string;
  gasLimit: bigint;
  supportMint: boolean;
  creator?: string; // Optional creator address
}

export async function calculateFluxusCrossChainAddress(
  chainId: number,
  params: ContractParams
): Promise<string> {
  console.log('calculateFluxusCrossChainAddress', chainId, params);
  try {
    // Get factory address for the chain
    const factoryAddress = FACTORY_ADDRESSES[chainId];
    if (!factoryAddress) {
      throw new Error(`Factory address not found for chain ID: ${chainId}`);
    }

    // Get RPC URL for the chain
    const rpcUrl = RPC_URLS[chainId];
    if (!rpcUrl) {
      throw new Error(`RPC URL not found for chain ID: ${chainId}`);
    }

    // Create provider
    const provider = new ethers.JsonRpcProvider(rpcUrl);

    // Create contract instance
    const factory = new ethers.Contract(factoryAddress, FACTORY_CONTRACT_ABI, provider);

    // Log detailed parameters for debugging
    console.log(`Calling calculateFluxusCrossChainAddress for chain ${chainId} with parameters:`, {
      factoryAddress,
      rpcUrl,
      creator: params.creator || 'using default msg.sender',
      params: {
        name: params.name,
        symbol: params.symbol,
        initialPrice: params.initialPrice.toString(),
        maxSupply: params.maxSupply.toString(),
        maxPrice: params.maxPrice.toString(),
        creatorFeePercent: params.creatorFeePercent.toString(),
        imageUrl: params.imageUrl,
        gatewayAddress: params.gatewayAddress,
        gasLimit: params.gasLimit,
        supportMint: params.supportMint
      }
    });

    // Call the calculateFluxusCrossChainAddress function with specified sender if provided
    let contractAddress: string;
    
    if (params.creator) {
      // If creator is specified, use staticCall with from address
      console.log(`🎯 Using creator address: ${params.creator}`);
      contractAddress = await factory.calculateFluxusCrossChainAddress.staticCall(
        params.name,
        params.symbol,
        params.initialPrice,
        params.maxSupply,
        params.maxPrice,
        params.creatorFeePercent,
        params.imageUrl,
        params.gatewayAddress,
        params.gasLimit,
        params.supportMint,
        { from: params.creator }
      );
    } else {
      // Default call without specifying sender
      console.log(`🔧 Using default msg.sender`);
      contractAddress = await factory.calculateFluxusCrossChainAddress(
        params.name,
        params.symbol,
        params.initialPrice,
        params.maxSupply,
        params.maxPrice,
        params.creatorFeePercent,
        params.imageUrl,
        params.gatewayAddress,
        params.gasLimit,
        params.supportMint
      );
    }

    console.log(`✅ Calculated contract address for chain ${chainId}:`, contractAddress);
    console.log(`Expected address: 0x1dFBF9189E5526FE30A21ff2e19057E3Bb2b2Eb9`);
    console.log(`Addresses match:`, contractAddress.toLowerCase() === "0x1dFBF9189E5526FE30A21ff2e19057E3Bb2b2Eb9".toLowerCase());
    
    return contractAddress;

  } catch (error) {
    console.error(`Error calculating contract address for chain ${chainId}:`, error);
    throw error;
  }
}

// Default parameters matching the script exactly
export const DEFAULT_CONTRACT_PARAMS: ContractParams = {
  name: "FluxusTest1",
  symbol: "FluxusTEST1",
  initialPrice: ethers.parseEther("0.001"), // BigInt directly
  maxSupply: BigInt(10000),
  maxPrice: ethers.parseEther("1"), // BigInt directly
  creatorFeePercent: ethers.parseEther("0.05"), // BigInt directly
  imageUrl: "https://ipfs.io/ipfs/bafybeiglzppf7i4ki4lwcny62q5uc43qto7elsqakixnwh57r3ipxvie2m",
  gatewayAddress: "0x0c487a766110c85d301d96e33579c5b317fa4995",
  gasLimit: BigInt(12000000),
  supportMint: true,
};
