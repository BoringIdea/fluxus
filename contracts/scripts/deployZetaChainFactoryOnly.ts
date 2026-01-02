import { ethers } from "hardhat";

async function main() {
  console.log("Deploy ZetaChain Factory contract...\n");
  
  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH\n");

  // Replace with your deployed contract addresses
  const REGISTRY_ADDRESS = "0x..."; // Replace with your Registry contract address
  const FEE_VAULT_ADDRESS = "0x..."; // Replace with your FeeVault contract address
  const PRICE_CONTRACT_ADDRESS = "0x..."; // Replace with your Price contract address
  const UNISWAP_ROUTER_ADDRESS = "0x..."; // Replace with your Uniswap Router contract address
  const UNIVERSAL_CONTRACT_ADDRESS = "0x..."; // Replace with your Universal Contract contract address

  // Verify if the addresses are set
  const addresses = [
    { name: "Registry", address: REGISTRY_ADDRESS },
    { name: "FeeVault", address: FEE_VAULT_ADDRESS },
    { name: "Price Contract", address: PRICE_CONTRACT_ADDRESS },
    { name: "Uniswap Router", address: UNISWAP_ROUTER_ADDRESS },
    { name: "Universal Contract", address: UNIVERSAL_CONTRACT_ADDRESS }
  ];

  for (const { name, address } of addresses) {
    if (address === "0x...") {
      throw new Error(`Please set the ${name} contract address`);
    }
    if (!ethers.isAddress(address)) {
      throw new Error(`${name} address format is invalid: ${address}`);
    }
  }

  console.log("Used dependent contract addresses:");
  addresses.forEach(({ name, address }) => {
    console.log(`  ${name}: ${address}`);
  });
  console.log();

  // Deploy Factory contract
  console.log("1. Deploy Factory contract...");
  const ZetaChainFactory = await ethers.getContractFactory("ZetaChainFactory");
  const zetaChainFactory = await ZetaChainFactory.deploy(
    REGISTRY_ADDRESS,
    FEE_VAULT_ADDRESS,
    PRICE_CONTRACT_ADDRESS,
    UNISWAP_ROUTER_ADDRESS,
    UNIVERSAL_CONTRACT_ADDRESS
  );
  await zetaChainFactory.waitForDeployment();
  const zetaChainFactoryAddress = await zetaChainFactory.getAddress();
  console.log("   ZetaChain Factory deployed address:", zetaChainFactoryAddress);

  // Verify contract status
  console.log("\n=== Verify Factory contract configuration ===");
  console.log("  Registry address:", await zetaChainFactory.registry());
  console.log("  FeeVault address:", await zetaChainFactory.feeVault());
  console.log("  Price contract address:", await zetaChainFactory.priceContract());
  console.log("  Uniswap Router address:", await zetaChainFactory.uniswapRouter());
  console.log("  Universal Contract address:", await zetaChainFactory.universalContract());

  // Deployment summary
  console.log("\n=== Deployment summary ===");
  console.log("Deployer account:", deployer.address);
  console.log("Network:", (await ethers.provider.getNetwork()).name || "unknown");
  console.log("Chain ID:", (await ethers.provider.getNetwork()).chainId);
  console.log("ZetaChain Factory address:", zetaChainFactoryAddress);

  // Verification commands
  console.log("\n=== Verification commands ===");
  console.log("Verify Factory:");
  console.log(`npx hardhat verify --network <network-name> ${zetaChainFactoryAddress} "${REGISTRY_ADDRESS}" "${FEE_VAULT_ADDRESS}" "${PRICE_CONTRACT_ADDRESS}" "${UNISWAP_ROUTER_ADDRESS}" "${UNIVERSAL_CONTRACT_ADDRESS}"`);

  return {
    zetaChainFactory: zetaChainFactoryAddress
  };
}

main()
  .then((deployedContracts) => {
    console.log("\n✅ Factory contract deployed successfully!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("\n❌ Deployment failed:", error);
    process.exit(1);
  }); 