<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row, wrap)
      v-flex(xs12)
        v-tabs(
          centered
          color="primary"
          dark
          icons-and-text
        )
          v-tabs-slider(color="white")
          v-tab
            | Настройки
            v-icon settings
          v-tab-item
            v-card
              v-card-text
                EditUser
</template>

<script>
import EditUser from '@/components/Edit/EditUser.vue'

export default {
  name: 'Settings',
  data: () => {
    return {}
  },
  components: {
    EditUser
  }
}
</script>
