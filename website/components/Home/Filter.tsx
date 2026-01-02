import { Button } from '@/components/ui/button';
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
  pagination,
  onTrendingFilterChange,
  onTimeFilterChange,
  onChainChange,
  onPageChange,
}: FilterProps) => {
  return (
    <div className="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-2 sm:gap-4 mb-2 sm:mb-4 px-1 sm:px-4">
      <Tabs value={trendingFilter} className="w-full sm:w-auto">
        <TabsList className="bg-[#171a1f] p-0 h-auto w-full sm:w-auto text-xs sm:text-base">
          <TabsTrigger
            value="trending"
            onClick={() => onTrendingFilterChange('trending')}
            className="flex-1 sm:flex-none px-4 py-1.5 rounded-xs data-[state=active]:bg-[#16A34A] data-[state=active]:text-black text-sm sm:text-base"
          >
            TRENDING
          </TabsTrigger>
          <TabsTrigger
            value="top"
            onClick={() => onTrendingFilterChange('top')}
            className="flex-1 sm:flex-none px-4 py-1.5 rounded-xs data-[state=active]:bg-[#16A34A] data-[state=active]:text-black text-sm sm:text-base"
          >
            TOP
          </TabsTrigger>
        </TabsList>
      </Tabs>
      <div className="flex flex-col sm:flex-row gap-1 sm:gap-2 w-full sm:w-auto">
        <div className="flex w-full sm:w-auto">
          {["1h", "24h", "7d"].map((time) => (
            <button
              key={time}
              className={`flex-1 sm:flex-none px-2 sm:px-3 py-1 sm:py-1.5 text-xs sm:text-sm ${activeTimeFilter === time ? "bg-[#2f343e] text-white" : "bg-[#171a1f] text-gray-400"}`}
              onClick={() => onTimeFilterChange(time as '1h' | '24h' | '7d')}
            >
              {time.toUpperCase()}
            </button>
          ))}
        </div>

        <select
          value={selectedChain}
          onChange={(e) => onChainChange(e.target.value)}
          className="w-full sm:w-auto h-9 sm:h-10 appearance-none bg-[#171a1f] border border-[#2f343e] rounded-lg px-3 sm:px-4 cursor-pointer hover:border-[#16A34A] focus:outline-none focus:border-[#16A34A] text-white text-xs sm:text-base"
        >
          {chains.map((chain) => (
            <option key={chain.id} value={chain.id} className="bg-[#171a1f] text-white">{chain.name}</option>
          ))}
        </select>
      </div>
    </div>
  );
};