import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Vue3Toastify from 'vue3-toastify'
import { createPinia } from 'pinia'

const app = createApp(App).
    use(router).
    use(Vue3Toastify, {
        autoClose: 3000,
    });

const pinia = createPinia()
app.use(pinia);
app.mount('#app')
