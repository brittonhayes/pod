let cssConfig = {};

if (process.env.NODE_ENV == "production") {
  cssConfig = {
    extract: {
      filename: "[name].css",
      chunkFilename: "[name].css",
    },
    loaderOptions: {
      sass: {
        additionalData: `@import "./src/assets/css/buefy.scss";`,
      },
    },
  };
}

module.exports = {
  configureWebpack: {
    output: {
      filename: "[name].js",
    },
    optimization: {
      splitChunks: false,
    },
    resolve: {
      alias: {
        vue$: "vue/dist/vue.esm.js",
        "@": "src",
      },
    },
  },
  css: {
    extract: {
      filename: "[name].css",
      chunkFilename: "[name].css",
    },
    loaderOptions: {
      sass: {
        additionalData: `@import "./src/assets/css/buefy.scss";`,
      },
    },
  },
  devServer: {
    disableHostCheck: true,
  },
  chainWebpack: (config) => {
    let limit = 9999999999999999;
    config.module
      .rule("images")
      .test(/\.(png|gif|jpg)(\?.*)?$/i)
      .use("url-loader")
      .loader("url-loader")
      .tap((options) => Object.assign(options, { limit: limit }));
    config.module
      .rule("fonts")
      .test(/\.(woff2?|eot|ttf|otf|svg)(\?.*)?$/i)
      .use("url-loader")
      .loader("url-loader")
      .options({
        limit: limit,
      });
  },
};
