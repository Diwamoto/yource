const fs = require('fs')
module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  configureWebpack: {
    devtool: 'source-map'
  },
  publicPath: "./",
  pages: {
    index: {
      entry: "src/main.js",
      title: "yource | ポートフォリオを簡単に公開する",
    }
  },
  // devServer: {
  //   port: 9092,
  //   disableHostCheck: true,
  //   https: {
  //     key: fs.readFileSync('./keys/server.key'),
  //     cert: fs.readFileSync('./keys/server.crt'),
  //   },
  //   public: 'https://localhost:9092/',
  //   watchOptions: {
  //     poll: true
  //   }
  // },
  pluginOptions: {
    // browserSync: {
    //   host: 'localhost',
    //   port: 3000,
    //   server: { baseDir: ['src'] }
    // }
  }
}