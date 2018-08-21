<template lang="pug">
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
      v-spacer
      v-btn(
        v-if="STATE.token"
        flat
        @click="logout"
      ) Выйти
</template>

<script>
export default {
  name: 'Header',
  data: () => {
    return {
    }
  },
  methods: {
    logout () {
      this.$store.commit('setToken', '')
      this.$store.commit('setRole', 'client')
      localStorage.removeItem('token')
      this.$router.push('/admin')
    },
    changeDrawer () {
      this.COMMIT('setDrawer', !this.STATE.drawer)
    }
  }
}
</script>
