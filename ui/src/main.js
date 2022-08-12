import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/router'
import store from '@/store'
import axios from 'axios'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'

axios.defaults.baseURL = 'https://uruapi-api.herokuapp.com/'
axios.defaults.withCredentials = true

loadFonts()

createApp(App)
  .use(vuetify)
  .use(store)
  .use(router)
  .mount('#app')
