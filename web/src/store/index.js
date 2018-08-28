import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: localStorage.getItem('token') || '',
    alert: {
      show: true,
      message: ''
    },
    role: 'client',
    drawer: false,
    loader: false
  },
  actions: {
  },
  mutations: {
    setAlert (state, show, message) {
      state.message = message
      state.alert = alert
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
