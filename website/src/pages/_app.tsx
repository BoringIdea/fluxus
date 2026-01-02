import '../styles/globals.css';
import { Inter, Orbitron, Rajdhani, Exo_2, Audiowide } from 'next/font/google'
import '@rainbow-me/rainbowkit/styles.css';
import type { AppProps } from 'next/app';
import { useRouter } from 'next/router';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { WagmiProvider } from 'wagmi';
import { RainbowKitProvider, Theme, darkTheme } from '@rainbow-me/rainbowkit';

import { config } from '../wagmi';

const client = new QueryClient();

// Main body font - clean and readable
const inter = Inter({
  subsets: ['latin'],
  weight: ['400', '500', '600', '700'],
  display: 'swap',
  variable: '--font-inter',
});

// Display font - futuristic and geometric with strong block feel
const orbitron = Orbitron({
  subsets: ['latin'],
  weight: ['400', '500', '600', '700', '800', '900'],
  display: 'swap',
  variable: '--font-orbitron',
});

// Heading font - bold and condensed with strong geometric feel
const rajdhani = Rajdhani({
  subsets: ['latin'],
  weight: ['300', '400', '500', '600', '700'],
  display: 'swap',
  variable: '--font-rajdhani',
});

// Accent font - modern sci-fi with angular feel
const exo2 = Exo_2({
  subsets: ['latin'],
  weight: ['400', '500', '600', '700', '800', '900'],
  display: 'swap',
  variable: '--font-exo2',
});

// Display font - retro-futuristic with strong block presence
const audiowide = Audiowide({
  subsets: ['latin'],
  weight: ['400'],
  display: 'swap',
  variable: '--font-audiowide',
});

import Head from 'next/head';
import Header from '@/components/Header';
import CommingSoon from '@/components/CommingSoon';
import Sidebar from '@/components/SideBar';
import Footer from '@/components/Footer';
import { PrimaryColor } from '../utils';

function MyApp({ Component, pageProps }: AppProps) {
  const router = useRouter();
  const isHomePage = router.pathname === '/home';

  return (
    <>
      <Head>
        <title>Fluxus: NFT Revolution</title>
        <link href="/fluxus.svg" rel="icon" />
        <meta
          name="description"
          content="Fluxus is a Revolutionary NFT Liquidity Solution" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <WagmiProvider config={config}>
        <QueryClientProvider client={client}>
          <RainbowKitProvider locale='en' theme={darkTheme({
            accentColor: '#16A34A',
            accentColorForeground: '#000'
          })}>
            {isHomePage ? (
              // Home page layout without sidebar/header/footer
              <main className={`${inter.variable} ${orbitron.variable} ${rajdhani.variable} ${exo2.variable} ${audiowide.variable} font-inter min-h-screen bg-background text-primary w-full`}>
                <Component {...pageProps} />
              </main>
            ) : (
              // Other pages with normal layout
              <main className={`${inter.variable} ${orbitron.variable} ${rajdhani.variable} ${exo2.variable} ${audiowide.variable} font-inter flex min-h-screen bg-background text-primary w-full`}>
                <Sidebar />
                <div className="flex-1 flex flex-col sm:ml-20 border-l border-primary/60">
                  <Header />
                  <Component {...pageProps} />
                  <Footer />
                </div>
              </main>
            )}
          </RainbowKitProvider>
        </QueryClientProvider>
      </WagmiProvider>
    </>
  );
}

export default MyApp;
