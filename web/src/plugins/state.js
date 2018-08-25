
const statePlugin = {
  install (Vue, { store }) {
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
        loader (value) {
          console.log(value)
          console.log(store)
        },
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
