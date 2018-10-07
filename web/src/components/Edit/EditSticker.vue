<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row, wrap)
      v-flex(xs12)
        h3.display-1 Редактировать стикер
      v-flex(xs12)
        v-form(v-model="valid")
          v-text-field(
            label="Описание"
            v-model="form.description"
            :rules="form.descriptionRules"
          )
      v-flex(xs12 v-if="form.images.length > 0")
        h4 Картинки
        draggable(:list="form.images", element="v-list")
          div(v-for="(image, index) in form.images", :key="image.id")
            v-list-tile(@click="")
              v-list-tile-avatar
                img(:src="getMedia(image.name)")
              v-list-tile-title
                | Картинка:
                |
                strong {{image.name}}
              v-spacer
              v-btn(color="error" icon @click="deleteSticker(image.id)")
                v-icon delete
        v-flex(xs12)
          v-layout(justify-end)
            v-btn(color="success", @click="updateCount") Обновить порядок
      v-flex(xs12)
        FileUpload(
          v-model="files"
          :data="{linked_id: sticker.id}"
        )
      v-flex(xs12)
        v-layout(justify-end)
          v-btn(color="success" @click="updateSticker") Обновить
            v-icon update
</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'
export default {
  name: '',
  data: () => {
    return {
      files: [],
      valid: false,
      form: {
        description: '',
        descriptionRules: [
          v => !!v || 'Не может быть пустым'
        ],
        images: []
      }
    }
  },
  watch: {
    sticker: {
      handler: function (val, oldVal) {
        this.preLoadProps()
      }
    }
  },
  methods: {
    async updateCount () {
      this.loader(true)
      let slides = R.clone(this.form.images)
      let index = 0
      let sending = R.map(el => {
        index++
        return {
          id: el.id,
          number: index
        }
      }, slides)
      await this.$api.send('put', '/app/image', {
        images: sending
      })
      this.loader(false)
    },
    deleteSticker: async function (id) {
      let result = await this.$api.send('delete', `/app/image/${id}`)
      if (result) {
        this.$emit('refresh')
      }
    },
    updateSticker: async function () {
      if (this.valid) {
        let result = await this.$api.send(
          'put',
          '/app/sticker',
          {
            ...R.pick(['description'], this.form),
            id: this.sticker.id
          }
        )
        if (result) {
          this.$emit('refresh')
        }
      }
    },
    preLoadProps () {
      const keys = ['description']
      let needProps = R.pick(keys, this.sticker)
      R.forEach((el, val) => {
        if (!R.isEmpty(needProps[el])) { this.form[el] = needProps[el] }
      }, keys)
      if (this.sticker.images) {
        this.$set(this.form, 'images', this.sticker.images)
      }
    }
  },
  mounted () {
    if (!R.isEmpty(this.sticker)) {
      this.preLoadProps()
    }
  },
  props: {
    sticker: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  components: {
    FileUpload
  }
}
</script>
