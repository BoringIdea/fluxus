import React from 'react';
import { useAccount, useChainId } from 'wagmi';
import { useRouter } from 'next/router';
import { useUserProfile } from '@/src/api/hooks';
import { sliceAddress, getChainSymbol } from '@/src/utils';
import { ethers } from 'ethers';
import Loading from '@/components/ui/Loading';

export default function UserPage() {
  const { address } = useAccount();
  const router = useRouter();
  const chainId = useChainId();
  const { totalCollections, totalNFTs, collections, isLoading } = useUserProfile(chainId, address as string);

  const handleCardClick = (collectionAddress: string) => {
    router.push(`/collection/${collectionAddress}`);
  };

  const formatFloorPrice = (price: string) => {
    if (!price || price === '0') return '0';
    return `${Number(ethers.formatEther(BigInt(price))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}`;
  }

  const statCards = [
    {
      label: 'Collections',
      value: totalCollections || 0,
      helper: 'Total collections you own',
      icon: '📚',
    },
    {
      label: 'NFTs',
      value: totalNFTs || 0,
      helper: 'Total NFTs in your wallet',
      icon: '🖼️',
    },
  ];

  return (
    <div className="w-full min-h-screen bg-transparent text-primary">
      <div className="mx-auto max-w-5xl space-y-8 px-4 pb-12 pt-8 sm:space-y-10 sm:pb-20 sm:pt-16">
        <div className="border-b border-black/10 pb-5">
          <div className="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
            <div>
              <p className="flux-kicker mb-3">Portfolio</p>
              <h1 className="flux-title text-[clamp(2.25rem,4.2vw,3.75rem)]">My Portfolio</h1>
            </div>
            <p className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
              {address ? sliceAddress(address) : 'Wallet not connected'}
            </p>
          </div>
        </div>

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          {statCards.map((card) => (
            <div
              key={card.label}
              className="flux-panel flex flex-col gap-4 px-5 py-5"
            >
              <div className="flex items-center justify-between">
                <div className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">{card.label}</div>
                <div className="text-xl">{card.icon}</div>
              </div>
              <div className="font-heading text-[42px] leading-none text-[color:var(--text-primary)]">{card.value}</div>
              <div className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">{card.helper}</div>
            </div>
          ))}
        </div>

        <div className="border-b border-black/10 pb-4">
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
            <div>
              <p className="flux-kicker mb-2">Holdings</p>
              <h2 className="flux-h2">Your Collections</h2>
            </div>
            <p className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
              {collections?.length || 0} tracked
            </p>
          </div>
        </div>

        {isLoading ? (
          <div className="flux-panel py-12 text-center">
            <div className="flex flex-col items-center gap-3">
              <Loading className="py-0" />
              <span className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
                Loading your collections...
              </span>
            </div>
          </div>
        ) : !collections || collections.length === 0 ? (
          <div className="flux-panel py-12 text-center">
            <div className="flex flex-col items-center gap-4">
              <div className="text-5xl">📦</div>
              <span className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
                No collections found
              </span>
              <p className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                Start collecting NFTs to see them here
              </p>
            </div>
          </div>
        ) : (
          <div className="grid grid-cols-1 gap-4 lg:grid-cols-2">
            {collections.map((col) => (
              <div
                key={col.address}
                onClick={() => handleCardClick(col.address)}
                className="flux-panel group flex cursor-pointer flex-col gap-5 px-5 py-5 transition-colors hover:bg-[color:var(--bg-card-hover)]"
              >
                <div className="flex items-start justify-between gap-3">
                  <div>
                    <p className="mb-2 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Collection</p>
                    <div className="font-heading text-[28px] leading-none text-[color:var(--text-primary)] transition-colors group-hover:text-[color:var(--color-primary)]">
                      {col.name}
                    </div>
                    <div className="mt-3 inline-flex items-center gap-2 border border-black/10 bg-[color:var(--bg-muted)] px-3 py-2 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                      {sliceAddress(col.address)}
                      <span className="text-[9px] tracking-[0.18em] text-[color:var(--color-primary)]">ID</span>
                    </div>
                  </div>
                  <div className="font-primary text-[14px] uppercase tracking-[0.14em] text-[color:var(--color-primary)] opacity-70 transition-opacity group-hover:opacity-100">
                    →
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-3 text-sm">
                  <div className="border border-black/10 bg-[color:var(--bg-muted)] px-4 py-3">
                    <div className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                      Holding NFTs
                    </div>
                    <div className="font-heading text-[26px] leading-none text-[color:var(--text-primary)]">{col.nftCount}</div>
                  </div>

                  <div className="border border-black/10 bg-[color:var(--bg-muted)] px-4 py-3">
                    <div className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                      Floor Price
                    </div>
                    <div className="font-heading text-[22px] leading-none text-[color:var(--text-primary)]">{formatFloorPrice(col.floorPrice)}</div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
