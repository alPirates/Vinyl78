
const statePlugin = {
  install (Vue) {
    Vue.mixin({
      computed: {
        STATE () {
          return this.$store.state
        },
        COMMIT () {
          return this.$store.commit
        },
        DISPATCH () {
          return this.$store.dispatch
        }
      },
      methods: {
      }
    })
  }
}

export default statePlugin
