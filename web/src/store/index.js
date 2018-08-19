import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: '',
    drawer: false
  },
  actions: {
  },
  mutations: {
    setToken (state, token) {
      state.token = token
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
    }
  }
})

export default store
