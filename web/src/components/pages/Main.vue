<template lang="pug">
  div
    v-carousel(v-if="carouselImages.length > 0")
      v-carousel-item(v-for="(image, index) in carouselImages", :key="index")
        img(:src="getMedia(image.name)")
    v-container(grid-list-lg)
      v-layout(row wrap)
        v-flex(xs12, sm6, md4 v-for="(el, index) in thumbnails", :key="index")
          v-card
            v-layout(row, justify-center, align-center)
              v-avatar(size="128").text-xs-center
                img(:src="getImage(el)")
            v-card-title
              div.text-xs-center.m-top
                h3.headline.mb-0.text-xs-center {{el.name}}
                p {{el.description}}
                v-btn(color="danger", :to="'/category/' + el.id")
                  | Больше
                v-btn(color="success", to="/form")
                  | Купить
      v-divider.mt-2.mb-2
      v-container(grid-list-sm)
        v-card
          v-toolbar(color="primary")
            v-toolbar-title.white--text Заголовок
          v-card-text.mt-2
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
      carouselImages: [],
      thumbnails: []
    }
  },
  async mounted () {
    let mainCategories = await this.$api.send('get', '/main_sidebar')
    this.$set(this, 'thumbnails', mainCategories.data.result)

    let carousel = await this.$api.send('get', '/property', null, {
      key: 'carousel_id'
    })
    console.log('carousel is ', carousel)
    let carouselImages = await this.$api.send('get', '/image', null, {
      linked_id: carousel.data.value
    })

    if (carouselImages) {
      this.$set(this, 'carouselImages', carouselImages.data.images)
    }
  }
}
</script>

<style scoped>
  .m-top {
    margin: 0 auto;
  }
</style>
