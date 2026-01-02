/**
 * Fluxus Design System Type Definitions
 * 
 * Provides TypeScript types for the Fluxus design system variables
 * and utility classes to ensure type safety and better DX.
 */

// Color System Types
export type FluxusColor = 
  | 'fluxus-primary'
  | 'fluxus-primary-dark'
  | 'fluxus-secondary'
  | 'fluxus-accent'
  | 'fluxus-warning';

export type BackgroundColor =
  | 'bg-primary'
  | 'bg-secondary' 
  | 'bg-tertiary'
  | 'bg-card'
  | 'bg-card-hover'
  | 'bg-glass'
  | 'bg-overlay';

export type TextColor =
  | 'text-primary'
  | 'text-secondary'
  | 'text-muted'
  | 'text-disabled';

export type BorderColor =
  | 'border-primary'
  | 'border-secondary'
  | 'border-accent';

// Typography System Types
export type FontFamily =
  | 'font-primary'
  | 'font-heading'
  | 'font-body'
  | 'font-accent';

export type FontSize = 
  | 'text-xs'
  | 'text-sm'
  | 'text-base'
  | 'text-lg'
  | 'text-xl'
  | 'text-2xl'
  | 'text-3xl'
  | 'text-4xl'
  | 'text-5xl'
  | 'text-6xl'
  | 'text-sm-desktop'
  | 'text-base-desktop'
  | 'text-lg-desktop'
  | 'text-xl-desktop'
  | 'text-2xl-desktop'
  | 'text-3xl-desktop'
  | 'text-4xl-desktop'
  | 'text-5xl-desktop'
  | 'text-6xl-desktop';

export type FontWeight =
  | 'font-normal'
  | 'font-medium'
  | 'font-semibold'
  | 'font-bold'
  | 'font-extrabold'
  | 'font-black';

// Spacing System Types
export type Spacing =
  | 'xs'
  | 'sm'
  | 'md'
  | 'lg'
  | 'xl'
  | '2xl'
  | '3xl'
  | '4xl'
  | '5xl';

export type Size =
  | 'xs'
  | 'sm'
  | 'md'
  | 'lg'
  | 'xl'
  | '2xl';

// Border System Types
export type BorderRadius =
  | 'rounded-sm'
  | 'rounded-md'
  | 'rounded-lg'
  | 'rounded-xl'
  | 'rounded-2xl'
  | 'rounded-full';

export type BorderWidth =
  | 'border-thin'
  | 'border-normal'
  | 'border-thick';

// Animation System Types
export type TransitionDuration =
  | 'duration-fast'
  | 'duration-normal'
  | 'duration-slow'
  | 'duration-slower';

export type TransitionTiming =
  | 'ease-linear'
  | 'ease-in'
  | 'ease-out'
  | 'ease-in-out';

// Component Class Types
export type TypographyClass =
  | 'heading-hero'
  | 'heading-hero-highlight'
  | 'text-body-hero'
  | 'heading-h1'
  | 'heading-h2'
  | 'heading-h3'
  | 'text-body-lg'
  | 'text-body'
  | 'text-body-sm'
  | 'text-caption';

export type BackgroundClass =
  | 'bg-primary'
  | 'bg-card'
  | 'bg-card-hover'
  | 'bg-glass';

export type ButtonClass =
  | 'btn-primary'
  | 'btn-secondary';

export type CardClass =
  | 'card'
  | 'card-feature';

export type AnimationClass =
  | 'animate-float'
  | 'animate-glow'
  | 'animate-shimmer';

export type GradientTextClass =
  | 'text-gradient-primary'
  | 'text-gradient-secondary'
  | 'text-gradient-accent';

export type UtilityClass =
  | 'glass-border'
  | 'glow-primary'
  | 'glow-secondary'
  | 'transition-smooth'
  | 'transition-slow';

// Combined Design System Class Type
export type DesignSystemClass = 
  | TypographyClass
  | BackgroundClass
  | ButtonClass
  | CardClass
  | AnimationClass
  | GradientTextClass
  | UtilityClass;

// Design System Configuration
export interface DesignSystemConfig {
  colors: {
    primary: FluxusColor;
    background: BackgroundColor;
    text: TextColor;
    border: BorderColor;
  };
  typography: {
    family: FontFamily;
    size: FontSize;
    weight: FontWeight;
  };
  spacing: {
    margin: Spacing;
    padding: Spacing;
  };
  borders: {
    radius: BorderRadius;
    width: BorderWidth;
  };
  animations: {
    duration: TransitionDuration;
    timing: TransitionTiming;
  };
}

// Component Props with Design System Support
export interface FluxusComponentProps {
  className?: string;
  variant?: 'primary' | 'secondary' | 'accent';
  size?: Size;
  rounded?: BorderRadius;
  children?: React.ReactNode;
}

// Theme Context Types
export interface FluxusTheme extends DesignSystemConfig {
  isDark: boolean;
  toggleTheme: () => void;
}

export type ThemeProvider = React.Context<FluxusTheme>;

// Utility function types
export type ClassNameBuilder = (...classes: (string | undefined | null | boolean)[]) => string;
export type VariantBuilder<T extends string> = (variant: T, size?: Size) => string;

// Design System Constants
export const DESIGN_SYSTEM_CONSTANTS = {
  BREAKPOINTS: {
    sm: '640px',
    md: '768px',
    lg: '1024px',
    xl: '1280px',
    '2xl': '1536px',
  },
  Z_INDEX: {
    dropdown: 1000,
    sticky: 1020,
    fixed: 1030,
    modalBackdrop: 1040,
    modal: 1050,
    popover: 1060,
    tooltip: 1070,
  },
  ANIMATION_DURATIONS: {
    fast: 150,
    normal: 300,
    slow: 500,
    slower: 700,
  },
} as const;

export type BreakpointKey = keyof typeof DESIGN_SYSTEM_CONSTANTS.BREAKPOINTS;
export type ZIndexKey = keyof typeof DESIGN_SYSTEM_CONSTANTS.Z_INDEX;
export type AnimationDurationKey = keyof typeof DESIGN_SYSTEM_CONSTANTS.ANIMATION_DURATIONS;
