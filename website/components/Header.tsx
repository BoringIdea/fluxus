"use client"

import { Search, Menu } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { ConnectButton } from "@rainbow-me/rainbowkit"
import { useState, useEffect, useRef } from "react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { FaRegUserCircle } from "react-icons/fa";
import { useAccount } from "wagmi"
import { cn } from "@/lib/utils"
import { useSearch } from "@/hooks/useSearch"
import SearchResults from "./SearchResults"

const inactiveStroke = "#4B4B4B";

export default function Header() {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const pathname = usePathname();
  const { isConnected } = useAccount();
  
  // Search functionality
  const searchRef = useRef<HTMLDivElement>(null);
  const {
    query,
    setQuery,
    results,
    loading,
    error,
    isOpen,
    clearSearch,
    closeSearch,
  } = useSearch(84532); // Default to Base Sepolia

  // Handle keyboard shortcuts
  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      // Cmd/Ctrl + K to focus search
      if ((event.metaKey || event.ctrlKey) && event.key === 'k') {
        event.preventDefault();
        const searchInput = searchRef.current?.querySelector('input');
        if (searchInput) {
          searchInput.focus();
        }
      }
      
      // Escape to close search
      if (event.key === 'Escape' && isOpen) {
        closeSearch();
        clearSearch();
      }
    };

    document.addEventListener('keydown', handleKeyDown);
    return () => document.removeEventListener('keydown', handleKeyDown);
  }, [isOpen, closeSearch, clearSearch]);

  // Close search when clicking outside
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (searchRef.current && !searchRef.current.contains(event.target as Node)) {
        closeSearch();
      }
    };

    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
      return () => document.removeEventListener('mousedown', handleClickOutside);
    }
  }, [isOpen, closeSearch]);

  // Mobile menu icons - using the same icons as SideBar
  const MobileHomeIcon = ({ isActive }: { isActive: boolean }) => (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path 
        d="M3 9L12 2L21 9V20C21 20.5304 20.7893 21.0391 20.4142 21.4142C20.0391 21.7893 19.5304 22 19 22H5C4.46957 22 3.96086 21.7893 3.58579 21.4142C3.21071 21.0391 3 20.5304 3 20V9Z" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
        fill={isActive ? "#16A34A" : 'none'}
        fillOpacity={isActive ? 0.1 : 0}
      />
      <path 
        d="M9 22V12H15V22" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
    </svg>
  );

  const MobileLaunchIcon = ({ isActive }: { isActive: boolean }) => (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path 
        d="M12 2L2 7L12 12L22 7L12 2Z" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
        fill={isActive ? "#16A34A" : 'none'}
        fillOpacity={isActive ? 0.1 : 0}
      />
      <path 
        d="M2 17L12 22L22 17" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
      <path 
        d="M2 12L12 17L22 12" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
    </svg>
  );

  const MobileGuideIcon = ({ isActive }: { isActive: boolean }) => (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path 
        d="M9 12L11 14L15 10" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
      <path 
        d="M21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12Z" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
        fill={isActive ? "#16A34A" : 'none'}
        fillOpacity={isActive ? 0.1 : 0}
      />
    </svg>
  );

  const MobileDocIcon = ({ isActive }: { isActive: boolean }) => (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path 
        d="M14 2H6C5.46957 2 4.96086 2.21071 4.58579 2.58579C4.21071 2.96086 4 3.46957 4 4V20C4 20.5304 4.21071 21.0391 4.58579 21.4142C4.96086 21.7893 5.46957 22 6 22H18C18.5304 22 19.0391 21.7893 19.4142 21.4142C19.7893 21.0391 20 20.5304 20 20V8L14 2Z" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
        fill={isActive ? "#16A34A" : 'none'}
        fillOpacity={isActive ? 0.1 : 0}
      />
      <path 
        d="M14 2V8H20" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
      <path 
        d="M16 13H8" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
      <path 
        d="M16 17H8" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
      <path 
        d="M10 9H8" 
        stroke={isActive ? "#16A34A" : inactiveStroke} 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
    </svg>
  );

  // SideBar 菜单项
  const routes = [
    { path: "/", label: "Home", icon: <MobileHomeIcon isActive={pathname === "/"} /> },
    { path: "/collection/create", label: "Launch", icon: <MobileLaunchIcon isActive={pathname === "/collection/create"} /> },
    { path: "/docs/creation-guide", label: "Guide", icon: <MobileGuideIcon isActive={pathname === "/docs/creation-guide"} /> },
    { path: "/litepaper", label: "Docs", icon: <MobileDocIcon isActive={pathname === "/litepaper"} /> },
  ];

  return (
    <header className="flex items-center justify-between border-b border-primary bg-background/95 px-2 sm:px-6 py-2 sticky top-0 z-[9999] w-full h-14 sm:h-16">
      {/* Left section - Logo/Menu */}
      <div className="w-[60px] sm:w-[200px] flex items-center min-w-0">
        <Button
          variant="ghost"
          size="icon"
          className="sm:hidden"
          onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
        >
          <Menu className="w-5 h-5" />
        </Button>
      </div>

      {/* Center section with search */}
      <div className="hidden sm:flex items-center flex-1 justify-center px-1 sm:px-4 min-w-0">
        <div ref={searchRef} className="relative w-full max-w-[240px] sm:max-w-[420px]">
          <Input
            type="text"
            placeholder="Search collections..."
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            className="w-full pl-10 pr-16 text-xs sm:text-sm bg-bg-tertiary text-primary placeholder:text-secondary/70"
          />
          <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
            <Search className="w-4 h-4 text-secondary" />
          </div>
          <div className="hidden sm:flex items-center gap-1 absolute inset-y-0 right-2 my-auto text-[10px] text-secondary border border-primary px-2 py-0.5 uppercase tracking-[0.2em] bg-background">
            <span className="font-bold">⌘</span>
            <span className="font-bold">K</span>
          </div>
          
          {/* Search Results */}
          <SearchResults
            results={results}
            loading={loading}
            error={error}
            isOpen={isOpen}
            onClose={closeSearch}
          />
        </div>
      </div>

      {/* Right section with wallet button */}
      <div className="sm:w-[200px] flex justify-end min-w-0 items-center gap-2">
        <ConnectButton showBalance={false} label="Connect" />
        {isConnected && (
          <Link href={`/user`}>
            <FaRegUserCircle className="w-6 h-6 text-secondary hover:text-fluxus-primary transition-colors cursor-pointer" />
          </Link>
        )}
      </div>

      {/* Mobile Menu */}
      {isMobileMenuOpen && (
        <div className="fixed top-0 left-0 w-screen h-screen bg-background sm:hidden z-[9999] overflow-y-auto border-t border-primary">
          <nav className="pt-12 p-4 flex flex-col gap-3">
            {routes.map(route => (
              <Link
                key={route.path}
                href={route.path}
                className={cn(
                  "flex items-center gap-3 px-4 py-3 border border-border text-primary font-bold uppercase tracking-[0.12em] hover:bg-bg-card-hover transition-colors",
                  pathname === route.path && "bg-bg-card text-fluxus-primary border-fluxus-primary"
                )}
                onClick={() => setIsMobileMenuOpen(false)}
              >
                {route.icon}
                <span>{route.label}</span>
              </Link>
            ))}
          </nav>
        </div>
      )}
    </header>
  )
}
