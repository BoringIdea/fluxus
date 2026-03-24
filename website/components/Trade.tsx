import { useState } from 'react';
import { cn } from '@/lib/utils';
import Buy from './Buy';
import Sell from './Sell';
import Activity from './Activity';
import Holders from './Holders';

interface TradeProps {
  contractAddress: string;
  collection?: any;
}

const tabs = [
  { key: 'buy', label: 'Buy' },
  { key: 'sell', label: 'Sell' },
  { key: 'activity', label: 'Activity' },
  { key: 'holders', label: 'Holders' },
] as const;

export default function Trade({ contractAddress, collection }: TradeProps) {
  const [activeTradeTab, setActiveTradeTab] = useState<typeof tabs[number]['key']>('buy');

  return (
    <div className="mt-4 px-3 pb-24 sm:px-6">
      <div className="flex flex-col border border-black/10 bg-[color:var(--bg-surface)] lg:flex-row">
        <aside className="border-b border-black/10 lg:w-40 lg:border-b-0 lg:border-r">
          {tabs.map((tab) => {
            const isActive = activeTradeTab === tab.key;
            return (
              <button
                key={tab.key}
                onClick={() => setActiveTradeTab(tab.key)}
                className={cn(
                  'w-full border-l-2 px-4 py-4 text-left font-primary text-[10px] uppercase tracking-[0.18em]',
                  isActive ? 'border-[color:var(--fg-strong)] bg-[color:var(--bg-muted)] text-[color:var(--text-primary)]' : 'border-transparent text-[color:var(--text-muted)]'
                )}
                type="button"
              >
                {tab.label}
              </button>
            );
          })}
        </aside>
        <section className="flex-1 p-4 lg:p-6">
          {activeTradeTab === 'buy' && <Buy contractAddress={contractAddress} collection={collection} />}
          {activeTradeTab === 'sell' && <Sell contractAddress={contractAddress} collection={collection} />}
          {activeTradeTab === 'activity' && <Activity collectionAddress={contractAddress} />}
          {activeTradeTab === 'holders' && <Holders collectionAddress={contractAddress} />}
        </section>
      </div>
    </div>
  );
}
