<template lang="pug">
  v-container(grid-list-sm)
    br
    v-layout(row, wrap)
      v-flex(xs12 sm8 offset-xs0 offset-sm2)
        v-card
          v-toolbar(color="primary")
            v-toolbar-title.white--text Заказать стикер
          v-container
            v-layout(row wrap)
            v-flex(xs12)
              h3 Форма
              v-form(
                v-model="valid"
                ref="form"
              )
                v-text-field(
                  label="Имя"
                  v-model="form.name"
                  :rules="form.nameRules"
                )
                v-text-field(
                  label="Телефон"
                  v-model="form.phone"
                  :rules="form.phoneRules"
                  mask="+7(###)-###-##-##"
                )
                v-text-field(
                  label="E-mail"
                  v-model="form.email"
                  :rules="form.emailRules"
                )
                v-textarea(
                  label="Сообщение"
                  v-model="form.message"
                  :rules="form.messageRules"
                )
              v-flex(xs12)
                v-layout(justify-end)
                  v-btn(color="success" @click="sendForm") Отправить
</template>

<script>
export default {
  name: 'Form',
  data: () => {
    return {
      valid: false,
      form: {
        name: '',
        nameRules: [
          v => !!v || 'Поле не может быть пустым'
        ],
        phone: '',
        phoneRules: [
          v => v.length === 11 || 'Введите все цифры номера'
        ],
        email: '',
        emailRules: [
          v => !!v || 'Введите E-mail',
          v => /.+@.+/.test(v) || 'Введите коректный email'
        ],
        message: ''
      }
    }
  },
  mounted () {
    this.baseForm = this.R.clone(this.form)
  },
  methods: {
    async sendForm () {
      if (this.valid) {
        let result = await this.$api.send(
          'post',
          '/application',
          this.R.pick(['name', 'phone', 'email', 'message'], this.form)
        )
        if (result) {
          this.form = this.baseForm
          this.$refs.form.reset()
        }
      }
    }
  }
}
</script>
