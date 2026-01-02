'use client'
import { Card, CardContent, CardFooter, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { CheckIcon } from "@radix-ui/react-icons"
import { useRouter } from "next/navigation";
import { cn } from "@/lib/utils";

interface NFTItemCardProps {
  token: {
    tokenId: number;
    name?: string;
    image: {
      originalUrl: string;
    };
  };
  collectionSymbol?: string;
  contractAddress?: string;
  isPending: boolean;
  onBuy: (tokenId: number) => void;
  onSelect: (tokenId: number, selected: boolean) => void;
  ref?: React.Ref<HTMLDivElement>;
  selected: boolean;
}

export const NFTItemCardBuy = ({ token, collectionSymbol, contractAddress, isPending, onBuy, onSelect, ref, selected: isSelected }: NFTItemCardProps) => {
  const router = useRouter();

  const handleSelect = () => {
    onSelect?.(token.tokenId, !isSelected)
  }

  const handleCardClick = (e: React.MouseEvent) => {
    // Don't navigate if clicking on checkbox or buy button
    if ((e.target as HTMLElement).closest('[data-radix-checkbox-root]') || 
        (e.target as HTMLElement).closest('button')) {
      return;
    }
    
    if (contractAddress) {
      router.push(`/nft/${contractAddress}/${token.tokenId}`);
    }
  }

  return (
    <Card
      className={cn(
        "group relative flex flex-col gap-0 cursor-pointer transition-all duration-150 hover:bg-bg-card-hover",
        isSelected && "outline outline-2 outline-fluxus-primary/80"
      )}
      ref={ref}
      onClick={handleCardClick}
    >
      <button
        type="button"
        aria-pressed={isSelected}
        className={cn(
          "absolute top-3 right-3 z-10 h-5 w-5 border border-border bg-background text-primary flex items-center justify-center rounded-none focus:outline-none focus:ring-2 focus:ring-fluxus-primary/60",
          isSelected && "bg-fluxus-primary border-fluxus-primary text-black"
        )}
        onClick={(event) => {
          event.stopPropagation();
          handleSelect();
        }}
      >
        <CheckIcon className={cn("h-3.5 w-3.5", isSelected ? "text-black" : "text-fluxus-primary")} />
      </button>
      <CardContent className="flex-grow p-0 border-b border-border">
        <div className="aspect-square w-full bg-bg-tertiary">
          <img
            src={token.image.originalUrl}
            alt={`${token.name} #${token.tokenId}`}
            className="h-full w-full object-cover"
          />
        </div>
      </CardContent>
      <CardFooter className="flex flex-col gap-2 px-4 py-3">
        <CardTitle className="text-xs font-bold text-secondary uppercase tracking-[0.2em] flex items-center justify-between">
          <span>{collectionSymbol ?? 'NFT'}</span>
          <span className="text-primary font-black tracking-tight">#{token.tokenId}</span>
        </CardTitle>
        <Button
          disabled={isPending}
          onClick={(event) => {
            event.stopPropagation();
            onBuy(token.tokenId);
          }}
          className="w-full text-xs tracking-[0.2em]"
        >
          Buy NFT
        </Button>
      </CardFooter>
    </Card>
  );
}
