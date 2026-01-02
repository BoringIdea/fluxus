import { useEffect, useMemo, useState } from 'react';
import { useParams } from 'next/navigation';
import { useRouter } from 'next/navigation';
import { useChainId, useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi';
import { useCollection } from '@/src/api/hooks';
import { getNFTMetadataByTokenId } from '@/src/api/alchemy';
import NFTTraits from '@/components/NFT/NFTTraits';
import NFTActivity from '@/components/NFT/NFTActivity';
import { PRICE_ABI, getTradeContractAddress } from '@/src/contract';
import { buildBuyTx } from '@/src/onchain/tradeTx';
import { SupportedChainId } from '@/src/alchemy';
import TransactionDialog from '@/components/shared/TransactionDialog';
import { getChainSymbol, sliceAddress } from '@/src/utils';
import { Button } from '@/components/ui/button';

export default function NFTDetailPage() {
  const params = useParams();
  const router = useRouter();
  const address = params?.address as string;
  const tokenIdParam = params?.tokenId as string;
  const tokenId = Number(tokenIdParam);
  const chainId = useChainId();

  const { data: collectionResp } = useCollection(chainId, address);
  const collection = collectionResp?.data;

  const [imageUrl, setImageUrl] = useState<string>('');
  const [traits, setTraits] = useState<any[]>([]);
  const [name, setName] = useState<string>('');

  // load metadata
  useEffect(() => {
    let ignore = false;
    const run = async () => {
      try {
        const meta = await getNFTMetadataByTokenId(chainId as SupportedChainId, address, tokenId);
        if (ignore) return;
        const img = meta?.image?.originalUrl || meta?.image?.pngUrl || '';
        if (img) setImageUrl(img);
        setTraits((meta?.raw?.metadata as any)?.attributes || []);
        setName(meta?.name || `${collection?.symbol || 'Token'} #${tokenId}`);
      } catch {
        // fallback: keep empty traits
      }
    };
    if (address && tokenIdParam) run();
    return () => { ignore = true; };
  }, [address, tokenIdParam, chainId, collection?.symbol]);

  // floor/buy price
  const { data: buyPrice } = useReadContract({
    address: collection?.price_contract as any,
    abi: PRICE_ABI,
    functionName: 'getBuyPriceAfterFee',
    args: [collection?.address],
    query: { enabled: !!collection?.price_contract && !!collection?.address }
  });

  const { data: hash, writeContract, isPending, isError: isWriteContractError, error: writeContractError } = useWriteContract();
  const { isLoading: isConfirming, isSuccess: isConfirmed } = useWaitForTransactionReceipt({ hash });

  const handleBuyNow = async () => {
    if (!buyPrice) return;
    await writeContract(
      buildBuyTx({
        tradeContractAddress: getTradeContractAddress(chainId as SupportedChainId),
        collectionAddress: address,
        tokenId,
        buyPriceWei: BigInt(buyPrice.toString()),
      }) as any
    );
  };

  const floorText = useMemo(() => {
    if (!buyPrice) return '0';
    const n = Number(buyPrice) / 1e18;
    return `${n.toFixed(4).replace(/\.?0+$/, '')} ${getChainSymbol(chainId)}`;
  }, [buyPrice, chainId]);

  const tokenLabel = name || `${collection?.symbol || 'Token'} #${tokenId}`;
  const collectionName = collection?.name || 'Collection';

  return (
    <div className="w-full min-h-screen bg-black text-primary">
      <div className="max-w-6xl mx-auto px-4 sm:px-8 py-8 space-y-6">
        <button
          onClick={() => router.back()}
          className="inline-flex items-center gap-2 text-[11px] uppercase tracking-[0.35em] text-secondary hover:text-white"
        >
          <span className="text-lg">&lt;</span>
          Back
        </button>

        <div className="border border-border bg-bg-card/40 p-4 sm:p-6 flex flex-col lg:flex-row gap-6">
          <div className="lg:w-1/2 space-y-4">
            <div className="border border-border bg-black/40 aspect-square flex items-center justify-center overflow-hidden">
              {imageUrl ? (
                // eslint-disable-next-line @next/next/no-img-element
                <img src={imageUrl} alt={tokenLabel} className="w-full h-full object-cover" />
              ) : (
                <div className="text-secondary uppercase tracking-[0.3em] text-xs">
                  No Image
                </div>
              )}
            </div>
            <div className="border border-border bg-black/40 p-4 space-y-3">
              <div>
                <p className="text-[11px] uppercase tracking-[0.35em] text-secondary">Token</p>
                <h1 className="text-2xl sm:text-3xl font-bold text-white">{tokenLabel}</h1>
                <p className="text-xs text-secondary tracking-[0.25em] uppercase">
                  {collectionName} • ID #{tokenId}
                </p>
              </div>
              <div className="flex flex-wrap items-baseline gap-2">
                <p className="text-[11px] uppercase tracking-[0.3em] text-secondary">
                  Floor Price
                </p>
                <span className="text-xl font-semibold text-white">{floorText}</span>
              </div>
              <Button
                onClick={handleBuyNow}
                disabled={!buyPrice || isPending}
                className="w-full border border-fluxus-primary bg-fluxus-primary text-black hover:bg-[#1FB455]"
              >
                {isPending ? 'Processing...' : 'Buy Now'}
              </Button>
            </div>
          </div>

          <div className="lg:w-1/2 flex flex-col gap-4">
            <div className="border border-border bg-black/40 p-4">
              <p className="text-[11px] uppercase tracking-[0.35em] text-secondary mb-2">Traits</p>
              <NFTTraits traits={traits as any[]} />
            </div>
            <div className="border border-border bg-black/40 p-4 flex-1">
              <p className="text-[11px] uppercase tracking-[0.35em] text-secondary mb-2">Activity</p>
              <NFTActivity collectionAddress={address} tokenId={tokenId} />
            </div>
          </div>
        </div>
      </div>

      <TransactionDialog
        isOpen={isPending}
        onOpenChange={() => {}}
        status={isPending ? 'pending' : isConfirmed ? 'success' : 'idle'}
        hash={hash}
        error={isWriteContractError ? writeContractError?.toString() : undefined}
        title="Buy NFT success"
        chainId={chainId}
      />
    </div>
  );
}
