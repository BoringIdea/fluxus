import { ethers } from "hardhat";

async function main() {
  console.log("Transfer NFT cross-chain from ZetaChain to EVM...\n");
  
  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ZETA\n");

  // Get network information
  const network = await ethers.provider.getNetwork();
  const chainId = Number(network.chainId);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);

  // Contract addresses - replace with your deployed addresses
  const ZETACHAIN_NFT_CONTRACT = "0x3122c1dEadB0321758217a9E396533f78c43D083"; // Replace with your ZetaChainFluxusCrossChain contract address
  
  // Transfer parameters - replace with your values
  const TOKEN_ID = 1; // The NFT token ID to transfer
  const RECEIVER_ADDRESS = "0x50753Ca349636CA8732762e8Ccf057d3999891A0"; // Receiver address on target chain
  const DESTINATION_ADDRESS = "0x236b0DE675cC8F46AE186897fCCeFe3370C9eDeD"; // ZRC20 address for Base Sepolia
  const GAS_FEE = ethers.parseEther("0.01"); // Gas fee for cross-chain transfer (in ZETA) 
  const GAS_LIMIT = ethers.parseEther("0.01");
  
  // Verify contract address
  if (!ethers.isAddress(ZETACHAIN_NFT_CONTRACT)) {
    throw new Error("Invalid ZetaChain NFT contract address format");
  }
  
  if (!ethers.isAddress(RECEIVER_ADDRESS)) {
    throw new Error("Invalid receiver address format");
  }
  
  if (!ethers.isAddress(DESTINATION_ADDRESS)) {
    throw new Error("Invalid destination address format");
  }

  console.log("ZetaChain NFT Contract:", ZETACHAIN_NFT_CONTRACT);
  console.log("Token ID:", TOKEN_ID);
  console.log("Receiver Address:", RECEIVER_ADDRESS);
  console.log("Destination Address (ZRC20):", DESTINATION_ADDRESS);
  console.log("Gas Fee:", ethers.formatEther(GAS_FEE), "ZETA");
  console.log();

  // Connect to the ZetaChain NFT contract
  const ZetaChainFluxusCrossChain = await ethers.getContractFactory("ZetaChainFluxusCrossChain");
  const zetaChainNFT = ZetaChainFluxusCrossChain.attach(ZETACHAIN_NFT_CONTRACT);

  // Verify contract ownership and token ownership
  // try {
  //   const owner = await zetaChainNFT.owner();
  //   console.log("Contract owner:", owner);
    
  //   const tokenOwner = await zetaChainNFT.ownerOf(TOKEN_ID);
  //   console.log("Token owner:", tokenOwner);
    
  //   if (tokenOwner !== deployer.address) {
  //     throw new Error(`Token ${TOKEN_ID} is owned by ${tokenOwner}, but deployer is ${deployer.address}`);
  //   }
  //   console.log("✅ Token ownership verified");
  // } catch (error) {
  //   console.error("❌ Failed to verify ownership:", error);
  //   process.exit(1);
  // }


  // Get token URI before transfer
  // try {
  //   const tokenURI = await zetaChainNFT.tokenURI(TOKEN_ID);
  //   console.log("Token URI:", tokenURI);
  // } catch (error) {
  //   console.log("⚠️  Could not get token URI:", error);
  // }

  // Call transferCrossChain function
  try {
    console.log("Calling transferCrossChain function...");
    
    const tx = await zetaChainNFT.transferCrossChain(
      TOKEN_ID,
      RECEIVER_ADDRESS,
      DESTINATION_ADDRESS,
      { value: GAS_FEE, gasLimit: 50000000 }
    );
    
    console.log("Transaction sent, waiting for confirmation...");
    console.log("Transaction hash:", tx.hash);
    // await for 2 seconds
    await new Promise(resolve => setTimeout(resolve, 2000));
    
    await tx.wait();
    console.log("✅ transferCrossChain transaction confirmed!");
    
  } catch (error) {
    console.error("❌ Failed to call transferCrossChain:", error);
    // Print the basic info of the zetaChainNFT contract
    // print the gateway address
    const gateway = await zetaChainNFT.gateway();
    console.log("Gateway address:", gateway);
    // print the uniswapRouter address
    const uniswapRouter = await zetaChainNFT.uniswapRouter();
    console.log("Uniswap Router address:", uniswapRouter);
    // print the gasLimitAmount
    const gasLimitAmount = await zetaChainNFT.gasLimitAmount();
    console.log("Gas Limit Amount:", gasLimitAmount);
    // print the connected address
    const connectedAddress = await zetaChainNFT.connected(DESTINATION_ADDRESS);
    console.log("Connected address:", connectedAddress);
    
    process.exit(1);
  }

  // Verify token is burned (should not exist anymore)
  try {
    console.log("\nVerifying token transfer...");
    await zetaChainNFT.ownerOf(TOKEN_ID);
    console.log("⚠️  Token still exists - transfer may have failed");
  } catch (error) {
    console.log("✅ Token successfully burned - transfer initiated");
  }

  console.log("\n=== Transfer Summary ===");
  console.log("ZetaChain NFT Contract:", ZETACHAIN_NFT_CONTRACT);
  console.log("Token ID:", TOKEN_ID);
  console.log("Receiver Address:", RECEIVER_ADDRESS);
  console.log("Destination Address (ZRC20):", DESTINATION_ADDRESS);
  console.log("Gas Fee:", ethers.formatEther(GAS_FEE), "ZETA");
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);
  
  console.log("\n✅ Cross-chain transfer initiated successfully!");
  console.log("\n=== Next Steps ===");
  console.log("1. Wait for the cross-chain transfer to complete (usually takes a few minutes)");
  console.log("2. Check the receiver address on the target chain for the transferred NFT");
  console.log("3. The NFT should appear in the connected contract on the target chain");
  console.log();
  console.log("=== Usage Instructions ===");
  console.log("To transfer to different destinations:");
  console.log("- Update DESTINATION_ADDRESS with the ZRC20 address of the target chain");
  console.log("- Update RECEIVER_ADDRESS with the recipient address on the target chain");
  console.log("- Adjust GAS_FEE based on the destination chain's gas requirements");
  console.log("- Make sure the destination is connected using setConnected() function");
  console.log();
  console.log("Common ZRC20 addresses for different chains (ZetaChain testnet):");
  console.log("- Sepolia ETH: 0x0000c304D2934c00Db1d51995b9f6996AffD17c0");
  console.log("- BSC Testnet BNB: 0x0000c9ec4042283e8139c74f4c64bcd1e0b9b54f");
  console.log("- Base Sepolia ETH: 0x236b0DE675cC8F46AE186897fCCeFe3370C9eDeD");
  console.log();
  console.log("Note: The transfer process involves:");
  console.log("1. Burning the NFT on ZetaChain");
  console.log("2. Swapping ZETA for the target chain's gas token");
  console.log("3. Sending a message to the target chain via Gateway");
  console.log("4. Minting the NFT on the target chain");
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
