import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-foundry";
import "@nomicfoundation/hardhat-verify";
import * as dotenv from "dotenv";

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.28",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          },
          viaIR: true,
          evmVersion: "paris"
        }
      },
      {
        version: "0.8.26",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          },
          viaIR: true,
          evmVersion: "paris"
        }
      }
    ]
  },
  networks: {
    "bsc-testnet": {
      url: `https://bsc-testnet-dataseed.bnbchain.org`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 97,
    },
    "zeta-testnet": {
      url: `https://zetachain-athens-evm.blockpi.network/v1/rpc/public`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 7001,
    },
    "base-sepolia": {
      url: `https://sepolia.base.org`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 84532,
    },
    "megaeth-testnet": {
      url: `https://carrot.megaeth.com/rpc`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 6342,
    },
    "pharos-testnet": {
      url: `https://testnet.dplabs-internal.com`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 688688,
      gas: 12000000,
    },
    "pharos-devnet": {
      url: `https://devnet.dplabs-internal.com`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 50002,
      gas: 12000000,
    },
    "optimism-sepolia": {
      url: `https://sepolia.optimism.io`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 11155420
    },
    "monad-testnet": {
      url: `https://testnet-rpc.monad.xyz`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 10143
    },
    "ethereum-sepolia": {
      url: `https://sepolia.drpc.org`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 11155111
    },
    "scroll-sepolia": {
      url: `https://sepolia-rpc.scroll.io`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 534351
    },
    "mantle-sepolia": {
      url: `https://rpc.sepolia.mantle.xyz`,
      accounts: process.env.WALLET_PRIVATE_KEY ? [process.env.WALLET_PRIVATE_KEY] : [],
      chainId: 5003
    }
  },
  sourcify: {
    enabled: true,
    apiUrl: "https://sourcify-api-monad.blockvision.org",
    browserUrl: "https://testnet.monadexplorer.com"
  },
  etherscan: {
    enabled: false,
  }
  // etherscan: {
  //   enabled: false,
  //   apiKey: {
  //     "optimism-sepolia": process.env.ETHERSCAN_API_KEY || "",
  //     "monad-testnet": process.env.MONAD_API_KEY || ""
  //   },
  //   customChains: [
  //     {
  //       network: "optimism-sepolia",
  //       chainId: 11155420,
  //       urls: {
  //         apiURL: "https://api-sepolia-optimistic.etherscan.io/api",
  //         browserURL: "https://sepolia-optimism.etherscan.io"
  //       }
  //     }
  // ]
  // }
};

export default config;