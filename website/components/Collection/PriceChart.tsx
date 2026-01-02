"use client";
import { useEffect, useState } from "react";
import { ethers } from "ethers";
import { Line } from 'react-chartjs-2';

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js';
import { Collection } from "@/src/api/types";
import { cn } from "@/lib/utils";
import { getChainSymbol, PrimaryColor } from "@/src/utils";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

interface PriceChartProps {
  collection: {
    max_supply?: number;
    initial_price?: string;
    current_supply?: number;
    floor_price?: string;
  };
  chainId: number;
  className?: string;
}

export default function PriceChart({ collection, chainId, className }: PriceChartProps) {
  const [priceChartData, setPriceChartData] = useState<{
    labels: string[];
    datasets: any[];
  }>({
    labels: [],
    datasets: []
  });

  useEffect(() => {
    if (!collection ||
      collection.max_supply === undefined ||
      collection.initial_price === undefined) {
      return;
    }

    const points = 50;
    const labels: string[] = [];
    const prices: string[] = [];

    for (let i = 0; i <= points; i++) {
      const supply = Math.floor((Number(collection.max_supply) * i) / points);
      labels.push(supply.toString());

      let price;
      if (supply === 0) {
        price = BigInt(collection.initial_price);
      } else {
        const sqrt1 = Math.sqrt(100 * supply / Number(collection.max_supply));
        const sqrt2 = Math.sqrt(10000 * supply * supply / (Number(collection.max_supply) * Number(collection.max_supply)));
        const multiplier = Math.floor(sqrt1 * sqrt2);
        price = BigInt(collection.initial_price) +
          (BigInt(collection.initial_price) * BigInt(2) * BigInt(multiplier));
      }
      prices.push(ethers.formatEther(price.toString()));
    }

    setPriceChartData({
      labels,
      datasets: [
        {
          label: 'Mint Price',
          data: prices,
          borderColor: '#fff',
          tension: 0.1,
          borderWidth: 2,
          fill: false,
          pointRadius: 1,
          pointBorderWidth: 1,
        },
        {
          label: 'Current Position',
          data: Array(points + 1).fill(null).map((_, i) => {
            if (collection.current_supply === undefined || collection.floor_price === undefined) {
              return null;
            }
            const supplyAtPoint = (Number(collection.max_supply) * i) / points;
            const currentSupplyPoint = Math.floor((Number(collection.current_supply) * points) / Number(collection.max_supply)) * Number(collection.max_supply) / points;

            return Math.abs(currentSupplyPoint - supplyAtPoint) < 0.1
              ? Number(ethers.formatEther(collection.floor_price.toString()))
              : null;
          }),
          pointBackgroundColor: '#fff',
          pointRadius: 6,
          showLine: false,
          pointStyle: 'circle',
          borderWidth: 2,
          borderColor: PrimaryColor,
        }
      ]
    });
  }, [collection]);

  return (
    <div className={cn("w-full h-full flex flex-col gap-3", className)}>
      <div className="flex items-center gap-6 text-[11px] uppercase tracking-[0.3em] text-secondary">
        <span className="flex items-center gap-2 text-secondary">
          <span className="w-6 h-[2px] bg-white" />
          Mint Price
        </span>
        <span className="flex items-center gap-2 text-secondary">
          <span className="w-6 h-[2px]" style={{ backgroundColor: PrimaryColor }} />
          Current Position
        </span>
      </div>
      <div className="flex-1">
        <Line
          data={priceChartData}
          options={{
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
              legend: {
                display: false,
              },
              tooltip: {
                displayColors: false,
                backgroundColor: 'rgba(0,0,0,0.7)',
                titleColor: '#fff',
                bodyColor: '#fff',
              },
            },
            elements: {
              point: {
                radius: 0,
              },
            },
            scales: {
              y: {
                type: 'linear',
                min: 0,
                ticks: {
                  color: '#aaa',
                  callback: (value) => Number(value).toFixed(2),
                },
                title: {
                  display: true,
                  text: `PRICE (${getChainSymbol(chainId)})`,
                  color: '#fff',
                },
                grid: {
                  color: 'rgba(22, 163, 74, 0.2)',
                  // drawBorder: false,
                },
              },
              x: {
                ticks: {
                  color: '#aaa',
                },
                title: {
                  display: true,
                  color: '#fff',
                  text: 'CURRENT SUPPLY',
                },
                grid: {
                  color: 'rgba(22, 163, 74, 0.15)',
                  // drawBorder: false,
                },
              },
            },
          }}
        />
      </div>
    </div>
  );
}
