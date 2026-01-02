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
    <div className="px-3 sm:px-6 pb-24 mt-4">
      <div className="border border-border bg-bg-card flex flex-col lg:flex-row">
        <aside className="lg:w-40 border-b lg:border-b-0 lg:border-r border-border">
          {tabs.map((tab) => {
            const isActive = activeTradeTab === tab.key;
            return (
              <button
                key={tab.key}
                onClick={() => setActiveTradeTab(tab.key)}
                className={cn(
                  'w-full text-left px-4 py-4 border-l-2 uppercase tracking-[0.3em] text-[11px]',
                  isActive ? 'border-fluxus-primary text-fluxus-primary bg-bg-tertiary' : 'border-transparent text-secondary'
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
