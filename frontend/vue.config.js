module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
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


