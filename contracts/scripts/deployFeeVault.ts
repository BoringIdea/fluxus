import { ethers } from "hardhat"; 

async function main() {
  const FeeVault = await ethers.getContractFactory("FeeVault");
  const feeVault = await FeeVault.deploy();
  await feeVault.waitForDeployment();

  console.log("FeeVault deployed to:", await feeVault.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });