import React from 'react';
import { useAccount, useChainId } from 'wagmi';
import { useRouter } from 'next/router';
import { useUserProfile } from '@/src/api/hooks';
import { sliceAddress, getChainSymbol } from '@/src/utils';
import { ethers } from 'ethers';

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
      <div className="max-w-5xl mx-auto px-4 pb-12 sm:pb-20 pt-8 sm:pt-16 space-y-8">
        {/* Header Section */}
        <div className="border border-border bg-black/60 px-5 py-6">
          <p className="text-[11px] uppercase tracking-[0.35em] text-secondary mb-2">Dashboard</p>
          <div className="flex flex-col sm:flex-row sm:items-end sm:justify-between gap-3">
            <h1 className="text-3xl sm:text-4xl font-bold text-white">MY PORFOLIO</h1>
            <p className="text-[11px] uppercase tracking-[0.25em] text-secondary">
              Track holdings & manage collections
            </p>
          </div>
        </div>

        {/* Stats Cards */}
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {statCards.map((card) => (
            <div
              key={card.label}
              className="border border-border bg-gradient-to-b from-black/70 to-bg-card/30 px-5 py-5 flex flex-col gap-3"
            >
              <div className="flex items-center justify-between">
                <div className="text-[11px] uppercase tracking-[0.35em] text-secondary">{card.label}</div>
                <div className="text-xl">{card.icon}</div>
              </div>
              <div className="text-4xl font-extrabold text-white">{card.value}</div>
              <div className="text-[11px] uppercase tracking-[0.25em] text-secondary/90">{card.helper}</div>
            </div>
          ))}
        </div>

        {/* Collections Section */}
        <div className="border border-border px-5 py-4">
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
            <div>
              <p className="text-[11px] uppercase tracking-[0.35em] text-secondary">Section</p>
              <h2 className="text-2xl font-bold text-white">YOUR COLLECTIONS</h2>
            </div>
            <p className="text-[11px] uppercase tracking-[0.25em] text-secondary">
              {collections?.length || 0} tracked
            </p>
          </div>
        </div>

        {/* Collections List */}
        {isLoading ? (
          <div className="border border-border bg-bg-card py-12 text-center">
            <div className="flex flex-col items-center gap-3">
              <div className="w-8 h-8 border-2 border-fluxus-primary border-t-transparent rounded-full animate-spin"></div>
              <span className="text-secondary uppercase tracking-[0.25em] text-[11px]">
                Loading your collections...
              </span>
            </div>
          </div>
        ) : !collections || collections.length === 0 ? (
          <div className="border border-border bg-bg-card py-12 text-center">
            <div className="flex flex-col items-center gap-4">
              <div className="text-5xl">📦</div>
              <span className="text-secondary uppercase tracking-[0.3em] text-[11px]">
                No collections found
              </span>
              <p className="text-secondary/70 text-[11px] tracking-[0.2em]">
                Start collecting NFTs to see them here
              </p>
            </div>
          </div>
        ) : (
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
            {collections.map((col) => (
              <div
                key={col.address}
                onClick={() => handleCardClick(col.address)}
                className="border border-border bg-gradient-to-b from-black via-bg-card to-bg-card/60 px-5 py-5 cursor-pointer hover:border-fluxus-primary/70 hover:bg-black/70 transition-all duration-200 group flex flex-col gap-5"
              >
                <div className="flex items-start justify-between gap-3">
                  <div>
                    <p className="text-[11px] uppercase tracking-[0.35em] text-secondary">Collection</p>
                    <div className="font-bold text-2xl text-white group-hover:text-fluxus-primary transition-colors">
                      {col.name}
                    </div>
                    <div className="mt-2 text-xs font-mono text-primary bg-black/50 border border-border px-3 py-1 inline-flex items-center gap-2">
                      {sliceAddress(col.address)}
                      <span className="text-[9px] uppercase tracking-[0.3em] text-secondary">ID</span>
                    </div>
                  </div>
                  <div className="text-2xl opacity-60 group-hover:opacity-100 transition-opacity text-fluxus-primary">
                    →
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-3 text-sm">
                  <div className="border border-border/70 bg-black/40 px-4 py-3">
                    <div className="text-[10px] uppercase tracking-[0.3em] text-secondary mb-1">
                      Holding NFTs
                    </div>
                    <div className="text-2xl font-bold text-white">{col.nftCount}</div>
                  </div>

                  <div className="border border-border/70 bg-black/40 px-4 py-3">
                    <div className="text-[10px] uppercase tracking-[0.3em] text-secondary mb-1">
                      Floor Price
                    </div>
                    <div className="text-xl font-bold text-white">{formatFloorPrice(col.floorPrice)}</div>
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
