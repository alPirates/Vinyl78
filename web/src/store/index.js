import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: ''
  },
  actions: {
  },
  mutations: {
    setToken (state, token) {
      state.token = token
    }
  },
  getters: {
    getToken (state) {
      return state.token
    }
  }
})

export default store
