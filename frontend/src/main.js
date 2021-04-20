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
Vue.prototype.$front = process.env.VUE_APP_FRONT_URL
Vue.prototype.$setTitle = (title, description) => {
  document.title = title
  document.querySelector("meta[property='og:title']").setAttribute('content', title)
  document.querySelector("meta[property='og:description']").setAttribute('content', description)
}

//urlを判断して、もしサブドメイン付きでアクセスしてきたらスペースへ直接飛ばす
var hostStrings = window.location.host.split('.');
//'.'でホスト文字列を分けて、一番始めがyourceでなければ→サブドメイン付きでアクセスしてくれば
//変数に記録しておき、後ほどの処理で使う
if (hostStrings[0] != "yource"){
  Vue.prototype.$external = true
}else {
  Vue.prototype.$external = false
}
Vue.prototype.$host = hostStrings;


//カスタムスクロールディレクティブの作成
Vue.directive('scroll', {
  inserted: function (el, binding) {
      let f = function (evt) {
          if (binding.value(evt, el)) {
              window.removeEventListener('scroll', f)
          }
      }
      window.addEventListener('scroll', f)
  }
})