import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: ''
  },
  actions: {
    setToken ({commit}, token) {
      commit('set', {type: 'token', item: token})
    }
  },
  mutations: {
  },
  getters: {
  }
})

export default store
