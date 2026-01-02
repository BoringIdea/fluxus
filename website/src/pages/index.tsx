import type { NextPage } from 'next';
import { useEffect, useMemo, useState } from 'react';
import { useRouter } from 'next/router';

import { Collection, Pagination as PaginationType } from '@/src/api/types';
import { useCollections, useRecommendedCollection } from '@/src/api/hooks';
import { useChainId } from 'wagmi';
import { RecommendCollection } from '@/components/Home/RecommendCollection';
import { NFTTable } from '@/components/Home/NFTTable';
import { Filter } from '@/components/Home/Filter';
import ComingSoon from '@/components/CommingSoon';
import { isHideHomeAndLaunchPage } from '../utils';
import Pagination from '@/components/ui/pagination';
import Loading from '@/components/ui/Loading';
import { fetchCollectionImage } from '@/lib/utils';
import { string } from 'zod';

const DEFAULT_PAGE_SIZE = 10;

const Home: NextPage = () => {
  const chainId = useChainId();
  const router = useRouter();
  const [trendingFilter, setTrendingFilter] = useState<'trending' | 'top'>('trending');
  const [activeFilter, setActiveFilter] = useState<'trending' | 'top'>('trending');
  const [activeTimeFilter, setActiveTimeFilter] = useState<'1h' | '24h' | '7d'>('24h');
  const [selectedChain, setSelectedChain] = useState<string>('all');
  const [pagination, setPagination] = useState<PaginationType>({
    hasMore: true,
    limit: DEFAULT_PAGE_SIZE,
    offset: 0,
    total: '0'
  })
  const { data: recommendedCollection } = useRecommendedCollection(chainId);
  const { data: collections, isLoading: isLoadingCollections } = useCollections(chainId, pagination.offset, pagination.limit, activeTimeFilter);
  const [collectionsImages, setCollectionsImages] = useState<string[]>([]);
  useEffect(() => {
    if (collections) {
      const fetchImages = async () => {
        const images = await Promise.all(
          collections.data.map(async (collection) => {
            return await fetchCollectionImage(collection) || '/fluxus.svg';
          })
        );
        setCollectionsImages(images);
      };
      fetchImages();
    }
  }, [collections]);
  const chains = [
    { id: 'all', name: 'ALL CHAINS' },
  ];

  const handleRowClick = (collectionAddress: string) => {
    router.push(`/collection/${collectionAddress}`);
  };


  return (
    isHideHomeAndLaunchPage ? (
      <ComingSoon />
    ) : (
      <div className="container w-full max-w-full overflow-x-hidden px-2 sm:px-6 lg:px-8 pb-8 sm:pb-16 pt-8 sm:pt-20">
        <div className='mb-10 sm:mb-20'>
          <RecommendCollection collections={recommendedCollection?.data || []} />
          {/* <RecommendCollection collections={[...recommendedCollection, ...recommendedCollection, ...recommendedCollection, ...recommendedCollection, ...recommendedCollection]} /> */}
        </div>
        <Filter
          trendingFilter={trendingFilter}
          activeTimeFilter={activeTimeFilter}
          selectedChain={selectedChain}
          chains={chains}
          pagination={pagination}
          onTrendingFilterChange={setTrendingFilter}
          onTimeFilterChange={setActiveTimeFilter}
          onChainChange={setSelectedChain}
          onPageChange={(p) => setPagination({ ...pagination, offset: (p - 1) * (pagination.limit || DEFAULT_PAGE_SIZE) })}
        />
        {
          isLoadingCollections ? (
            <Loading />
          ) : collections && collections.data.length > 0 && (
            <div className="overflow-x-auto">
              <NFTTable
                collections={collections.data}
                onRowClick={handleRowClick}
                collectionsImages={collectionsImages}
              />
            </div>
          )
        }
        <div className="flex justify-center mt-4">
          <Pagination
            totalPages={(collections?.pagination?.limit || pagination.limit) > 0
              ? Math.max(1, Math.ceil(Number((collections?.pagination?.total ?? pagination.total) || 0) / (collections?.pagination?.limit || pagination.limit)))
              : 1}
            offset={Math.floor(((collections?.pagination?.offset ?? pagination.offset) || 0) / ((collections?.pagination?.limit || pagination.limit) || DEFAULT_PAGE_SIZE))}
            onPageChange={(p) => {
              const limit = collections?.pagination?.limit || pagination.limit || DEFAULT_PAGE_SIZE;
              setPagination(prev => ({ ...prev, limit, offset: (p - 1) * limit }));
            }}
          />
        </div>
      </div>
    )
  );
};

export default Home;