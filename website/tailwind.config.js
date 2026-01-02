/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: ["class"],
    content: [
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
 
    // Or if using `src` directory:
    "./src/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
  	extend: {
      /* Fluxus Design System Extensions */
      fontFamily: {
        caveat: ['var(--font-caveat)'],
        primary: ['var(--font-primary)'],
        heading: ['var(--font-heading)'],
        body: ['var(--font-body)'],
        accent: ['var(--font-accent)'],
        display: ['var(--font-display)'],
        inter: ['var(--font-inter)'],
        orbitron: ['var(--font-orbitron)'],
        rajdhani: ['var(--font-rajdhani)'],
        exo2: ['var(--font-exo2)'],
        audiowide: ['var(--font-audiowide)'],
      },
      
      fontSize: {
        'xs': 'var(--text-xs)',
        'sm': 'var(--text-sm)',
        'base': 'var(--text-base)',
        'lg': 'var(--text-lg)',
        'xl': 'var(--text-xl)',
        '2xl': 'var(--text-2xl)',
        '3xl': 'var(--text-3xl)',
        '4xl': 'var(--text-4xl)',
        '5xl': 'var(--text-5xl)',
        '6xl': 'var(--text-6xl)',
        'sm-desktop': 'var(--text-sm-desktop)',
        'base-desktop': 'var(--text-base-desktop)',
        'lg-desktop': 'var(--text-lg-desktop)',
        'xl-desktop': 'var(--text-xl-desktop)',
        '2xl-desktop': 'var(--text-2xl-desktop)',
        '3xl-desktop': 'var(--text-3xl-desktop)',
        '4xl-desktop': 'var(--text-4xl-desktop)',
        '5xl-desktop': 'var(--text-5xl-desktop)',
        '6xl-desktop': 'var(--text-6xl-desktop)',
      },
      
      fontWeight: {
        normal: 'var(--font-normal)',
        medium: 'var(--font-medium)',
        semibold: 'var(--font-semibold)',
        bold: 'var(--font-bold)',
        extrabold: 'var(--font-extrabold)',
        black: 'var(--font-black)',
      },
      
      lineHeight: {
        tight: 'var(--leading-tight)',
        normal: 'var(--leading-normal)',
        relaxed: 'var(--leading-relaxed)',
        loose: 'var(--leading-loose)',
      },
      
      spacing: {
        'xs': 'var(--space-xs)',
        'sm': 'var(--space-sm)',
        'md': 'var(--space-md)',
        'lg': 'var(--space-lg)',
        'xl': 'var(--space-xl)',
        '2xl': 'var(--space-2xl)',
        '3xl': 'var(--space-3xl)',
        '4xl': 'var(--space-4xl)',
        '5xl': 'var(--space-5xl)',
      },
      
      width: {
        'xs': 'var(--size-xs)',
        'sm': 'var(--size-sm)',
        'md': 'var(--size-md)',
        'lg': 'var(--size-lg)',
        'xl': 'var(--size-xl)',
        '2xl': 'var(--size-2xl)',
      },
      
      height: {
        'xs': 'var(--size-xs)',
        'sm': 'var(--size-sm)',
        'md': 'var(--size-md)',
        'lg': 'var(--size-lg)',
        'xl': 'var(--size-xl)',
        '2xl': 'var(--size-2xl)',
      },
      
  		borderRadius: {
        'sm': 'var(--radius-sm)',
        'md': 'var(--radius-md)',
        'lg': 'var(--radius-lg)',
        'xl': 'var(--radius-xl)',
        '2xl': 'var(--radius-2xl)',
        'full': 'var(--radius-full)',
  		},
      
      borderWidth: {
        thin: 'var(--border-thin)',
        normal: 'var(--border-normal)',
        thick: 'var(--border-thick)',
      },
      
      transitionDuration: {
        fast: 'var(--duration-fast)',
        normal: 'var(--duration-normal)',
        slow: 'var(--duration-slow)',
        slower: 'var(--duration-slower)',
      },
      
      transitionTimingFunction: {
        'ease-linear': 'var(--ease-linear)',
        'ease-in': 'var(--ease-in)',
        'ease-out': 'var(--ease-out)',
        'ease-in-out': 'var(--ease-in-out)',
      },
      
      zIndex: {
        dropdown: 'var(--z-dropdown)',
        sticky: 'var(--z-sticky)',
        fixed: 'var(--z-fixed)',
        'modal-backdrop': 'var(--z-modal-backdrop)',
        modal: 'var(--z-modal)',
        popover: 'var(--z-popover)',
        tooltip: 'var(--z-tooltip)',
      },
      
      boxShadow: {
        'sm': 'var(--shadow-sm)',
        'md': 'var(--shadow-md)',
        'lg': 'var(--shadow-lg)',
        'xl': 'var(--shadow-xl)',
        '2xl': 'var(--shadow-2xl)',
      },
      
      backgroundImage: {
        'gradient-primary': 'var(--gradient-primary)',
        'gradient-secondary': 'var(--gradient-secondary)',
        'gradient-accent': 'var(--gradient-accent)',
        'gradient-warning': 'var(--gradient-warning)',
        'gradient-background': 'var(--gradient-background)',
      },
      
  		colors: {
        /* Fluxus Color System */
        'bg-primary': 'hsl(var(--bg-primary))',
        'bg-secondary': 'hsl(var(--bg-secondary))',
        'bg-tertiary': 'hsl(var(--bg-tertiary))',
        'bg-card': 'hsl(var(--bg-card))',
        'bg-card-hover': 'hsl(var(--bg-card-hover))',
        'bg-glass': 'hsl(var(--bg-glass))',
        'bg-overlay': 'hsl(var(--bg-overlay))',
        
        'fluxus-primary': 'hsl(var(--color-primary))',
        'fluxus-primary-dark': 'hsl(var(--color-primary-dark))',
        'fluxus-secondary': 'hsl(var(--color-secondary))',
        'fluxus-accent': 'hsl(var(--color-accent))',
        'fluxus-warning': 'hsl(var(--color-warning))',
        
        'text-primary': 'hsl(var(--text-primary))',
        'text-secondary': 'hsl(var(--text-secondary))',
        'text-muted': 'hsl(var(--text-muted))',
        'text-disabled': 'hsl(var(--text-disabled))',
        
        'border-primary': 'hsl(var(--border-primary))',
        'border-secondary': 'hsl(var(--border-secondary))',
        'border-accent': 'hsl(var(--border-accent))',
        
        /* Legacy Compatibility */
  			background: 'hsl(var(--background))',
  			foreground: 'hsl(var(--foreground))',
  			card: {
  				DEFAULT: 'hsl(var(--card))',
  				foreground: 'hsl(var(--card-foreground))'
  			},
  			primary: {
  				DEFAULT: 'hsl(var(--primary))',
  				foreground: 'hsl(var(--primary-foreground))'
  			},
  			border: 'hsl(var(--border))',
  			ring: 'hsl(var(--ring))',
  		}
  	}
  },
  plugins: [require("tailwindcss-animate"), require('@tailwindcss/typography')],
}

