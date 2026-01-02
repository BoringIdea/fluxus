"use client"

import type React from "react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { ReactNode, useState } from "react"
import { cn } from "@/lib/utils"
import { PrimaryColor } from "@/src/utils"

const inactiveStroke = "#4B4B4B";

interface Route {
  path: string
  icon: React.ElementType
  label: string
}



export default function Sidebar() {

  const routes: Route[] = [
    {
      path: "/",
      icon: HomeIcon,
      label: "Home",
    },
    {
      path: "/collection/create",
      icon: LaunchIcon,
      label: "Launch",
    },
    {
      path: "/docs/creation-guide",
      icon: GuideIcon,
      label: "Guide",
    },
    {
      path: "/litepaper",
      icon: DocIcon,
      label: "Docs",
    },
  ]
  const pathname = usePathname()
  const [hoveredPath, setHoveredPath] = useState<string | null>(null)

  return (
    <div className="hidden sm:flex w-[80px] bg-background flex-col items-center border-r border-primary fixed left-0 top-0 bottom-0 z-20">
      <div className="flex flex-col items-center w-full border-b border-primary py-6 gap-3">
        <div className="w-9 h-9 border border-primary bg-fluxus-primary" />
        <div className="text-fluxus-primary text-[11px] font-bold uppercase tracking-[0.4em]">
          Beta
        </div>
      </div>

      <div className="flex flex-col gap-6 items-center mt-10">
        {routes.map((route) => {
          const isActive = pathname === route.path
          const isHovered = hoveredPath === route.path

          return (
            <div key={route.path} className="relative group">
              <Link
                href={route.path}
                className="block"
                onMouseEnter={() => setHoveredPath(route.path)}
                onMouseLeave={() => setHoveredPath(null)}
              >
                <div
                  className={cn(
                    "w-12 h-12 border border-border flex items-center justify-center transition-colors duration-150 rounded-none",
                    isActive ? "bg-bg-card outline outline-2 outline-fluxus-primary/80" : isHovered ? "bg-bg-card-hover" : "bg-background",
                  )}
                >
                  <route.icon isActive={isActive} />
                </div>
              </Link>

              {/* Tooltip */}
              <div className="absolute left-full ml-2 px-2 py-1 bg-background border border-border text-[11px] font-semibold uppercase tracking-[0.2em] whitespace-nowrap opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 text-secondary">
                {route.label}
              </div>
            </div>
          )
        })}
      </div>
    </div>
  )
}

const HomeIcon = ({ isActive }: { icon: ReactNode; isActive: boolean }) => {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke={isActive ? PrimaryColor : inactiveStroke} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M4 12L12 4L20 12" />
      <path d="M6 12V20H18V12" />
      <path d="M9 20V14H15V20" />
    </svg>
  )
}
const LaunchIcon = ({ isActive }: { icon: ReactNode; isActive: boolean }) => {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke={isActive ? PrimaryColor : inactiveStroke} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M12 21L15 16L19 15L12 2L5 15L9 16L12 21Z" />
      <path d="M7 11H17" />
      <path d="M12 11V6" />
    </svg>
  )
}
const GuideIcon = ({ isActive }: { icon: ReactNode; isActive: boolean }) => {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke={isActive ? PrimaryColor : inactiveStroke} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M12 6C9.23858 6 7 8.23858 7 11C7 13.7614 9.23858 16 12 16C14.7614 16 17 13.7614 17 11C17 8.23858 14.7614 6 12 6Z" />
      <path d="M12 15V18" />
      <path d="M12 20H12.01" />
    </svg>
  )
}
const DocIcon = ({ isActive }: { icon: ReactNode; isActive: boolean }) => {
  return (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke={isActive ? PrimaryColor : inactiveStroke} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M7 22H17C18.1046 22 19 21.1046 19 20V6C19 4.89543 18.1046 4 17 4H13L7 10V20C7 21.1046 7.89543 22 9 22Z" />
      <path d="M7 10H13V4" />
      <path d="M9 15H15" />
      <path d="M9 18H15" />
    </svg>
  )
}
