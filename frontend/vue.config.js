module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  publicPath: "./",
  pages: {
    index: {
      entry: "src/main.js",
      title: "yource | ポートフォリオを簡単に公開する",
    }
  },
  devServer: {
    port: 9092,
    disableHostCheck: true,
    watchOptions: {
      poll: true
    }
  },
  pluginOptions: {
    browserSync: {
      host: 'localhost',
      port: 3000,
      server: { baseDir: ['src'] }
    }
  }
}


