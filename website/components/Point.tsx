'use client'

import { useEffect, useState } from 'react';
import { customMonadTestnet, customMainnet, customBerachainTestnet } from '@/src/wagmi';
import { useAccount } from 'wagmi';

interface Reward {
  totalRewards: number;
  chains: {
    ethereum: number | 0;
    monad: number | 0;
    berachain: number | 0;
  };
}

interface Transaction {
  tx_hash: string;
  chain: number;
  tx_type: number;
  sender: string;
  price: number;
  token_id: number;
  creator_fee: number;
  created_at: string;
}

export default function Point() {
  const { address } = useAccount();
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
  const [rewards, setRewards] = useState<Reward | null>(null);
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        if (!address) {
          console.log('User address not found');
          setLoading(false);
          return;
        }
        const userAddress = address; 

        const rewardsResponse = await fetch(`${backendUrl}/rewards/${userAddress}`);
        const rewardsData = await rewardsResponse.json();
        console.log('rewardsData', rewardsData);

        const txsResponse = await fetch(
          `${backendUrl}/txs/${userAddress}?page=1&pageSize=10`
        );
        const txsData = await txsResponse.json();
        console.log('txsData', txsData);

        setRewards({
          totalRewards: rewardsData,
          chains: {
            ethereum: 0,
            monad: rewardsData,
            berachain: 0,
          },
        });
        setTransactions(txsData);
      } catch (err) {
        setError('Failed to fetch data');
        console.error('Error fetching data:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return <div className="p-8 text-center font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Loading...</div>;
  }

  if (error) {
    return <div className="p-8 text-center text-red-600">{error}</div>;
  }

  return (
    <div className="mx-auto max-w-2xl space-y-6 p-4">
      <div className="border border-black/10 bg-[color:var(--bg-surface)] p-8">
        <div className="flex flex-col items-center">
          <div className="flex justify-center items-center">
            <img src="/fluxus.svg" alt="Point" width={100} height={100} />
          </div>

          <p className="flux-kicker mb-3 mt-6">Rewards</p>
          <h3 className="font-heading text-[34px] leading-none text-[color:var(--text-primary)]">My Points</h3>
          <p className="mt-4 font-heading text-[56px] leading-none text-[color:var(--text-primary)]">{rewards?.totalRewards || 0}</p>
        </div>

        <div className="mt-6 grid grid-cols-3 gap-4">
        <div className="border border-black/10 bg-[color:var(--bg-muted)] p-4">
            <div className="flex flex-col items-center">
              <div className="mb-2 border border-black/10 bg-[color:var(--bg-surface)] p-2">
                <img src={customMonadTestnet.iconUrl} alt="Monad" width={24} height={24} />
              </div>
              <p className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Monad</p>
              <p className="mt-2 font-heading text-[24px] leading-none text-[color:var(--text-primary)]">{rewards?.chains.monad || 0}</p>
            </div>
          </div>

          <div className="border border-black/10 bg-[color:var(--bg-muted)] p-4">
            <div className="flex flex-col items-center">
              <div className="mb-2 border border-black/10 bg-[color:var(--bg-surface)] p-2">
                <img src={customMainnet.iconUrl} alt="Ethereum" width={24} height={24} />
              </div>
              <p className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Ethereum</p>
              <p className="mt-2 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-primary)]">
                <span>Coming Soon</span>
              </p>
            </div>
          </div>

          <div className="border border-black/10 bg-[color:var(--bg-muted)] p-4">
            <div className="flex flex-col items-center">
              <div className="mb-2 border border-black/10 bg-[color:var(--bg-surface)] p-2">
                <img src={customBerachainTestnet.iconUrl} alt="Berachain" width={24} height={24} />
              </div>
              <p className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Berachain</p>
              <p className="mt-2 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-primary)]">
              <span>Coming Soon</span>
              </p>
            </div>
          </div>
        </div>
      </div>

      <div className="border border-black/10 bg-[color:var(--bg-surface)]">
        <div className="border-b border-black/10 p-4">
          <h3 className="font-heading text-[28px] leading-none text-[color:var(--text-primary)]">Recent Transactions</h3>
        </div>
        <div className="divide-y">
          {transactions.map((tx) => (
            <div key={tx.tx_hash} className="flex items-center justify-between p-3 hover:bg-[color:var(--bg-muted)]">
              <div className="flex items-center space-x-3">
                <div className={`border border-black/10 p-2 ${
                  tx.tx_type === 1 ? 'bg-[color:var(--bg-muted)]' : 
                  tx.tx_type === 2 ? 'bg-[color:var(--color-primary)]/10' : 'bg-[color:var(--fg-strong)]/6'
                }`}>
                  <svg 
                    className={`w-5 h-5 ${
                      tx.tx_type === 1 ? 'text-[color:var(--text-primary)]' : 
                      tx.tx_type === 2 ? 'text-[color:var(--color-primary)]' : 'text-[color:var(--fg-strong)]'
                    }`} 
                    fill="none" 
                    viewBox="0 0 24 24" 
                    stroke="currentColor"
                  >
                    {tx.tx_type === 1 ? (
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                    ) : tx.tx_type === 2 ? (
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                    ) : (
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 12H4" />
                    )}
                  </svg>
                </div>
                <div>
                  <h4 className="text-sm text-[color:var(--text-primary)]">
                    {tx.tx_type === 1 ? 'Mint' : tx.tx_type === 2 ? 'Buy' : 'Sell'} TokenID #{tx.token_id}
                  </h4>
                  <p className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                    for {(Number(tx.price) / 1e18).toFixed(4)} ETH
                  </p>
                </div>
              </div>
              <div className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                {new Date(tx.created_at).toLocaleDateString()}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
