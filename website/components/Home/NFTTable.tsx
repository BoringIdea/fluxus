import { Collection } from '@/src/api/types';
import { ethers } from 'ethers';
import Image from 'next/image';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { useChainId } from 'wagmi';
import { getChainSymbol } from '@/src/utils';

interface NFTTableProps {
  collections: Collection[];
  onRowClick: (collectionId: string) => void;
  collectionsImages: string[];
}

export const NFTTable = ({ collections, onRowClick, collectionsImages }: NFTTableProps) => {
  const chainId = useChainId();

  const formatValue = (value?: string | number | bigint | null) => {
    if (!value) return '0';
    return `${Number(ethers.formatEther(BigInt(value.toString()))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}`;
  };

  return (
    <div className="w-full px-0 pb-4 sm:px-0">
      <div className="hidden overflow-hidden border border-black/10 bg-[color:var(--bg-surface)] sm:block">
        <Table>
          <TableHeader className="bg-[color:var(--bg-muted)]">
            <TableRow className="border-0 hover:bg-transparent">
              <TableHead className="text-[color:var(--text-secondary)]">Collection</TableHead>
              <TableHead className="hidden text-[color:var(--text-secondary)] sm:table-cell">Floor Price</TableHead>
              <TableHead className="hidden text-[color:var(--text-secondary)] sm:table-cell">Volume</TableHead>
              <TableHead className="hidden text-[color:var(--text-secondary)] sm:table-cell">Owners</TableHead>
              <TableHead className="text-[color:var(--text-secondary)]">Minted</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {collections?.map((collection, index) => (
              <TableRow key={collection.address} className="cursor-pointer border-t border-black/10" onClick={() => onRowClick(collection.address)}>
                <TableCell className="text-[color:var(--text-secondary)]">
                  <div className="flex items-center gap-3">
                    <div className="h-[50px] w-[50px] overflow-hidden border border-black/10 bg-[color:var(--bg-muted)]">
                      <Image
                        src={collectionsImages[index] || '/fluxus.svg'}
                        alt={collection?.name || ''}
                        width={50}
                        height={50}
                        className="h-full w-full object-cover"
                      />
                    </div>
                    <div className="flex flex-col gap-1">
                      <span className="font-serif text-[18px] font-semibold leading-none tracking-[-0.02em] text-[color:var(--text-primary)]">{collection?.name || ''}</span>
                      <span className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)] sm:hidden">
                        Floor {formatValue(collection?.floor_price)}
                      </span>
                    </div>
                  </div>
                </TableCell>
                <TableCell className="hidden font-medium text-[color:var(--text-secondary)] sm:table-cell">{formatValue(collection?.floor_price)}</TableCell>
                <TableCell className="hidden font-medium text-[color:var(--text-secondary)] sm:table-cell">{formatValue(collection?.period_volume)}</TableCell>
                <TableCell className="hidden font-medium text-[color:var(--text-secondary)] sm:table-cell">{collection?.owners || '0'}</TableCell>
                <TableCell className="font-medium text-[color:var(--text-secondary)]">{collection?.total_supply || '0'} / {collection?.max_supply || '0'}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <div className="flex flex-col gap-2 sm:hidden">
        {collections?.map((collection, index) => (
          <div key={collection.address} className="border border-black/10 bg-[color:var(--bg-surface)] p-4" onClick={() => onRowClick(collection.address)}>
            <div className="mb-3 flex items-center gap-3">
              <div className="h-[40px] w-[40px] overflow-hidden border border-black/10 bg-[color:var(--bg-muted)]">
                <Image
                  src={collectionsImages[index] || '/fluxus.svg'}
                  alt={collection?.name || ''}
                  width={40}
                  height={40}
                  className="h-full w-full object-cover"
                />
              </div>
              <div className="font-serif text-[18px] font-semibold leading-none tracking-[-0.02em] text-[color:var(--text-primary)]">{collection?.name || ''}</div>
            </div>
            <div className="grid grid-cols-2 gap-x-4 gap-y-2 font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
              <div>Floor: <span className="font-medium text-[color:var(--text-secondary)]">{formatValue(collection?.floor_price)}</span></div>
              <div>Volume: <span className="font-medium text-[color:var(--text-secondary)]">{formatValue(collection?.period_volume)}</span></div>
              <div>Owners: <span className="font-medium text-[color:var(--text-secondary)]">{collection?.owners || '0'}</span></div>
              <div>Minted: <span className="font-medium text-[color:var(--text-secondary)]">{collection?.total_supply || '0'} / {collection?.max_supply || '0'}</span></div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
