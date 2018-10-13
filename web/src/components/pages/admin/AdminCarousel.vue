<template lang="pug">
  v-container
    v-layout(row, wrap)
      v-flex(xs12)
        v-list
          draggable(v-model="carouselImages", element="v-list", :options="{group:'id'}").dragg
            div(v-for="(item, index) in carouselImages")
              v-list-tile(avatar @click="")
                v-list-tile-avatar
                  img(:src="getMedia(item.image[0].name)" v-if="item.image.length > 0")
                v-list-tile-title {{item.name}}
                v-spacer
                v-btn(@click="showAdding(item.id)" color="success") Добавить картинку
                v-btn(@click="showEditing(index)" icon)
                  v-icon edit
                v-btn(icon @click="removeItem(item.id)" color="error")
                  v-icon remove
              v-divider
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

        //- dialogs
        v-dialog(v-model="dialogImages.show" fullscreen hide-overlay transition="dialog-bottom-transition")
          v-card
            v-toolbar(color="primary")
              v-toolbar-title.white--text Добавить
              v-spacer
              v-btn(icon @click="dialogImages.show = false").white--text
                v-icon close
            v-container
              v-layout(row)
                v-flex(xs12)
                  FileUpload(
                    v-model="form.file"
                    :data="{linked_id: this.dialogImages.data}"
                  )
        v-dialog(v-model="dialogEdit.show" fullscreen hide-overlay transition="dialog-bottom-transition")
          v-card
            v-toolbar(color="primary")
              v-toolbar-title.white--text Редактировать
              v-spacer
              v-btn(icon @click="dialogEdit.show = false").white--text
                v-icon close
            v-container
              v-layout(row)
                v-flex(xs12)
                  v-text-field(
                    v-model="dialogEdit.data.name"
                  )
                v-flex(xs12)
                  v-layout(justify-end)
                    v-btn(color="success" @click="updateEdit") Обновить
</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'

export default {
  data () {
    return {
      dialog: false,
      carouselImages: [],
      dialogImages: {
        show: false,
        data: ''
      },
      form: {
        href: '',
        file: null
      },
      dialogEdit: {
        show: false,
        data: {
          name: ''
        }
      }
    }
  },
  methods: {
    showAdding (id) {
      this.dialogImages.data = id
      this.dialogImages.show = true
    },
    showEditing (index) {
      this.dialogEdit.data = this.carouselImages[index]
      this.dialogEdit.show = true
    },
    async updateEdit () {
      let result = await this.$api.send('put', `/app/carousel`, {
        ...R.pick(['name', 'id'], this.dialogEdit.data)
      })
      if (result) {
        this.update()
      }
      this.dialogEdit = false
    },
    async removeItem (id) {
      await this.$api.send('delete', `/app/carousel/${id}`)
      await this.update()
    },
    async addNewImage () {
      if (R.isEmpty(this.form.href)) {
        return
      }
      await this.$api.send('post', '/app/carousel', {
        name: this.form.href
      })
      this.form.href = ''
      await this.update()
    },
    async update () {
      let carouselImages = await this.$api.send('get', '/carousel', null, {
        limit: 100,
        skip: 0
      })
      console.log('images is', carouselImages.data.result)
      if (carouselImages) {
        console.log('setting')
        this.$set(this, 'carouselImages', carouselImages.data.result)
      }
    }
  },
  async mounted () {
    await this.update()
  },
  components: {
    FileUpload
  }
}
</script>
