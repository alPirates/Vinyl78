<template lang="pug">
  v-toolbar(fixed dark color="primary" height="70")
    v-toolbar-side-icon
    v-toolbar-title Vinil78
    v-spacer
    v-toolbar-items.hidden-sm-and-down
      v-btn(flat) One
      v-btn(flat) Two
      v-btn(flat) Three
</template>

<script>
export default {
  name: 'Header',
  data: () => {
    return {
    }
  }
}
</script>
