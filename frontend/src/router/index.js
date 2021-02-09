import Vue from 'vue'
import VueRouter from 'vue-router'
import NotFoundComponent from '@/components/NotFoundComponent.vue'
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
    path: '/mypage',
    name: 'Mypage',
    component: () => import('../views/Mypage.vue')
  },
  //ブラウザバック対策
  { path: '*', component: NotFoundComponent }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
