<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row, wrap)
      v-flex(xs12)
        v-form(v-model="valid")
          pre {{valid}}
          v-layout(row, wrap)
            v-flex(xs12)
              v-text-field(
                label="Имя"
                v-model="form.name"
              )
              v-text-field(
                label="Иконка"
                v-model="form.icon"
              )
              v-textarea(
                label="Описание стикера"
                v-model="form.description"
              )
            v-flex(xs12)
              v-layout(justify-end)
                v-btn(color="success" @click="updateCategory") Обновить
                  v-icon(right) update
</template>

<script>
export default {
  name: 'EditCategory',
  data: () => {
    return {
      valid: false,
      form: {
        icon: '',
        name: '',
        description: ''
      }
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
        console.log('result is', result)
        if (result) {
          let data = result.data.result
          const keys = ['icon', 'name', 'description']
          console.log('data is ', data)
          R.forEach(key => {
            if (data[key]) {
              this.form[key] = data[key]
            }
          }, keys)
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
  }
}
</script>
