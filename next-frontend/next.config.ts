/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'images.unsplash.com',
        port: '',
        pathname: '/**',
      },
    ],
  },
  // ใช้ standalone เมื่อ NODE_ENV=production และ DOCKER_BUILD=true
  output: (process.env.NODE_ENV === 'production' && process.env.DOCKER_BUILD === 'true') ? 'standalone' : undefined,
  
  // Hot reload configuration for Docker development
  ...(process.env.NODE_ENV === 'development' && {
    webpack: (config: any, { dev }: { dev: boolean }) => {
      if (dev) {
        config.watchOptions = {
          poll: 1000,
          aggregateTimeout: 300,
        }
      }
      return config
    },
  }),
}

export default nextConfig
