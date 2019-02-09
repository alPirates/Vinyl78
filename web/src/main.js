// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'

// Utils zone
import moment from 'moment'
import * as R from 'ramda'
import draggable from 'vuedraggable'
import InfiniteLoading from 'vue-infinite-loading'

import Api from './plugins/api'
import Filters from './plugins/filters'
import State from './plugins/state'

import 'fotorama/fotorama.css'
import 'fotorama/fotorama.js'

import VueCarousel from 'vue-carousel'

import Vuetify from 'vuetify'
import SvgFiller from 'vue-svg-filler'

import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@fortawesome/fontawesome-free/css/all.css'

Vue.config.productionTip = false

// todo remove from latest commits
Vue.prototype.R = R

Vue.use(Vuetify, {
  theme: {
    primary: '#2a2a2a',
    secondary: '#424242',
    accent: '#C62828',
    error: '#B71C1C',
    warning: '#FDD835',
    info: '#80CBC4',
    success: '#5f6574',
    custom: '#e9e9e9'
  },
  iconfont: 'fa'
})

// eslint-disable-next-line
import '@/assets/main.scss'

Vue.use(VueCarousel)
Vue.use(Api, { store, R })
Vue.use(Filters, { moment })
Vue.use(State, { store })
Vue.component('draggable', draggable)
Vue.component('InfiniteLoading', InfiniteLoading)
Vue.component('svg-filler', SvgFiller)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: {
    App
  },
  template: '<App/>'
})
