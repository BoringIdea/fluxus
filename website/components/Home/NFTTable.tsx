import { Collection } from '@/src/api/types';
import { ethers } from 'ethers';
import Image from 'next/image';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"

import { useChainId } from 'wagmi';
import { getChainSymbol } from '@/src/utils';

interface NFTTableProps {
  collections: Collection[];
  onRowClick: (collectionId: string) => void;
  collectionsImages: string[];
}

export const NFTTable = ({ collections, onRowClick, collectionsImages }: NFTTableProps) => {
  const chainId = useChainId();
  return (
    <div className="w-full overflow-x-auto px-0 sm:px-4 pb-4">
      {/* 桌面端表格 */}
      <div className="rounded-lg min-w-[480px] sm:min-w-0 overflow-hidden border border-[#171a1f] hidden sm:block">
        <Table>
          <TableHeader className="bg-[#171a1f]">
            <TableRow className="border-0 hover:bg-transparent">
              <TableHead className="text-gray-400 font-normal text-sm sm:text-base">COLLECTION</TableHead>
              <TableHead className="text-gray-400 font-normal text-sm sm:text-base hidden sm:table-cell">FLOOR PRICE</TableHead>
              <TableHead className="text-gray-400 font-normal text-sm sm:text-base hidden sm:table-cell">VOLUME</TableHead>
              <TableHead className="text-gray-400 font-normal text-sm sm:text-base hidden sm:table-cell">OWNERS</TableHead>
              <TableHead className="text-gray-400 font-normal text-sm sm:text-base">MINTED</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {collections?.map((collection, index) => (
              <TableRow
                key={collection.address}
                className="border-t border-[#171a1f] hover:bg-[#171a1f]/50 cursor-pointer"
                onClick={() => onRowClick(collection.address)}
              >
                <TableCell className="text-gray-400 font-normal text-sm sm:text-base">
                  <div className="flex flex-row items-center gap-3">
                    <div className="w-[50px] h-[50px] rounded-lg overflow-hidden">
                      <Image
                        src={collectionsImages[index] || '/fluxus.svg'}
                        alt={collection?.name || ''}
                        width={50}
                        height={50}
                        className="w-full h-full object-cover"
                      />
                    </div>
                    <div className="flex flex-col">
                      <span>{collection?.name || ''}</span>
                      <span className="sm:hidden text-xs mt-1">
                        Floor: {collection?.floor_price
                          ? `${Number(ethers.formatEther(BigInt(collection.floor_price.toString()))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}`
                          : '0'
                        }
                      </span>
                    </div>
                  </div>
                </TableCell>
                <TableCell className="text-gray-400 font-normal text-sm sm:text-base hidden sm:table-cell">
                  {collection?.floor_price
                    ? `${Number(ethers.formatEther(BigInt(collection.floor_price.toString()))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}`
                    : '0'
                  }
                </TableCell>
                <TableCell className="text-gray-400 font-normal text-sm sm:text-base hidden sm:table-cell">
                  {collection?.period_volume
                    ? `${Number(ethers.formatEther(BigInt(collection.period_volume.toString()))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}`
                    : '0'
                  }
                </TableCell>
                <TableCell className="text-gray-400 font-normal text-sm sm:text-base hidden sm:table-cell">{collection?.owners || '0'}</TableCell>
                <TableCell className="text-gray-400 font-normal text-sm sm:text-base">{collection?.total_supply || '0'} / {collection?.max_supply || '0'}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
      {/* 移动端卡片列表 */}
      <div className="flex flex-col gap-2 sm:hidden">
        {collections?.map((collection, index) => (
          <div
            key={collection.address}
            className="rounded-lg border border-[#171a1f] bg-[#181a1f] p-4 shadow hover:bg-[#23272f] transition cursor-pointer"
            onClick={() => onRowClick(collection.address)}
          >
            <div className="flex flex-row items-center gap-3 mb-2">
              <div className="w-[40px] h-[40px] rounded-lg overflow-hidden">
                <Image
                  src={collectionsImages[index] || '/fluxus.svg'}
                  alt={collection?.name || ''}
                  width={40}
                  height={40}
                  className="w-full h-full object-cover"
                />
              </div>
              <div className="font-semibold text-white text-base mb-2">{collection?.name || ''}</div>
            </div>
            <div className="flex flex-wrap gap-y-1 text-xs text-gray-400">
              <div className="w-1/2">FLOOR PRICE: <span className="text-white">{collection?.floor_price ? `${Number(ethers.formatEther(BigInt(collection.floor_price.toString()))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}` : '0'}</span></div>
              <div className="w-1/2">VOLUME: <span className="text-white">{collection?.period_volume ? `${Number(ethers.formatEther(BigInt(collection.period_volume.toString()))).toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}` : '0'}</span></div>
              <div className="w-1/2">OWNERS: <span className="text-white">{collection?.owners || '0'}</span></div>
              <div className="w-1/2">MINTED: <span className="text-white">{collection?.total_supply || '0'} / {collection?.max_supply || '0'}</span></div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};