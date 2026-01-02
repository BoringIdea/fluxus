import { ethers } from "hardhat";

async function main() {
  console.log("Transfer NFT cross-chain from EVM to ZetaChain...\n");
  
  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH\n");

  // Get network information
  const network = await ethers.provider.getNetwork();
  const chainId = Number(network.chainId);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);

  const TRADE_CONTRACT = "0x2EcbBeaB651348CbE32677dE6BcB2e7e4a94b26f";
  // Contract addresses - replace with your deployed addresses
  const EVM_NFT_CONTRACT = "0x2D78d51d9Bb7e08dE4ef9e3192cbf106353ca22b"; // Replace with your EVMFluxusCrossChain contract address
  
  // Transfer parameters - replace with your values
  const TOKEN_ID = 2; // The NFT token ID to transfer
  const RECEIVER_ADDRESS = "0x50753Ca349636CA8732762e8Ccf057d3999891A0"; // Receiver address on ZetaChain
  const DESTINATION_ADDRESS = "0x236b0DE675cC8F46AE186897fCCeFe3370C9eDeD"; // destination chain zrc20 address
  // const DESTINATION_ADDRESS = "0x0000000000000000000000000000000000000000"; // Use address(0) to transfer to ZetaChain
  const GAS_FEE = ethers.parseEther("0.001"); // Gas fee for cross-chain transfer
  
  // Verify contract address
  if (!ethers.isAddress(EVM_NFT_CONTRACT)) {
    throw new Error("Invalid EVM NFT contract address format");
  }
  
  if (!ethers.isAddress(RECEIVER_ADDRESS)) {
    throw new Error("Invalid receiver address format");
  }
  
  if (!ethers.isAddress(DESTINATION_ADDRESS)) {
    throw new Error("Invalid destination address format");
  }

  console.log("EVM NFT Contract:", EVM_NFT_CONTRACT);
  console.log("Trade Contract:", TRADE_CONTRACT);
  console.log("Token ID:", TOKEN_ID);
  console.log("Receiver Address:", RECEIVER_ADDRESS);
  console.log("Destination Address:", DESTINATION_ADDRESS);
  console.log("Gas Fee:", ethers.formatEther(GAS_FEE), "ETH");
  console.log();

  // Connect to the EVM NFT contract
  const EVMFluxusCrossChain = await ethers.getContractFactory("EVMFluxusCrossChain");
  const evmNFT = EVMFluxusCrossChain.attach(EVM_NFT_CONTRACT);

  // Connect to the Trade contract
  const Trade = await ethers.getContractFactory("Trade");
  const tradeContract = Trade.attach(TRADE_CONTRACT);

  // Verify contract ownership and token ownership
  try {
    const owner = await evmNFT.owner();
    console.log("Contract owner:", owner);
    
    const tokenOwner = await evmNFT.ownerOf(TOKEN_ID);
    console.log("Token owner:", tokenOwner);
    
    if (tokenOwner !== deployer.address) {
      throw new Error(`Token ${TOKEN_ID} is owned by ${tokenOwner}, but deployer is ${deployer.address}`);
    }
    console.log("✅ Token ownership verified");
  } catch (error) {
    console.error("❌ Failed to verify ownership:", error);
    process.exit(1);
  }

  // Check if universal address is set
  try {
    const universalAddress = await evmNFT.universal();
    console.log("Universal address:", universalAddress);
    
    if (universalAddress === "0x0000000000000000000000000000000000000000") {
      console.log("⚠️  Universal address is not set. Please set it using setUniversal() function first.");
      console.log("You can use the setUniversal.ts script to set the Universal NFT contract address on ZetaChain.");
      process.exit(1);
    }
    console.log("✅ Universal address is set");
  } catch (error) {
    console.error("❌ Failed to get universal address:", error);
    process.exit(1);
  }

  // Check gateway address
  try {
    const gatewayAddress = await evmNFT.gateway();
    console.log("Gateway address:", gatewayAddress);
    console.log("✅ Gateway address is set");
  } catch (error) {
    console.error("❌ Failed to get gateway address:", error);
    process.exit(1);
  }

  // Get token URI before transfer
  try {
    const tokenURI = await evmNFT.tokenURI(TOKEN_ID);
    console.log("Token URI:", tokenURI);
  } catch (error) {
    console.log("⚠️  Could not get token URI:", error);
  }

  // Check if Trade contract is approved for the token
  try {
    const approvedAddress = await evmNFT.getApproved(TOKEN_ID);
    console.log("Currently approved address:", approvedAddress);
    
    if (approvedAddress !== TRADE_CONTRACT) {
      console.log("Approving Trade contract for token transfer...");
      const approveTx = await evmNFT.approve(TRADE_CONTRACT, TOKEN_ID);
      console.log("Approve transaction sent:", approveTx.hash);
      await approveTx.wait();
      console.log("✅ Trade contract approved for token", TOKEN_ID);
    } else {
      console.log("✅ Trade contract already approved for token", TOKEN_ID);
    }
  } catch (error) {
    console.error("❌ Failed to approve Trade contract:", error);
    process.exit(1);
  }

  // Call transferCrossChain function
  try {
    console.log("Calling transferCrossChain function...");
    
    const tx = await tradeContract.transferCrossChain(
      EVM_NFT_CONTRACT,
      TOKEN_ID,
      RECEIVER_ADDRESS,
      DESTINATION_ADDRESS,
      { value: GAS_FEE }
    );
    
    console.log("Transaction sent, waiting for confirmation...");
    console.log("Transaction hash:", tx.hash);
    
    await tx.wait();
    console.log("✅ transferCrossChain transaction confirmed!");
    
  } catch (error) {
    console.error("❌ Failed to call transferCrossChain:", error);
    process.exit(1);
  }

  // Verify token is burned (should not exist anymore)
  try {
    console.log("\nVerifying token transfer...");
    await evmNFT.ownerOf(TOKEN_ID);
    console.log("⚠️  Token still exists - transfer may have failed");
  } catch (error) {
    console.log("✅ Token successfully burned - transfer initiated");
  }

  console.log("\n=== Transfer Summary ===");
  console.log("EVM NFT Contract:", EVM_NFT_CONTRACT);
  console.log("Token ID:", TOKEN_ID);
  console.log("Receiver Address:", RECEIVER_ADDRESS);
  console.log("Destination Address:", DESTINATION_ADDRESS);
  console.log("Gas Fee:", ethers.formatEther(GAS_FEE), "ETH");
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);
  
  console.log("\n✅ Cross-chain transfer initiated successfully!");
  console.log("\n=== Next Steps ===");
  console.log("1. Wait for the cross-chain transfer to complete (usually takes a few minutes)");
  console.log("2. Check the receiver address on ZetaChain for the transferred NFT");
  console.log("3. The NFT should appear in the Universal NFT contract on ZetaChain");
  console.log();
  console.log("=== Usage Instructions ===");
  console.log("To transfer to different destinations:");
  console.log("- To ZetaChain: Use DESTINATION_ADDRESS = 0x0000000000000000000000000000000000000000");
  console.log("- To other EVM chains: Use the ZRC20 address of the target chain");
  console.log("- Update RECEIVER_ADDRESS with the recipient address on the destination chain");
  console.log("- Adjust GAS_FEE based on the destination chain's gas requirements");
  console.log();
  console.log("Common ZRC20 addresses for different chains:");
  console.log("- Sepolia ETH: 0x0000c304D2934c00Db1d51995b9f6996AffD17c0");
  console.log("- BSC Testnet BNB: 0x0000c9ec4042283e8139c74f4c64bcd1e0b9b54f");
  console.log("- Base Sepolia ETH: 0x236b0DE675cC8F46AE186897fCCeFe3370C9eDeD");
}

main()
  .then(() => {
    console.log("\n✅ Cross-chain transfer script completed successfully!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("\n❌ Transfer failed:", error);
    process.exit(1);
  });
