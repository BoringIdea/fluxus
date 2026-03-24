'use client'
import React from 'react';

export interface TraitItem {
  trait_type?: string;
  value?: string | number;
}

export default function NFTTraits({ traits }: { traits?: TraitItem[] }) {
  const hasTraits = Array.isArray(traits) && traits.length > 0;
  return (
    <div className="h-72 overflow-y-auto border border-black/10 bg-[color:var(--bg-surface)] p-4">
      {!hasTraits ? (
        <div className="flex h-full w-full items-center justify-center font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">No traits</div>
      ) : (
        <div className="grid grid-cols-2 sm:grid-cols-3 gap-3">
          {traits!.map((t, idx) => (
            <div key={idx} className="border border-black/10 bg-[color:var(--bg-muted)] p-3">
              <div className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">{t.trait_type || 'Trait'}</div>
              <div className="break-all text-sm text-[color:var(--text-primary)]">{String(t.value ?? '-') }</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

