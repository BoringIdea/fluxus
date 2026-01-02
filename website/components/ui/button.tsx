import * as React from "react"
import { Slot } from "@radix-ui/react-slot"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const buttonVariants = cva(
  "inline-flex items-center justify-center whitespace-nowrap border border-border bg-transparent px-5 py-2 text-sm font-bold uppercase tracking-[0.08em] text-primary transition-colors duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-border focus-visible:ring-offset-2 focus-visible:ring-offset-background disabled:pointer-events-none disabled:opacity-40 rounded-none",
  {
    variants: {
      variant: {
        default:
          "bg-fluxus-primary text-black hover:bg-[#1FB455] focus-visible:ring-border focus-visible:ring-offset-background",
        secondary:
          "text-primary hover:bg-bg-tertiary/80",
        outline:
          "bg-transparent text-primary border border-border hover:bg-bg-card-hover",
        ghost:
          "border-transparent text-secondary hover:text-primary hover:bg-bg-card-hover/80",
        danger:
          "border-[#FB7185] bg-[#FB7185] text-black hover:bg-[#F43F5E]",
        link: "border-transparent px-0 py-0 text-primary underline-offset-4 hover:underline",
      },
      size: {
        default: "h-11 px-5",
        sm: "h-9 px-4 text-xs",
        lg: "h-12 px-8 text-base",
        icon: "h-10 w-10 px-0",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  },
)

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
  VariantProps<typeof buttonVariants> {
  asChild?: boolean
}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, asChild = false, ...props }, ref) => {
    const Comp = asChild ? Slot : "button"
    return <Comp className={cn(buttonVariants({ variant, size, className }))} ref={ref} {...props} />
  },
)
Button.displayName = "Button"

export { Button, buttonVariants }
