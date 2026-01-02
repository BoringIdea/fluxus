import { ethers } from "hardhat";
import { Contract } from "ethers";


async function main() {
  console.log("Start deploying ZetaChainFluxusCrossChain implementation contract...\n");
  
  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH\n");
  
  // Get network information
  const network = await ethers.provider.getNetwork();
  const chainId = Number(network.chainId);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);

  // ZetaChain network configuration
  const zetaChainConfig = getZetaChainConfig(chainId);
  if (!zetaChainConfig) {
    throw new Error(`Unsupported ZetaChain network, Chain ID: ${chainId}`);
  }

  console.log("ZetaChain Configuration:");
  console.log("  Gateway Address:", zetaChainConfig.gateway);
  console.log("  Uniswap Router Address:", zetaChainConfig.uniswapRouter);
  console.log("  Default Gas Limit:", zetaChainConfig.gasLimit);
  console.log();

  // 1. Deploy FeeVault contract
  console.log("1. Deploy FeeVault contract...");
  const FeeVault = await ethers.getContractFactory("FeeVault");
  
  let feeVault;
  try {
    feeVault = await FeeVault.deploy();
    console.log("   FeeVault deployment transaction sent, waiting for confirmation...");
    console.log("   FeeVault deployment transaction hash:", feeVault.hash);
    await new Promise(resolve => setTimeout(resolve, 2000));
    
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
    await new Promise(resolve => setTimeout(resolve, 2000));

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

  // 3. Deploy ZetaChainFluxusCrossChain implementation contract
  console.log("\n3. Deploy ZetaChainFluxusCrossChain implementation contract...");
  const ZetaChainFluxusCrossChain = await ethers.getContractFactory("ZetaChainFluxusCrossChain");
  
  let zetaChainFluxusCrossChainImplementation;
  try {
    zetaChainFluxusCrossChainImplementation = await ZetaChainFluxusCrossChain.deploy();
    console.log("   ZetaChainFluxusCrossChain implementation deployment transaction sent, waiting for confirmation...");
    console.log("   ZetaChainFluxusCrossChain implementation deployment transaction hash:", zetaChainFluxusCrossChainImplementation.hash);
    await new Promise(resolve => setTimeout(resolve, 2000));

    try {
      await zetaChainFluxusCrossChainImplementation.waitForDeployment();
      console.log("   ✅ ZetaChainFluxusCrossChain implementation deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", zetaChainFluxusCrossChainImplementation.hash);
    }
  } catch (error: any) {
    console.log("   ❌ ZetaChainFluxusCrossChain implementation deployment failed");
    console.log("   Error:", error?.message || error);
    throw error;
  }
  
  const zetaChainFluxusCrossChainImplementationAddress = await zetaChainFluxusCrossChainImplementation.getAddress();
  console.log("   ZetaChainFluxusCrossChain Implementation deployed address:", zetaChainFluxusCrossChainImplementationAddress);

  // 4. Prepare initialization parameters
  console.log("\n4. Prepare initialization parameters...");
  
  // Sample initialization parameters
  const initParams = {
    name: "Fluxus Universal NFT",
    symbol: "FUZNF",
    initialPrice: ethers.parseEther("0.001"),
    maxSupply: 10000,
    maxPrice: ethers.parseEther("1"),
    creatorFeePercent: ethers.parseEther("0.05"), // 5%
    feeVault: feeVaultAddress, // Use deployed FeeVault
    priceContract: priceContractAddress, // Use deployed Price contract
    creator: deployer.address,
    baseURI: "https://api.example.com/metadata/",
    gatewayAddress: zetaChainConfig.gateway,
    uniswapRouterAddress: zetaChainConfig.uniswapRouter, 
    gasLimit: zetaChainConfig.gasLimit,
    supportMint: true
  };

  console.log("Initialization parameters:");
  console.log("  Name:", initParams.name);
  console.log("  Symbol:", initParams.symbol);
  console.log("  Initial Price:", ethers.formatEther(initParams.initialPrice), "ZETA");
  console.log("  Max Supply:", initParams.maxSupply);
  console.log("  Max Price:", ethers.formatEther(initParams.maxPrice), "ZETA");
  console.log("  Creator Fee:", ethers.formatEther(initParams.creatorFeePercent), "(", Number(ethers.formatEther(initParams.creatorFeePercent)) * 100, "%)");
  console.log("  Creator:", initParams.creator);
  console.log("  Base URI:", initParams.baseURI);
  console.log("  Gateway Address:", initParams.gatewayAddress);
  console.log("  Uniswap Router Address:", initParams.uniswapRouterAddress);
  console.log("  Gas Limit:", initParams.gasLimit);
  console.log("  Support Mint:", initParams.supportMint);

  // 5. Deploy ERC1967Proxy for ZetaChainFluxusCrossChain
  console.log("\n5. Deploy ERC1967Proxy for ZetaChainFluxusCrossChain...");
  const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
  
  // Prepare initialization data
  const initData = zetaChainFluxusCrossChainImplementation.interface.encodeFunctionData("initialize", [
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
    initParams.uniswapRouterAddress,
    initParams.gasLimit,
    initParams.supportMint
  ]);
  
  let zetaChainFluxusCrossChainProxy;
  try {
    zetaChainFluxusCrossChainProxy = await ERC1967Proxy.deploy(
      zetaChainFluxusCrossChainImplementationAddress,
      initData
    );
    console.log("   ERC1967Proxy deployment transaction sent, waiting for confirmation...");
    console.log("   ERC1967Proxy deployment transaction hash:", zetaChainFluxusCrossChainProxy.hash);
    await new Promise(resolve => setTimeout(resolve, 2000));

    try {
      await zetaChainFluxusCrossChainProxy.waitForDeployment();
      console.log("   ✅ ERC1967Proxy deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", zetaChainFluxusCrossChainProxy.hash);
    }
  } catch (error: any) {
    console.log("   ❌ ERC1967Proxy deployment failed");
    console.log("   Error:", error?.message || error);
    throw error;
  }
  
  const zetaChainFluxusCrossChainProxyAddress = await zetaChainFluxusCrossChainProxy.getAddress();
  console.log("   ZetaChainFluxusCrossChain Proxy deployed address:", zetaChainFluxusCrossChainProxyAddress);
  
  // 6. Verify proxy deployment
  console.log("\n6. Verify proxy deployment...");
  
  try {
    // Check if proxy contract has code
    const code = await ethers.provider.getCode(zetaChainFluxusCrossChainProxyAddress);
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
    ZetaChainFluxusCrossChainImplementation: zetaChainFluxusCrossChainImplementationAddress,
    ZetaChainFluxusCrossChainProxy: zetaChainFluxusCrossChainProxyAddress
  };

  console.log("\nDeployed contract addresses:");
  Object.entries(deployedContracts).forEach(([name, address]) => {
    console.log(`  ${name}: ${address}`);
  });

  // ZetaChain specific configuration guide
  console.log("\n=== ZetaChain CrossChain Configuration Guide ===");
  console.log("All contracts have been deployed and initialized with sample parameters.");
  console.log("Parameters used for initialization:");
  console.log(`  FeeVault Address: ${feeVaultAddress}`);
  console.log(`  Price Contract Address: ${priceContractAddress}`);
  console.log(`  Uniswap Router Address: ${zetaChainConfig.uniswapRouter}`);
  console.log(`  Gas Limit: ${zetaChainConfig.gasLimit}`);
  console.log();
  console.log("To use this proxy contract, use the proxy address:");
  console.log(`  ZetaChainFluxusCrossChain Proxy: "${zetaChainFluxusCrossChainProxyAddress}"`);
  console.log();
  console.log("Note: The proxy contract is now ready to be used for ZetaChain CrossChain NFT operations.");
  console.log("The proxy contract is already initialized and can be called directly.");

  return deployedContracts;
}

/**
 * Get ZetaChain network configuration
 * @param chainId Chain ID
 * @returns ZetaChain configuration object or null
 */
function getZetaChainConfig(chainId: number) {
  const configs: Record<number, {
    gateway: string;
    uniswapRouter: string;
    gasLimit: number;
    name: string;
  }> = {
    // ZetaChain Mainnet
    7000: {
      gateway: "0x6c533f7fe93fae114d0954697069df33c9b74fd7",
      uniswapRouter: "0x2ca7d64A7EFE2D62A725E2B35Cf7230D6677FfEe",
      gasLimit: 1000000,
      name: "zeta-mainnet"
    },
    // ZetaChain Testnet
    7001: {
      gateway: "0x6c533f7fe93fae114d0954697069df33c9b74fd7",
      uniswapRouter: "0x2ca7d64A7EFE2D62A725E2B35Cf7230D6677FfEe",
      gasLimit: 1000000,
      name: "zeta-testnet"
    }
  };

  return configs[chainId] || null;
}

main()
  .then((deployedContracts) => {
    console.log("\n✅ All ZetaChain contracts deployed and initialized successfully!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("\n❌ Deployment failed:", error);
    process.exit(1);
  });
