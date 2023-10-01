const {createVanillaExtractPlugin} = require('@vanilla-extract/next-plugin');
const withVanillaExtract = createVanillaExtractPlugin();

/** @type {import('next').NextConfig} */
const nextConfig = {
  poweredByHeader: false,
};

module.exports = withVanillaExtract(nextConfig);
