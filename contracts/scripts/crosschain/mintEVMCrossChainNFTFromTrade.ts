import { ethers } from "hardhat";

const fluxusAddress = "0x5876a33b9F6bdaafDA81dF58eD6E3C24faBecEB3";
const priceAddress = "0x3299a867798D8C6bF64181883402641c9b30756E";
const tradeAddress = "0x334A3f51f09704D21ac1d8180ff81Cdcb27C9860";

async function main() {
  const Fluxus = await ethers.getContractFactory("EVMFluxusCrossChain");
  const fluxus = Fluxus.attach(fluxusAddress);

  const priceContract = await ethers.getContractAt("Price", priceAddress);
  const tradeContract = await ethers.getContractAt("Trade", tradeAddress);

  const buyPrice = await priceContract.getBuyPrice(fluxusAddress);
  const sellPrice = await priceContract.getSellPrice(fluxusAddress);

  console.log(`Buy Price: ${ethers.formatEther(buyPrice)} ETH`);
  console.log(`Sell Price: ${ethers.formatEther(sellPrice)} ETH`);

  console.log("Start Minting Fluxus token...");
  
  try {
    const [signer] = await ethers.getSigners();
    const signerAddress = await signer.getAddress();
        
    const balance = await ethers.provider.getBalance(signerAddress);
    console.log(`Current account address: ${signerAddress}`);
    console.log(`Account balance: ${ethers.formatEther(balance)} ETH`);

    const buyPrice = await priceContract.getBuyPriceAfterFee(fluxusAddress);
    console.log(`Mint Price: ${ethers.formatEther(buyPrice)} ETH`);

    if (balance < buyPrice) {
      console.log("Insufficient balance to mint");
      return;
    }

    console.log("Minting NFT...");
    const tx = await tradeContract.mint(
      fluxusAddress,
      {
      value: buyPrice,
      // gasLimit: 300000
    });

    console.log("Transaction hash:", tx.hash);
    
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