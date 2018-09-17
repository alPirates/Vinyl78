<template lang="pug">
  v-container(grid-list-lg)
    v-layout(row, wrap)
      v-flex(xs12)
        v-form(v-model="valid")
          v-layout(row, wrap)
            v-flex(xs12)
              v-text-field(
                label="Имя"
                v-model="form.name"
              )
              v-layout(row)
                v-flex(xs6)
                  v-text-field(
                    label="Иконка"
                    v-model="form.icon"
                  )
                v-flex(xs6)
                  | Превью иконки
                  v-btn(icon)
                    v-icon {{form.icon}}
              label.mt-2 Превью раздела меню на главной
              v-textarea(
                label="Описание стикера"
                v-model="form.description"
              )
            v-flex(xs12)
              h4 Картинка
              v-list
                div(v-for="(el, index) in form.images")
                  v-list-tile(@click="")
                    v-list-tile-avatar
                      img(:src="getMedia(el.name)")
                    v-list-tile-title {{el.name}}
                    v-spacer
                    v-btn(color="error" icon @click="deleteImage(el.id)")
                      v-icon remove
              FileUpload(
                v-model="files"
                :data="{linked_id: id}"
              )
            v-flex(xs12)
              v-layout(justify-end)
                v-btn(color="success" @click="updateCategory") Обновить
                  v-icon(right) update
</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'
export default {
  name: 'EditCategory',
  data: () => {
    return {
      valid: false,
      form: {
        icon: '',
        name: '',
        description: '',
        images: []
      },
      files: []
    }
  },
  watch: {
    id: {
      handler: function (val, oldVal) {
        this.updateProps()
      }
    }
  },
  methods: {
    async updateProps () {
      if (!R.isEmpty(this.id)) {
        this.loader(true)
        let result = await this.$api.send('get', `/category/${this.id}`)
        if (result) {
          let data = result.data.result
          const keys = ['icon', 'name', 'description']
          console.log('data is ', data)
          R.forEach(key => {
            if (data[key]) {
              this.form[key] = data[key]
            }
          }, keys)
          this.$set(this.form, 'images', data.images)
        }
        this.loader(false)
      }
    },
    updateCategory: async function () {
      this.loader(true)
      try {
        const fields = ['icon', 'name', 'description']
        let result = await this.$api.send('put', '/app/category', {
          ...R.pick(fields, this.form),
          id: this.id
        })

        if (result) {
          await this.updateProps()
          this.$emit('refresh')
        }
      } catch (err) {
        this.loader(false)
      }
      this.loader(false)
    },
    async deleteImage (id) {
      let result = await this.$api.send('delete', `/app/image/${id}`)
      if (result) {
        this.updateCategory()
      }
    }
  },
  mounted () {
    this.updateProps()
  },
  props: {
    id: {
      type: String,
      default: ''
    }
  },
  components: {
    FileUpload
  }
}
</script>
