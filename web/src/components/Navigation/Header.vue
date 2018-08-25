<template lang="pug">
    v-toolbar(
      :clipped-left="$vuetify.breakpoint.lgAndUp"
      color="primary"
      :height="64"
      dark
      app
      fixed
    )
      v-toolbar-title
        v-toolbar-side-icon(@click.stop="changeDrawer()")
        img(src="@/assets/logo.svg" height="30px").logo-image
      v-spacer
      v-btn(icon
        v-if="isAdmin()"
        to="/admin"
      )
        v-icon verified_user
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
      localStorage.removeItem('role')
      this.$router.push('/admin')
    },
    changeDrawer () {
      this.COMMIT('setDrawer', !this.STATE.drawer)
    }
  }
}
</script>

<style>
.logo-image {
  position: absolute;
  top: 20px;
  left: calc(50% - 63px);
}
</style>
