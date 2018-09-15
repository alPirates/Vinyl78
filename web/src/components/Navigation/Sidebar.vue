<template lang="pug">
  v-navigation-drawer(
    v-model="drawer"
    v-resize="onResize"
    :clipped="$vuetify.breakpoint.lgAndUp"
    fixed
    :width="windowSize"
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
      v-list-group(
        no-action
      )
        v-list-tile(slot="activator")
          v-list-tile-avatar
            v-icon info
          v-list-tile-title Контакты
        v-list-tile(@click="")
          v-list-tile-avatar
            v-icon phone
          v-list-tile-title
            strong Телефон:
            |  +7 931 308 73 77
        v-list-tile(@click="")
          v-list-tile-avatar
            v-icon email
          v-list-tile-title
            strong E-mail:
            |  vinyl78official@gmail.com
</template>

<script>
export default {
  name: 'Sidebar',
  data: () => {
    return {
      categories: [],
      windowSize: 0
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
    },
    onResize () {
      this.windowSize = window.innerWidth
    }
  },
  mounted () {
    this.update()
  }
}
</script>

<style>
  .v-navigation-drawer--is-mobile {
    /* replace styles */
    margin-top: 64px !important;
    /* add new styles */
    opacity: 0.96;
  }
  .v-overlay--active {
    display: none !important;
    background: none !important;
  }
</style>
