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
        "absolute left-1/2 -translate-x-1/2 -top-12 bg-black/90 text-white px-6 py-3 rounded-lg shadow-lg transition-all duration-300",
        "border border-white/20 backdrop-blur-sm min-w-[200px]",
        isVisible ? "opacity-100 translate-y-0" : "opacity-0 translate-y-2 pointer-events-none"
      )}
    >
      <div className="flex items-center gap-3 justify-center">
        <svg className="w-5 h-5 text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
        </svg>
        <span className="text-sm font-medium">{message}</span>
      </div>
    </div>
  )
} 