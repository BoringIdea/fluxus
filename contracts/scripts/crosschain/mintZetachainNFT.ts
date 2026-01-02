import { ethers } from "hardhat";

const fluxusCrossChainAddress = "0x3122c1dEadB0321758217a9E396533f78c43D083";
const priceAddress = "0x3EAfFbbB500778568B9eFF76aF16e4155A809C67";

async function main() {
  console.log("Minting NFT on ZetaChain...");
  const Fluxus = await ethers.getContractFactory("ZetaChainFluxusCrossChain");
  const fluxus = Fluxus.attach(fluxusCrossChainAddress);

  const priceContract = await ethers.getContractAt("Price", priceAddress);

  const buyPrice = await priceContract.getBuyPrice(fluxusCrossChainAddress);
  const sellPrice = await priceContract.getSellPrice(fluxusCrossChainAddress);

  console.log(`Buy Price: ${ethers.formatEther(buyPrice)} ETH`);
  console.log(`Sell Price: ${ethers.formatEther(sellPrice)} ETH`);

  console.log("Start Minting Fluxus token...");
  
  try {
    const [signer] = await ethers.getSigners();
    const signerAddress = await signer.getAddress();
        
    const balance = await ethers.provider.getBalance(signerAddress);
    console.log(`Current account address: ${signerAddress}`);
    console.log(`Account balance: ${ethers.formatEther(balance)} ETH`);

    const buyPrice = await priceContract.getBuyPriceAfterFee(fluxusCrossChainAddress);
    console.log(`Mint Price: ${ethers.formatEther(buyPrice)} ETH`);

    if (balance < buyPrice) {
      console.log("Insufficient balance to mint");
      return;
    }

    console.log("Minting NFT...");
    const tx = await fluxus.mint(
      {
      value: buyPrice,
      // gasLimit: 300000
    });

    console.log("Transaction hash:", tx.hash);
    await new Promise(resolve => setTimeout(resolve, 2000));
    
    const receipt = await tx.wait();
    
    const event = receipt?.logs.find(
      (log: any) => log.topics[0] === fluxus.interface.getEvent("TokenMinted(address,uint256,uint256)")?.topicHash
    );

    if (event) {
      const decodedEvent = fluxus.interface.parseLog({
        topics: event.topics,
        data: event.data
      });
      console.log("Minting successful! TokenID:", decodedEvent?.args[1].toString());
      console.log("Payment amount:", ethers.formatEther(decodedEvent?.args[2]), "ETH");
      console.log("Creator fee:", ethers.formatEther(decodedEvent?.args[3]), "ETH");
    }
  } catch (error) {
    console.error("Minting failed:", error);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });