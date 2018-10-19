<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row, wrap)
      v-flex(xs12 v-if="isAdmin()")
        h2 Добавить стикер
        v-btn(@click="category.show = true" v-if="isAdmin()") Радактировать категорию
          v-icon(right) edit
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
      v-flex(xs12, sm4, lg4, xl4, v-for="(el, index) in stickers", :key="index")
        v-card(
            v-on:mouseover="mouseOnSticker(index)"
            v-on:mouseleave="mouseOutSticker(index)"
            @click="gotoSticker(el.id)"
        ).mb-2
          v-toolbar(v-if="isAdmin()" color="primary")
            v-toolbar-title.white--text Редактировать
            v-spacer
            v-btn(@click="edit(index)" icon).white--text
              v-icon edit
            v-btn(@click="remove(el.id)" icon).white--text
              v-icon remove
          router-link(:to="'/sticker/' + el.id").link
            v-card-media(
              height="200px"
              :src="getMedia(getImagePath(el))"
            ).ccard-media
              v-container(fill-height fluid).low-index
                v-layout(fill-height)
                  v-flex(xs12 align-end flexbox fill-height).centered
                    div.center-block.white--text.invisible {{el.description}}

      InfiniteLoading(
        @infinite="loadNew"
        v-if="disableScroll"
      )
        v-layout(row, justify-center, slot="spinner")
          v-progress-circular(size="50" indeterminate color="primary").mt-2

    // edit sticker dialog
    v-dialog(v-model="dialog.show" fullscreen hide-overlay transition="dialog-bottom-transition")
      v-card
        v-toolbar(color="primary")
          v-toolbar-title.white--text Редактировать
          v-spacer
          v-btn(icon).white--text
            v-icon(@click="dialog.show = false") close
        EditSticker(
          :sticker="dialog.data"
          @refresh="refreshData"
        )
    // edit category
    v-dialog(v-model="category.show" fullscreen hide-overlay transition="dialog-bottom-transition")
      v-card
        v-toolbar(color="primary")
          v-toolbar-title.white--text Редактировать
          v-spacer
          v-btn(icon).white--text
            v-icon(@click="category.show = false") close
        EditCategory(
          :id="$route.params.id"
          @refresh="changeData"
        )

</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'
import EditSticker from '@/components/Edit/EditSticker'
import EditCategory from '@/components/Edit/EditCategory'

export default {
  name: 'Stickers',
  data: () => {
    return {
      dialog: {
        show: false,
        data: {}
      },
      // category dialog
      category: {
        show: false
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
      page: 0,
      disableScroll: true
    }
  },
  watch: {
    '$route.params.id': function (id) {
      this.update()
    }
  },
  methods: {
    mouseOnSticker (index) {
      let stickers = document.getElementsByClassName('invisible')
      if (stickers) {
        stickers[index].classList.add('visible')
      }
    },
    mouseOutSticker (index) {
      let stickers = document.getElementsByClassName('invisible')
      if (stickers) {
        stickers[index].classList.remove('visible')
      }
    },
    async loadNew ($state) {
      let result = await this.$api.send('get', '/sticker', null, {
        limit: '12',
        skip: this.page * 12,
        category_id: this.$route.params.id
      })
      if (result.data.status === 'success') {
        this.page++
        if (R.isEmpty(result.data.result)) {
          this.disableScroll = false
        }
        let data = this.stickers.concat(result.data.result)
        this.$set(this, 'stickers', data)
        $state.loaded()
      }
    },
    refreshData () {
      this.update()
      this.dialog.show = false
    },
    async remove (id) {
      this.loader(true)
      try {
        let result = await this.$api.send('delete', `/app/sticker/${id}`)
        if (result) {
          this.update()
        }
      } catch (err) {
        this.loader(false)
      }
      this.loader(false)
    },
    changeData () {
      this.category.show = false
      this.update()
    },
    getImagePath (sticker) {
      return R.path(['name'], this.R.head(this.R.path(['images'], sticker))) || ''
    },
    edit (index) {
      this.dialog.data = this.stickers[index]
      this.dialog.show = true
    },
    async addNewSticker () {
      if (this.valid) {
        let result = await this.$api.send('post', '/app/sticker', {
          description: this.sticker.description,
          category_id: this.$route.params.id
        })
        if (result) {
          this.update()
        }
      }
    },
    async update () {
      let result = await this.$api.send('get', '/sticker', null, {
        limit: '12',
        skip: 0,
        category_id: this.$route.params.id
      })
      this.page++
      if (result.data.status === 'success') {
        this.$set(this, 'stickers', result.data.result)
      }
    }
  },
  mounted () {
    this.stickers = []
    // this.update()
  },
  components: {
    FileUpload,
    EditSticker,
    EditCategory
  }
}
</script>

<style scoped>
  .infinite-loading-container {
    width: 100%
  }
  .ccard-media:hover:after {
    content: "";
    position: absolute;
    left: 50%;
    top: 50%;
    z-index: 1;
    box-shadow: 0 0 0 300px rgba(0,0,0,0.4);
  }
  .low-index {
    z-index: 10;
  }
  .invisible {
    visibility: hidden;
  }
  .visible {
    visibility: visible !important;
  }
  .link {
    text-decoration: none;
  }
  .centered {
    text-align: center !important;
  }
  .center-block {
    width: 100%;
    height: 100%;
    line-height: 130px;
    font-size: 1.8em;
  }
 </style>
