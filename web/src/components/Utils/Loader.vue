<template lang="pug">
  v-dialog(
    persistent
    :width="300"
    v-model="STATE.loader"
  )
    v-card(color="primary")
      v-card-text.text-xs-center
        h3.white--text Пожалуйста подождите...
        v-progress-linear(indeterminate color="white").mb-0
</template>

<script>
export default {
  name: 'Loader',
  data: () => {
    return {
    }
  }
}
</script>
