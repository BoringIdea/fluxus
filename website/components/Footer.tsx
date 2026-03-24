import Link from "next/link"
import { Github } from "lucide-react"

export default function Footer() {
  return (
    <footer className="border-t border-black/10 bg-[color:var(--bg-page)] text-[color:var(--text-muted)]">
      <div className="mx-auto flex max-w-6xl items-center justify-between px-4 py-4 sm:px-6">
        <div className="font-primary text-[10px] uppercase tracking-[0.18em]">Fluxus NFT liquidity protocol</div>
        <div className="flex items-center gap-3">
          <Link href="https://github.com/HashIdea/fluxus" className="flex h-8 w-8 items-center justify-center border border-black/10 bg-[color:var(--bg-surface)] transition-colors hover:border-[color:var(--color-primary)]/20 hover:bg-[color:var(--color-primary)]/8 hover:text-[color:var(--color-primary)]">
            <Github className="h-4 w-4" />
          </Link>
          <Link href="https://x.com/BoringIdea" className="flex h-8 w-8 items-center justify-center border border-black/10 bg-[color:var(--bg-surface)] transition-colors hover:border-[color:var(--color-primary)]/20 hover:bg-[color:var(--color-primary)]/8 hover:text-[color:var(--color-primary)]">
            <img className="w-4" src="/x.svg" alt="x" />
          </Link>
          <Link href="https://fluxus-docs.vercel.app/" className="flex h-8 w-8 items-center justify-center border border-black/10 bg-[color:var(--bg-surface)] transition-colors hover:border-[color:var(--color-primary)]/20 hover:bg-[color:var(--color-primary)]/8 hover:text-[color:var(--color-primary)]">
            <img className="w-4" src="/docs.svg" alt="docs" />
          </Link>
        </div>
      </div>
    </footer>
  )
}
