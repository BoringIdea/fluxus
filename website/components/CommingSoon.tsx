import React from 'react';

const ComingSoon: React.FC = () => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-black">
      <div className="text-center p-8">
        <h1 className="text-4xl md:text-6xl font-bold text-green-600 mb-4">
          Coming Soon
        </h1>
        <p className="text-xl text-green-500 mb-8">
          We are working on it, please stay tuned!
        </p>
      </div>
    </div>
  );
};

export default ComingSoon;
