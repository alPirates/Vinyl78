<template lang="pug">
  div
    v-carousel
      v-carousel-item(v-for="(image, index) in carouselImages", :key="index")
        img(:src="getMedia(image.name)")
    v-container(grid-list-lg)
      v-layout(row wrap)
        v-flex(xs12, sm6, md4)
          v-card
            v-layout(row, justify-center, align-center)
              v-avatar(size="128").text-xs-center
                img(src="@/assets/logo.png")
            v-card-title
              div.text-xs-center
                h3.headline.mb-0.text-xs-center Заголовок
                p Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur
                v-btn(color="danger")
                  | Больше
                v-btn(color="success", to="/form")
                  | Купить
        v-flex(xs12, sm6, md4)
          v-card
            v-layout(row, justify-center, align-center)
              v-avatar(size="128").text-xs-center
                img(src="@/assets/logo.png")
            v-card-title
              div.text-xs-center
                h3.headline.mb-0.text-xs-center Заголовок
                p Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur
                v-btn(color="danger")
                  | Больше
                v-btn(color="success", to="/form")
                  | Купить
        v-flex(xs12, sm6, md4)
          v-card
            v-layout(row, justify-center, align-center)
              v-avatar(size="128").text-xs-center
                img(src="@/assets/logo.png")
            v-card-title
              div.text-xs-center
                h3.headline.mb-0.text-xs-center Заголовок
                p Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur
                v-btn(color="danger")
                  | Больше
                v-btn(color="success", to="/form")
                  | Купить
      v-divider.mt-2.mb-2
      v-container(grid-list-sm)
        v-card
          v-toolbar(color="primary")
            v-toolbar-title.white--text Заголовок
          v-card-content.mt-2
            v-container(grid-list-sm)
              v-layout(row, wrap)
                v-flex(xs12)
                  h3.display-1 Наши преимущества
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
