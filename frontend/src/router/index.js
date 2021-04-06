import Vue from 'vue'
import VueRouter from 'vue-router'
import NotFoundComponent from '@/components/NotFoundComponent.vue'

Vue.use(VueRouter)

var routes = []

//メンテナンス中の場合はメンテナンスページに飛ばし、それ以外には行かせない
if (process.env.VUE_APP_IS_MAINTENANCE == "true") {
  routes = [{
    path: '/',
    name: 'maintenance',
    component: () => import('../views/Maintenance.vue')
  }, ]

} else {

  routes = [{
      path: '/',
      name: 'Top',
      component: () => import('../views/Top.vue')
    },
    {
      path: '/about',
      name: 'About',
      component: () => import('../views/About.vue')
    },
    {
      path: '/privacy',
      component: () => import('../views/PrivacyPolicy.vue'),
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/home',
      component: () => import('../views/Home.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/new',
      component: () => import('../views/Create.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/register',
      component: () => import('../views/Register.vue'),
    },
    {
      path: '/verify',
      component: () => import('../views/Verify.vue'),
    },
    {
      path: '/maintenance',
      component: () => import('../views/Maintenance.vue'),
    },
    {
      path: '/user',
      component: () => import('../views/User.vue'),
      children: [{
          // /user/:id/profile がマッチした時に
          // UserProfile は User の <router-view> 内部で描画されます
          path: 'profile',
          component: () => import('@/components/user/Profile.vue'),
        },
        {
          path: '',
          component: () => import('@/components/user/Main.vue')
        },
      ]
    },
    //ブラウザバック対策
    {
      path: '*',
      name: "404",
      component: NotFoundComponent
    }
  ]

}

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
      Vue.$cookies.set("msg", "続けるにはログインが必要です。", {
        expires: "1H",
      })
      next({
        path: '/login',
        query: {
          redirect: to.fullPath
        }
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

function isLogin() {
  return Vue.$cookies.get("token") != null
}


export default router