<template lang="pug">
  v-app
    v-navigation-drawer(
      v-model="drawer"
      :clipped="$vuetify.breakpoint.lgAndUp"
      fixed
      app
    )
      v-list(dense subheader)
        v-subheader Menu
        v-list-tile
          v-list-tile-avatar
            v-icon add
          v-list-tile-content
            v-list-tile-title
              | but stickers
        v-divider
        v-list-tile
          v-list-tile-avatar
            v-icon delete
          v-list-tile-content
            v-list-tile-title
              | Category 2
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
      router-view

</template>

<script>
import Header from '@/components/Navigation/Header.vue'

export default {
  name: 'App',
  data: () => {
    return {
      drawer: false
    }
  },
  components: {
    Header
  }
}
</script>
