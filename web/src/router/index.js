import Vue from 'vue'
import Router from 'vue-router'

import Main from '@/components/pages/Main'
import Form from '@/components/pages/Form'
import Admin from '@/components/pages/Admin'
import AdminCategory from '@/components/pages/admin/AdminCategory'
import Stickers from '@/components/pages/Stickers'
import Settings from '@/components/pages/Settings'

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
    },
    {
      path: '/admin',
      name: 'Admin',
      component: Admin
    },
    {
      path: '/admin/category',
      name: 'Admin category',
      component: AdminCategory
    },
    {
      path: '/category/:id',
      name: 'Stickers',
      component: Stickers
    },
    {
      path: '/settings',
      name: 'Settings',
      component: Settings
    }
  ]
})
