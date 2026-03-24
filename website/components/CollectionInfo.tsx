import { Collection } from "@/src/api"
import { getChainSymbol, sliceAddress } from "@/src/utils"
import { formatEther } from "ethers"
import { useChainId } from "wagmi"
import Image from "next/image"
import { FaGlobe, FaTwitter, FaDiscord, FaCopy, FaInfoCircle } from 'react-icons/fa'
import { useEffect, useState } from "react"
import { Toast } from "@/components/ui/toast"
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogAction,
} from "@/components/ui/alert-dialog"
import { fetchCollectionImage } from "@/lib/utils"

export default function CollectionInfo({ collection }: { collection?: Collection }) {

  const chainId = useChainId();
  const [showToast, setShowToast] = useState(false)
  const [showInfoDialog, setShowInfoDialog] = useState(false)
  const [collectionImage, setCollectionImage] = useState<string | null>(null);
  useEffect(() => {
    if (collection) {
      fetchCollectionImage(collection).then((image) => {
        setCollectionImage(image);
      }).catch((error) => {
        console.error('Error fetching collection image:', error);
        setCollectionImage(null);
      });
    }
  }, [collection]);

  const calculateListedPercentage = () => {
    if (!collection || collection.current_supply === 0) return "0%"
    const percentage = (1 - (collection.current_supply / collection.total_supply)) * 100
    return `${percentage.toFixed(0)}%`
  }

  const calculateListedCount = () => {
    if (!collection || collection.current_supply === 0) return "0"
    return (collection.total_supply - collection.current_supply).toString()
  }

  // If collection data is not loaded, show loading state
  if (!collection) {
    return (
      <div className="mb-6 border border-black/10 bg-[color:var(--bg-surface)] p-4 sm:p-6">
        <div className="flex flex-col gap-6 animate-pulse">
          <div className="flex items-center gap-4">
            <div className="h-16 w-16 border border-black/10 bg-[color:var(--bg-muted)]" />
            <div className="space-y-2 w-full">
              <div className="h-6 w-40 bg-[color:var(--bg-muted)]" />
              <div className="h-3 w-24 bg-[color:var(--bg-muted)]" />
            </div>
          </div>
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
            {[...Array(8)].map((_, idx) => (
              <div key={idx} className="space-y-2">
                <div className="h-3 bg-[color:var(--bg-muted)]" />
                <div className="h-4 bg-[color:var(--bg-muted)]" />
              </div>
            ))}
          </div>
        </div>
      </div>
    )
  }

  const chainSymbol = getChainSymbol(chainId)
  const stats = [
    { label: "Floor Price", value: `${Number(formatEther(collection?.floor_price?.toString() || '0')).toFixed(4)} ${chainSymbol}` },
    { label: "Creator Fee", value: `${(Number(formatEther(collection?.creator_fee?.toString() || '0')) * 100).toFixed(2)}%` },
    { label: "1D Volume", value: `${Number(formatEther(BigInt(collection?.volume_1d?.toString() || '0'))).toFixed(2)} ${chainSymbol}` },
    { label: "All Volume", value: `${Number(formatEther(BigInt(collection?.total_volume?.toString() || '0'))).toFixed(2)} ${chainSymbol}` },
    { label: "1D Sales", value: `${collection?.sales_1d || 0}` },
    { label: "Owners", value: `${collection.owners}` },
    { label: "Listed", value: `${calculateListedCount()} (${calculateListedPercentage()})` },
    { label: "Minted", value: `${collection.total_supply} / ${collection.max_supply}` },
  ]

  return (
    <div className="mb-6 border border-black/10 bg-[color:var(--bg-surface)] px-4 py-5 text-primary sm:px-6">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-6">
        <div className="flex flex-col gap-4">
          <div className="flex items-center gap-4">
            <div className="h-16 w-16 overflow-hidden border border-black/10 bg-[color:var(--bg-muted)]">
            <Image
              src={collectionImage || '/fluxus.svg'}
              alt={collection?.name}
              className="object-cover w-full h-full"
              width={64}
              height={64}
            />
          </div>
            <div className="flex flex-col">
              <p className="mb-2 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Collection</p>
              <h1 className="font-heading text-[34px] leading-none text-[color:var(--text-primary)]">{collection.name}</h1>
              <div className="mt-2 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
                {chainSymbol} • {collection.max_supply} Supply
              </div>
            </div>
          </div>
          <div className="space-y-2">
            <div className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Contract</div>
            <div className="flex flex-col sm:flex-row sm:items-center gap-2">
              <span className="inline-flex items-center gap-2 border border-black/10 bg-[color:var(--bg-muted)] px-3 py-2 font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-primary)]">
                {sliceAddress(collection.address)}
                <button
                  onClick={() => {
                    navigator.clipboard.writeText(collection.address);
                    setShowToast(true);
                  }}
                  className="text-[color:var(--text-muted)] transition-colors hover:text-[color:var(--color-primary)]"
                  aria-label="Copy contract address"
                >
                  <FaCopy className="w-3 h-3" />
                </button>
              </span>
              <div className="flex items-center gap-2 text-lg text-[color:var(--text-muted)]">
                <SocialIcon href={collection?.meta_data?.website} icon={<FaGlobe />} label="Website" />
                <SocialIcon href={collection?.meta_data?.twitter} icon={<FaTwitter />} label="Twitter" />
                <SocialIcon href={collection?.meta_data?.discord} icon={<FaDiscord />} label="Discord" />
                <button
                  onClick={() => setShowInfoDialog(true)}
                  className="flex h-8 w-8 items-center justify-center border border-black/10 text-[color:var(--text-muted)] transition-colors hover:border-[color:var(--color-primary)] hover:text-[color:var(--color-primary)]"
                  aria-label="Collection description"
                >
                  <FaInfoCircle />
                </button>
              </div>
            </div>
          </div>
        </div>

        <div className="grid w-full grid-cols-2 sm:grid-cols-4 gap-4">
          {stats.map((stat) => (
            <StatItem key={stat.label} label={stat.label} value={stat.value} />
          ))}
        </div>
      </div>

      <Toast
        message="Address copied"
        isVisible={showToast}
        onClose={() => setShowToast(false)}
      />

      <AlertDialog open={showInfoDialog} onOpenChange={setShowInfoDialog}>
        <AlertDialogContent className="border-black/10 bg-[color:var(--bg-surface)]">
          <AlertDialogHeader>
            <AlertDialogTitle className="font-primary text-[11px] uppercase tracking-[0.18em] text-[color:var(--text-primary)]">Collection Info</AlertDialogTitle>
            <AlertDialogDescription className="whitespace-pre-wrap text-[color:var(--text-secondary)]">
              {collection?.meta_data?.description || "No description available"}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogAction 
              className="border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-primary)] hover:bg-[color:var(--bg-card-hover)]"
              onClick={() => setShowInfoDialog(false)}
            >
              Close
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}

interface StatItemProps {
  label: string
  value: string
}

function StatItem({ label, value }: StatItemProps) {
  return (
    <div className="flex flex-col gap-1 border border-black/10 bg-[color:var(--bg-muted)] px-3 py-3">
      <p className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">{label}</p>
      <p className="font-heading text-[18px] leading-none text-[color:var(--text-primary)]">{value}</p>
    </div>
  )
}

function SocialIcon({ href, icon, label }: { href?: string; icon: React.ReactNode; label: string }) {
  const content = (
    <span className="inline-flex h-8 w-8 items-center justify-center border border-black/10 text-[color:var(--text-muted)] transition-colors hover:border-[color:var(--color-primary)] hover:text-[color:var(--color-primary)]">
      {icon}
    </span>
  )

  if (!href) {
    return (
      <span className="inline-flex h-8 w-8 cursor-not-allowed items-center justify-center border border-black/10 text-[color:var(--text-muted)]/40">
        {icon}
      </span>
    )
  }

  return (
    <a href={href} target="_blank" rel="noopener noreferrer" aria-label={label}>
      {content}
    </a>
  )
}
