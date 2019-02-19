import Vue from 'vue'
import Router from 'vue-router'

import Main from '@/components/pages/Main'
import Form from '@/components/pages/Form'
import Admin from '@/components/pages/Admin'
import AdminCategory from '@/components/pages/admin/AdminCategory'
import AdminSeo from '@/components/pages/admin/AdminSeo'
import Stickers from '@/components/pages/Stickers'
import StickerView from '@/components/pages/StickerView'
import Settings from '@/components/pages/Settings'

Vue.use(Router)

const ADMIN_SECURE_LINK = 'nXnA1E1jzYzSDfqvWDkez0DDlxg52QYD'

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
      path: `/${ADMIN_SECURE_LINK}`,
      name: 'Admin',
      component: Admin
    },
    {
      path: '/admin/category',
      name: 'Admin category',
      component: AdminCategory
    },
    {
      path: '/admin/seo',
      name: 'Admin seo',
      component: AdminSeo
    },
    {
      path: '/sticker/:id',
      name: 'Sticker full view',
      component: StickerView
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
