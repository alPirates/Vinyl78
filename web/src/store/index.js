import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: localStorage.getItem('token') || '',
    role: 'client',
    drawer: false,
    loader: true
  },
  actions: {
  },
  mutations: {
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
    getRole: state => state.role
  }
})

export default store
