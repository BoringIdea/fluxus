import { NextPage } from 'next'
import Link from 'next/link'
import { ArrowRight, TrendingUp, Zap, Shield, Globe } from 'lucide-react'

const features = [
  {
    icon: TrendingUp,
    title: 'Bonding Curve Liquidity',
    description: 'Price discovery stays visible, structured, and tradable from the first mint to the last bid.',
  },
  {
    icon: Zap,
    title: 'Instant Market Access',
    description: 'Collections can move from launch to active trading without waiting for fragmented marketplace depth.',
  },
  {
    icon: Shield,
    title: 'Contract-Native Trading',
    description: 'Liquidity logic lives in protocol rails rather than platform-controlled order flow.',
  },
  {
    icon: Globe,
    title: 'Cross-Chain Extension',
    description: 'Designed for multi-chain NFT movement without turning the interface into a bridge dashboard.',
  },
]

const HomePage: NextPage = () => {
  return (
    <div className="flux-shell min-h-screen">
      <main className="mx-auto max-w-7xl px-6 pb-24 pt-8 lg:px-12">
        <section className="grid gap-10 border border-black/10 bg-[color:var(--bg-surface)] px-8 py-10 lg:grid-cols-[1.2fr_0.8fr] lg:px-12 lg:py-14">
          <div className="space-y-8">
            <div className="flux-kicker">NFT Liquidity Infrastructure</div>
            <div className="space-y-5">
              <h1 className="flux-title">Liquidity for NFTs, designed as market structure.</h1>
              <p className="flux-copy max-w-2xl text-[16px]">
                Fluxus is an NFT liquidity protocol with a restrained interface and a green signal language. The product focus is not spectacle. It is pricing, launch, and tradability.
              </p>
            </div>
            <div className="flex flex-col gap-3 sm:flex-row">
              <Link href="/" className="inline-flex items-center justify-center gap-2 border border-transparent bg-[color:var(--color-primary)] px-6 py-3 font-primary text-[11px] uppercase tracking-[0.18em] text-white transition-colors hover:bg-[color:var(--color-primary-dark)]">
                Explore Market
                <ArrowRight className="h-4 w-4" />
              </Link>
              <Link href="/collection/create" className="inline-flex items-center justify-center gap-2 border border-black/10 bg-[color:var(--bg-muted)] px-6 py-3 font-primary text-[11px] uppercase tracking-[0.18em] text-[color:var(--text-primary)] transition-colors hover:bg-[color:var(--bg-card-hover)]">
                Launch Collection
              </Link>
            </div>
          </div>

          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-1">
            <div className="flux-panel p-5">
              <div className="flux-kicker">Pricing Layer</div>
              <div className="mt-3 font-heading text-[32px] leading-none text-[color:var(--text-primary)]">Curve-based</div>
              <p className="flux-copy mt-3 text-[13px]">Green is used as signal, not decoration. The interface stays neutral so pricing stays first.</p>
            </div>
            <div className="flux-panel-muted p-5">
              <div className="flux-kicker">Market Mode</div>
              <div className="mt-3 font-heading text-[28px] leading-none text-[color:var(--text-primary)]">Launch to Trade</div>
              <p className="flux-copy mt-3 text-[13px]">From creation flow to collection depth, the system is meant to feel like an exchange surface for NFTs.</p>
            </div>
          </div>
        </section>

        <section className="mt-14 grid gap-4 md:grid-cols-2 xl:grid-cols-4">
          {features.map((feature) => {
            const Icon = feature.icon
            return (
              <article key={feature.title} className="flux-panel p-6">
                <div className="flex h-10 w-10 items-center justify-center border border-[color:var(--color-primary)]/18 bg-[color:var(--color-primary)]/8 text-[color:var(--color-primary)]">
                  <Icon className="h-4 w-4" />
                </div>
                <h2 className="mt-5 font-heading text-[26px] leading-[1] text-[color:var(--text-primary)]">{feature.title}</h2>
                <p className="flux-copy mt-4 text-[13px]">{feature.description}</p>
              </article>
            )
          })}
        </section>

        <section className="mt-14 grid gap-6 lg:grid-cols-[0.9fr_1.1fr]">
          <div className="flux-panel-muted p-6">
            <div className="flux-kicker">Why Fluxus</div>
            <h2 className="flux-h2 mt-4">A protocol interface, not a generic NFT landing page.</h2>
          </div>
          <div className="flux-panel p-6">
            <div className="grid gap-6 md:grid-cols-3">
              <div>
                <div className="flux-kicker">01</div>
                <p className="flux-copy mt-3 text-[13px]">Collections should launch into liquidity, not wait for volume to appear elsewhere.</p>
              </div>
              <div>
                <div className="flux-kicker">02</div>
                <p className="flux-copy mt-3 text-[13px]">Traders need price rails, not decorative dashboards.</p>
              </div>
              <div>
                <div className="flux-kicker">03</div>
                <p className="flux-copy mt-3 text-[13px]">Cross-chain should extend the market, not dominate the page.</p>
              </div>
            </div>
          </div>
        </section>
      </main>
    </div>
  )
}

export default HomePage
