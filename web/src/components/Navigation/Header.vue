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
        router-link(to="/")
          img(src="@/assets/logo.svg" height="30px" to="/").logo-image
      v-spacer
      v-btn(
        icon
        v-if="!isAdmin()"
        @click="openPhone"
      )#phone-btn
        //- v-icon phone
        img(
          width="30"
          src="@/assets/phone.svg"
        )
      v-btn(
        icon
        v-if="isAdmin()"
        to="/admin"
      )
        v-icon verified_user
      v-btn(
        v-if="STATE.token"
        icon
        flat
        @click="logout")
        v-icon fa-sign-out-alt

</template>

<script>
export default {
  name: 'Header',
  data: () => {
    return {
    }
  },
  methods: {
    openPhone () {
      window.location.href = 'tel:79313087377'
    },
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
#phone-btn {
  width: 50px;
  height: 50px;
}
</style>
