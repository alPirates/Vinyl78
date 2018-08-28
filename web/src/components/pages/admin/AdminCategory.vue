<template lang="pug">
  v-container(grid-list-lg)
    AdminNavigation()
    v-layout(row, wrap)
      v-flex(xs12)
        v-card()
          v-toolbar(color="primary").white--text
            v-toolbar-title Категории
          v-card-text
            v-layout(row, wrap)
              v-flex(xs6)
                h5.display-1 Все
              v-flex(xs6)
                h5.display-1 На главной

</template>

<script>
import AdminNavigation from '@/components/Navigation/AdminNavigation'

export default {
  name: 'AdminCategory',
  data: () => {
    return {
    }
  },
  components: {
    AdminNavigation
  },
  methods: {
    async update () {
      let result = await this.$api.send('get', '/sidebar')
      let categories = result.data.result
      // let inHome = R.reduce((acc, el) => {
      // }, {}, categories)
    }
  },
  async mounted () {
    this.update()
  }
}
</script>
