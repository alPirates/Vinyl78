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
</template>

<script>
export default {
  name: 'Header',
  data: () => {
    return {
    }
  },
  methods: {
    changeDrawer () {
      this.COMMIT('setDrawer', !this.STATE.drawer)
    }
  }
}
</script>
