<template lang="pug">
  v-container(fluid)
    v-carousel
      v-carousel-item(v-for="(image, index) in carouselImages")
        img(:src="getMedia(image.name)")
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
      carouselImages: []
    }
  },
  async mounted () {
    let carousel = await this.$api.send('get', '/app/property', null, {
      key: 'carousel_id'
    })
    let carouselImages = await this.$api.send('get', '/image', null, {
      linked_id: carousel.data.value
    })
    console.log(carouselImages.data.images)
    if (carouselImages) {
      this.$set(this, 'carouselImages', carouselImages.data.images)
    }
  }
}
</script>
