<template lang="pug">
  v-container
    h3 hello admin
    v-layout(row wrap)
      v-flex(xs12 lg6)
        v-card
          v-toolbar(color="success")
            v-toolbar-title Категории
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
</template>

<script>
export default {
  name: 'AdminPanel',
  data: () => {
    return {
      valid: false,
      newCategory: '',
      newCategoryRules: [
        v => !!v || 'Не может быть пустым'
      ],
      categories: ['kek', 'cheburek']
    }
  },
  methods: {
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
      let categories = await this.$api.send('get', '/sidebar')
      if (categories.data.status === 'success') {
        console.log('setting')
        this.$set(this, 'categories', categories.data.result)
      }
    }
  },
  async mounted () {
    this.update()
  }
}
</script>
