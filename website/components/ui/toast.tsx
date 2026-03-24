import * as React from "react"
import { cn } from "@/lib/utils"

interface ToastProps {
  message: string
  isVisible: boolean
  onClose: () => void
}

export function Toast({ message, isVisible, onClose }: ToastProps) {
  React.useEffect(() => {
    if (isVisible) {
      const timer = setTimeout(() => {
        onClose()
      }, 2000)
      return () => clearTimeout(timer)
    }
  }, [isVisible, onClose])

  return (
    <div
      className={cn(
        "absolute left-1/2 -top-12 min-w-[200px] -translate-x-1/2 border border-black/10 bg-[color:var(--bg-surface)] px-6 py-3 shadow-[0_18px_40px_rgba(17,24,39,0.08)] transition-all duration-300",
        isVisible ? "opacity-100 translate-y-0" : "opacity-0 translate-y-2 pointer-events-none"
      )}
    >
      <div className="flex items-center justify-center gap-3">
        <svg className="h-5 w-5 text-[color:var(--color-primary)]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
        </svg>
        <span className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-primary)]">{message}</span>
      </div>
    </div>
  )
}
