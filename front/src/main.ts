import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router'
import { createPinia } from 'pinia'

import 'vfonts/Lato.css'
import './styles/common.css'
import { i18n } from './locale/index'

import axios from 'axios'

axios.defaults.baseURL = (import.meta as any).env.VITE_APP_BACKEND_BASE_URL;
axios.defaults.withCredentials = true

createApp(App).use(router).use(i18n).use(createPinia()).mount('#app')