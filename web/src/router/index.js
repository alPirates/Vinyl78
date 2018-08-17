import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/pages/Main'
import Form from '@/components/pages/Form'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main
    },
    {
      path: '/form',
      name: 'Form',
      component: Form
    }
  ]
})
