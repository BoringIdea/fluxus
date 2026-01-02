import { ethers } from "hardhat";

// Factory contract address, replace with your own
const factoryAddress = "0x8a048D14b32eF2944D8F39adA9B3D1A3ff6F211B";

async function main() {
  // Contract params
  const params = {
    name: "FluxusTest",
    symbol: "FluxusTEST",
    initialPrice: ethers.parseEther("0.001"),
    maxSupply: 10000n,
    creatorFeePercent: ethers.parseEther("0.05"),  // 5%
    imageUrl: "https://ipfs.io/ipfs/bafybeiglzppf7i4ki4lwcny62q5uc43qto7elsqakixnwh57r3ipxvie2m",
  };

  // Get Factory contract instance
  const Factory = await ethers.getContractFactory("Factory");
  const factory = Factory.attach(factoryAddress);

  // create a nft collection
  console.log("Creating new nft collection...");
  const tx = await factory.createFluxus(
    params.name,
    params.symbol,
    params.initialPrice,
    params.maxSupply,
    params.creatorFeePercent,
    params.imageUrl,
    {
      gasLimit: 12000000,
    }
  );

  console.log("Transaction hash:", tx.hash);
  
  const receipt = await tx.wait();
  
  const event = receipt?.logs.find(
    (log: { topics: any[]; }) => log.topics[0] === factory.interface.getEvent("FluxusCreated")?.topicHash
  );

  if (event) {
    const decodedEvent = factory.interface.parseLog({
      topics: event.topics,
      data: event.data
    });
    console.log("New Fluxus created at:", decodedEvent?.args[1]);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });