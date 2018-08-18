<template lang="pug">
  v-container
    v-list(dense subheader)
    v-subheader Menu
    v-list-tile
        v-list-tile-avatar
        v-icon add
        v-list-tile-content
        v-list-tile-title
            | but stickers
    v-divider
    v-list-tile
        v-list-tile-avatar
        v-icon delete
        v-list-tile-content
        v-list-tile-title
            | Category 2
</template>

<script>
export default {
  name: 'Sidebar',
  data: () => {
    return {
    }
  }
}
</script>
