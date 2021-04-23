/**
 * @type {import('@vue/cli-service').ProjectOptions}
 */
 module.exports = {
  outputDir: '../../../web/assets',
  devServer: {
    writeToDisk: true,
    proxy: {
      '/': {
        target: 'http://localhost:5000',
        ws: true,
      },
    }
  },
  // disable hashes in filenames
  filenameHashing: false,
  // delete HTML related webpack plugins
  chainWebpack: config => {
    config.plugins.delete('html')
    config.plugins.delete('preload')
    config.plugins.delete('prefetch')
  }
 }