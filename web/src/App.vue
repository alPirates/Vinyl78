<template lang="pug">
  v-app
    Sidebar
    Header
    v-container(fluid)
      router-view

</template>

<script>
import Header from '@/components/Navigation/Header.vue'
import Sidebar from '@/components/Navigation/Sidebar.vue'

export default {
  name: 'App',
  data: () => {
    return {
      drawer: false
    }
  },
  mounted () {
    const token = localStorage.getItem('token')
    if (token) {
      this.COMMIT('setToken', token)
    }
    const role = localStorage.getItem('role')
    if (role) {
      this.COMMIT('setRole', role)
    }
  },
  components: {
    Header,
    Sidebar
  }
}
</script>
