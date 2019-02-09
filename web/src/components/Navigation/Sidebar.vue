<template lang="pug">
  v-navigation-drawer(
    v-model="drawer"
    v-resize="onResize"
    :clipped="$vuetify.breakpoint.lgAndUp"
    fixed,
    :width="windowSize <= 1264  ? windowSize : 350"
    app
  )
    v-list(dense subheader)#side-list
      span(v-for="(el, index) in categories")
        v-list-tile(:to="'/category/'+el.id" @click="")
            //- v-list-tile-avatar
            //-   v-icon {{el.icon}}
            v-list-tile-content
            v-list-tile-title.tile-title-custom
                | {{el.name}}
      v-list-group(
        no-action
      )
        v-list-tile(slot="activator")
          v-list-tile-title.tile-title-contact КОНТАКТЫ
        v-list-tile(@click="openPhone").tile-phone
          v-list-tile-title.tile-big-title
            v-icon.lefted phone
            strong Телефон:
            |  +7 931 308 73 77
        v-list-tile(@click="openEmail").tile-email
          v-list-tile-title.tile-big-title
            v-icon.lefted email
            strong E-mail:
            |  vinyl78official@gmail.com
        v-list-tile()
          v-layout(row, wrap, justify-center)#links
            v-flex(xs12)
            a.links.big-icon.mr-4(target="_blank" href="https://www.instagram.com/vinyl_78").custom--text
              v-icon fab fa-instagram
            a.links.big-icon(target="_blank" href="https://vk.com/vinyl_78")
              v-icon fab fa-vk
</template>

<script>
export default {
  name: 'Sidebar',
  data: () => {
    return {
      categories: [],
      windowSize: 0
    }
  },
  computed: {
    drawer: {
      get () {
        return this.STATE.drawer
      },
      set (value) {
        this.COMMIT('setDrawer', value)
      }
    }
  },
  methods: {
    openEmail () {
      window.location.href = 'mailto:vinyl78official@gmail.com'
    },
    openPhone () {
      window.location.href = 'tel:79313087377'
    },
    async update () {
      let categories = await this.$api.send('get', '/sidebar')
      console.log('categories', categories)
      if (categories.data.status === 'success') {
        this.$set(this, 'categories', categories.data.result)
      }
    },
    onResize () {
      this.windowSize = window.innerWidth
    }
  },
  mounted () {
    this.update()
  }
}
</script>

<style>
  .v-navigation-drawer .v-list {
    background: none !important;
  }

  .theme--light .v-navigation-drawer {
    background-color: rgba(33, 33, 33, .96) !important;
    z-index: 1000;
  }
  .v-navigation-drawer {
    max-height: calc(100% - 10px) !important;
    font-size: 1.2em !important;
  }

  .tile-title-custom,
  .tile-title-contact,
  .tile-email,
  .tile-phone,
  .theme--light .v-icon {
    color: #e9e9e9 !important;
  }

  #side-list span div a{
    font-size: 1em;
  }
  .tile-title-contact {
    font-size: 1.3em;
  }
  .tile-phone a div {
    font-size: 1.2em;
  }
  .tile-email a div {
    font-size: 1.2em;
  }
  .links {
    text-decoration: none;
  }
  .big-icon i {
    font-size: 32px !important;
    color: #e9e9e9 !important;
  }

  .v-navigation-drawer--is-mobile {
    /* replace styles */
    margin-top: 64px !important;
    /* add new styles */
    opacity: 0.96;
  }

  .v-overlay--active {
    display: none !important;
    background: none !important;
  }

  .v-list__group__items--no-action .v-list__tile {
    padding-left: 26px !important;
  }
  .lefted {
    margin-right: 7px;
  }

  @media (max-width: 600px) {
    #side-list span div a{
      font-size: 4em;
    }
    .tile-title-contact {
      font-size: 1.2em;
    }
    .tile-phone a div {
      font-size: 1.2em;
    }
    .tile-email a div {
      font-size: 1.2em;
    }
  }

  @media (min-width: 600px) and (max-width: 1264px) {
    #side-list span div a{
      font-size: 1.1em !important;
    }
    .v-list__tile__title {
      height: 40px !important;
      line-height: 40px !important;
    }
    .tile-title-contact {
      font-size: 2em !important;
    }
    .tile-phone a div {
      font-size: 1.8em !important;
    }
    .tile-email a div {
      font-size: 1.8em !important;
    }
  }
  @media (max-width: 1264px) {
    #side-list span div a{
      font-size: 0.8em;
    }
    #side-list {
      margin-top: calc(10% + 24px);
    }
    .v-navigation-drawer--is-mobile {
      opacity: 1;
    }
    .lefted {
      margin-right: 10px;
    }
    .tile-title-custom {
      text-align: center !important;
      font-size: 1.4em;
    }
    .tile-title-contact {
      padding-left: 53px;
      text-align: center !important;
      font-size: 1.4em;
    }
    .tile-big-title {
      font-size: 1.4em;
      text-align: center !important;
    }
    .theme--light .v-list .v-list__group--active:before {
      display: none !important;
    }
    .theme--light .v-list .v-list__group--active:after {
      display: none !important;
    }
    .v-navigation-drawer>.v-list .v-list__tile {
      margin-top: 7px;
      letter-spacing: 0.037em;
    }
    .big-icon i {
      font-size: 40px !important;
      color: #e9e9e9 !important;
      /*color: #000 !important;*/
    }
  }
</style>
