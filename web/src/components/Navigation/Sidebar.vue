<template lang="pug">
  v-navigation-drawer(
    v-model="drawer"
    :clipped="$vuetify.breakpoint.lgAndUp"
    fixed
    app
  )
    v-list(dense subheader)
    v-subheader Menu
    v-list-tile(:to="'/'")
        v-list-tile-avatar
          v-icon add
        v-list-tile-content
        v-list-tile-title
            | Главная
    v-list-tile(:to="'/form'")
        v-list-tile-avatar
          v-icon add
        v-list-tile-content
        v-list-tile-title
            | Заказать стикер
    v-divider
    span(v-for="(el, index) in categories")
      v-list-tile(:to="'/category/'+el.ID").ml-4
          //- v-list-tile-avatar
            v-icon {{el.icon}}
          v-list-tile-content
          v-list-tile-title
              | {{el.name}}
      v-divider(v-if="index + 1 !== categories.length")
</template>

<script>
export default {
  name: 'Sidebar',
  data: () => {
    return {
      categories: []
    }
  },
  computed: {
    drawer: {
      get () {
        return this.STATE.drawer
      },
      set (value) {
        this.COMMIT('setDrawer', value)
      }
    }
  },
  methods: {
    async update () {
      let categories = await this.$api.send('get', '/sidebar')
      console.log('categories', categories)
      if (categories.data.status === 'success') {
        this.$set(this, 'categories', categories.data.result)
      }
    }
  },
  mounted () {
    this.update()
  }
}
</script>
