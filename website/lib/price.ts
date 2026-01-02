export const getBatchBuyPrice = (maxSupply: string, currentSupply: string, initialPrice: string, amount: number, feePercent: string) => {
    const maxSupplyBigInt = BigInt(maxSupply);
    const currentSupplyBigInt = BigInt(currentSupply);
    const initialPriceBigInt = BigInt(initialPrice);
    
    let totalPrice = BigInt(0);
    for (let i = 0; i < amount; i++) {
        const supply = currentSupplyBigInt + BigInt(i);
        
        if (supply === BigInt(0)) {
            totalPrice = totalPrice + initialPriceBigInt;
            continue;
        }
        
        // Calculate price according to contract formula:
        // price = initialPrice + (initialPrice * 2 * sqrt(100 * supply * maxSupply) * sqrt(10000 * supply * supply)) / (maxSupply * maxSupply)
        
        // Calculate sqrt(100 * supply * maxSupply)
        const sqrt1Input = BigInt(100) * supply * maxSupplyBigInt;
        const sqrt1 = Math.sqrt(Number(sqrt1Input.toString()));
        
        // Calculate sqrt(10000 * supply * supply)
        const sqrt2Input = BigInt(10000) * supply * supply;
        const sqrt2 = Math.sqrt(Number(sqrt2Input.toString()));
        
        // Calculate final price with precision handling
        const price = initialPriceBigInt + 
            (initialPriceBigInt * BigInt(2) * 
            BigInt(Math.floor(sqrt1 * 1e6)) * 
            BigInt(Math.floor(sqrt2 * 1e6))) / 
            (maxSupplyBigInt * maxSupplyBigInt * BigInt(1e12));
        
        totalPrice = totalPrice + price;
    }

    totalPrice = totalPrice + (totalPrice * BigInt(feePercent)) / BigInt(1e18)
    return totalPrice;
};

export const getBatchSellPrice = (maxSupply: string, currentSupply: string, initialPrice: string, amount: number, feePercent: string) => {
    const maxSupplyBigInt = BigInt(maxSupply);
    const currentSupplyBigInt = BigInt(currentSupply);
    const initialPriceBigInt = BigInt(initialPrice);

    let totalPrice = BigInt(0);
    for (let i = 0; i < amount; i++) {
        const supply = currentSupplyBigInt - BigInt(i) - BigInt(1); // Decrease from current supply
        
        if (supply < BigInt(0)) {
            break; // Stop if supply becomes negative
        }
        
        if (supply === BigInt(0)) {
            totalPrice = totalPrice + initialPriceBigInt;
            continue;
        }
        
        // Calculate price according to contract formula:
        // price = initialPrice + (initialPrice * 2 * sqrt(100 * supply * maxSupply) * sqrt(10000 * supply * supply)) / (maxSupply * maxSupply)
        
        // Calculate sqrt(100 * supply * maxSupply)
        const sqrt1Input = BigInt(100) * supply * maxSupplyBigInt;
        const sqrt1 = Math.sqrt(Number(sqrt1Input.toString()));
        
        // Calculate sqrt(10000 * supply * supply)
        const sqrt2Input = BigInt(10000) * supply * supply;
        const sqrt2 = Math.sqrt(Number(sqrt2Input.toString()));
        
        // Calculate final price with precision handling
        const price = initialPriceBigInt + 
            (initialPriceBigInt * BigInt(2) * 
            BigInt(Math.floor(sqrt1 * 1e6)) * 
            BigInt(Math.floor(sqrt2 * 1e6))) / 
            (maxSupplyBigInt * maxSupplyBigInt * BigInt(1e12));
        
        totalPrice = totalPrice + price;
    }

    totalPrice = totalPrice - (totalPrice * BigInt(feePercent)) / BigInt(1e18)
    console.log('total sell price', totalPrice);
    
    return totalPrice;
};

