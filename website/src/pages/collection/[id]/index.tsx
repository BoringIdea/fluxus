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
    <div className="w-full max-w-full overflow-x-hidden px-2 sm:px-12 py-4 sm:py-10 pb-24 sm:pb-0">
      <div className="w-full space-y-6">
        <CollectionInfo collection={collection} />
        <div className="border border-border bg-bg-card">
          {/* Desktop Tab Bar */}
          <div className="hidden sm:flex w-full justify-between border-b border-border bg-background rounded-none">
            {desktopTabs.map((tab) => (
              <button
                key={tab.key}
                className={`flex-1 px-6 py-4 text-xs font-bold uppercase tracking-[0.2em] border-b-2 transition-colors ${
                  activeTab === tab.key
                    ? 'text-fluxus-primary border-fluxus-primary'
                    : 'text-secondary border-transparent'
                }`}
                onClick={() => setActiveTab(tab.key)}
              >
                {tab.label}
              </button>
            ))}
          </div>

          {/* Mobile TabBar */}
          <div className="sm:hidden">
            <div className="fixed bottom-0 left-0 right-0 z-40 border-t border-border bg-background/95">
              <div className="grid grid-cols-4 w-full">
                {[
                  { key: 'mint', label: 'Mint', icon: '🎨' },
                  { key: 'trade', label: 'Trade', icon: '💱' },
                  { key: 'crosschain', label: 'Cross', icon: '🌉' },
                  { key: 'chat', label: 'Chat', icon: '💬' },
                ].map((tab) => (
                  <button
                    key={tab.key}
                    aria-label={tab.label}
                    className={`flex flex-col items-center justify-center gap-1 py-3 text-[10px] font-bold uppercase tracking-[0.2em] border-t-2 ${
                      activeTab === tab.key
                        ? 'text-fluxus-primary border-fluxus-primary bg-bg-card'
                        : 'text-secondary border-transparent'
                    }`}
                    onClick={() => setActiveTab(tab.key as typeof activeTab)}
                  >
                    <span className="text-lg leading-none">{tab.icon}</span>
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
        {/* Bottom spacer for mobile to avoid overlap with fixed TabBar */}
        <div className="sm:hidden h-20" />
      </div>
    </div>
  );
}
