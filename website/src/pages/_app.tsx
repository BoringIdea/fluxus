import '../styles/globals.css';
import { Inter } from 'next/font/google'
import '@rainbow-me/rainbowkit/styles.css';
import type { AppProps } from 'next/app';
import { useRouter } from 'next/router';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { WagmiProvider } from 'wagmi';
import { RainbowKitProvider, lightTheme } from '@rainbow-me/rainbowkit';

import { config } from '../wagmi';

const client = new QueryClient();

// Main body font - clean and readable
const inter = Inter({
  subsets: ['latin'],
  weight: ['400', '500', '600', '700'],
  display: 'swap',
  variable: '--font-inter',
});

import Head from 'next/head';
import Header from '@/components/Header';
import Sidebar from '@/components/SideBar';
import Footer from '@/components/Footer';

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
          <RainbowKitProvider locale='en' theme={lightTheme({
            accentColor: '#16A34A',
            accentColorForeground: '#ffffff'
          })}>
            {isHomePage ? (
              <main className={`${inter.variable} flux-shell min-h-screen w-full text-foreground`}>
                <Component {...pageProps} />
              </main>
            ) : (
              <main className={`${inter.variable} flux-shell flex min-h-screen w-full text-foreground`}>
                <Sidebar />
                <div className="flex flex-1 flex-col sm:ml-20 border-l border-black/10">
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
