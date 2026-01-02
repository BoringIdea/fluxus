import { ethers } from "hardhat";
import { Contract } from "ethers";

async function main() {
  console.log("Start deploying Fluxus system contracts...\n");

  const UNISWAP_ROUTER_ADDRESS = "0x..."; // Replace with your Uniswap Router contract address
  
  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH\n");

  // 1. Deploy FeeVault contract
  console.log("1. Deploy FeeVault contract...");
  const FeeVault = await ethers.getContractFactory("FeeVault");
  const feeVault = await FeeVault.deploy();
  await feeVault.waitForDeployment();
  const feeVaultAddress = await feeVault.getAddress();
  console.log("   FeeVault deployed address:", feeVaultAddress);

  // 2. Deploy Registry contract
  console.log("\n2. Deploy Registry contract...");
  const Registry = await ethers.getContractFactory("Registry");
  const registry = await Registry.deploy();
  await registry.waitForDeployment();
  const registryAddress = await registry.getAddress();
  console.log("   Registry deployed address:", registryAddress);

  // 3. Deploy Price contract
  console.log("\n3. Deploy Price contract...");
  const Price = await ethers.getContractFactory("Price");
  const priceContract = await Price.deploy();
  await priceContract.waitForDeployment();
  const priceContractAddress = await priceContract.getAddress();
  console.log("   Price deployed address:", priceContractAddress);

  // 4. Deploy ZetaChain Fluxus implementation contract
  console.log("\n4. Deploy ZetaChain Fluxus implementation contract...");
  const ZetaChainFluxus = await ethers.getContractFactory("ZetaChainFluxusCrossChainOptimized");
  const zetaChainFluxusImplementation = await ZetaChainFluxus.deploy();
  await zetaChainFluxusImplementation.waitForDeployment();
  const zetaChainFluxusImplementationAddress = await zetaChainFluxusImplementation.getAddress();
  console.log("   ZetaChain Fluxus Implementation deployed address:", zetaChainFluxusImplementationAddress);

  // 5. Deploy ZetaChain Factory contract
  console.log("\n5. Deploy ZetaChain Factory contract...");
  const ZetaChainFactory = await ethers.getContractFactory("ZetaChainFactory");
  const zetaChainFactory = await ZetaChainFactory.deploy(
    registryAddress,
    feeVaultAddress,
    priceContractAddress,
    UNISWAP_ROUTER_ADDRESS,
    zetaChainFluxusImplementationAddress
  );
  await zetaChainFactory.waitForDeployment();
  const zetaChainFactoryAddress = await zetaChainFactory.getAddress();
  console.log("   ZetaChain Factory deployed address:", zetaChainFactoryAddress);

  // Verify contract status
  console.log("\n=== Deployment completed, verify contract status ===");
  
  console.log("\nFactory contract configuration:");
  console.log("  Registry address:", await zetaChainFactory.registry());
  console.log("  FeeVault address:", await zetaChainFactory.feeVault());
  console.log("  Price contract address:", await zetaChainFactory.priceContract());
  console.log("  ZetaChain Fluxus Implementation address:", await zetaChainFactory.zetaChainFluxusImplementation());

  console.log("\nFeeVault contract configuration:");
  console.log("  Admin address:", await feeVault.admin());
  console.log("  Protocol fee rate:", ethers.formatEther(await feeVault.protocolFeePercent()), "(", Number(ethers.formatEther(await feeVault.protocolFeePercent())) * 100, "%)");

  // Deployment summary
  console.log("\n=== Deployment summary ===");
  console.log("Deployer account:", deployer.address);
  console.log("Network:", (await ethers.provider.getNetwork()).name || "unknown");
  console.log("Chain ID:", (await ethers.provider.getNetwork()).chainId);
  
  const deployedContracts = {
    FeeVault: feeVaultAddress,
    Registry: registryAddress,
    Price: priceContractAddress,
    ZetaChainFluxusImplementation: zetaChainFluxusImplementationAddress,
    ZetaChainFactory: zetaChainFactoryAddress
  };

  console.log("\nDeployed contract addresses:");
  Object.entries(deployedContracts).forEach(([name, address]) => {
    console.log(`  ${name}: ${address}`);
  });

  return deployedContracts;
}

main()
  .then((deployedContracts) => {
    console.log("\n✅ All contracts deployed successfully!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("\n❌ Deployment failed:", error);
    process.exit(1);
  }); 