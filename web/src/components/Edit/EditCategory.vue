<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row, wrap)
      v-flex(xs12)
        v-form(v-model="valid")
          v-layout(row, wrap)
            v-flex(xs12)
              v-text-field(
                label="Имя"
                v-model="form.name"
              )
            v-flex(xs12)
              v-layout(row)
                v-flex(xs12)
                  v-text-field(
                    label="Иконка"
                    v-model="form.icon"
                  )
            v-flex(xs12)
              v-layout(justify-end)
                v-btn(color="success") Обновить
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
        name: ''
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
        // get one category todo create update and create methods
        this.loader(false)
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
  }
}
</script>
