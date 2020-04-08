/* eslint-disable */
// const dotenvLoad = require('dotenv-load');
const Dotenv = require("dotenv-webpack");
const path = require("path");

// dotenvLoad();

const withSass = require("@zeit/next-sass");

module.exports = withSass({
  cssModules: true,
  devIndicators: {
    autoPrerender: false
  },
  webpack: config => {
    config.plugins = config.plugins || [];

    const envPlugin =
      process.env.NODE_ENV === "prod"
        ? new Dotenv({
            path: path.join(__dirname, ".env.prod"),
            systemvars: true
          })
        : new Dotenv({
            path: path.join(__dirname, ".env.dev"),
            systemvars: true
          });

    config.plugins = [
      ...config.plugins,
      // Read the .env file
      envPlugin
    ];
    // Fixes npm packages that depend on `fs` module
    config.node = {
      fs: "empty"
    };

    return config;
  }
});
