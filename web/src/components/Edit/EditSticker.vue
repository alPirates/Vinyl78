<template lang="pug">
  v-container(grid-list-sm)
    | edit here
    pre {{sticker}}
</template>

<script>
export default {
  name: '',
  data: () => {
    return {}
  },
  methods: {
  },
  props: {
    sticker: {
      type: Object,
      default: () => {
        return {}
      }
    }
  }
}
</script>
