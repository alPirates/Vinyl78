<template lang="pug">
  v-container(grid-list-lg)
    AdminNavigation(v-if="isAdmin()")
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
import EditUser from '@/components/Edit/EditUser'
import AdminNavigation from '@/components/Navigation/AdminNavigation'

export default {
  name: 'Settings',
  data: () => {
    return {}
  },
  components: {
    EditUser,
    AdminNavigation
  }
}
</script>
