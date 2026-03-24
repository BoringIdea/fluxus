import { Pagination as PaginationType } from '@/src/api/types';
import { Tabs, TabsList, TabsTrigger } from '../ui/tabs';

interface FilterProps {
  trendingFilter: 'trending' | 'top';
  activeTimeFilter: '1h' | '24h' | '7d';
  selectedChain: string;
  chains: { id: string; name: string }[];
  pagination: PaginationType;
  onTrendingFilterChange: (filter: 'trending' | 'top') => void;
  onTimeFilterChange: (filter: '1h' | '24h' | '7d') => void;
  onChainChange: (chainId: string) => void;
  onPageChange: (page: number) => void;
}

export const Filter = ({
  trendingFilter,
  activeTimeFilter,
  selectedChain,
  chains,
  onTrendingFilterChange,
  onTimeFilterChange,
  onChainChange,
}: FilterProps) => {
  return (
    <div className="mb-4 flex flex-col gap-3 border border-black/10 bg-[color:var(--bg-surface)] px-4 py-4 sm:flex-row sm:items-center sm:justify-between sm:px-6">
      <div className="flex flex-col gap-3 sm:flex-row sm:items-center">
        <div>
          <div className="flux-kicker mb-2">Rank</div>
          <Tabs value={trendingFilter} className="w-full sm:w-auto">
            <TabsList className="w-full sm:w-auto">
              <TabsTrigger value="trending" onClick={() => onTrendingFilterChange('trending')}>
                Trending
              </TabsTrigger>
              <TabsTrigger value="top" onClick={() => onTrendingFilterChange('top')}>
                Top
              </TabsTrigger>
            </TabsList>
          </Tabs>
        </div>

        <div>
          <div className="flux-kicker mb-2">Window</div>
          <div className="flex border border-black/10 bg-[color:var(--bg-muted)]">
            {['1h', '24h', '7d'].map((time) => (
              <button
                key={time}
                className={`px-3 py-2 font-primary text-[11px] uppercase tracking-[0.16em] transition-colors ${activeTimeFilter === time ? 'border border-black/10 bg-[color:var(--bg-surface)] text-[color:var(--text-primary)]' : 'text-[color:var(--text-muted)] hover:text-[color:var(--text-primary)]'}`}
                onClick={() => onTimeFilterChange(time as '1h' | '24h' | '7d')}
              >
                {time}
              </button>
            ))}
          </div>
        </div>
      </div>

      <div>
        <div className="flux-kicker mb-2">Chain</div>
        <select
          value={selectedChain}
          onChange={(e) => onChainChange(e.target.value)}
          className="h-10 min-w-[180px] appearance-none border border-black/10 bg-[color:var(--bg-surface)] px-3 font-primary text-[11px] uppercase tracking-[0.16em] text-[color:var(--text-primary)] outline-none transition-colors hover:bg-[color:var(--bg-card-hover)] focus:border-black/20"
        >
          {chains.map((chain) => (
            <option key={chain.id} value={chain.id}>
              {chain.name}
            </option>
          ))}
        </select>
      </div>
    </div>
  );
};
