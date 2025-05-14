// /** @type {import('next').NextConfig} */
// const nextConfig = {
//   images: {
//     domains: ['localhost:3100', 'picsum.photos', 'loremfaces.net'],
//   },
// };

// export default nextConfig;

/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    unoptimized: true,
    remotePatterns: [
      {
        protocol: "http",
        hostname: "localhost",
        port: "3000",
        pathname: "/**",
      },
      {
        protocol: "http",
        hostname: "localhost",
        port: "8080",
        pathname: "/**",
      },
      {
        protocol: "http",
        hostname: "localhost",
        port: "3100",
        pathname: "/**",
      },
      {
        protocol: "https",
        hostname: "picsum.photos",
        port: "",
        pathname: "/**",
      },
      {
        protocol: "https",
        hostname: "loremfaces.net",
        port: "",
        pathname: "/**",
      },
    ],
  },
  env: {
    BACKEND_URL: process.env.BACKEND_URL,
    MEDIA_URL: process.env.MEDIA_URL,
    FAKE_URL: process.env.FAKE_URL,
  },
};

export default nextConfig;
