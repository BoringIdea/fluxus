import { Network, Alchemy } from "alchemy-sdk";
import {
  mainnet,
  sepolia,
  optimism,
  optimismSepolia,
  arbitrum,
  arbitrumSepolia,
  base,
  baseSepolia,
  mantle,
  mantleSepoliaTestnet,
  scroll,
  scrollSepolia,
  berachain,
  berachainTestnetbArtio,
  monadTestnet,
  bscTestnet,
} from 'wagmi/chains';
import { customPharosDevnet, customPharosTestnet, megaethTestnet } from "./wagmi";

export type SupportedChainId = 
  | typeof mainnet.id 
  | typeof sepolia.id 
  | typeof optimism.id 
  | typeof optimismSepolia.id 
  | typeof arbitrum.id 
  | typeof arbitrumSepolia.id 
  | typeof base.id 
  | typeof baseSepolia.id 
  | typeof mantle.id 
  | typeof mantleSepoliaTestnet.id 
  | typeof scroll.id 
  | typeof scrollSepolia.id 
  | typeof berachain.id 
  | typeof berachainTestnetbArtio.id 
  | typeof monadTestnet.id
  | typeof megaethTestnet.id
  | typeof customPharosTestnet.id
  | typeof customPharosDevnet.id
  | typeof bscTestnet.id;

export const ALCHEMY_NETWORK = {
  [mainnet.id]: Network.ETH_MAINNET,
  [sepolia.id]: Network.ETH_SEPOLIA,
  [optimism.id]: Network.OPT_MAINNET,
  [optimismSepolia.id]: Network.OPT_SEPOLIA,
  [arbitrum.id]: Network.ARB_MAINNET,
  [arbitrumSepolia.id]: Network.ARB_SEPOLIA,
  [base.id]: Network.BASE_MAINNET,
  [baseSepolia.id]: Network.BASE_SEPOLIA,
  [berachain.id]: Network.BERACHAIN_MAINNET,
  [mantle.id]: Network.MANTLE_MAINNET,
  [mantleSepoliaTestnet.id]: Network.MANTLE_SEPOLIA,
  [scroll.id]: Network.SCROLL_MAINNET,
  [scrollSepolia.id]: Network.SCROLL_SEPOLIA,
  [berachainTestnetbArtio.id]: Network.BERACHAIN_BARTIO,
  [monadTestnet.id]: 'Monad Testnet' as Network,
  [bscTestnet.id]: Network.BNB_TESTNET,
};

export const getAlchemy = (chainId: SupportedChainId) => {
  if (chainId === customPharosTestnet.id || chainId === customPharosDevnet.id || chainId === megaethTestnet.id) {
    return null;
  }
  return new Alchemy({
    apiKey: process.env.NEXT_PUBLIC_ALCHEMY_API_KEY,
    network: ALCHEMY_NETWORK[chainId],
  });
};