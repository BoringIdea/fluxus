import * as React from "react"

import { cn } from "@/lib/utils"

const Input = React.forwardRef<HTMLInputElement, React.ComponentProps<"input">>(
  ({ className, type, ...props }, ref) => {
    return (
      <input
        type={type}
        className={cn(
          "flex h-10 w-full rounded-none border border-border bg-bg-tertiary px-4 py-2 text-sm font-semibold text-primary placeholder:text-secondary/80 placeholder:font-semibold transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-border focus-visible:ring-offset-0 focus-visible:border-border disabled:cursor-not-allowed disabled:opacity-40 file:border-0 file:bg-transparent file:text-sm file:font-semibold file:text-primary",
          className
        )}
        ref={ref}
        {...props}
      />
    )
  }
)
Input.displayName = "Input"

export { Input }
