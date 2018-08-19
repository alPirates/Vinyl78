<template lang="pug">
  v-container(grid-list-sm)
    br
    br
    pre {{count}}
    v-layout(row, wrap)
      v-flex(xs12)
        v-layout(row, wrap, justify-end)
          v-btn(color="success", @click="dialog = !dialog") Добавить
            v-icon(right) add
      v-flex(xs12, sm6, lg4, xl3)
        v-card(v-for="(el, index) in stickers", :key="index")
          v-toolbar
            v-tolbar-title.white--text {{el}}
    v-dialog(v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition")
      v-card
        v-toolbar(color="primary")
          v-toolbar-title.white--text Добавить стикер
          v-spacer
          v-btn(icon @click="dialog = false").white--text
            v-icon close
        v-container
          v-layout(row, wrap)
            v-flex(xs12)
</template>

<script>
export default {
  name: 'Stickers',
  data: () => {
    return {
      dialog: false,
      stickers: [],
      count: 0
    }
  },
  methods: {
    async update () {
      let result = await this.$api.send('get', '/sticker', null, {
        limit: '10',
        skip: '0',
        category: this.$route.params.id
      })

      if (result.data.status === 'success') {
        this.count = result.data.count
        this.$set(this, 'stickers', result.data.result)
      }
    }
  },
  mounted () {
    this.update()
  }
}
</script>
