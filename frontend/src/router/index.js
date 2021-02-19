import Vue from 'vue'
import VueRouter from 'vue-router'
import NotFoundComponent from '@/components/NotFoundComponent.vue'
import Main from '@/components/mypage/Main.vue'

Vue.use(VueRouter)


const routes = [
  {
    path: '/',
    name: 'Top',
    component: () => import('../views/Home.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/space/:id', component: () => import('../views/Space.vue'),
    meta: { requiresAuth: true },
      children: [
        // /user/:id がマッチした時に
        // UserHome は User の <router-view> 内部で描画されます
        { path: '', component: Main },

        // 他のサブルートも同様に...
      ]
  },
  {
    path: '/new',
    component: () => import('../views/Create.vue'),
    meta: { requiresAuth: true },
  },
  //ブラウザバック対策
  { path: '*', component: NotFoundComponent }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // このルートはログインされているかどうか認証が必要です。
    // もしされていないならば、ログインページにリダイレクトします。
    if (!isLogin()) {

      ////////////////////////////////
      //
      // TODO:エラーメッセージを作る
      //
      /////////////////////////////////
      Vue.$cookies.set("msg", "続けるにはログインが必要です。", 3600, "/", "localhost", true, "None")
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

function isLogin(){
  return Vue.$cookies.get("token")
}


export default router
