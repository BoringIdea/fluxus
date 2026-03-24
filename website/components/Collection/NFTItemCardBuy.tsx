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
          "absolute right-3 top-3 z-10 flex h-5 w-5 items-center justify-center border border-black/10 bg-[color:var(--bg-surface)] text-[color:var(--text-primary)] focus:outline-none focus:ring-2 focus:ring-[color:var(--color-primary)]/40",
          isSelected && "border-[color:var(--color-primary)] bg-[color:var(--color-primary)] text-white"
        )}
        onClick={(event) => {
          event.stopPropagation();
          handleSelect();
        }}
      >
        <CheckIcon className={cn("h-3.5 w-3.5", isSelected ? "text-white" : "text-[color:var(--color-primary)]")} />
      </button>
      <CardContent className="flex-grow border-b border-black/10 p-0">
        <div className="aspect-square w-full bg-[color:var(--bg-muted)]">
          <img
            src={token.image.originalUrl}
            alt={`${token.name} #${token.tokenId}`}
            className="h-full w-full object-cover"
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
            onBuy(token.tokenId);
          }}
          className="w-full text-[10px] tracking-[0.18em]"
        >
          Buy NFT
        </Button>
      </CardFooter>
    </Card>
  );
}
