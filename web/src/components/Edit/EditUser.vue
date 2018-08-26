<template lang="pug">
  v-container(grid-list-sm v-if="needShow")
    v-layout(row, wrap)
      v-flex(xs12)
        v-form(v-model="validEmail")
          v-text-field(
            label="Новый email"
            v-model="emailForm.email"
            :rules="emailForm.emailRules"
          )
          v-layout(row, wrap, justify-end)
            v-btn(color="success" @click="updateEmail") Обновить
              v-icon refresh
      v-flex(xs12)
        v-divider
        br
      v-flex(xs12)
        v-form(v-model="validPassword" )
          v-text-field(
            label="Введите новый пароль"
            v-model="passwordForm.password1"
            :rules="passwordForm.password1Rule"
          )
          v-text-field(
            label="Введите новый пароль"
            v-model="passwordForm.password2"
            :rules="passwordForm.password2Rule"
          )
          v-layout(row, wrap, justify-end)
            v-btn(color="success" @click="updatePassword") Обновить
              v-icon refresh

</template>

<script>
export default {
  name: 'EditUser',
  data: () => {
    return {
      needShow: false,
      validEmail: false,
      emailForm: {
        email: '',
        emailRules: [
          v => !!v || 'Поля обязательно'
        ]
      },
      validPassword: false,
      passwordForm: {
        password1: '',
        password1Rule: [
          v => !!v || 'Повторите пароль'
        ],
        password2: '',
        password2Rule: [
          v => !!v || 'Повторите пароль'
        ]
      }
    }
  },
  methods: {
    async updateEmail () {
      if (this.validEmail) {
        this.loader(true)
        try {
          let result = await this.$api.send(
            'put',
            '/app/user',
            R.pick(['email'], this.emailForm)
          )
          if (result) {
            this.emailForm.email = ''
          }
        } catch (err) {
          this.loader(false)
        }
        this.loader(false)
      }
    },
    async updatePassword () {
      if (this.validPassword) {
        this.loader(true)
        try {
          let result = await this.$api.send('put', '/app/user', {
            password: this.passwordForm.password1
          })
          if (result) {
            this.passwordForm.password1 = ''
            this.passwordForm.password2 = ''
          }
        } catch (err) {
          this.loader(false)
        }
        this.loader(false)
      }
    }
  },
  mounted () {
    if (!R.isEmpty(this.STATE.token)) {
      this.needShow = true
      this.update()
    }
  },
  components: {
  }
}
</script>
