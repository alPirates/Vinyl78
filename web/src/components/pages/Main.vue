<template lang="pug">
  v-container(fluid)
    v-carousel
      v-carousel-item()
      v-carousel-item()
      v-carousel-item()
    v-container
      v-layout(row wrap)
        v-flex(xs12)
          h2 Main page
          p lorem text
          v-btn(:to="'/buy-sticker'")
            | купить стикеры
          v-btn(:to="'/form'")
            | Заказать стикер

</template>

<script>
export default {
  name: 'Main',
  data: () => {
    return {
    }
  }
}
</script>
