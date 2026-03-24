"use client"

import { Search, Menu, X } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { ConnectButton } from "@rainbow-me/rainbowkit"
import { useState, useEffect, useRef } from "react"
import Link from "next/link"
import { useRouter } from "next/router"
import { FaRegUserCircle } from "react-icons/fa"
import { useAccount } from "wagmi"
import { cn } from "@/lib/utils"
import { useSearch } from "@/hooks/useSearch"
import SearchResults from "./SearchResults"

export default function Header() {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false)
  const router = useRouter()
  const pathname = router.pathname
  const { isConnected } = useAccount()
  const searchRef = useRef<HTMLDivElement>(null)
  const { query, setQuery, results, loading, error, isOpen, clearSearch, closeSearch } = useSearch(84532)

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if ((event.metaKey || event.ctrlKey) && event.key === "k") {
        event.preventDefault()
        const searchInput = searchRef.current?.querySelector("input")
        if (searchInput) searchInput.focus()
      }
      if (event.key === "Escape" && isOpen) {
        closeSearch()
        clearSearch()
      }
    }

    document.addEventListener("keydown", handleKeyDown)
    return () => document.removeEventListener("keydown", handleKeyDown)
  }, [isOpen, closeSearch, clearSearch])

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (searchRef.current && !searchRef.current.contains(event.target as Node)) {
        closeSearch()
      }
    }

    if (isOpen) {
      document.addEventListener("mousedown", handleClickOutside)
      return () => document.removeEventListener("mousedown", handleClickOutside)
    }
  }, [isOpen, closeSearch])

  const routes = [
    { path: "/", label: "Market" },
    { path: "/collection/create", label: "Launch" },
    { path: "/docs/creation-guide", label: "Guide" },
    { path: "/litepaper", label: "Litepaper" },
  ]

  return (
    <header className="sticky top-0 z-[9999] h-16 border-b border-black/10 bg-[color:var(--bg-page)]/95 px-3 py-3 backdrop-blur-sm sm:px-6">
      <div className="flex items-center justify-between gap-3">
        <div className="flex min-w-0 items-center gap-3 sm:w-[220px]">
          <Button variant="ghost" size="icon" className="sm:hidden" onClick={() => setIsMobileMenuOpen((prev) => !prev)}>
            {isMobileMenuOpen ? <X className="h-5 w-5" /> : <Menu className="h-5 w-5" />}
          </Button>
          <Link href="/" className="flex items-center gap-3">
            <div className="font-serif text-[22px] font-semibold leading-none tracking-[-0.03em] text-[color:var(--text-primary)] sm:text-[24px]">
              Fluxus
            </div>
          </Link>
        </div>

        <div className="hidden flex-1 justify-center px-4 sm:flex">
          <div ref={searchRef} className="relative w-full max-w-[440px]">
            <Input
              type="text"
              placeholder="Search collections"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              className="w-full pl-10 pr-16"
            />
            <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-[color:var(--text-muted)]">
              <Search className="h-4 w-4" />
            </div>
            <div className="absolute inset-y-0 right-2 my-auto hidden items-center gap-1 border border-black/10 bg-[color:var(--bg-muted)] px-2 py-0.5 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)] sm:flex">
              <span>⌘</span>
              <span>K</span>
            </div>
            <SearchResults results={results} loading={loading} error={error} isOpen={isOpen} onClose={closeSearch} />
          </div>
        </div>

        <div className="flex min-w-0 items-center justify-end gap-2 sm:w-[220px]">
          <ConnectButton showBalance={false} label="Connect" />
          {isConnected && (
            <Link href="/user" className="text-[color:var(--text-muted)] transition-colors hover:text-[color:var(--text-primary)]">
              <FaRegUserCircle className="h-6 w-6" />
            </Link>
          )}
        </div>
      </div>

      {isMobileMenuOpen && (
        <div className="fixed inset-x-0 top-16 z-[9998] border-b border-black/10 bg-[color:var(--bg-page)] sm:hidden">
          <nav className="flex flex-col gap-2 p-3">
            {routes.map((route) => (
              <Link
                key={route.path}
                href={route.path}
                className={cn(
                  "border px-4 py-3 font-primary text-[11px] uppercase tracking-[0.18em] transition-colors",
                  pathname === route.path
                    ? "border-black/12 bg-[color:var(--fg-strong)] text-white"
                    : "border-black/10 bg-[color:var(--bg-surface)] text-[color:var(--text-secondary)]"
                )}
                onClick={() => setIsMobileMenuOpen(false)}
              >
                {route.label}
              </Link>
            ))}
          </nav>
        </div>
      )}
    </header>
  )
}
