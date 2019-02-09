<template lang="pug">
  v-app(
    sidebar-under-toolbar
  )
    Sidebar
    Header
    v-content
      div.wrapper
        router-view
      Footer
    Loader
    Alert

</template>

<script>
import Header from '@/components/Navigation/Header'
import Sidebar from '@/components/Navigation/Sidebar'
import Footer from '@/components/Navigation/Footer'
import Loader from '@/components/Utils/Loader'
import Alert from '@/components/Utils/Alert'

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
    Footer,
    Alert
  }
}
</script>

<style>
</style>
