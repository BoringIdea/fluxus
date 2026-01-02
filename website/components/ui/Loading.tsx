import React from 'react';

export const Loading: React.FC<{ className?: string }> = ({ className = '' }) => (
  <div className={`flex flex-col items-center justify-center py-8 ${className}`} role="status">
    <svg
      className="animate-spin h-10 w-10 text-green-500"
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
    >
      <circle
        className="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        strokeWidth="4"
      />
      <path
        className="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"
      />
    </svg>
    <span className="mt-4 text-green-600 text-sm font-medium">Loading...</span>
  </div>
);

export default Loading; 