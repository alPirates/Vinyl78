
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
        getMedia (value) {
          return `/api/media/${value}`
        },
        isAdmin () {
          if (this.STATE.role === 'admin') {
            return true
          }
          return false
        }
      }
    })
  }
}

export default statePlugin
