<template lang="pug">
  v-app
    Sidebar
    Header
    v-content
      div.wrapper
        router-view
    Loader

</template>

<script>
import Header from '@/components/Navigation/Header.vue'
import Sidebar from '@/components/Navigation/Sidebar.vue'
import Footer from '@/components/Navigation/Footer.vue'
import Loader from '@/components/Utils/Loader.vue'

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
    Sidebar,
    Loader,
    Footer
  }
}
</script>
