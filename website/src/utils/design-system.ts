/**
 * Fluxus Design System Utilities
 * 
 * Provides utility functions for working with the Fluxus design system
 */

import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
import {
  DesignSystemClass,
  Size,
  BorderRadius,
  FontSize,
  Spacing,
  FluxusColor,
  BackgroundColor,
  TextColor
} from '@/src/types/design-system';

/**
 * Combines class names with tailwind-merge for optimal CSS
 */
export function cn(...inputs: ClassValue[]): string {
  return twMerge(clsx(inputs));
}

/**
 * Generates responsive font size classes
 */
export function responsiveText(
  mobileSize: FontSize,
  desktopSize?: FontSize
): string {
  if (!desktopSize) return mobileSize;
  return cn(mobileSize, `md:${desktopSize}`);
}

/**
 * Generates spacing classes with consistent naming
 */
export function spacing(
  type: 'p' | 'm' | 'px' | 'py' | 'pt' | 'pb' | 'pl' | 'pr' | 'mx' | 'my' | 'mt' | 'mb' | 'ml' | 'mr',
  size: Spacing
): string {
  return `${type}-${size}`;
}

/**
 * Generates button variant classes
 */
export function buttonVariant(
  variant: 'primary' | 'secondary' | 'ghost' = 'primary',
  size: Size = 'md'
): string {
  const baseClasses = 'inline-flex items-center justify-center font-heading font-bold transition-slow';
  
  const variantClasses = {
    primary: 'btn-primary',
    secondary: 'btn-secondary',
    ghost: 'bg-transparent text-fluxus-primary hover:bg-fluxus-primary/10'
  };
  
  const sizeClasses = {
    xs: 'px-sm py-xs text-sm',
    sm: 'px-md py-sm text-sm',
    md: 'px-xl py-md text-base',
    lg: 'px-2xl py-lg text-lg',
    xl: 'px-3xl py-xl text-xl',
    '2xl': 'px-4xl py-2xl text-2xl'
  };
  
  return cn(baseClasses, variantClasses[variant], sizeClasses[size]);
}

/**
 * Generates card variant classes
 */
export function cardVariant(
  variant: 'default' | 'feature' | 'glass' = 'default',
  interactive = false
): string {
  const baseClasses = 'card';
  
  const variantClasses = {
    default: '',
    feature: 'card-feature',
    glass: 'bg-glass'
  };
  
  const interactiveClasses = interactive ? 'cursor-pointer hover:bg-card-hover' : '';
  
  return cn(baseClasses, variantClasses[variant], interactiveClasses);
}

/**
 * Generates text variant classes
 */
export function textVariant(
  variant: 'hero' | 'hero-highlight' | 'hero-body' | 'h1' | 'h2' | 'h3' | 'body-lg' | 'body' | 'body-sm' | 'caption',
  color?: TextColor,
  gradient?: 'primary' | 'secondary' | 'accent'
): string {
  const variantClasses = {
    hero: 'heading-hero',
    'hero-highlight': 'heading-hero-highlight',
    'hero-body': 'text-body-hero',
    h1: 'heading-h1',
    h2: 'heading-h2',
    h3: 'heading-h3',
    'body-lg': 'text-body-lg',
    body: 'text-body',
    'body-sm': 'text-body-sm',
    caption: 'text-caption'
  };
  
  const colorClass = color ? color : '';
  const gradientClass = gradient ? `text-gradient-${gradient}` : '';
  
  return cn(variantClasses[variant], colorClass, gradientClass);
}

/**
 * Generates background variant classes
 */
export function backgroundVariant(
  variant: 'primary' | 'card' | 'glass' | 'overlay' = 'primary'
): string {
  const variantClasses = {
    primary: 'bg-primary',
    card: 'bg-card',
    glass: 'bg-glass',
    overlay: 'bg-overlay'
  };
  
  return variantClasses[variant];
}

/**
 * Generates border variant classes
 */
export function borderVariant(
  radius: BorderRadius = 'rounded-md',
  color: 'primary' | 'secondary' | 'accent' = 'primary',
  width: 'thin' | 'normal' | 'thick' = 'thin'
): string {
  const colorClasses = {
    primary: 'border-primary',
    secondary: 'border-secondary',
    accent: 'border-accent'
  };
  
  const widthClasses = {
    thin: 'border-thin',
    normal: 'border-normal',
    thick: 'border-thick'
  };
  
  return cn(radius, colorClasses[color], widthClasses[width]);
}

/**
 * Generates animation classes
 */
export function animationVariant(
  type: 'float' | 'glow' | 'shimmer' | 'none' = 'none',
  duration: 'fast' | 'normal' | 'slow' | 'slower' = 'normal'
): string {
  if (type === 'none') return '';
  
  const typeClasses = {
    float: 'animate-float',
    glow: 'animate-glow',
    shimmer: 'animate-shimmer'
  };
  
  return cn(typeClasses[type], `duration-${duration}`);
}

/**
 * Generates glow effect classes
 */
export function glowEffect(
  intensity: 'primary' | 'secondary' | 'none' = 'none'
): string {
  if (intensity === 'none') return '';
  
  const glowClasses = {
    primary: 'glow-primary',
    secondary: 'glow-secondary'
  };
  
  return glowClasses[intensity];
}

/**
 * Generates grid layout classes
 */
export function gridLayout(
  cols: 1 | 2 | 3 | 4 | 6 | 12,
  gap: Spacing = 'md',
  responsive = true
): string {
  const baseGrid = `grid gap-${gap}`;
  
  if (!responsive) {
    return cn(baseGrid, `grid-cols-${cols}`);
  }
  
  // Responsive grid breakpoints
  const responsiveClasses = {
    1: 'grid-cols-1',
    2: 'grid-cols-1 md:grid-cols-2',
    3: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3',
    4: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-4',
    6: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-6',
    12: 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-12'
  };
  
  return cn(baseGrid, responsiveClasses[cols]);
}

/**
 * Generates flex layout classes
 */
export function flexLayout(
  direction: 'row' | 'col' = 'row',
  align: 'start' | 'center' | 'end' | 'stretch' = 'center',
  justify: 'start' | 'center' | 'end' | 'between' | 'around' | 'evenly' = 'center',
  gap: Spacing = 'md'
): string {
  const directionClass = direction === 'col' ? 'flex-col' : 'flex-row';
  const alignClass = `items-${align}`;
  const justifyClass = `justify-${justify}`;
  const gapClass = `gap-${gap}`;
  
  return cn('flex', directionClass, alignClass, justifyClass, gapClass);
}

/**
 * Generates z-index utility class
 */
export function zIndex(
  level: 'dropdown' | 'sticky' | 'fixed' | 'modal-backdrop' | 'modal' | 'popover' | 'tooltip'
): string {
  return `z-${level}`;
}

/**
 * Generates container classes with max width
 */
export function container(
  size: 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full' = 'xl',
  centered = true
): string {
  const sizeClasses = {
    sm: 'max-w-screen-sm',
    md: 'max-w-screen-md',
    lg: 'max-w-screen-lg',
    xl: 'max-w-screen-xl',
    '2xl': 'max-w-screen-2xl',
    full: 'max-w-full'
  };
  
  const centerClass = centered ? 'mx-auto' : '';
  
  return cn('w-full', sizeClasses[size], centerClass);
}

/**
 * Validates if a class name is part of the design system
 */
export function isDesignSystemClass(className: string): className is DesignSystemClass {
  const designSystemClasses = [
    'heading-hero', 'heading-h1', 'heading-h2', 'heading-h3',
    'text-body-lg', 'text-body', 'text-body-sm', 'text-caption',
    'bg-primary', 'bg-card', 'bg-card-hover', 'bg-glass',
    'btn-primary', 'btn-secondary',
    'card', 'card-feature',
    'animate-float', 'animate-glow', 'animate-shimmer',
    'text-gradient-primary', 'text-gradient-secondary', 'text-gradient-accent',
    'glass-border', 'glow-primary', 'glow-secondary', 'transition-smooth', 'transition-slow'
  ];
  
  return designSystemClasses.includes(className as DesignSystemClass);
}

/**
 * Generates a complete component class string from design system tokens
 */
export function buildComponentClasses(config: {
  base?: string;
  variant?: string;
  size?: Size;
  color?: FluxusColor | TextColor;
  background?: BackgroundColor;
  border?: BorderRadius;
  spacing?: Spacing;
  animation?: 'float' | 'glow' | 'shimmer';
  glow?: 'primary' | 'secondary';
  responsive?: boolean;
  interactive?: boolean;
  custom?: string;
}): string {
  const {
    base = '',
    variant = '',
    size,
    color,
    background,
    border,
    spacing: spacingValue,
    animation,
    glow,
    responsive = true,
    interactive = false,
    custom = ''
  } = config;
  
  const classes = [
    base,
    variant,
    size && `size-${size}`,
    color,
    background,
    border,
    spacingValue && `p-${spacingValue}`,
    animation && animationVariant(animation),
    glow && glowEffect(glow),
    interactive && 'cursor-pointer transition-smooth hover:scale-105',
    custom
  ].filter(Boolean);
  
  return cn(...classes);
}

// Export commonly used combinations
export const commonStyles = {
  button: {
    primary: () => buttonVariant('primary'),
    secondary: () => buttonVariant('secondary'),
    ghost: () => buttonVariant('ghost')
  },
  card: {
    default: () => cardVariant('default'),
    feature: () => cardVariant('feature', true),
    glass: () => cardVariant('glass')
  },
  text: {
    hero: () => textVariant('hero', 'text-primary'),
    heading: () => textVariant('h1', 'text-primary'),
    subheading: () => textVariant('h2', 'text-secondary'),
    body: () => textVariant('body', 'text-secondary'),
    caption: () => textVariant('caption', 'text-muted')
  },
  layout: {
    container: () => container('xl', true),
    grid: () => gridLayout(4, 'lg'),
    flex: () => flexLayout('row', 'center', 'between', 'md')
  }
};
