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
      <div className="mb-6 border border-border bg-black/40 p-4 sm:p-6">
        <div className="flex flex-col gap-6 animate-pulse">
          <div className="flex items-center gap-4">
            <div className="w-16 h-16 border border-border bg-bg-card" />
            <div className="space-y-2 w-full">
              <div className="h-6 w-40 bg-bg-card" />
              <div className="h-3 w-24 bg-bg-card" />
            </div>
          </div>
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
            {[...Array(8)].map((_, idx) => (
              <div key={idx} className="space-y-2">
                <div className="h-3 bg-bg-card" />
                <div className="h-4 bg-bg-card" />
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
    <div className="mb-6 border border-border bg-gradient-to-b from-black via-bg-card to-bg-card/70 px-4 sm:px-6 py-5 text-primary">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-6">
        <div className="flex flex-col gap-4">
          <div className="flex items-center gap-4">
            <div className="w-16 h-16 border border-border overflow-hidden bg-bg-card">
            <Image
              src={collectionImage || '/fluxus.svg'}
              alt={collection?.name}
              className="object-cover w-full h-full"
              width={64}
              height={64}
            />
          </div>
            <div className="flex flex-col">
              <p className="text-[11px] uppercase tracking-[0.4em] text-secondary">Collection</p>
              <h1 className="text-2xl font-bold text-white">{collection.name}</h1>
              <div className="text-[11px] uppercase tracking-[0.3em] text-secondary mt-1">
                {chainSymbol} • {collection.max_supply} Supply
              </div>
            </div>
          </div>
          <div className="space-y-2">
            <div className="text-[10px] uppercase tracking-[0.3em] text-secondary">Contract</div>
            <div className="flex flex-col sm:flex-row sm:items-center gap-2">
              <span className="font-mono text-xs text-primary bg-black/40 border border-border px-3 py-1 inline-flex items-center gap-2">
                {sliceAddress(collection.address)}
                <button
                  onClick={() => {
                    navigator.clipboard.writeText(collection.address);
                    setShowToast(true);
                  }}
                  className="text-secondary hover:text-fluxus-primary transition-colors"
                  aria-label="Copy contract address"
                >
                  <FaCopy className="w-3 h-3" />
                </button>
              </span>
              <div className="flex items-center gap-2 text-secondary text-lg">
                <SocialIcon href={collection?.meta_data?.website} icon={<FaGlobe />} label="Website" />
                <SocialIcon href={collection?.meta_data?.twitter} icon={<FaTwitter />} label="Twitter" />
                <SocialIcon href={collection?.meta_data?.discord} icon={<FaDiscord />} label="Discord" />
                <button
                  onClick={() => setShowInfoDialog(true)}
                  className="w-8 h-8 border border-border text-secondary hover:text-fluxus-primary hover:border-fluxus-primary transition-colors flex items-center justify-center"
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
        <AlertDialogContent className="bg-black border-white/20">
          <AlertDialogHeader>
            <AlertDialogTitle className="text-white">COLLECTION INFO</AlertDialogTitle>
            <AlertDialogDescription className="whitespace-pre-wrap text-white/70">
              {collection?.meta_data?.description || "No description available"}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogAction 
              className="bg-white/10 text-white border border-white/20 hover:bg-white/20"
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
    <div className="border border-border/70 bg-black/30 px-3 py-3 flex flex-col gap-1">
      <p className="text-[10px] uppercase tracking-[0.3em] text-secondary">{label}</p>
      <p className="text-white font-semibold text-sm">{value}</p>
    </div>
  )
}

function SocialIcon({ href, icon, label }: { href?: string; icon: React.ReactNode; label: string }) {
  const content = (
    <span className="w-8 h-8 border border-border text-secondary hover:text-fluxus-primary hover:border-fluxus-primary transition-colors inline-flex items-center justify-center">
      {icon}
    </span>
  )

  if (!href) {
    return (
      <span className="w-8 h-8 border border-border/40 text-secondary/40 inline-flex items-center justify-center cursor-not-allowed">
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
