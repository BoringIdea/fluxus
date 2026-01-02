import { ethers } from "hardhat";

async function main() {
  console.log("Start setting up universal contract connection for EVM CrossChain NFT contract...\n");
  
  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH\n");

  // Get network information
  const network = await ethers.provider.getNetwork();
  const chainId = Number(network.chainId);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);

  // Contract addresses - replace with your deployed addresses
  // Example: Base Sepolia
  const EVM_NFT_CONTRACT = "0x2C66a9B3AF88Cf4C670273c7760eD797d79692b6"; // Replace with your EVMFluxusCrossChain contract address
  const ZETACHAIN_UNIVERSAL_CONTRACT = "0xE0A1Dc23f928A206821921270DA4D70027ADdE37"; // Replace with your ZetaChain Universal NFT contract address
  
  // Verify contract addresses
  if (!ethers.isAddress(EVM_NFT_CONTRACT)) {
    throw new Error("Invalid EVM NFT contract address format");
  }
  if (!ethers.isAddress(ZETACHAIN_UNIVERSAL_CONTRACT)) {
    throw new Error("Invalid ZetaChain Universal contract address format");
  }

  console.log("EVM NFT Contract:", EVM_NFT_CONTRACT);
  console.log("ZetaChain Universal Contract:", ZETACHAIN_UNIVERSAL_CONTRACT);
  console.log();

  // Connect to the EVM NFT contract
  const EVMFluxusCrossChain = await ethers.getContractFactory("EVMFluxusCrossChain");
  const evmNFT = EVMFluxusCrossChain.attach(EVM_NFT_CONTRACT);

  // Verify contract ownership
  try {
    const owner = await evmNFT.owner();
    if (owner !== deployer.address) {
      throw new Error(`Contract owner is ${owner}, but deployer is ${deployer.address}`);
    }
    console.log("✅ Contract ownership verified");
  } catch (error) {
    console.error("❌ Failed to verify contract ownership:", error);
    process.exit(1);
  }

  // Get current universal address
  try {
    const currentUniversal = await evmNFT.universal();
    console.log("Current Universal Address:", currentUniversal);
    
    if (currentUniversal !== ethers.ZeroAddress) {
      console.log("⚠️  Universal address is already set. This will update it.");
    }
  } catch (error) {
    console.log("Could not read current universal address");
  }

  // Set universal contract address
  console.log("\nSetting universal contract address...");
  try {
    const tx = await evmNFT.setUniversal(ZETACHAIN_UNIVERSAL_CONTRACT);
    console.log("Transaction sent:", tx.hash);
    
    await tx.wait();
    console.log("✅ Universal contract address set successfully!");
    
    // Verify the setting
    const newUniversal = await evmNFT.universal();
    if (newUniversal === ZETACHAIN_UNIVERSAL_CONTRACT) {
      console.log("✅ Universal address verification successful");
    } else {
      console.log("❌ Universal address verification failed");
      console.log("Expected:", ZETACHAIN_UNIVERSAL_CONTRACT);
      console.log("Actual:", newUniversal);
    }
    
  } catch (error) {
    console.error("❌ Failed to set universal contract address:", error);
    process.exit(1);
  }

  // Display contract configuration
  console.log("\n=== Contract Configuration ===");
  try {
    const gateway = await evmNFT.gateway();
    const gasLimit = await evmNFT.gasLimitAmount();
    const universal = await evmNFT.universal();
    
    console.log("Gateway Address:", gateway);
    console.log("Gas Limit:", gasLimit.toString());
    console.log("Universal Address:", universal);
    
  } catch (error) {
    console.log("Could not read some contract configuration");
  }

  // Display setup summary
  console.log("\n=== Setup Summary ===");
  console.log("EVM NFT Contract:", EVM_NFT_CONTRACT);
  console.log("ZetaChain Universal Contract:", ZETACHAIN_UNIVERSAL_CONTRACT);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);
  
  console.log("\n✅ Universal contract connection setup completed!");
  console.log("\nCross-chain flow:");
  console.log("EVM NFT → Gateway → ZetaChain Universal NFT");
  console.log("     ↓         ↓              ↓");
  console.log("  Burn NFT  Send Message   Mint NFT");
}


main()
  .then(() => {
    console.log("\n✅ Universal setup completed successfully!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("\n❌ Setup failed:", error);
    process.exit(1);
  });
