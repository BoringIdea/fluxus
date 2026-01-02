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
} from 'wagmi/chains';

import { customPharosDevnet, customPharosTestnet, megaethTestnet } from './wagmi';

export const isHideHomeAndLaunchPage = process.env.NEXT_PUBLIC_HIDE_HOME_AND_LAUNCH_PAGE === 'true';

export function getChainSymbol(chainId: number) {
  switch (chainId) {
    case mainnet.id:
      return mainnet.nativeCurrency.symbol;
    case sepolia.id:
      return sepolia.nativeCurrency.symbol;
    case optimismSepolia.id:
      return optimismSepolia.nativeCurrency.symbol;
    case mantle.id:
      return mantle.nativeCurrency.symbol;
    case mantleSepoliaTestnet.id:
      return mantleSepoliaTestnet.nativeCurrency.symbol;
    case monadTestnet.id:
      return monadTestnet.nativeCurrency.symbol;
    case megaethTestnet.id:
      return megaethTestnet.nativeCurrency.symbol;
    case customPharosDevnet.id:
      return customPharosDevnet.nativeCurrency.symbol;
    case customPharosTestnet.id:
      return customPharosTestnet.nativeCurrency.symbol;
    default:
      return "ETH";
  }
}

export function getExplorerUrl(chainId: number) {
  switch (chainId) {
    case mainnet.id:
      return mainnet.blockExplorers.default.url;
    case sepolia.id:
      return sepolia.blockExplorers.default.url;
    case optimism.id:
      return optimism.blockExplorers.default.url;
    case optimismSepolia.id:
      return optimismSepolia.blockExplorers.default.url;
    case arbitrum.id:
      return arbitrum.blockExplorers.default.url;
    case arbitrumSepolia.id:
      return arbitrumSepolia.blockExplorers.default.url;
    case base.id:
      return base.blockExplorers.default.url;
    case baseSepolia.id:
      return baseSepolia.blockExplorers.default.url;
    case mantle.id:
      return mantle.blockExplorers.default.url;
    case mantleSepoliaTestnet.id:
      return mantleSepoliaTestnet.blockExplorers.default.url;
    case scroll.id:
      return scroll.blockExplorers.default.url;
    case scrollSepolia.id:
      return scrollSepolia.blockExplorers.default.url;
    case berachain.id:
      return berachain.blockExplorers.default.url;
    case berachainTestnetbArtio.id:
      return berachainTestnetbArtio.blockExplorers.default.url;
    case monadTestnet.id:
      return monadTestnet.blockExplorers.default.url;
    case megaethTestnet.id:
      return megaethTestnet.blockExplorers.default.url;
    case customPharosDevnet.id:
      return customPharosDevnet.blockExplorers.default.url;
    case customPharosTestnet.id:
      return customPharosTestnet.blockExplorers.default.url;
    default:
      return "https://etherscan.io";
  }
}

export const PrimaryColor = '#16A34A'

export const sliceAddress = (address: string) => {
  return address.slice(0, 6) + '...' + address.slice(-4)
}

// Format number with max decimals and rounding
export const formatNumberWithMaxDecimals = (value: string | number, maxDecimals: number = 2): string => {
  const num = typeof value === 'string' ? parseFloat(value) : value;
  
  if (isNaN(num) || num === 0) {
    return '0';
  }
  
  // if it is an integer, return it
  if (num % 1 === 0) {
    return num.toString();
  }
  
  // convert to string to handle decimal places
  const numStr = num.toString();
  const [intPart, decPart] = numStr.split('.');
  
  if (!decPart) {
    return intPart;
  }
  
  // find the first maxDecimals non-zero decimal places
  let nonZeroCount = 0;
  let resultDecimals = '';
  let lastNonZeroIndex = -1;
  
  for (let i = 0; i < decPart.length; i++) {
    const digit = decPart[i];
    resultDecimals += digit;
    
    if (digit !== '0') {
      nonZeroCount++;
      lastNonZeroIndex = i;
      if (nonZeroCount >= maxDecimals) {
        break;
      }
    }
  }
  
  // If we have more digits than needed, round the last significant digit
  if (resultDecimals.length > lastNonZeroIndex + 1) {
    const significantPart = resultDecimals.substring(0, lastNonZeroIndex + 1);
    const nextDigit = parseInt(resultDecimals[lastNonZeroIndex + 1] || '0');
    
    if (nextDigit >= 5) {
      // Round up
      const roundedPart = (parseInt(significantPart) + 1).toString();
      resultDecimals = roundedPart.padStart(significantPart.length, '0');
    } else {
      resultDecimals = significantPart;
    }
  }
  
  // remove trailing zeros
  resultDecimals = resultDecimals.replace(/0+$/, '');
  
  return resultDecimals ? `${intPart}.${resultDecimals}` : intPart;
}

// Format number with max decimals and rounding (improved version)
export const formatNumberWithMaxDecimalsAndRounding = (value: string | number, maxDecimals: number = 2): string => {
  const num = typeof value === 'string' ? parseFloat(value) : value;
  
  if (isNaN(num) || num === 0) {
    return '0';
  }
  
  // if it is an integer, return it
  if (num % 1 === 0) {
    return num.toString();
  }
  
  // convert to string to handle decimal places
  const numStr = num.toString();
  const [intPart, decPart] = numStr.split('.');
  
  if (!decPart) {
    return intPart;
  }
  
  // find the first maxDecimals non-zero decimal places
  let nonZeroCount = 0;
  let resultDecimals = '';
  let lastNonZeroIndex = -1;
  
  for (let i = 0; i < decPart.length; i++) {
    const digit = decPart[i];
    resultDecimals += digit;
    
    if (digit !== '0') {
      nonZeroCount++;
      lastNonZeroIndex = i;
      if (nonZeroCount >= maxDecimals) {
        break;
      }
    }
  }
  
  // If we have more digits than needed, round the last significant digit
  if (resultDecimals.length > lastNonZeroIndex + 1) {
    const significantPart = resultDecimals.substring(0, lastNonZeroIndex + 1);
    const nextDigit = parseInt(resultDecimals[lastNonZeroIndex + 1] || '0');
    
    if (nextDigit >= 5) {
      // Round up
      const roundedPart = (parseInt(significantPart) + 1).toString();
      resultDecimals = roundedPart.padStart(significantPart.length, '0');
    } else {
      resultDecimals = significantPart;
    }
  }
  
  // remove trailing zeros
  resultDecimals = resultDecimals.replace(/0+$/, '');
  
  return resultDecimals ? `${intPart}.${resultDecimals}` : intPart;
}

export const getTokenImage = (baseUri: string, tokenId: string) => {
  return `${baseUri}/${tokenId}.json`;
}

export const getCollectionImage = (baseUri: string) => {
  return `${baseUri}/collection.json`;
}
