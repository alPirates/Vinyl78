<template lang="pug">
  v-container(grid-list-sm)
    .wrapper
      v-card
        v-toolbar(color="primary")
          v-toolbar-title.white--text Обратная связь
        v-container
          v-layout(row wrap)
          v-flex(xs12)
            //- h3 Обратная связь
            v-form(
              v-model="valid"
              ref="form"
            )
              v-text-field(
                label="* Имя"
                v-model="form.name"
                validate-on-blur
                :rules="form.nameRules"
              )
              v-text-field(
                @click.native="addFirstPhoneLetter()"
                label="Телефон"
                v-model="form.phone"
                :rules="form.phoneRules"
                mask="+7(###)-###-##-##"
                validate-on-blur
              )
              v-text-field(
                label="E-mail"
                v-model="form.email"
                :rules="form.emailRules"
                validate-on-blur
              )
              v-textarea(
                label="* Сообщение"
                v-model="form.message"
                :rules="form.messageRules"
                validate-on-blur
              )
            v-flex(xs12)
              p.error--text(v-show="formEmailError") Необходимо ввести телефон или E-mail
              v-layout(justify-end)
                .my-btn(@click="sendForm") Отправить
</template>

<script>
export default {
  name: 'Form',
  data: () => {
    return {
      valid: false,
      formEmailError: false,
      form: {
        name: '',
        nameRules: [
          v => !!v || 'Поле не может быть пустым'
        ],
        phone: '',
        phoneRules: [
          v => (!v || v.length === 11) || 'Введите все цифры номера'
        ],
        email: '',
        emailRules: [
          v => (!v || /.+@.+/.test(v)) || 'Введите коректный email'
        ],
        message: '',
        messageRules: [
          v => !!v || 'Поле не может быть пустым'
        ]
      }
    }
  },
  mounted () {
    this.baseForm = this.R.clone(this.form)
  },
  methods: {
    async sendForm () {
      this.$refs.form.validate()
      if (this.valid) {
        // exit if email and phone not setting up
        if (this.R.isEmpty(this.form.email) && this.R.isEmpty(this.form.phone)) {
          this.formEmailError = true
          return
        }
        console.log('sending')
        let result = await this.$api.send(
          'post',
          '/application',
          this.R.pick(['name', 'phone', 'email', 'message'], this.form)
        )
        if (result) {
          this.form = this.baseForm
          this.$refs.form.reset()
        }
      } else {
        if (this.R.isEmpty(this.form.email) && this.R.isEmpty(this.form.phone)) {
          this.formEmailError = true
        }
      }
    },
    addFirstPhoneLetter () {
      if (R.isEmpty(this.form.phone)) {
        this.form.phone = '7'
      }
    }
  }
}
</script>

<style scoped>
  .wrapper {
    display: block;
    max-width: 600px;
    margin: 0 auto;
  }
  .my-btn {
    background-color: #777c89 ;
    border-color: #777c89;
    color: #fff;
    padding-left: 16px;
    padding-right: 16px;
    cursor: pointer;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    border-radius: 2px;
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    height: 36px;
    -webkit-box-flex: 0;
    -ms-flex: 0 0 auto;
    flex: 0 0 auto;
    font-size: 14px;
    font-weight: 500;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    margin: 6px 8px;
    min-width: 88px;
    outline: 0;
    text-transform: uppercase;
    text-decoration: none;
    -webkit-transition: .3s cubic-bezier(.25,.8,.5,1),color 1ms;
    transition: .3s cubic-bezier(.25,.8,.5,1),color 1ms;
    position: relative;
    vertical-align: middle;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    -webkit-box-shadow: 0 3px 1px -2px rgba(0,0,0,.2),0 2px 2px 0 rgba(0,0,0,.14),0 1px 5px 0 rgba(0,0,0,.12);
    box-shadow: 0 3px 1px -2px rgba(0,0,0,.2),0 2px 2px 0 rgba(0,0,0,.14),0 1px 5px 0 rgba(0,0,0,.12);
  }

  .my-btn:hover {
    background-color: rgb(95, 101, 116);
  }
</style>
