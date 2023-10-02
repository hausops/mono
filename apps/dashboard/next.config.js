const {createVanillaExtractPlugin} = require('@vanilla-extract/next-plugin');
const withVanillaExtract = createVanillaExtractPlugin();

/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    typedRoutes: true,
  },
  poweredByHeader: false,
};

module.exports = withVanillaExtract(nextConfig);
