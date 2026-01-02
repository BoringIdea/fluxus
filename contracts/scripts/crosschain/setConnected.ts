import { ethers } from "hardhat";

async function main() {
  console.log("Setting connected address for cross-chain transfer...\n");

  const [deployer] = await ethers.getSigners();
  console.log("Deployer account:", deployer.address);
  console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ZETA\n");

  // Get network information
  const network = await ethers.provider.getNetwork();
  const chainId = Number(network.chainId);
  console.log("Network:", network.name || "unknown");
  console.log("Chain ID:", chainId);

  // Contract addresses - replace with your deployed addresses
  const ZETACHAIN_NFT_CONTRACT = "0x70429D1246c02BF15F5a38277DAB63e4D3C30682"; // Replace with your ZetaChainFluxusCrossChain contract address

  const connectedAddressConfig: { zrc20Address: string[], targetContractAddress: string[] } = await getConnectedAddress();
  for (let i = 0; i < connectedAddressConfig.zrc20Address.length; i++) {
    const ZRC20_ADDRESS = connectedAddressConfig.zrc20Address[i];
    const TARGET_CONTRACT_ADDRESS = connectedAddressConfig.targetContractAddress[i];

    // Verify contract address
    if (!ethers.isAddress(ZETACHAIN_NFT_CONTRACT)) {
      throw new Error("Invalid ZetaChain NFT contract address format");
    }

    if (!ethers.isAddress(ZRC20_ADDRESS)) {
      throw new Error("Invalid ZRC20 address format");
    }

    if (!ethers.isAddress(TARGET_CONTRACT_ADDRESS)) {
      throw new Error("Invalid target contract address format");
    }

    console.log("ZetaChain NFT Contract:", ZETACHAIN_NFT_CONTRACT);
    console.log("ZRC20 Address:", ZRC20_ADDRESS);
    console.log("Target Contract Address:", TARGET_CONTRACT_ADDRESS);
    console.log();

    // Connect to the already deployed ZetaChain NFT contract
    const zetaChainNFT = new ethers.Contract(
      ZETACHAIN_NFT_CONTRACT,
      [
        "function owner() view returns (address)",
        "function setConnected(address zrc20, address contractAddress) external",
        "function connected(address zrc20) view returns (address)"
      ],
      deployer
    );

    // Verify contract ownership
    try {
      const owner = await zetaChainNFT.owner();
      console.log("Contract owner:", owner);

      if (owner !== deployer.address) {
        throw new Error(`Contract is owned by ${owner}, but deployer is ${deployer.address}`);
      }
      console.log("✅ Contract ownership verified");
    } catch (error) {
      console.error("❌ Failed to verify ownership:", error);
      process.exit(1);
    }

    // Call setConnected function
    try {
      console.log("Calling setConnected function...");

      const tx = await zetaChainNFT.setConnected(ZRC20_ADDRESS, TARGET_CONTRACT_ADDRESS);

      // await for 2 seconds
      await new Promise(resolve => setTimeout(resolve, 10000));

      console.log("Transaction sent, waiting for confirmation...");
      console.log("Transaction hash:", tx.hash);

      await tx.wait();
      console.log("✅ setConnected transaction confirmed!");

    } catch (error) {
      console.error("❌ Failed to call setConnected:", error);
      process.exit(1);
    }

    // query connected address
    const connectedAddress = await zetaChainNFT.connected(ZRC20_ADDRESS);
    console.log("Connected address:", connectedAddress);
    if (connectedAddress !== TARGET_CONTRACT_ADDRESS) {
      throw new Error("Connected address does not match target contract address");
    }
    console.log("✅ Connected address verified");

    console.log("\n=== Connection Summary ===");
    console.log("ZetaChain NFT Contract:", ZETACHAIN_NFT_CONTRACT);
    console.log("ZRC20 Address:", ZRC20_ADDRESS);
    console.log("Target Contract Address:", TARGET_CONTRACT_ADDRESS);
    console.log("Network:", network.name || "unknown");
    console.log("Chain ID:", chainId);
  }
  console.log("\n✅ Cross-chain connection set successfully!");
  console.log("\n=== Next Steps ===");
  console.log("1. Now you can use zetachainTransferCrossChain.ts to transfer NFTs");
  console.log("2. The transfer will go from ZetaChain to the connected contract on Base Sepolia");
  console.log("3. Make sure the target contract on Base Sepolia is properly deployed and initialized");
  console.log();
  console.log("=== Usage Instructions ===");
  console.log("To connect to different chains:");
  console.log("- Update ZRC20_ADDRESS with the ZRC20 address of the target chain");
  console.log("- Update TARGET_CONTRACT_ADDRESS with the NFT contract address on the target chain");
  console.log("- Make sure the target contract supports cross-chain operations");
  console.log();
  console.log("Common ZRC20 addresses for different chains (ZetaChain testnet):");
  console.log("- Sepolia ETH: 0x0000c304D2934c00Db1d51995b9f6996AffD17c0");
  console.log("- BSC Testnet BNB: 0x0000c9ec4042283e8139c74f4c64bcd1e0b9b54f");
  console.log("- Base Sepolia ETH: 0x236b0DE675cC8F46AE186897fCCeFe3370C9eDeD");
}

async function getConnectedAddress() {
  return {
    zrc20Address: [
      "0xd97B1de3619ed2c6BEb3860147E30cA8A7dC9891", // BSC Testnet
      "0x236b0DE675cC8F46AE186897fCCeFe3370C9eDeD", // Base Sepolia
    ],
    targetContractAddress: [
      "0x84FE8d6E7Af2fda771b7574411Bb277810a801E3", // BSC Testnet
      "0x5876a33b9F6bdaafDA81dF58eD6E3C24faBecEB3", // Base Sepolia
    ]
  }
}

main()
  .then(() => {
    console.log("\n✅ setConnected script completed successfully!");
    process.exit(0);
  })
  .catch((error) => {
    console.error("\n❌ setConnected failed:", error);
    process.exit(1);
  });
