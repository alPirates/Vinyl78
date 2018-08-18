<template lang="pug">
  v-app
    v-navigation-drawer(
      v-model="drawer"
      :clipped="$vuetify.breakpoint.lgAndUp"
      fixed
      app
    )
      Sidebar
    v-toolbar(
      :clipped-left="$vuetify.breakpoint.lgAndUp"
      color="primary"
      dark
      app
      fixed
    )
      v-toolbar-title
        v-toolbar-side-icon(@click.stop="drawer = !drawer")
        span.hidden-sm-and-down Vinil78
    v-container(fluid full-height)
      br
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
      console.log(token)
      this.COMMIT('setToken', token)
    }
  },
  components: {
    Header,
    Sidebar
  }
}
</script>
