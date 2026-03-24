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

  useEffect(() => {
    setCollection(collections[0]);
  }, [collections]);

  const handleViewCollection = () => {
    router.push(`/collection/${collection?.address}`);
  };

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
    <section className="border border-black/10 bg-[color:var(--bg-surface)] p-6 sm:p-8">
      <div className="grid gap-8 lg:grid-cols-[1fr_0.95fr] lg:items-end">
        <div className="space-y-5">
          <div className="flux-kicker">Featured Collection</div>
          <div>
            <h2 className="flux-h2">{collection?.name || 'Collection'}</h2>
            <p className="mt-3 font-primary text-[11px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
              {collection?.symbol?.toUpperCase()} • {collection?.max_supply || 0} supply
            </p>
          </div>
          <div className="overflow-hidden border border-black/10 bg-[color:var(--bg-muted)]">
            <img
              src={collection?.meta_data?.banner_image || '/Drug-banner.png'}
              className="h-[220px] w-full object-cover"
              alt={collection?.name || 'Featured Collection'}
            />
          </div>
        </div>

        <div className="space-y-4">
          <div className="grid gap-3 sm:grid-cols-2">
            {stats.map((stat) => (
              <div key={stat.label} className="border border-black/10 bg-[color:var(--bg-surface)] px-4 py-4">
                <p className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">{stat.label}</p>
                <p className="mt-2 font-heading text-[24px] leading-none text-[color:var(--text-primary)]">{stat.value}</p>
              </div>
            ))}
          </div>
          <Button onClick={handleViewCollection} className="h-11 w-full sm:w-auto px-6">
            View Collection
          </Button>
        </div>
      </div>
    </section>
  );
};
