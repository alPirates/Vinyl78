import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: localStorage.getItem('token') || '',
    alert: {
      show: false,
      message: 'kek'
    },
    role: 'client',
    drawer: false,
    loader: false
  },
  actions: {
  },
  mutations: {
    setAlert (state, message, time = 3000) {
      if (state.alert.show) {
        state.alert.show = false
      }
      state.alert.message = message
      state.alert.show = true
      setTimeout(() => {
        state.alert.show = false
      }, time)
    },
    setLoader (state, loader) {
      state.loader = loader
    },
    setToken (state, token) {
      state.token = token
    },
    setRole (state, role) {
      state.role = role
    },
    setDrawer (state, drawer) {
      state.drawer = drawer
    }
  },
  getters: {
    getToken: state => state.token,
    getDrawer: state => state.drawer,
    getRole: state => state.role,
    getAlert: state => state.alert
  }
})

export default store
