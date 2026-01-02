import { ethers } from "hardhat";

const fluxusAddress = "0x4bbe0b924931f7ce30b5f684316a65480840dbc4";
const priceAddress = "0xe11138b84939393F9badDf923Ff0778019f98a5c";
const tradeAddress = "0x971BE052417f30418284C25b31164249cbd7Ac3f";

async function main() {
  const Fluxus = await ethers.getContractFactory("Fluxus");
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
      (log: any) => log.topics[0] === fluxus.interface.getEvent("TokenMinted")?.topicHash
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