<template lang="pug">
  v-container(grid-list-sm)
    br
    br
    v-layout(row, wrap)
      v-flex(xs12 v-if="isAdmin()")
        h2 Добавить стикер
      v-flex(xs12 v-else)
        h3.display-2 {{Стикеры}}
      v-flex(xs12)
        v-layout(row,  justify-end, v-if="isAdmin()")
          v-flex(xs12)
            v-form(v-model="valid")
              v-text-field(
                label="Описание"
                v-model="sticker.description"
                :rules="sticker.descriptionRules"
                @key.enter="addNewSticker"
              )
          v-btn(color="success" @click="addNewSticker") Добавить
            v-icon(right) add
      v-flex(xs12, sm6, lg4, xl3, v-for="(el, index) in stickers", :key="index")
        v-card.mb-2
          v-toolbar(v-if="isAdmin()" color="primary")
            v-toolbar-title.white--text Редактировать
            v-spacer
            v-btn(@click="edit(index)" icon).white--text
              v-icon edit
          v-card-media(
            height="200px"
            :src="getMedia(getImagePath(el))"
          )
            v-container(fill-height fluid)
              v-layout(fill-height)
                v-flex(xs12 align-end flexbox)
                  span.headline.white--text {{el.description}}

    // edit sticker dialog
    v-dialog(v-model="dialog.show")
      v-card
        v-toolbar(color="primary")
          v-toolbar-title.white--text Редактировать
        EditSticker(
          :sticker="dialog.data"
        )

</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'
import EditSticker from '@/components/Edit/EditSticker'
export default {
  name: 'Stickers',
  data: () => {
    return {
      dialog: {
        show: false,
        data: {}
      },
      valid: false,
      stickers: [],
      sticker: {
        description: '',
        descriptionRules: [
          v => !!v || 'Не может быть пустым'
        ],
        files: []
      },
      count: 0
    }
  },
  watch: {
    '$route.params.id': function (id) {
      this.update()
    }
  },
  methods: {
    getImagePath (sticker) {
      return this.R.path(['name'], this.R.head(this.R.path(['images'], sticker))) || ''
    },
    edit (index) {
      this.dialog.data = this.stickers[index]
      this.dialog.show = true
    },
    async addNewSticker () {
      let result = await this.$api.send('post', '/app/sticker', {
        description: this.sticker.description,
        category_id: this.$route.params.id
      })
      if (result) {
        this.update()
      }
    },
    async update () {
      let result = await this.$api.send('get', '/sticker', null, {
        limit: '10',
        skip: '0',
        category_id: this.$route.params.id
      })

      if (result.data.status === 'success') {
        this.count = result.data.count
        this.$set(this, 'stickers', result.data.result)
      }
    }
  },
  mounted () {
    this.update()
  },
  components: {
    FileUpload, EditSticker
  }
}
</script>
