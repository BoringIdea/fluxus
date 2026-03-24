'use client'
import { useEffect, useState } from 'react'

export default function MobileWarning() {
  const [isMobile, setIsMobile] = useState(false)

  useEffect(() => {
    const checkMobile = () => {
      setIsMobile(window.innerWidth <= 768)
    }
    
    checkMobile()
    window.addEventListener('resize', checkMobile)
    
    return () => {
      window.removeEventListener('resize', checkMobile)
    }
  }, [])

  if (!isMobile) return null

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-[color:var(--fg-strong)]/80 p-4">
      <div className="max-w-md border border-white/15 bg-[color:var(--bg-surface)] p-6 text-center">
        <p className="flux-kicker mb-3">Display</p>
        <h2 className="mb-4 font-heading text-[30px] leading-none text-[color:var(--text-primary)]">Please use desktop browser</h2>
        <p className="mb-4 text-sm text-[color:var(--text-secondary)]">
          To get the best experience, please use a desktop browser to visit our website.
        </p>
        <div className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
          Mobile support is under development...
        </div>
      </div>
    </div>
  )
} 
