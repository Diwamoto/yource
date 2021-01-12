import Vue from 'vue'
import Router from 'vue-router'
import test1 from '@/components/test1'
import test2 from '@/components/test2'
import test3 from '@/components/test3'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: test1
    },
    {
      path: '/test2',
      component: test2
    },
    {
      path: '/test3',
      component: test3
    }
  ]
})
