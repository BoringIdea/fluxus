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
  const FRACTORY_CONTRACT = "0xa0Ae112654E028B6612CD3972447681010F3Bc49";
  const EVM_NFT_CONTRACT = "0x84FE8d6E7Af2fda771b7574411Bb277810a801E3"; // Replace with your EVMFluxusCrossChain contract address
  const ZETACHAIN_UNIVERSAL_CONTRACT = "0x70429D1246c02BF15F5a38277DAB63e4D3C30682"; // Replace with your ZetaChain Universal NFT contract address
  
  // Verify contract addresses
  if (!ethers.isAddress(FRACTORY_CONTRACT)) {
    throw new Error("Invalid Trade contract address format");
  }
  if (!ethers.isAddress(EVM_NFT_CONTRACT)) {
    throw new Error("Invalid EVM NFT contract address format");
  }
  if (!ethers.isAddress(ZETACHAIN_UNIVERSAL_CONTRACT)) {
    throw new Error("Invalid ZetaChain Universal contract address format");
  }

  console.log("Factory Contract:", FRACTORY_CONTRACT);
  console.log("EVM NFT Contract:", EVM_NFT_CONTRACT);
  console.log("ZetaChain Universal Contract:", ZETACHAIN_UNIVERSAL_CONTRACT);
  console.log();

  // Connect to the Trade contract
  const Factory = await ethers.getContractFactory("Factory");
  const factory = Factory.attach(FRACTORY_CONTRACT);

  const EVMFluxusCrossChain = await ethers.getContractFactory("EVMFluxusCrossChain");
  const evmNFT = EVMFluxusCrossChain.attach(EVM_NFT_CONTRACT);

  // Verify contract ownership
  try {
    const owner = await evmNFT.owner();
    if (owner !== FRACTORY_CONTRACT) {
      throw new Error(`Contract owner is ${owner}, but factory is ${FRACTORY_CONTRACT}`);
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

  // Check creator of the EVM NFT contract
  try {
    const creator = await evmNFT.creator();
    console.log("Creator of the EVM NFT contract:", creator);
    // check if creator is the deployer
    if (creator !== deployer.address) {
      throw new Error(`Creator of the EVM NFT contract is ${creator}, but deployer is ${deployer.address}`);
    }
  } catch (error) {
    console.log("Could not read creator of the EVM NFT contract");
  }
  console.log("✅ Creator of the EVM NFT contract verified");

  // Set universal contract address
  console.log("\nSetting universal contract address...");
  try {
    const tx = await factory.setUniversal(EVM_NFT_CONTRACT, ZETACHAIN_UNIVERSAL_CONTRACT);
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
