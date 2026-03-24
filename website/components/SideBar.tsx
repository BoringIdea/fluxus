"use client"

import Link from "next/link"
import { useRouter } from "next/router"
import { cn } from "@/lib/utils"
import { BarChart3, Rocket, BookOpen, FileText } from "lucide-react"

interface Route {
  path: string
  label: string
  icon: React.ComponentType<{ className?: string }>
}

export default function Sidebar() {
  const router = useRouter()
  const pathname = router.pathname

  const routes: Route[] = [
    { path: "/", label: "Market", icon: BarChart3 },
    { path: "/collection/create", label: "Launch", icon: Rocket },
    { path: "/docs/creation-guide", label: "Guide", icon: BookOpen },
    { path: "/litepaper", label: "Docs", icon: FileText },
  ]

  return (
    <aside className="fixed bottom-0 left-0 top-0 z-20 hidden w-20 border-r border-black/10 bg-[color:var(--bg-page)] sm:flex sm:flex-col sm:items-center">
      <div className="flex w-full flex-col items-center gap-3 border-b border-black/10 px-4 py-6">
        <div className="flex h-10 w-10 items-center justify-center border border-[#16A34A]/20 bg-[#16A34A] font-heading text-[18px] text-white">
          F
        </div>
      </div>

      <div className="mt-8 flex flex-col gap-4">
        {routes.map((route) => {
          const isActive = pathname === route.path
          const Icon = route.icon
          return (
            <div key={route.path} className="group relative">
              <Link href={route.path} className="block">
                <div
                  className={cn(
                    "flex h-12 w-12 items-center justify-center border font-primary text-[11px] uppercase tracking-[0.18em] transition-colors",
                    isActive
                      ? "border-black/12 bg-[color:var(--bg-muted)] text-[color:var(--text-primary)]"
                      : "border-black/10 bg-[color:var(--bg-surface)] text-[color:var(--text-muted)] hover:text-[color:var(--text-primary)]"
                  )}
                >
                  <Icon className="h-4 w-4" />
                </div>
              </Link>
              <div className="invisible absolute left-full top-1/2 ml-3 -translate-y-1/2 border border-black/10 bg-[color:var(--bg-surface)] px-2 py-1 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)] opacity-0 transition-all duration-150 group-hover:visible group-hover:opacity-100">
                {route.label}
              </div>
            </div>
          )
        })}
      </div>
    </aside>
  )
}
