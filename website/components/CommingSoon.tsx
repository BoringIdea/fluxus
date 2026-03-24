import React from 'react';

const ComingSoon: React.FC = () => {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="border border-black/10 bg-[color:var(--bg-surface)] p-8 text-center">
        <p className="flux-kicker mb-3">Status</p>
        <h1 className="mb-4 font-heading text-[44px] leading-none text-[color:var(--text-primary)] md:text-[68px]">
          Coming Soon
        </h1>
        <p className="text-base text-[color:var(--text-secondary)] md:text-lg">
          We are working on it, please stay tuned!
        </p>
      </div>
    </div>
  );
};

export default ComingSoon;
