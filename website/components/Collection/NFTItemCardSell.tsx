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
        "group relative flex flex-col gap-0 cursor-pointer transition-colors duration-150 hover:bg-bg-card-hover",
        isSelected && "outline outline-2 outline-fluxus-primary/80"
      )}
      ref={ref}
      onClick={handleCardClick}
    >
      <button
        type="button"
        aria-pressed={isSelected}
        className={cn(
          "absolute top-3 right-3 z-10 h-5 w-5 border border-fluxus-primary text-fluxus-primary flex items-center justify-center rounded-none bg-background focus:outline-none focus:ring-2 focus:ring-fluxus-primary/60",
          isSelected && "bg-fluxus-primary text-black"
        )}
        onClick={(event) => {
          event.stopPropagation();
          handleSelect();
        }}
      >
        <CheckIcon className="h-3.5 w-3.5" />
      </button>

      <CardContent className="flex-grow p-0 border-b border-border">
        <div className="aspect-square w-full bg-bg-tertiary">
          <img
            src={img}
            alt={`${token.name} #${token.tokenId}`}
            className="h-full w-full object-cover"
            loading="lazy"
          />
        </div>
      </CardContent>
      <CardFooter className="flex flex-col gap-2 px-4 py-3">
        <CardTitle className="text-xs font-bold text-secondary uppercase tracking-[0.2em] flex items-center justify-between">
          <span>{collectionSymbol ?? 'NFT'}</span>
          <span className="text-primary tracking-tight font-black">#{token.tokenId}</span>
        </CardTitle>
        <Button
          disabled={isPending}
          onClick={(event) => {
            event.stopPropagation();
            onSell(token.tokenId);
          }}
          className="w-full text-xs tracking-[0.2em]"
        >
          {buttonLabel}
        </Button>
      </CardFooter>
    </Card>
  );
}
