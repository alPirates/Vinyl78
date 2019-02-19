<template lang="pug">
  v-container(grid-list-lg)
    AdminNavigation()
    v-layout(row wrap)
      v-flex(xs12)
        v-card
          v-toolbar(color="secondary")
            v-toolbar-title.white--text Разделы меню
          v-container
            v-layout(row, wrap)
              v-flex(xs12)
                v-layout(row)
                  v-flex(xs12)
                    v-form(v-model="valid")
                      v-text-field(
                        label="Добавить раздел меню"
                        :rules="newCategoryRules"
                        v-model="newCategory"
                        @key.enter="addCategory"
                      )
                  v-flex
                    v-btn(color="primary" @click="addCategory") Добавить
                      v-icon(right) add
              v-flex(xs12)
                v-list(v-for="(el, index) in categories", :key="el.id")
                  v-list-tile(id="sticker-card")
                    v-list-tile-title {{el.name}}
                    v-spacer
                    v-btn(icon color="success" @click="showEditCategory(el.id)")
                     v-icon edit
                    v-btn(icon color="grey lighter-3", @click="deleteCategory(el.id)")
                      v-icon delete
      v-flex(xs12)
        v-card
          v-toolbar(color="secondary")
            v-toolbar-title.white--text Карусель
          v-container
            v-layout(row, wrap)
              v-flex(xs12)
                AdminCarousel()
      v-flex(xs12)
       v-card
          v-toolbar(color="secondary")
            v-toolbar-title.white--text Заявки
          v-container
            v-layout(row, wrap)
              v-flex(xs12)
                v-layout(row)
                  v-flex(xs12)
                    v-data-table(:headers="headers", :items="info.applications").elevation-1
                      template(slot="items" slot-scope="el")
                        td {{el.item.name}}
                        td {{el.item.phone | checkEmpty}}
                        td {{el.item.email | checkEmpty}}
                        td {{el.item.message | checkEmpty}}
                        td {{el.item.time | date}}
                        td {{el.item.status | status}}
                        td
                          v-btn(icon color="success" @click="changeStatus(el.item, 1)" v-if="el.item.status === 0")
                            v-icon check
                          v-btn(icon color="error" @click="changeStatus(el.item, 0)" v-if="el.item.status === 1")
                            v-icon remove
      v-dialog(v-model="editCateg.show" fullscreen hide-overlay transition="dialog-bottom-transition")
        v-card
          v-toolbar(color="primary")
            v-toolbar-title.white--text Редактировать категорию
            v-spacer
            v-btn(icon @click="editCateg.show = false").white--text
              v-icon close
          v-card-title
            EditCategory(:id="editCateg.id")

</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'
import EditCategory from '@/components/Edit/EditCategory'
import AdminNavigation from '@/components/Navigation/AdminNavigation'
import AdminCarousel from '@/components/pages/admin/AdminCarousel'

export default {
  name: 'AdminPanel',
  data: () => {
    return {
      editCateg: {
        show: false,
        id: ''
      },
      valid: false,
      files: [],
      headers: [
        {text: 'Имя', value: 'name'},
        {text: 'Телефон', value: 'phone'},
        {text: 'E-mail', value: 'email'},
        {text: 'Сообщение', value: 'message'},
        {text: 'Дата', value: 'time'},
        {text: 'Cтатус', value: 'status'},
        {text: 'Опции', sortable: 'false'}
      ],
      carousel_id: '',
      newCategory: '',
      newCategoryRules: [
        v => !!v || 'Не может быть пустым'
      ],
      categories: [],
      info: {},
      carouselImages: []
    }
  },
  filters: {
    checkEmpty (value) {
      if (R.isEmpty(value)) {
        return '-'
      }
      return value
    }
  },
  methods: {
    showEditCategory (id) {
      this.editCateg.id = id
      this.editCateg.show = true
      console.log('showing')
    },
    async refreshCarouselSlides () {
      this.loader(true)
      let index = 0
      let sending = R.map(el => {
        index++
        return {
          id: el.id,
          number: index
        }
      }, this.carouselImages)
      let result = await this.$api.send('put', '/app/image', {
        images: sending
      })
      if (result) {
        this.update()
      }
      this.loader(false)
    },
    async dropImage (param) {
      console.log('param', this.carouselImages)
    },
    async removeImage (imageId) {
      console.log(imageId)
      let result = await this.$api.send('delete', `/app/image/${imageId}`)
      if (result.data.status === 'success') {
        this.update()
      }
    },
    async addCategory () {
      if (this.valid) {
        let result = await this.$api.send('post', '/app/category', {
          name: this.newCategory,
          icon: 'add'
        })
        if (result.data.status === 'success') {
          this.update()
        }
      }
    },
    async deleteCategory (id) {
      let result = await this.$api.send('delete', `/app/category/${id}`)
      console.log(result)
      if (result.data.status === 'success') {
        this.update()
      }
    },
    async changeStatus (application, status) {
      this.loader(true)
      try {
        let data = application
        data.status = status
        let result = await this.$api.send('put', `/app/application`, data)
        if (result) {
          await this.update()
        }
      } catch (err) {
        console.error(err)
      }
      this.loader(false)
    },
    async update () {
      let adminPanel = await this.$api.send('get', '/app/admin')
      if (adminPanel) {
        this.$set(this, 'info', adminPanel.data)
      }
      console.log(adminPanel)
      let categories = await this.$api.send('get', '/sidebar')
      if (categories.data.status === 'success') {
        console.log('setting')
        this.$set(this, 'categories', categories.data.result)
      }

      let carousel = await this.$api.send('get', '/app/property', null, {
        key: 'carousel_id'
      })
      this.carousel_id = carousel.data.value
      let carouselImages = await this.$api.send('get', '/image', null, {
        linked_id: carousel.data.value
      })
      if (carouselImages) {
        console.log('1111', carouselImages.data.images)
        this.$set(this, 'carouselImages', carouselImages.data.images)
      }
    }
  },
  async mounted () {
    this.update()
  },
  components: {
    FileUpload,
    EditCategory,
    AdminNavigation,
    AdminCarousel
  }
}
</script>

<style scoped>
#preview-carousel {
  max-height: 300px;
}
#red-badge {
  top: -2px !important;
  right: 1px !important;
}
</style>
