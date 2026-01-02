'use client'
import React from 'react';

export interface TraitItem {
  trait_type?: string;
  value?: string | number;
}

export default function NFTTraits({ traits }: { traits?: TraitItem[] }) {
  const hasTraits = Array.isArray(traits) && traits.length > 0;
  return (
    <div className="rounded-xl border border-neutral-800 bg-neutral-900 p-4 h-72 overflow-y-auto">
      {/* <div className="text-sm font-semibold text-gray-300 mb-3">TRAITS</div> */}
      {!hasTraits ? (
        <div className="w-full h-full flex items-center justify-center text-gray-500">No traits</div>
      ) : (
        <div className="grid grid-cols-2 sm:grid-cols-3 gap-3">
          {traits!.map((t, idx) => (
            <div key={idx} className="rounded-lg border border-neutral-800 bg-neutral-800/50 p-3">
              <div className="text-xs text-gray-400 mb-1">{t.trait_type || 'Trait'}</div>
              <div className="text-sm text-white break-all">{String(t.value ?? '-') }</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}


