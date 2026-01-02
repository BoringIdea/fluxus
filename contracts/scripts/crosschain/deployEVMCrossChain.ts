import { ethers } from "hardhat";
import { Contract } from "ethers";

async function main() {
  console.log("Start deploying EVMFluxusCrossChain implementation contract...\n");
  
  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH\n");

  // Get network information
  const network = await ethers.provider.getNetwork();
  const chainId = Number(network.chainId);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);

  // EVM network configuration
  const evmConfig = getEVMConfig(chainId);
  if (!evmConfig) {
    throw new Error(`Unsupported EVM network, Chain ID: ${chainId}`);
  }

  console.log("EVM Configuration:");
  console.log("  Gateway Address:", evmConfig.gateway);
  console.log("  Default Gas Limit:", evmConfig.gasLimit);
  console.log();

  // 1. Deploy FeeVault contract
  console.log("1. Deploy FeeVault contract...");
  const FeeVault = await ethers.getContractFactory("FeeVault");
  
  let feeVault;
  try {
    feeVault = await FeeVault.deploy();
    console.log("   FeeVault deployment transaction sent, waiting for confirmation...");
    console.log("   FeeVault deployment transaction hash:", feeVault.hash);
    
    try {
      await feeVault.waitForDeployment();
      console.log("   ✅ FeeVault deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", feeVault.hash);
    }
  } catch (error: any) {
    console.log("   ❌ FeeVault deployment failed");
    console.log("   Error:", error?.message || error);
    throw error;
  }
  
  const feeVaultAddress = await feeVault.getAddress();
  console.log("   FeeVault deployed address:", feeVaultAddress);

  // 2. Deploy Price contract
  console.log("\n2. Deploy Price contract...");
  const Price = await ethers.getContractFactory("Price");
  
  let priceContract;
  try {
    priceContract = await Price.deploy();
    console.log("   Price deployment transaction sent, waiting for confirmation...");
    console.log("   Price deployment transaction hash:", priceContract.hash);
    try {
      await priceContract.waitForDeployment();
      console.log("   ✅ Price deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", priceContract.hash);
    }
  } catch (error: any) {
    console.log("   ❌ Price deployment failed");
    console.log("   Error:", error?.message || error);
    throw error;
  }
  
  const priceContractAddress = await priceContract.getAddress();
  console.log("   Price deployed address:", priceContractAddress);

  // 3. Deploy EVMFluxusCrossChain implementation contract
  console.log("\n3. Deploy EVMFluxusCrossChain implementation contract...");
  const EVMFluxusCrossChain = await ethers.getContractFactory("EVMFluxusCrossChain");
  
  let evmFluxusCrossChainImplementation;
  try {
    evmFluxusCrossChainImplementation = await EVMFluxusCrossChain.deploy();
    console.log("   EVMFluxusCrossChain implementation deployment transaction sent, waiting for confirmation...");
    console.log("   EVMFluxusCrossChain implementation deployment transaction hash:", evmFluxusCrossChainImplementation.hash);

    try {
      await evmFluxusCrossChainImplementation.waitForDeployment();
      console.log("   ✅ EVMFluxusCrossChain implementation deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", evmFluxusCrossChainImplementation.hash);
    }
  } catch (error: any) {
    console.log("   ❌ EVMFluxusCrossChain implementation deployment failed");
    console.log("   Error:", error?.message || error);
    throw error;
  }
  
  const evmFluxusCrossChainImplementationAddress = await evmFluxusCrossChainImplementation.getAddress();
  console.log("   EVMFluxusCrossChain Implementation deployed address:", evmFluxusCrossChainImplementationAddress);

  // 4. Prepare initialization parameters
  console.log("\n4. Prepare initialization parameters...");
  
  // Sample initialization parameters
  const initParams = {
    name: "Fluxus Universal EVM NFT2",
    symbol: "FUENFT",
    initialPrice: ethers.parseEther("0.001"),
    maxSupply: 10000,
    maxPrice: ethers.parseEther("1"),
    creatorFeePercent: ethers.parseEther("0.05"), // 5%
    feeVault: feeVaultAddress, // Use deployed FeeVault
    priceContract: priceContractAddress, // Use deployed Price contract
    creator: deployer.address,
    baseURI: "https://ipfs.io/ipfs/bafybeiglzppf7i4ki4lwcny62q5uc43qto7elsqakixnwh57r3ipxvie2m",
    gatewayAddress: evmConfig.gateway,
    gasLimit: evmConfig.gasLimit,
    supportMint: true
  };

  console.log("Initialization parameters:");
  console.log("  Name:", initParams.name);
  console.log("  Symbol:", initParams.symbol);
  console.log("  Initial Price:", ethers.formatEther(initParams.initialPrice), "ETH");
  console.log("  Max Supply:", initParams.maxSupply);
  console.log("  Max Price:", ethers.formatEther(initParams.maxPrice), "ETH");
  console.log("  Creator Fee:", ethers.formatEther(initParams.creatorFeePercent), "(", Number(ethers.formatEther(initParams.creatorFeePercent)) * 100, "%)");
  console.log("  Creator:", initParams.creator);
  console.log("  Base URI:", initParams.baseURI);
  console.log("  Gateway Address:", initParams.gatewayAddress);
  console.log("  Gas Limit:", initParams.gasLimit);
  console.log("  Support Mint:", initParams.supportMint);

  // 5. Deploy ERC1967Proxy for EVMFluxusCrossChain
  console.log("\n5. Deploy ERC1967Proxy for EVMFluxusCrossChain...");
  const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
  
  // Prepare initialization data
  const initData = evmFluxusCrossChainImplementation.interface.encodeFunctionData("initialize", [
    initParams.name,
    initParams.symbol,
    initParams.initialPrice,
    initParams.maxSupply,
    initParams.maxPrice,
    initParams.creatorFeePercent,
    initParams.feeVault,
    initParams.priceContract,
    initParams.creator,
    initParams.baseURI,
    initParams.gatewayAddress,
    initParams.gasLimit,
    initParams.supportMint
  ]);
  
  let evmFluxusCrossChainProxy;
  try {
    evmFluxusCrossChainProxy = await ERC1967Proxy.deploy(
      evmFluxusCrossChainImplementationAddress,
      initData
    );
    console.log("   ERC1967Proxy deployment transaction sent, waiting for confirmation...");
    console.log("   ERC1967Proxy deployment transaction hash:", evmFluxusCrossChainProxy.hash);
    
    try {
      await evmFluxusCrossChainProxy.waitForDeployment();
      console.log("   ✅ ERC1967Proxy deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", evmFluxusCrossChainProxy.hash);
    }
  } catch (error: any) {
    console.log("   ❌ ERC1967Proxy deployment failed");
    console.log("   Error:", error?.message || error);
    throw error;
  }
  
  const evmFluxusCrossChainProxyAddress = await evmFluxusCrossChainProxy.getAddress();
  console.log("   EVMFluxusCrossChain Proxy deployed address:", evmFluxusCrossChainProxyAddress);

  // 6. Verify proxy deployment
  console.log("\n6. Verify proxy deployment...");
  
  try {
    // Check if proxy contract has code
    const code = await ethers.provider.getCode(evmFluxusCrossChainProxyAddress);
    if (code !== "0x") {
      console.log("   ✅ Proxy deployment verified - contract has code");
    } else {
      console.log("   ❌ Proxy deployment failed - no code at address");
    }
  } catch (error: any) {
    console.log("   ⚠️  Proxy verification failed, but deployment might be successful");
    console.log("   Error:", error?.message || error);
  }

  // Deployment summary
  console.log("\n=== Deployment summary ===");
  console.log("Deployer account:", deployer.address);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);
  
  const deployedContracts = {
    FeeVault: feeVaultAddress,
    Price: priceContractAddress,
    EVMFluxusCrossChainImplementation: evmFluxusCrossChainImplementationAddress,
    EVMFluxusCrossChainProxy: evmFluxusCrossChainProxyAddress
  };

  console.log("\nDeployed contract addresses:");
  Object.entries(deployedContracts).forEach(([name, address]) => {
    console.log(`  ${name}: ${address}`);
  });

  // EVM specific configuration guide
  console.log("\n=== EVM CrossChain Configuration Guide ===");
  console.log("All contracts have been deployed and initialized with sample parameters.");
  console.log("Parameters used for initialization:");
  console.log(`  FeeVault Address: ${feeVaultAddress}`);
  console.log(`  Price Contract Address: ${priceContractAddress}`);
  console.log(`  Gateway Address: ${evmConfig.gateway}`);
  console.log(`  Gas Limit: ${evmConfig.gasLimit}`);
  console.log();
  console.log("To use this proxy contract, use the proxy address:");
  console.log(`  EVMFluxusCrossChain Proxy: "${evmFluxusCrossChainProxyAddress}"`);
  console.log();
  console.log("Note: The proxy contract is now ready to be used for EVM CrossChain NFT operations.");
  console.log("The proxy contract is already initialized and can be called directly.");
  console.log();
  console.log("IMPORTANT: You need to set the Universal contract address using setUniversal() function");
  console.log("after deploying the Universal NFT contract on ZetaChain.");

  return deployedContracts;
}

/**
 * Get EVM network configuration
 * @param chainId Chain ID
 * @returns EVM configuration object or null
 */
function getEVMConfig(chainId: number) {
  const configs: Record<number, {
    gateway: string;
    gasLimit: number;
    name: string;
  }> = {
    // Base Sepolia
    84532: {
      gateway: "0x0c487a766110c85d301d96e33579c5b317fa4995",
      gasLimit: 500000,
      name: "base_sepolia"
    },
    // BSC Testnet
    97: {
      gateway: "0x0c487a766110c85d301d96e33579c5b317fa4995",
      gasLimit: 500000,
      name: "bsc_testnet"
    }
  };

  return configs[chainId] || null;
}

main()
  .then((deployedContracts) => {
    console.log("\n✅ All EVM contracts deployed and initialized successfully!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("\n❌ Deployment failed:", error);
    process.exit(1);
  });
