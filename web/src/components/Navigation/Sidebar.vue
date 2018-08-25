<template lang="pug">
  v-navigation-drawer(
    v-model="drawer"
    :clipped="$vuetify.breakpoint.lgAndUp"
    fixed
    app
  )
    v-list(dense subheader)
      v-subheader Меню
      v-list-tile(:to="'/'" @click="")
        v-list-tile-avatar
          v-icon home
        v-list-tile-content
        v-list-tile-title
            | Главная
      v-list-tile(:to="'/form'" @click="")
        v-list-tile-avatar
          v-icon turned_in
        v-list-tile-content
        v-list-tile-title
            | Заказать стикер
      v-divider
      span(v-for="(el, index) in categories")
        v-list-tile(:to="'/category/'+el.id" @click="")
            v-list-tile-avatar
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
