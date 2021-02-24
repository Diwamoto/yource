import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import vuetify from './plugins/vuetify';
import router from './router'
import axios from 'axios'

Vue.config.productionTip = false
Vue.use(require('vue-cookies'))
Vue.use(require('vue-moment'));
Vue.use(VueRouter)

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')

Vue.prototype.$http = axios
Vue.prototype.$api = process.env.VUE_APP_API_URL