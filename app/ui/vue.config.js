if (process.env.NODE_ENV === 'production') {
  process.env.VUE_APP_API_URL = 'https://secret.enazarov.online';
} else {
  process.env.VUE_APP_API_URL = 'http://localhost:8003';
}

const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})
