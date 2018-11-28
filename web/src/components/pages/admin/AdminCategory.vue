<template lang="pug">
  v-container(grid-list-lg)
    AdminNavigation()
    v-layout(row, wrap)
      v-flex(xs12)
        v-card()
          v-toolbar(color="primary").white--text
            v-toolbar-title Категории
          v-card-text
            v-layout(row, wrap)
              v-flex(xs12 sm6)
                h5.display-1 Все
                draggable(v-model="free", element="v-list", :options="{group:'id'}").dragg
                  div(v-for="(cat, index) in free", :key="index")
                    v-list-tile(@click="")
                      v-list-tile-title {{cat.name}}
                    v-divider(v-if="index + 1 !== free.length")
                  v-card(v-if="free.length === 0")
                    v-list-tile.text-xs-center
                      v-list-tile-content
                        v-list-tile-title.text-xs-center В этой группе ничего нету
                        v-list-tile-sub-title.text-xs-center Перетащи сюда для добавления
              v-flex(xs12 sm6)
                h5.display-1 На главной
                draggable(v-model="home", element="v-list", :options="{group:'id'}").dragg
                  div(v-for="(cat, index) in home", :key="index")
                    v-list-tile(@click="")
                      v-list-tile-title {{cat.name}}
                    v-divider(v-if="index + 1 !== free.length")
                  v-card(v-if="home.length === 0")
                    v-list-tile.text-xs-center
                      v-list-tile-content
                        v-list-tile-title.text-xs-center В этой группе ничего нету
                        v-list-tile-sub-title.text-xs-center Перетащи сюда для добавления
            v-layout(row, wrap)
              v-flex(xs12)
                small * Примечание: ссылки используются для перехода по меню:
                br
                small 1 - Брендирование автомобилей
                br
                small 2 - Декорирование витрин
                br
                small 3 - Виниловые наклейки
            v-layout(row, justify-end)
              v-btn(color="success" @click="send") Обновить
                v-icon(right) update

</template>

<script>
import AdminNavigation from '@/components/Navigation/AdminNavigation'

export default {
  name: 'AdminCategory',
  data: () => {
    return {
      home: [],
      free: []
    }
  },
  components: {
    AdminNavigation
  },
  methods: {
    async send () {
      let index = 1
      let home = R.map(el => {
        el.number = index
        index++
        return R.pick(['id', 'number'], el)
      }, R.clone(this.home))
      let free = R.map(el => {
        el.number = 0
        return R.pick(['id', 'number'], el)
      }, R.clone(this.free))

      let sending = {
        categories: [ ...home, ...free ]
      }
      console.log('sending', sending)
      let result = await this.$api.send('put', '/app/categories', sending)
      if (result) {
        console.log('success')
        this.update()
      }
    },
    async update () {
      let result = await this.$api.send('get', '/sidebar')
      let categories = result.data.result
      let inHome = R.reduce((acc, el) => {
        if (el.number === 0) {
          acc.free.push(el)
        } else {
          acc.home.push(el)
        }
        return acc
      }, { home: [], free: [] }, categories)
      this.$set(this, 'home', inHome.home)
      this.$set(this, 'free', inHome.free)
    }
  },
  async mounted () {
    this.update()
  }
}
</script>

<style scoped>
  .dragg {
    min-height: 70px;
  }
</style>
