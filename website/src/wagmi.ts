import { getDefaultConfig } from '@rainbow-me/rainbowkit';
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
import { http } from 'wagmi';
import { defineChain } from 'viem';

export const customMainnet = {
  ...mainnet,
  iconUrl: '/ethereum.png',
}

export const customMantleSepoliaTestnet = {
  ...mantleSepoliaTestnet,
  iconUrl: '/mantle.svg',
}

export const customBerachainTestnet = {
  ...berachainTestnetbArtio,
  iconUrl: '/berachain.png',
}

export const customMonadTestnet = {
  ...monadTestnet,
  iconUrl: '/monad.png',
}

export const customScrollSepolia = {
  ...scrollSepolia,
  iconUrl: '/scroll.svg',
}

export const customMantleSepolia = {
  ...mantleSepoliaTestnet,
  iconUrl: '/mantle.svg',
}

export const customPharosTestnet = {
  ...defineChain({
    id: 688688,
    name: 'Pharos Testnet',
    nativeCurrency: { name: 'PHRS', symbol: 'PHRS', decimals: 18 },
    rpcUrls: {
      default: { http: ['https://testnet.dplabs-internal.com'] },
    },
    blockExplorers: {
      default: { name: 'Pharos Explorer', url: 'https://testnet.pharosscan.xyz' },
    },
    testnet: true,
  }),
  iconUrl: '/pharos.jpg',
}

export const customPharosDevnet = {
  ...defineChain({
    id: 50002,
    name: 'Pharos Devnet',
    nativeCurrency: { name: 'ETH', symbol: 'ETH', decimals: 18 },
    rpcUrls: {
      default: { http: ['https://devnet.dplabs-internal.com'] },
    },
    blockExplorers: {
      default: { name: 'Pharos Explorer', url: 'https://pharosscan.xyz/' },
    },
    testnet: true,
  }),
  iconUrl: '/pharos.jpg',
}

export const megaethTestnet = {
  ...defineChain({
    id: 6342,
    name: 'MEGA Testnet',
    nativeCurrency: { name: 'ETH', symbol: 'ETH', decimals: 18 },
    rpcUrls: {
      default: { http: ['https://carrot.megaeth.com/rpc'] },
    },
    blockExplorers: {
      default: { name: 'MegaETH Explorer', url: 'https://megaexplorer.xyz' },
    },
    testnet: true,
  }),
  iconUrl: '/megaeth.svg',
}

export const config = getDefaultConfig({
  appName: 'Fluxus',
  projectId: 'Fluxus',
  chains: [
    baseSepolia,
    // bscTestnet,
  ],
  transports: {
    [baseSepolia.id]: http(`https://sepolia.base.org`),
    // [bscTestnet.id]: http(`https://data-seed-prebsc-1-s1.bnbchain.org:8545`),
  },
  ssr: true,
});