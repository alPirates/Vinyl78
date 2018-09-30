<template lang="pug">
  v-footer(
    height="auto"
    color="primary"
    app
    absolute
  )
    //- a(href="https://www.instagram.com/vinyl_78")
    //-   v-btn(icon).white--text
    //-     v-icon fab fa-instagram
    //- a(href="https://vk.com/vinyl_78")
    //-   v-btn(icon).white--text
    //-     v-icon fab fa-vk
    //- v-card(
    //-   flat, tile
    //-   flex
    //-   color="success"
    //- )
    v-flex.white--text.text-xs-center
      | &copy;2018 —
      strong Vinyl78
</template>

<script>
export default {
  name: 'Footer',
  data: () => {
    return {}
  }
}
</script>

<style>
  .v-footer {
    bottom: -53px !important;
  }
</style>
