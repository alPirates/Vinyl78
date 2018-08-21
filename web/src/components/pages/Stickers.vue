<template lang="pug">
  v-container(grid-list-sm)
    br
    br
    v-layout(row, wrap)
      v-flex(xs12)
        h2 Добавить стикер
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
      v-flex(xs12, sm6, lg4, xl3)
        v-card(v-for="(el, index) in stickers", :key="index").mb-1
          v-toolbar(color="green")
            v-tolbar-title.white--text {{el.description}}
            v-spacer
            v-btn(icon v-if="isAdmin()" @click="edit(index)").white--text
              v-icon edit
    // edit sticker dialog
    v-dialog(v-model="dialog.show")
      v-card
        v-toolbar(color="primary")
          v-toolbar-title.white--text Редактировать
        EditSticker(:sticker="dialog.data")

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
    edit (index) {
      this.dialog.data = this.stickers[index]
      this.dialog.show = true
    },
    async addNewSticker () {
      let result = await this.$api.send('post', '/app/sticker', {
        description: this.sticker.description,
        category: parseInt(this.$route.params.id)
      })
      if (result) {
        this.update()
      }
    },
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
  },
  components: {
    FileUpload, EditSticker
  }
}
</script>
