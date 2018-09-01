<template lang="pug">
  v-footer(
    height="auto"
    color="primary"
    app
    absolute
  )
    v-btn(icon).white--text
      v-icon fab fa-instagram
    v-btn(icon).white--text
      v-icon fab fa-vk
    v-spacer
    v-card(
      flat, tile
      flex
      color="success"
    )
      v-card-text.white--text
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
    bottom: -53px;
  }
</style>