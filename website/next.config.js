/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  // Remove deprecated experimental flags for Next 14+
  experimental: {},
  // Ensure build does not fail due to ESLint in CI
  eslint: {
    ignoreDuringBuilds: true,
  },
  // Explicitly disable SWC for minification due to environment binary issues
  swcMinify: false,
  images: {
    domains: ['ipfs.io', 'red-naval-coyote-720.mypinata.cloud'],
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'ipfs.io',
      }
      ,
      {
        protocol: 'https',
        hostname: 'red-naval-coyote-720.mypinata.cloud',
      }
    ],
  },
  // Workaround for third-party worker asset minification issues on some build platforms
  webpack: (config, { dev }) => {
    if (!dev) {
      // Disable JS minification to avoid Terser parsing ESM worker assets (e.g., HeartbeatWorker)
      // If you prefer keeping minification, remove this and investigate the offending dependency.
      if (config.optimization) {
        config.optimization.minimize = false;
      }
    }
    return config;
  },
};

module.exports = nextConfig;
