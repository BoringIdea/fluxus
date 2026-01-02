import { Collection } from '@/src/api/types';
import { getChainSymbol } from '@/src/utils';
import { ethers } from 'ethers';
import { useRouter } from 'next/router';
import { useState, useEffect } from 'react';
import { useChainId } from 'wagmi';
import { Button } from '@/components/ui/button';

interface RecommendCollectionProps {
  collections: Collection[];
}

export const RecommendCollection = ({ collections }: RecommendCollectionProps) => {
  const chainId = useChainId();
  const router = useRouter();
  const [collection, setCollection] = useState<Collection | undefined>(collections[0]);
  const handleViewCollection = () => {
    router.push(`/collection/${collection?.address}`);
  };

  useEffect(() => {
    setCollection(collections[0]);
  }, [collections]);

  const stats = [
    {
      label: 'Floor Price',
      value: collection?.floor_price
        ? `${Number(ethers.formatEther(collection.floor_price)).toFixed(4)} ${getChainSymbol(chainId)}`
        : '0',
    },
    {
      label: 'Volume',
      value: collection?.total_volume
        ? `${Number(ethers.formatEther(BigInt(collection.total_volume.toString()))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}`
        : '0',
    },
    { label: 'Owners', value: collection?.owners?.toString() || '0' },
    {
      label: 'Minted',
      value: `${collection?.total_supply || 0} / ${collection?.max_supply || 0}`,
    },
  ];

  return (
    <div className="px-4 sm:px-6 md:px-10">
      <div className="relative w-full overflow-hidden border border-border bg-black/40 rounded-3xl">
        <img
          src={collection?.meta_data?.banner_image || "/Drug-banner.png"}
          className="absolute inset-0 w-full h-full object-cover"
          alt={collection?.name || 'Featured Collection'}
        />
        <div className="absolute inset-0 bg-gradient-to-t from-black via-black/60 to-black/10" />
        <div className="relative z-10 p-6 md:p-10 flex flex-col gap-6 text-primary">
          <div>
            <p className="text-[11px] uppercase tracking-[0.35em] text-secondary mb-1">Featured Collection</p>
            <h2 className="text-3xl sm:text-4xl font-bold text-white">{collection?.name || 'Collection'}</h2>
            <p className="text-secondary text-sm tracking-[0.2em]">
              {collection?.symbol?.toUpperCase()} • {collection?.max_supply || 0} Supply
            </p>
          </div>
          <div className="flex flex-col lg:flex-row items-start lg:items-center gap-4">
            <div className="flex-1 grid grid-cols-2 lg:grid-cols-4 gap-3">
              {stats.map((stat) => (
                <div key={stat.label} className="border border-white/20 bg-black/20 px-4 py-3 backdrop-blur">
                  <p className="text-[10px] uppercase tracking-[0.3em] text-secondary">{stat.label}</p>
                  <p className="text-white text-lg font-semibold">{stat.value}</p>
                </div>
              ))}
            </div>
            <div>
              <Button
                onClick={handleViewCollection}
                className="bg-fluxus-primary text-black border border-fluxus-primary px-6 py-6 rounded-none hover:bg-[#1FB455]"
              >
                View Collection →
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
