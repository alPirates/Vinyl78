// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import moment from 'moment'
import * as R from 'ramda'

import Api from './plugins/api'
import Filters from './plugins/filters'
import State from './plugins/state'

import Vuetify from 'vuetify'

import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'

Vue.config.productionTip = false

Vue.prototype.R = R

Vue.use(Api, { store, R })
Vue.use(Filters, { moment })
Vue.use(Vuetify)
Vue.use(State)

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
