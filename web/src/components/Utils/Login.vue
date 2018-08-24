<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row wrap)
      v-flex(xs12)
        v-card
          v-toolbar(color="primary")
            v-toolbar-title.text-xs-center.white--text Вход
          v-card-content
            v-container(grid-list-sm)
              v-layout(row, wrap)
                v-flex(xs12)
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
          const role = result.data.role
          this.$store.commit('setToken', token)
          this.$store.commit('setRole', role)
          localStorage.setItem('token', token)
          localStorage.setItem('role', role)
        }
      }
    }
  }
}
</script>
