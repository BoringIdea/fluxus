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
    <div className="min-h-screen w-full text-primary">
      <div className="mx-auto max-w-6xl space-y-6 px-4 py-8 sm:px-8">
        <button
          onClick={() => router.back()}
          className="inline-flex items-center gap-2 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)] hover:text-[color:var(--text-primary)]"
        >
          <span className="text-lg">&lt;</span>
          Back
        </button>

        <div className="flex flex-col gap-6 border border-black/10 bg-[color:var(--bg-surface)] p-4 sm:flex-row sm:p-6">
          <div className="space-y-4 lg:w-1/2">
            <div className="flex aspect-square items-center justify-center overflow-hidden border border-black/10 bg-[color:var(--bg-muted)]">
              {imageUrl ? (
                // eslint-disable-next-line @next/next/no-img-element
                <img src={imageUrl} alt={tokenLabel} className="w-full h-full object-cover" />
              ) : (
                <div className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                  No Image
                </div>
              )}
            </div>
            <div className="space-y-3 border border-black/10 bg-[color:var(--bg-muted)] p-4">
              <div>
                <p className="mb-2 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Token</p>
                <h1 className="font-heading text-[34px] leading-none text-[color:var(--text-primary)]">{tokenLabel}</h1>
                <p className="mt-2 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                  {collectionName} • ID #{tokenId}
                </p>
              </div>
              <div className="flex flex-wrap items-baseline gap-2">
                <p className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                  Floor Price
                </p>
                <span className="font-heading text-[24px] leading-none text-[color:var(--text-primary)]">{floorText}</span>
              </div>
              <Button
                onClick={handleBuyNow}
                disabled={!buyPrice || isPending}
                className="w-full"
              >
                {isPending ? 'Processing...' : 'Buy Now'}
              </Button>
            </div>
          </div>

          <div className="flex flex-col gap-4 lg:w-1/2">
            <div className="border border-black/10 bg-[color:var(--bg-muted)] p-4">
              <p className="flux-kicker mb-2">Traits</p>
              <NFTTraits traits={traits as any[]} />
            </div>
            <div className="flex-1 border border-black/10 bg-[color:var(--bg-muted)] p-4">
              <p className="flux-kicker mb-2">Activity</p>
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
