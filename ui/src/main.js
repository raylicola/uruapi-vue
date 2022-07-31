import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import router from './router'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:8888'
axios.defaults.withCredentials = true

loadFonts()

createApp(App)
  .use(vuetify)
  .use(router)
  .mount('#app')
