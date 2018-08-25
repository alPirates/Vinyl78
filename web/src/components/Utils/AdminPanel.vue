<template lang="pug">
  v-container(grid-list-lg)
    br
    v-layout(row wrap)
      v-flex(xs12 lg6)
        v-card
          v-toolbar(color="success")
            v-toolbar-title.white--text Категории
          v-container
            v-layout(row, wrap)
              v-flex(xs12)
                v-layout(row)
                  v-flex(xs12)
                    v-form(v-model="valid")
                      v-text-field(
                        label="Добавить категорию"
                        :rules="newCategoryRules"
                        v-model="newCategory"
                        @key.enter="addCategory"
                      )
                  v-flex
                    v-btn(color="primary" @click="addCategory") Добавить
                      v-icon(right) add
              v-flex(xs12)
                v-list(v-for="(el, index) in categories", :key="index")
                  v-list-tile
                    v-list-tile-title {{el.name}}
                    v-spacer
                    v-btn(icon color="success")
                     v-icon edit
                    v-btn(icon color="grey lighter-3", @click="deleteCategory(el.ID)")
                      v-icon delete
      v-flex(xs12 lg6)
       v-card
          v-toolbar(color="success")
            v-toolbar-title.white--text Карусель
          v-container
            v-layout(row, wrap)
              v-flex(xs12)
                v-layout(row)
                  v-flex(xs12)
                    v-carousel(id="preview-carousel")
                      v-carousel-item(v-for="(image, index) in carouselImages", :key="index")
                        img(:src="getMedia(image.name)")
                    h3.mt-2.mb-2.display-1 Укажите новую последовательность
                    draggable(:list="carouselImages", element="v-list" @end="dropImage")
                      div(v-for="(image, index) in carouselImages", :key="index")
                        v-list-tile(@click="")
                          v-list-tile-avatar
                            img(:src="getMedia(image.name)")
                          v-list-tile-title
                            | Картинка:
                            |
                            strong {{image.name}}
                          v-spacer
                          v-btn(icon color="error")
                            v-icon delete
                        v-divider
                    v-layout(row, wrap)
                      v-flex(xs12).text-xs-right
                        v-btn(color="primary" @click="refreshCarouselSlides()") Обновить
                          v-icon(right) refresh
                    h3.mt-2.mb-2.display-1 Добавить картинки
                    FileUpload(
                      :data="{linked_id: carousel_id}"
                    )
      v-flex(xs12)
       v-card
          v-toolbar(color="success")
            v-toolbar-title.white--text заявки
          v-container
            v-layout(row, wrap)
              v-flex(xs12)
                v-layout(row)
                  v-flex(xs12)
                    v-data-table(:headers="headers", :items="info.applications").elevation-1
                      template(slot="items" slot-scope="el")
                        td {{el.item.name}}
                        td {{el.item.phone}}
                        td {{el.item.email}}
                        td {{el.item.message}}
                        td {{el.item.time | date}}
                        td {{el.item.status | status}}
                        td
                          v-btn(icon color="success")
                            v-icon check

</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'

export default {
  name: 'AdminPanel',
  data: () => {
    return {
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
  methods: {
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
      let result = await this.$api.send('put', '/app/image', sending)
      if (result) {
      }
      this.loader(false)
    },
    async dropImage (param) {
      console.log('param', this.carouselImages)
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
      console.log('value is ', carousel.data.value)
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
    FileUpload
  }
}
</script>

<style scoped>
#preview-carousel {
  max-height: 300px;
}
</style>
