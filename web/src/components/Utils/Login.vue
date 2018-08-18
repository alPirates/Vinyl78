<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row wrap)
      v-flex(xs12)
        br
        h3 Вход
        v-form(v-model="valid")
          v-text-field(
            label="E-mail"
            v-model="form.email"
            :rules="form.emailRules"
          )
          v-text-field(
            label="Пароль"
            v-model="form.password"
            :rules="form.passwordRules"
          )
          v-layout(row, wrap, justify-end)
            v-btn(color="success" @click="loginUser") Войти
</template>

<script>
export default {
  name: 'Login',
  data: () => {
    return {
      valid: false,
      form: {
        email: 'admin@mail.ru',
        emailRules: [
          v => !!v || 'Требуется пароль',
          v => /.+@.+/.test(v) || 'Введите коректный email'
        ],
        password: 'admin',
        passwordRules: [
          v => !!v || 'Требуется пароль'
        ]
      }
    }
  },
  methods: {
    async loginUser () {
      if (this.valid) {
        let result = await this.$api.send(
          'post',
          '/authorization',
          this.R.pick(['email', 'password'], this.form)
        )
        console.log('result is,', result)
        if (result.data.status === 'success') {
          const token = result.data.token
          this.$store.commit('setToken', token)
          localStorage.setItem('token', token)
        }
      }
    }
  }
}
</script>
