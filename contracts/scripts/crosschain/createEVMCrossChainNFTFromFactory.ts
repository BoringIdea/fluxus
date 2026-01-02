import { ethers } from "hardhat";

// Factory contract address, replace with your own
const factoryAddress = "0xa0Ae112654E028B6612CD3972447681010F3Bc49";

async function main() {
  // Contract params
  const params = {
    name: "Art Drug",
    symbol: "DRUG",
    initialPrice: ethers.parseEther("0.001"),
    maxSupply: 10000n,
    maxPrice: ethers.parseEther("1"),
    creatorFeePercent: ethers.parseEther("0.05"),  // 5%
    imageUrl: "https://red-naval-coyote-720.mypinata.cloud/ipfs/bafybeib3lobsuomwhmljxmdqvmritqwmb62gn6cvoq3nsuhm7ufbzsce4y",
    gatewayAddress: "0x0c487a766110c85d301d96e33579c5b317fa4995", // base sepolia
    gasLimit: 12000000,
    supportMint: true,
  };

  // Get Factory contract instance
  const Factory = await ethers.getContractFactory("Factory");
  const factory = Factory.attach(factoryAddress);

  // Calculate the address of the new nft collection
  const newNFTAddress = await factory.calculateFluxusCrossChainAddress(
    params.name,
    params.symbol,
    params.initialPrice,
    params.maxSupply,
    params.maxPrice,
    params.creatorFeePercent,
    params.imageUrl,
    params.gatewayAddress,
    params.gasLimit,
    params.supportMint,
  );

  console.log("New Fluxus CrossChain address:", newNFTAddress);

  // create a nft collection
  console.log("Creating new evm crosschain nft collection...");
  const tx = await factory.createFluxusCrossChain(
    params.name,
    params.symbol,
    params.initialPrice,
    params.maxSupply,
    params.maxPrice,
    params.creatorFeePercent,
    params.imageUrl,
    params.gatewayAddress,
    params.gasLimit,
    params.supportMint,
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
    if (decodedEvent?.args[1] !== newNFTAddress) {
      throw new Error("New Fluxus CrossChain created at is not the same as the calculated address");
    }
  }

  console.log("✅ New Fluxus CrossChain created at:", newNFTAddress);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });