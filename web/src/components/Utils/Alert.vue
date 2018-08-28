<template lang="pug">
  v-snackbar(
    :bottom="true"
    :multi-line="true"
    v-model="STATE.alert.show"
  ) {{SATTE.alert.message}}
    v-btn(
      color="pink"
      flat
      @click="STATE.alert.show = false"
    )
</template>

<script>
export default {
  name: 'Alert',
  data: () => {
    return {
    }
  }
}
</script>
