<template lang="pug">
  v-container(grid-list-sm)
    v-layout(row, wrap)  
      v-flex(xs12)
        | Hello
</template>

<script>
  export default {
    name: 'EditUser',
    data: () => {
      return {}
    },
    mounted () {
      if (R.isEmpty(this.STATE.token)) {
        console.log('user is unauthorized')
        return
      }
    }
  }
</script>