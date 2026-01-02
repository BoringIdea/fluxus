import { NextPage } from 'next';
import { useState, useEffect } from 'react';
import Link from 'next/link';
import { ArrowRight, Zap, Shield, TrendingUp, Palette, Globe, ChevronDown } from 'lucide-react';
import { cn, textVariant, buttonVariant, cardVariant, backgroundVariant, responsiveText, spacing, flexLayout, gridLayout, container } from '@/src/utils/design-system';

const HomePage: NextPage = () => {
  const [isVisible, setIsVisible] = useState(false);
  const [currentFeature, setCurrentFeature] = useState(0);

  useEffect(() => {
    setIsVisible(true);
    const interval = setInterval(() => {
      setCurrentFeature((prev) => (prev + 1) % 4);
    }, 3000);

    return () => {
      clearInterval(interval);
    };
  }, []);

  const features = [
    {
      icon: <TrendingUp className="w-8 h-8" />,
      title: "Bonding Curve",
      description: "Smart pricing algorithm for fair NFT valuation",
      color: "from-green-400 to-emerald-500"
    },
    {
      icon: <Zap className="w-8 h-8" />,
      title: "Instant Liquidity",
      description: "Trade NFTs anytime through smart contracts",
      color: "from-blue-400 to-cyan-500"
    },
    {
      icon: <Shield className="w-8 h-8" />,
      title: "Smart Contract",
      description: "Decentralized trading without platform risks",
      color: "from-purple-400 to-pink-500"
    },
    {
      icon: <Globe className="w-8 h-8" />,
      title: "Cross Chain",
      description: "Seamless NFT transfers across EVM networks",
      color: "from-orange-400 to-red-500"
    }
  ];


  return (
    <div className="min-h-screen relative overflow-hidden bg-[#141414]">
      {/* Animated Background Elements */}
      <div className="absolute inset-0">
        <div className="absolute top-20 left-10 w-72 h-72 bg-green-500 rounded-full mix-blend-multiply filter blur-xl opacity-10 animate-pulse"></div>
        <div className="absolute bottom-20 right-10 w-96 h-96 bg-blue-500 rounded-full mix-blend-multiply filter blur-xl opacity-10 animate-pulse delay-1000"></div>
        <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-80 h-80 bg-purple-500 rounded-full mix-blend-multiply filter blur-xl opacity-10 animate-pulse delay-2000"></div>
      </div>

      {/* Floating Planets */}
      {[...Array(8)].map((_, i) => {
        const size = 40 + Math.random() * 80;
        const colors = [
          'from-green-400 to-emerald-600',
          'from-blue-400 to-cyan-600', 
          'from-purple-400 to-pink-600',
          'from-orange-400 to-red-600',
          'from-indigo-400 to-purple-600',
          'from-cyan-400 to-blue-600'
        ];
        const colorClass = colors[Math.floor(Math.random() * colors.length)];
        
        return (
          <div
            key={i}
            className="absolute opacity-20 pointer-events-none"
            style={{
              left: `${Math.random() * 100}%`,
              top: `${Math.random() * 100}%`,
              animation: `planetFloat ${15 + Math.random() * 10}s ease-in-out infinite`,
              animationDelay: `${Math.random() * 5}s`
            }}
          >
            {/* Planet Body */}
            <div 
              className={`relative bg-gradient-to-br ${colorClass} rounded-full shadow-2xl`}
              style={{ width: size, height: size }}
            >
              {/* Planet Ring (for some planets) */}
              {Math.random() > 0.6 && (
                <div 
                  className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 border-2 border-current opacity-30 rounded-full"
                  style={{ 
                    width: size * 1.6, 
                    height: size * 0.3,
                    borderRadius: '50%'
                  }}
                />
              )}
              
              {/* Planet Surface Details */}
              <div className="absolute top-2 left-2 w-2 h-2 bg-white/30 rounded-full"></div>
              <div className="absolute bottom-3 right-3 w-1 h-1 bg-white/20 rounded-full"></div>
              {size > 60 && (
                <div className="absolute top-1/2 left-1/3 w-3 h-3 bg-white/10 rounded-full"></div>
              )}
              
              {/* Planet Glow */}
              <div className={`absolute inset-0 bg-gradient-to-br ${colorClass} rounded-full blur-sm opacity-50 scale-110`}></div>
            </div>
          </div>
        );
      })}

      {/* Floating Bubbles */}
      {[...Array(15)].map((_, i) => {
        const size = 20 + Math.random() * 40;
        const glowColors = [
          'shadow-green-400/20',
          'shadow-blue-400/20',
          'shadow-purple-400/20',
          'shadow-cyan-400/20',
          'shadow-pink-400/20'
        ];
        const glowClass = glowColors[Math.floor(Math.random() * glowColors.length)];
        
        return (
          <div
            key={i}
            className="absolute pointer-events-none"
            style={{
              left: `${Math.random() * 100}%`,
              top: `${Math.random() * 100}%`,
              animation: `bubbleFloat ${8 + Math.random() * 6}s ease-in-out infinite`,
              animationDelay: `${Math.random() * 3}s`
            }}
          >
            <div 
              className={`relative bg-gradient-to-br from-white/5 to-white/15 rounded-full backdrop-blur-sm border border-white/10 ${glowClass} shadow-xl`}
              style={{ width: size, height: size }}
            >
              {/* Bubble Highlight */}
              <div 
                className="absolute top-2 left-2 bg-white/40 rounded-full"
                style={{ width: size * 0.3, height: size * 0.3 }}
              ></div>
              
              {/* Inner Shimmer */}
              <div className="absolute inset-2 bg-gradient-to-br from-white/10 to-transparent rounded-full"></div>
            </div>
          </div>
        );
      })}

      {/* Cosmic Dust Particles */}
      {[...Array(25)].map((_, i) => {
        const sparkleColors = [
          'bg-green-400',
          'bg-blue-400', 
          'bg-purple-400',
          'bg-cyan-400',
          'bg-pink-400',
          'bg-yellow-400'
        ];
        const colorClass = sparkleColors[Math.floor(Math.random() * sparkleColors.length)];
        
        return (
          <div
            key={i}
            className={`absolute w-1 h-1 ${colorClass} rounded-full opacity-40 pointer-events-none`}
            style={{
              left: `${Math.random() * 100}%`,
              top: `${Math.random() * 100}%`,
              animation: `sparkle ${3 + Math.random() * 4}s ease-in-out infinite`,
              animationDelay: `${Math.random() * 2}s`
            }}
          />
        );
      })}

      {/* Constellation Lines */}
      <svg className="absolute inset-0 w-full h-full opacity-10 pointer-events-none" xmlns="http://www.w3.org/2000/svg">
        <defs>
          <linearGradient id="constellationGradient" x1="0%" y1="0%" x2="100%" y2="100%">
            <stop offset="0%" stopColor="#4ade80" stopOpacity="0.3"/>
            <stop offset="50%" stopColor="#3b82f6" stopOpacity="0.2"/>
            <stop offset="100%" stopColor="#8b5cf6" stopOpacity="0.1"/>
          </linearGradient>
        </defs>
        
        {/* Random constellation lines */}
        <path 
          d="M100,150 Q200,100 300,200 T500,150" 
          stroke="url(#constellationGradient)" 
          strokeWidth="1" 
          fill="none"
          className="animate-pulse"
        />
        <path 
          d="M600,300 Q700,250 800,350 T1000,300" 
          stroke="url(#constellationGradient)" 
          strokeWidth="1" 
          fill="none"
          className="animate-pulse"
          style={{ animationDelay: '1s' }}
        />
        <path 
          d="M200,500 Q400,450 600,550 T900,500" 
          stroke="url(#constellationGradient)" 
          strokeWidth="1" 
          fill="none"
          className="animate-pulse"
          style={{ animationDelay: '2s' }}
        />
      </svg>

      {/* Header */}
      <header className={cn("relative z-50", flexLayout('row', 'center', 'between'), 'px-8 lg:px-16 xl:px-20 py-6')}>
          <div className="flex items-center space-x-2">
            <div className="w-10 h-10 bg-gradient-primary rounded-xl flex items-center justify-center">
              <Palette className="w-6 h-6 text-white" />
            </div>
            <span className={cn(textVariant('h3'), 'text-gradient-primary tracking-wider')}>
              Fluxus
            </span>
          </div>
        
        <nav className="hidden md:flex items-center space-x-8">
          <Link href="/docs/creation-guide" className="text-secondary hover:text-fluxus-primary transition-colors">Guide</Link>
          <Link href="/litepaper" className="text-secondary hover:text-fluxus-primary transition-colors">Docs</Link>
          <Link href="/">
            <button className={buttonVariant('primary', 'md')}>
              Launch App
            </button>
          </Link>
        </nav>
      </header>

      {/* Hero Section */}
      <main className="relative z-10 px-6 lg:px-12">
        <div className={container('2xl', true)}>
          {/* Main Content */}
          <div className={cn(
            'text-center pt-20 pb-16 transition-all duration-1000',
            isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-10'
          )}>
            <h1 className="mb-10">
              <span className="block font-audiowide text-4xl md:text-6xl lg:text-7xl font-normal mb-2 text-white/90 tracking-wider">
                Collect & Sell
              </span>
              <span className="block font-orbitron text-4xl md:text-6xl lg:text-7xl font-black mb-2 bg-gradient-to-r from-green-400 via-emerald-500 to-green-600 bg-clip-text text-transparent tracking-wide">
                Super Rare
              </span>
              <span className="block font-audiowide text-4xl md:text-6xl lg:text-7xl font-normal text-white/90 tracking-wider">
                Digital Art
              </span>
            </h1>

            <p className="font-rajdhani text-xl md:text-2xl font-medium mb-12 max-w-4xl mx-auto text-white/80 leading-relaxed">
              The world's first revolutionary NFT liquidity solution with{' '}
              <span className="text-fluxus-primary font-bold bg-fluxus-primary/10 px-3 py-2 rounded-lg">bonding curve pricing</span>,{' '}
              <span className="text-fluxus-secondary font-bold bg-fluxus-secondary/10 px-3 py-2 rounded-lg">instant liquidity</span>, and{' '}
              <span className="text-fluxus-accent font-bold bg-fluxus-accent/10 px-3 py-2 rounded-lg">cross-chain functionality</span>
            </p>

            {/* CTA Buttons */}
            <div className={cn(flexLayout('row', 'center', 'center', '2xl'), 'flex-col sm:flex-row mb-20 gap-6')}>
              <Link href="/">
                <button className={cn(
                  buttonVariant('primary', 'lg'),
                  'group relative flex items-center space-x-3 overflow-hidden min-w-[200px]'
                )}>
                  {/* Shimmer Effect */}
                  <div className="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent transform -skew-x-12 -translate-x-full group-hover:translate-x-full transition-transform duration-700"></div>
                  
                  <span className="relative z-10">Explore Marketplace</span>
                  <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform relative z-10" />
                </button>
              </Link>
              <Link href="/collection/create">
                <button className={cn(
                  buttonVariant('secondary', 'lg'),
                  'group relative overflow-hidden min-w-[200px]'
                )}>
                  {/* Ripple Effect */}
                  <div className="absolute inset-0 bg-green-400 transform scale-0 group-hover:scale-100 transition-transform duration-300 origin-center rounded-full"></div>
                  
                  {/* Pulse Border */}
                  <div className="absolute inset-0 border-2 border-green-400 rounded-full animate-pulse opacity-0 group-hover:opacity-50 transition-opacity duration-300"></div>
                  
                  <span className="relative z-10">Create Collection</span>
                </button>
              </Link>
            </div>

          </div>

          {/* Features Section */}
          <div className={cn(
            'transition-all duration-1000 delay-300 pt-12',
            isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-10'
          )}>
            <div className="text-center mb-16">
              <h2 className="font-orbitron text-3xl md:text-4xl lg:text-5xl font-bold mb-6 text-white tracking-wider">
                Core Innovations
              </h2>
              <p className={cn(textVariant('body-lg', 'text-muted'), 'max-w-3xl mx-auto')}>
                Breaking NFT market limitations through revolutionary technology
              </p>
            </div>

            {/* Feature Cards */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 lg:gap-12 xl:gap-16 mb-16">
              {features.map((feature, index) => (
                <div
                  key={index}
                  className={cn(
                    cardVariant('feature', true),
                    'group relative transition-all duration-700 transform',
                    currentFeature === index 
                      ? 'card-feature active' 
                      : ''
                  )}
                  style={{
                    animationDelay: `${index * 150}ms`,
                    animation: isVisible ? 'slideInUp 0.8s ease-out forwards' : 'none',
                    opacity: isVisible ? 1 : 0
                  }}
                >
                  {/* Animated Background Pattern */}
                  <div className="absolute inset-0 opacity-5 overflow-hidden rounded-2xl">
                    <div className={`absolute -top-4 -right-4 w-24 h-24 bg-gradient-to-r ${feature.color} rounded-full blur-xl transition-all duration-700 ${
                      currentFeature === index ? 'scale-150 opacity-30' : 'scale-100 opacity-10'
                    }`}></div>
                    <div className={`absolute -bottom-4 -left-4 w-16 h-16 bg-gradient-to-r ${feature.color} rounded-full blur-lg transition-all duration-700 ${
                      currentFeature === index ? 'scale-125 opacity-20' : 'scale-100 opacity-5'
                    }`}></div>
                  </div>

                  {/* Floating Icon */}
                  <div className={`relative inline-flex p-4 rounded-2xl bg-gradient-to-r ${feature.color} mb-6 transition-all duration-500 ${
                    currentFeature === index 
                      ? 'scale-125 shadow-lg shadow-current/25 animate-pulse' 
                      : 'group-hover:scale-110 group-hover:rotate-3'
                  }`}>
                    {feature.icon}
                    
                    {/* Icon Glow Effect */}
                    <div className={`absolute inset-0 rounded-2xl bg-gradient-to-r ${feature.color} blur-md opacity-0 transition-opacity duration-500 ${
                      currentFeature === index ? 'opacity-30' : 'group-hover:opacity-20'
                    }`}></div>
                  </div>

                  {/* Animated Title */}
                  <h3 className={cn(
                    'text-lg md:text-xl font-rajdhani font-bold tracking-wide',
                    'mb-4 transition-all duration-500',
                    currentFeature === index 
                      ? 'text-fluxus-primary transform scale-105' 
                      : 'text-primary group-hover:text-fluxus-primary'
                  )}>
                    {feature.title}
                  </h3>

                  {/* Description */}
                  <p className={cn(
                    textVariant('body'),
                    'leading-relaxed transition-all duration-500',
                    currentFeature === index 
                      ? 'text-secondary' 
                      : 'text-muted group-hover:text-secondary'
                  )}>
                    {feature.description}
                  </p>

                  {/* Progress Indicator */}
                  {currentFeature === index && (
                    <div className="absolute bottom-4 left-8 right-8">
                      <div className="h-1 bg-gray-700 rounded-full overflow-hidden">
                        <div 
                          className="h-full bg-gradient-to-r from-green-400 to-emerald-500 rounded-full transition-all duration-3000 ease-linear"
                          style={{
                            width: '100%',
                            animation: 'progressBar 3s linear infinite'
                          }}
                        ></div>
                      </div>
                    </div>
                  )}

                  {/* Interactive Hover Effect */}
                  <div className={`absolute inset-0 bg-gradient-to-r from-green-400/5 to-emerald-500/5 rounded-2xl transition-opacity duration-500 ${
                    currentFeature === index ? 'opacity-100' : 'opacity-0 group-hover:opacity-100'
                  }`}></div>

                  {/* Floating Particles */}
                  {currentFeature === index && (
                    <>
                      {[...Array(6)].map((_, i) => (
                        <div
                          key={i}
                          className="absolute w-1 h-1 bg-green-400 rounded-full opacity-60"
                          style={{
                            left: `${20 + Math.random() * 60}%`,
                            top: `${20 + Math.random() * 60}%`,
                            animation: `float ${2 + Math.random() * 2}s ease-in-out infinite`,
                            animationDelay: `${Math.random() * 2}s`
                          }}
                        />
                      ))}
                    </>
                  )}

                  {/* Corner Accent */}
                  <div className={`absolute top-0 right-0 w-8 h-8 border-t-2 border-r-2 rounded-tr-2xl transition-all duration-500 ${
                    currentFeature === index 
                      ? 'border-green-400 opacity-100' 
                      : 'border-gray-600 opacity-0 group-hover:opacity-50'
                  }`}></div>
                </div>
              ))}
            </div>

            {/* Feature Navigation Dots */}
            <div className="flex justify-center space-x-4">
              {features.map((_, index) => (
                <button
                  key={index}
                  onClick={() => setCurrentFeature(index)}
                  className={`w-4 h-4 rounded-full transition-all duration-300 ${
                    currentFeature === index 
                      ? 'bg-fluxus-primary scale-125 shadow-lg shadow-fluxus-primary/50' 
                      : 'bg-slate-700 hover:bg-slate-600 hover:scale-110'
                  }`}
                />
              ))}
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="relative z-10 mt-32 backdrop-blur-sm">
        <div className="border-t border-secondary">
          <div className="max-w-7xl mx-auto px-8 lg:px-16 xl:px-20 py-12">
            <div className={cn(flexLayout('row', 'center', 'between', 'xl'), 'flex-col md:flex-row')}>
              <div className="flex items-center space-x-3 mb-6 md:mb-0">
                <div className="w-10 h-10 rounded-xl flex items-center justify-center border border-secondary">
                  <Palette className="w-6 h-6 text-fluxus-primary" />
                </div>
                <span className={cn(textVariant('h3'), 'text-gradient-primary tracking-wider')}>
                  Fluxus
                </span>
              </div>
              
              <div className="flex space-x-8">
                <Link href="https://github.com/HashIdea/art-fluxus" className="text-muted hover:text-fluxus-primary transition-colors">
                  <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"/>
                  </svg>
                </Link>
                <Link href="https://x.com/fluxus_io" className="text-muted hover:text-fluxus-primary transition-colors">
                  <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/>
                  </svg>
                </Link>
                <Link href="https://www.docs.fluxusnft.xyz/" className="text-muted hover:text-fluxus-primary transition-colors">
                  <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
                  </svg>
                </Link>
              </div>
            </div>
            
          </div>
          <div className="mt-8 pt-6 border-t border-secondary text-center">
            <p className={textVariant('body-sm', 'text-muted')}>&copy; 2025 Fluxus. All rights reserved.</p>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default HomePage;
