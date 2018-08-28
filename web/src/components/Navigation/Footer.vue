<template lang="pug">
  v-footer(
    height="auto"
    app
  )
    v-card(
      flat, tile
      flex
    ).text-xs-center
      v-card-text
        v-btn(icon) Vk
        v-btn(icon) Ist
      v-divider
      v-card-text
        | &copy;2018 —
        strong Vuetify
</template>

<script>
export default {
  name: 'Footer',
  data: () => {
    return {}
  }
}
</script>
