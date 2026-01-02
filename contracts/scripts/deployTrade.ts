import { ethers } from "hardhat";

async function main() {
  const trade = await ethers.deployContract("Trade");
  await trade.waitForDeployment();

  console.log("Trade deployed to:", await trade.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });