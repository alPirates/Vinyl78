<template lang="pug">
  v-container(grid-list-sm)
    pre token is {{STATE.token}}
    Login(v-if="!STATE.token")
    AdminPanel(v-else)
</template>

<script>
import Login from '@/components/Utils/Login'
import AdminPanel from '@/components/Utils/AdminPanel'

export default {
  name: 'Admin',
  data: () => {
    return {
    }
  },
  components: {
    Login,
    AdminPanel
  }
}
</script>
