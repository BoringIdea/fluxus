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
    <div className="fixed inset-0 bg-black/90 z-50 flex items-center justify-center p-4">
      <div className="bg-[#171a1f] p-6 rounded-lg max-w-md text-center">
        <h2 className="text-xl font-bold text-white mb-4">Please use desktop browser</h2>
        <p className="text-gray-400 mb-4">
          To get the best experience, please use a desktop browser to visit our website.
        </p>
        <div className="text-sm text-gray-500">
          Mobile support is under development...
        </div>
      </div>
    </div>
  )
} 