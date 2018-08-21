import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: '',
    role: 'client',
    drawer: false
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
    getToken (state) {
      return state.token
    },
    getDrawer (state) {
      return state.drawer
    },
    getRole (state) {
      return state.role
    }
  }
})

export default store
