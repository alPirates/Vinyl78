<template lang="pug">
  v-app
    Sidebar
    v-toolbar(
      :clipped-left="$vuetify.breakpoint.lgAndUp"
      color="primary"
      dark
      app
      fixed
    )
      v-toolbar-title
        v-toolbar-side-icon(@click.stop="changeDrawer()")
        span.hidden-sm-and-down Vinil78
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
  methods: {
    changeDrawer () {
      this.COMMIT('setDrawer', !this.STATE.drawer)
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
