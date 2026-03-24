'use client'
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { useEffect, useState } from "react"
import { CheckIcon } from "@radix-ui/react-icons"
import { useRouter } from "next/navigation";
import { cn } from "@/lib/utils"

interface NFTItemCardSellProps {
  token: {
    tokenId: number;
    name?: string;
    tokenUri?: string;
    image?: {
      originalUrl: string;
    };
  };
  collectionSymbol?: string;
  contractAddress?: string;
  isPending: boolean;
  onSell: (tokenId: number) => void;
  onSelect: (tokenId: number, selected: boolean) => void;
  ref?: React.Ref<HTMLDivElement>;
  selected: boolean;
  isApproved?: boolean;
  pendingApprovalTokenId?: number | null;
  isImageLoaded?: boolean;
}

export const NFTItemCardSell = ({ token, collectionSymbol, contractAddress, isPending, onSell, onSelect, ref, selected: isSelected, isApproved, pendingApprovalTokenId, isImageLoaded }: NFTItemCardSellProps) => {
  const router = useRouter();
  const [img, setImg] = useState('/fluxus-logo.png')
  
  const handleSelect = () => {
    onSelect?.(token.tokenId, !isSelected)
  }

  const handleCardClick = (e: React.MouseEvent) => {
    // Don't navigate if clicking on checkbox or sell button
    if ((e.target as HTMLElement).closest('button')) {
      return;
    }
    
    if (contractAddress) {
      router.push(`/nft/${contractAddress}/${token.tokenId}`);
    }
  }

  useEffect(() => {
    const fetchImage = async () => {
      // If there is a preloaded image, use it
      if (token.image?.originalUrl) {
        setImg(token.image.originalUrl);
        return;
      }

      // If there is no preloaded image, but there is a tokenUri, fetch the image
      if (token.tokenUri && !isImageLoaded) {
        try {
          const response = await fetch(token.tokenUri);
          const data = await response.json();
          setImg(data.image);
        } catch (error) {
          console.error("Error fetching image:", error);
        }
      }
    };
    fetchImage();
  }, [token, isImageLoaded])

  const buttonLabel = isPending
    ? pendingApprovalTokenId === token.tokenId
      ? 'Approving...'
      : 'Selling...'
    : isApproved
      ? 'Sell NFT'
      : 'Approve First';

  return (
    <Card
      className={cn(
        "group relative flex cursor-pointer flex-col gap-0 border-black/10 bg-[color:var(--bg-surface)] transition-colors duration-150 hover:bg-[color:var(--bg-card-hover)]",
        isSelected && "outline outline-2 outline-[color:var(--color-primary)]/70"
      )}
      ref={ref}
      onClick={handleCardClick}
    >
      <button
        type="button"
        aria-pressed={isSelected}
        className={cn(
          "absolute right-3 top-3 z-10 flex h-5 w-5 items-center justify-center border border-[color:var(--color-primary)] text-[color:var(--color-primary)] bg-[color:var(--bg-surface)] focus:outline-none focus:ring-2 focus:ring-[color:var(--color-primary)]/40",
          isSelected && "bg-[color:var(--color-primary)] text-white"
        )}
        onClick={(event) => {
          event.stopPropagation();
          handleSelect();
        }}
      >
        <CheckIcon className="h-3.5 w-3.5" />
      </button>

      <CardContent className="flex-grow border-b border-black/10 p-0">
        <div className="aspect-square w-full bg-[color:var(--bg-muted)]">
          <img
            src={img}
            alt={`${token.name} #${token.tokenId}`}
            className="h-full w-full object-cover"
            loading="lazy"
          />
        </div>
      </CardContent>
      <CardFooter className="flex flex-col gap-2 px-4 py-3">
        <CardTitle className="flex items-center justify-between font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
          <span>{collectionSymbol ?? 'NFT'}</span>
          <span className="font-heading text-[18px] leading-none text-[color:var(--text-primary)]">#{token.tokenId}</span>
        </CardTitle>
        <Button
          disabled={isPending}
          onClick={(event) => {
            event.stopPropagation();
            onSell(token.tokenId);
          }}
          className="w-full text-[10px] tracking-[0.18em]"
        >
          {buttonLabel}
        </Button>
      </CardFooter>
    </Card>
  );
}
