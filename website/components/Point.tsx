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
    return <div className="text-center p-8">Loading...</div>;
  }

  if (error) {
    return <div className="text-center p-8 text-red-500">{error}</div>;
  }

  return (
    <div className="space-y-6 max-w-2xl mx-auto p-4">
      {/* User's points */}
      <div className="bg-white rounded-xl shadow-lg p-8">
        <div className="flex flex-col items-center">
          <div className="flex justify-center items-center">
            <img src="/fluxus.svg" alt="Point" width={100} height={100} />
          </div>
          
          {/* User's points */}
          <h3 className="text-2xl font-semibold text-gray-800">My Points</h3>
          <p className="text-5xl font-bold text-primary mt-4">{rewards?.totalRewards || 0}</p>
        </div>
        
        {/* User's points on three chains */}
        <div className="grid grid-cols-3 gap-4 mt-6">
        <div className="bg-gray-50 rounded-lg p-4">
            <div className="flex flex-col items-center">
              <div className="bg-purple-100 p-2 rounded-full mb-2">
                {/* <svg className="w-6 h-6 text-purple-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg> */}
                <img src={customMonadTestnet.iconUrl} alt="Monad" width={24} height={24} />
              </div>
              <p className="text-sm text-gray-600">Monad</p>
              <p className="text-lg font-bold text-gray-800">{rewards?.chains.monad || 0}</p>
            </div>
          </div>

          <div className="bg-gray-50 rounded-lg p-4">
            <div className="flex flex-col items-center">
              <div className="bg-blue-100 p-2 rounded-full mb-2">
                {/* <svg className="w-6 h-6 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg> */}
                <img src={customMainnet.iconUrl} alt="Ethereum" width={24} height={24} />
              </div>
              <p className="text-sm text-gray-600">Ethereum</p>
              <p className="text-lg font-bold text-gray-800">
                <span className="text-xs font-medium text-black">Coming Soon</span>
              </p>
            </div>
          </div>

          <div className="bg-gray-50 rounded-lg p-4">
            <div className="flex flex-col items-center">
              <div className="bg-yellow-100 p-2 rounded-full mb-2">
                {/* <svg className="w-6 h-6 text-yellow-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 11c0 3.517-1.009 6.799-2.753 9.571m-3.44-2.04l.054-.09A13.916 13.916 0 008 11a4 4 0 118 0c0 1.017-.07 2.019-.203 3m-2.118 6.844A21.88 21.88 0 0015.171 17m3.839 1.132c.645-2.266.99-4.659.99-7.132A8 8 0 008 4.07M3 15.364c.64-1.319 1-2.8 1-4.364 0-1.457.39-2.823 1.07-4" />
                </svg> */}
                <img src={customBerachainTestnet.iconUrl} alt="Berachain" width={24} height={24} />
              </div>
              <p className="text-sm text-gray-600">Berachain</p>
              <p className="text-lg font-bold text-gray-800">
              <span className="text-xs font-medium text-black">Coming Soon</span>
              </p>
            </div>
          </div>
        </div>
      </div>

      {/* User's recent transactions */}
      <div className="bg-white rounded-xl shadow-lg">
        <div className="p-4 border-b">
          <h3 className="text-lg font-semibold text-gray-800">Recent Transactions</h3>
        </div>
        <div className="divide-y">
          {transactions.map((tx) => (
            <div key={tx.tx_hash} className="p-3 flex items-center justify-between hover:bg-gray-50">
              <div className="flex items-center space-x-3">
                <div className={`p-2 rounded-full ${
                  tx.tx_type === 1 ? 'bg-purple-100' : 
                  tx.tx_type === 2 ? 'bg-green-100' : 'bg-red-100'
                }`}>
                  <svg 
                    className={`w-5 h-5 ${
                      tx.tx_type === 1 ? 'text-purple-500' : 
                      tx.tx_type === 2 ? 'text-green-500' : 'text-red-500'
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
                  <h4 className="text-sm font-medium text-gray-800">
                    {tx.tx_type === 1 ? 'Mint' : tx.tx_type === 2 ? 'Buy' : 'Sell'} TokenID #{tx.token_id}
                  </h4>
                  <p className="text-xs text-gray-500">
                    for {(Number(tx.price) / 1e18).toFixed(4)} ETH
                  </p>
                </div>
              </div>
              <div className="text-xs text-gray-500">
                {new Date(tx.created_at).toLocaleDateString()}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}