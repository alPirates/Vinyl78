<template lang="pug">
  v-container
    v-layout(row, wrap)
      v-flex(xs12)
        | Carousel
        pre here {{carouselImages}}
      v-flex(xs12)
        v-layout(justify-end)
          v-dialog(v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition")
            v-btn(color="success" slot="activator") Добавить
            v-card
              v-toolbar(color="primary")
                v-toolbar-title.white--text Добавить
                v-spacer
                v-btn(icon @click="dialog = false").white--text
                  v-icon close
              v-container
                v-layout(row)
                  v-flex(xs12)
                    v-text-field(v-model="form.href", label="ссылка")
                v-layout(justify-end)
                  v-btn(color="success" @click="addNewImage") Добавить
          v-btn(color="error") Обновить
</template>

<script>
export default {
  data () {
    return {
      dialog: false,
      carouselImages: [],
      form: {
        href: ''
      }
    }
  },
  methods: {
    async addNewImage () {
      if (R.isEmpty(this.form.href)) {
        return
      }
      await this.$api.send('post', '/app/carousel', {
        name: this.form.href
      })

      this.form.href = ''
    },
    async update () {
      let carouselImages = await this.$api.send('get', '/carousel', null, {
        limit: 100,
        skip: 0
      })

      if (carouselImages) {
        this.$set(this, 'carouselImages', carouselImages.data.images)
      }
      console.log('images is', carouselImages)
    }
  },
  async mounted () {
    await this.update()
  }
}
</script>
