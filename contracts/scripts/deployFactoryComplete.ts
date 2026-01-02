import { ethers } from "hardhat";
import { Contract } from "ethers";

async function main() {
  console.log("Start deploying Fluxus system contracts...\n");
  
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

  // 4. Deploy Fluxus implementation contract
  console.log("\n4. Deploy Fluxus implementation contract...");
  const Fluxus = await ethers.getContractFactory("Fluxus");
  const fluxusImplementation = await Fluxus.deploy();
  await fluxusImplementation.waitForDeployment();
  const fluxusImplementationAddress = await fluxusImplementation.getAddress();
  console.log("   Fluxus Implementation deployed address:", fluxusImplementationAddress);

  // 5. Deploy FluxusCrossChain implementation contract
  console.log("\n5. Deploy FluxusCrossChain implementation contract...");
  const FluxusCrossChain = await ethers.getContractFactory("EVMFluxusCrossChain");
  const fluxusCrossChainImplementation = await FluxusCrossChain.deploy();
  await fluxusCrossChainImplementation.waitForDeployment();
  const fluxusCrossChainImplementationAddress = await fluxusCrossChainImplementation.getAddress();
  console.log("   FluxusCrossChain Implementation deployed address:", fluxusCrossChainImplementationAddress);

  // 6. Deploy Factory contract
  console.log("\n6. Deploy Factory contract...");
  const Factory = await ethers.getContractFactory("Factory");
  const factory = await Factory.deploy(
    registryAddress,
    feeVaultAddress,
    priceContractAddress,
    fluxusImplementationAddress,
    fluxusCrossChainImplementationAddress
  );
  await factory.waitForDeployment();
  const factoryAddress = await factory.getAddress();
  console.log("   Factory deployed address:", factoryAddress);

  // Verify contract status
  console.log("\n=== Deployment completed, verify contract status ===");
  
  console.log("\nFactory contract configuration:");
  console.log("  Registry address:", await factory.registry());
  console.log("  FeeVault address:", await factory.feeVault());
  console.log("  Price contract address:", await factory.priceContract());
  console.log("  Fluxus Implementation address:", await factory.Fluxus_IMPLEMENTATION());
  console.log("  FluxusCrossChain Implementation address:", await factory.Fluxus_CROSS_CHAIN_IMPLEMENTATION());

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
    FluxusImplementation: fluxusImplementationAddress,
    FluxusCrossChainImplementation: fluxusCrossChainImplementationAddress,
    Factory: factoryAddress
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