import { useMemo, useState } from 'react';
import Mint from '@/components/Mint';
import Trade from '@/components/Trade';
import CrossChain from '@/components/CrossChain';
import Chat from '@/components/Chat';
import { useParams } from 'next/navigation';
import { useChainId } from 'wagmi';
import { useCollection } from '@/src/api';
import CollectionInfo from '@/components/CollectionInfo';

export default function CollectionDetailPage() {
  const params = useParams();
  const id = params?.id as string;
  const chainId = useChainId();

  const { data, mutate } = useCollection(chainId, id);
  const [activeTab, setActiveTab] = useState<'mint' | 'trade' | 'crosschain' | 'chat'>('mint');
  const collection = useMemo(() => {
    if (data?.data) {
      return data.data;
    }
    return;
  }, [data])

  // useEffect(() => {
  //   const fetchMetadata = async () => {
  //     if (collection?.base_uri) {
  //       try {
  //         const response = await fetch(collection.base_uri + '/collection.json');
  //         const metadata: CollectionMetadata = await response.json();
  //         setCollectionMetadata(metadata);
  //       } catch (error) {
  //         console.error('Failed to fetch collection metadata:', error);
  //       }
  //     }
  //   };

  //   fetchMetadata();
  // }, [collection]);

  const desktopTabs = [
    { key: 'mint', label: 'Mint' },
    { key: 'trade', label: 'Trade' },
    { key: 'crosschain', label: 'CrossChain' },
    { key: 'chat', label: 'Chat' },
  ] as const;

  return (
    <div className="w-full max-w-full overflow-x-hidden px-2 py-4 pb-24 sm:px-12 sm:py-10 sm:pb-0">
      <div className="w-full space-y-6">
        <CollectionInfo collection={collection} />
        <div className="border border-black/10 bg-[color:var(--bg-surface)]">
          <div className="hidden w-full justify-between border-b border-black/10 bg-[color:var(--bg-surface)] sm:flex">
            {desktopTabs.map((tab) => (
              <button
                key={tab.key}
                className={`flex-1 border-b px-6 py-4 font-primary text-[10px] uppercase tracking-[0.18em] transition-colors ${
                  activeTab === tab.key
                    ? 'border-[color:var(--fg-strong)] text-[color:var(--text-primary)]'
                    : 'border-transparent text-[color:var(--text-muted)]'
                }`}
                onClick={() => setActiveTab(tab.key)}
              >
                {tab.label}
              </button>
            ))}
          </div>

          <div className="sm:hidden">
            <div className="fixed bottom-0 left-0 right-0 z-40 border-t border-black/10 bg-[color:var(--bg-surface)]/95 backdrop-blur">
              <div className="grid grid-cols-4 w-full">
                {[
                  { key: 'mint', label: 'Mint', icon: 'MN' },
                  { key: 'trade', label: 'Trade', icon: 'TR' },
                  { key: 'crosschain', label: 'Cross', icon: 'CC' },
                  { key: 'chat', label: 'Chat', icon: 'AI' },
                ].map((tab) => (
                  <button
                    key={tab.key}
                    aria-label={tab.label}
                    className={`flex flex-col items-center justify-center gap-1 border-t py-3 font-primary text-[10px] uppercase tracking-[0.16em] ${
                      activeTab === tab.key
                        ? 'border-[color:var(--fg-strong)] bg-[color:var(--bg-muted)] text-[color:var(--text-primary)]'
                        : 'border-transparent text-[color:var(--text-muted)]'
                    }`}
                    onClick={() => setActiveTab(tab.key as typeof activeTab)}
                  >
                    <span className="text-[10px] leading-none">{tab.icon}</span>
                    <span className="leading-none">{tab.label}</span>
                  </button>
                ))}
              </div>
            </div>
          </div>

          <div className="m-0">
            {activeTab === 'mint' && <Mint contractAddress={id} collection={collection} />}
            {activeTab === 'trade' && <Trade contractAddress={id} collection={collection} />}
            {activeTab === 'crosschain' && (
              <CrossChain contractAddress={id} collection={collection} />
            )}
            {activeTab === 'chat' && (
              <Chat contractAddress={id} collection={collection} chainId={chainId} />
            )}
          </div>
        </div>
        <div className="sm:hidden h-20" />
      </div>
    </div>
  );
}
