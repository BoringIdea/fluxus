import { ethers } from "hardhat";
import { Contract } from "ethers";

async function main() {
  console.log("Start deploying ZetaChainFluxusCrossChainOptimized contract...\n");
  
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

  // 3. Deploy libraries first
  console.log("\n3. Deploy libraries...");
  
  // Deploy ZetaChainCrossChainLib
  console.log("3.1. Deploy ZetaChainCrossChainLib...");
  const ZetaChainCrossChainLib = await ethers.getContractFactory("ZetaChainCrossChainLib");
  
  let zetaChainCrossChainLib;
  try {
    zetaChainCrossChainLib = await ZetaChainCrossChainLib.deploy();
    console.log("   ZetaChainCrossChainLib deployment transaction sent, waiting for confirmation...");
    console.log("   ZetaChainCrossChainLib deployment transaction hash:", zetaChainCrossChainLib.hash);
    await new Promise(resolve => setTimeout(resolve, 2000));

    try {
      await zetaChainCrossChainLib.waitForDeployment();
      console.log("   ✅ ZetaChainCrossChainLib deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", zetaChainCrossChainLib.hash);
    }
  } catch (error: any) {
    console.log("   ❌ ZetaChainCrossChainLib deployment failed");
    console.log("   Error:", error?.message || error);
    throw error;
  }
  
  const zetaChainCrossChainLibAddress = await zetaChainCrossChainLib.getAddress();
  console.log("   ZetaChainCrossChainLib deployed address:", zetaChainCrossChainLibAddress);

  // 4. Deploy ZetaChainFluxusCrossChainOptimized implementation contract with library links
  console.log("\n3.2. Deploy ZetaChainFluxusCrossChainOptimized implementation contract...");
  
  // Create contract factory with library links
  const ZetaChainFluxusCrossChainOptimized = await ethers.getContractFactory("ZetaChainFluxusCrossChainOptimized", {
    libraries: {
      ZetaChainCrossChainLib: zetaChainCrossChainLibAddress
    }
  });
  
  let zetaChainFluxusCrossChainOptimizedImplementation;
  try {
    zetaChainFluxusCrossChainOptimizedImplementation = await ZetaChainFluxusCrossChainOptimized.deploy();
    console.log("   ZetaChainFluxusCrossChainOptimized implementation deployment transaction sent, waiting for confirmation...");
    console.log("   ZetaChainFluxusCrossChainOptimized implementation deployment transaction hash:", zetaChainFluxusCrossChainOptimizedImplementation.hash);
    await new Promise(resolve => setTimeout(resolve, 2000));

    try {
      await zetaChainFluxusCrossChainOptimizedImplementation.waitForDeployment();
      console.log("   ✅ ZetaChainFluxusCrossChainOptimized implementation deployment confirmed");
    } catch (error: any) {
      console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
      console.log("   Error:", error?.message || error);
      console.log("   Transaction hash:", zetaChainFluxusCrossChainOptimizedImplementation.hash);
    }
  } catch (error: any) {
    console.log("   ❌ ZetaChainFluxusCrossChainOptimized implementation deployment failed");
    console.log("   Error:", error?.message || error);
    
    if (error.message.includes("max code size exceeded")) {
      console.log("\n   🔧 Contract size is still too large even after splitting.");
      console.log("   Consider further optimizations:");
      console.log("   1. Remove more unused functions");
      console.log("   2. Use more libraries");
      console.log("   3. Split into even smaller contracts");
    }
    throw error;
  }
  
  if (zetaChainFluxusCrossChainOptimizedImplementation) {
    const zetaChainFluxusCrossChainOptimizedImplementationAddress = await zetaChainFluxusCrossChainOptimizedImplementation.getAddress();
    console.log("   ZetaChainFluxusCrossChainOptimized Implementation deployed address:", zetaChainFluxusCrossChainOptimizedImplementationAddress);

    // 3. Prepare initialization parameters
    console.log("\n3.3. Prepare initialization parameters...");
    
    const initParams = {
      name: "Fluxus Universal NFT Optimized2",
      symbol: "FUNS",
      initialPrice: ethers.parseEther("0.001"),
      maxSupply: 10000,
      maxPrice: ethers.parseEther("1"),
      creatorFeePercent: ethers.parseEther("0.05"), // 5%
      feeVault: feeVaultAddress,
      priceContract: priceContractAddress,
      creator: deployer.address,
      baseURI: "https://ipfs.io/ipfs/bafybeiglzppf7i4ki4lwcny62q5uc43qto7elsqakixnwh57r3ipxvie2m",
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

    // 4. Deploy ERC1967Proxy for ZetaChainFluxusCrossChainSplit
    console.log("\n3.4. Deploy ERC1967Proxy for ZetaChainFluxusCrossChainSplit...");
    const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
    
    const initData = zetaChainFluxusCrossChainOptimizedImplementation.interface.encodeFunctionData("initialize", [
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
    
    let zetaChainFluxusCrossChainOptimizedProxy;
    try {
      zetaChainFluxusCrossChainOptimizedProxy = await ERC1967Proxy.deploy(
        zetaChainFluxusCrossChainOptimizedImplementationAddress,
        initData
      );
      console.log("   ERC1967Proxy deployment transaction sent, waiting for confirmation...");
      console.log("   ERC1967Proxy deployment transaction hash:", zetaChainFluxusCrossChainOptimizedProxy.hash);
      await new Promise(resolve => setTimeout(resolve, 2000));

      try {
        await zetaChainFluxusCrossChainOptimizedProxy.waitForDeployment();
        console.log("   ✅ ERC1967Proxy deployment confirmed");
      } catch (error: any) {
        console.log("   ⚠️  waitForDeployment failed, but deployment might be successful");
        console.log("   Error:", error?.message || error);
        console.log("   Transaction hash:", zetaChainFluxusCrossChainOptimizedProxy.hash);
      }
    } catch (error: any) {
      console.log("   ❌ ERC1967Proxy deployment failed");
      console.log("   Error:", error?.message || error);
      throw error;
    }
    
    if (zetaChainFluxusCrossChainOptimizedProxy) {
      const zetaChainFluxusCrossChainOptimizedProxyAddress = await zetaChainFluxusCrossChainOptimizedProxy.getAddress();
      console.log("   ZetaChainFluxusCrossChainOptimized Proxy deployed address:", zetaChainFluxusCrossChainOptimizedProxyAddress);
      
      console.log("\n✅ Deployment summary:");
      console.log("   FeeVault:", feeVaultAddress);
      console.log("   Price:", priceContractAddress);
      console.log("   ZetaChainCrossChainLib:", zetaChainCrossChainLibAddress);
      console.log("   ZetaChainFluxusCrossChainOptimized Implementation:", zetaChainFluxusCrossChainOptimizedImplementationAddress);
      console.log("   ZetaChainFluxusCrossChainOptimized Proxy:", zetaChainFluxusCrossChainOptimizedProxyAddress);
      console.log("\n🎉 Successfully deployed optimized version of ZetaChainFluxusCrossChain!");
    }
  }
}

function getZetaChainConfig(chainId: number) {
  const configs = {
    7001: { // zeta-testnet
      gateway: "0x6c533f7fe93fae114d0954697069df33c9b74fd7",
      uniswapRouter: "0x2ca7d64A7EFE2D62A725E2B35Cf7230D6677FfEe",
      gasLimit: 1000000
    }
  };
  
  return configs[chainId as keyof typeof configs];
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("\n❌ Deployment failed:", error);
    process.exit(1);
  });
