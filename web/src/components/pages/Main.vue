<template lang="pug">
  v-container(grid-list-sm).fixed-container
    //- v-carousel(v-if="carouselImages.length > 0" hide-delimiters)
    //-   v-carousel-item(v-for="(image, index) in carouselImages", :key="index")
    //-     a(:href="image.name")
    //-       img(:src="getMedia(image.image[0].name)" v-if="image.image.length > 0")
    div(ref="sl")
    carousel(
      v-if="carouselImages.length > 0",
      :perPage="1",
      :autoplay="true",
      :autoplayTimeout="4000",
      :loop="true"
    )
      slide(v-for="(image, index) in carouselImages", :key="index")
        a(:href="image.name")
          img(:src="getMedia(image.image[0].name)" v-if="image.image.length > 0", :width="w", :height="h")

    v-container(grid-list-lg)
      Icons(
        v-if="!R.isEmpty(thumbnails)"
        :links="getLinks(thumbnails)"
      )
      v-divider.mt-3.mb-2
      Form
      v-layout(row, wrap, justify-center)#links
        v-flex(xs12)
        a.links.big-icon.mr-4(target="_blank" href="https://www.instagram.com/vinyl_78")
          v-icon fab fa-instagram
        a.links.big-icon(target="_blank" href="https://vk.com/vinyl_78")
          v-icon fab fa-vk
</template>

<script>
import Form from '@/components/pages/Form'
import Icons from '@/components/pages/Icons'

export default {
  name: 'Main',
  data: () => {
    return {
      carouselImages: [],
      thumbnails: [],
      w: 0,
      h: 0
    }
  },
  computed: {
  },
  async mounted () {
    console.log(this.$refs)
    let mainCategories = await this.$api.send('get', '/main_sidebar')
    this.$set(this, 'thumbnails', mainCategories.data.result)

    let carouselImages = await this.$api.send('get', '/carousel', null, {
      limit: 100,
      skip: 0
    })

    if (carouselImages) {
      this.$set(this, 'carouselImages', carouselImages.data.result)
    }
    console.log('images', carouselImages)
  },
  methods: {
    getLinks () {
      if (R.isEmpty(this.thumbnails)) {
        return []
      }
      return R.map(el => el.id, this.thumbnails)
    },
    getSize () {
      this.w = this.$refs.sl.clientWidth
      this.h = this.$refs.sl.clientWidth / 16 * 9
    }
  },
  components: {
    Form, Icons
  }
}
</script>

<style scoped>
  .fixed-container {
    /* width: 940px; */
  }
  .m-top {
    margin: 0 auto;
  }
  .low-text {
    font-weight: 300;
    font-size: 27px !important;
  }
  .links {
    text-decoration: none;
  }
  .big-icon i {
    font-size: 70px !important;
    color: black;
  }
</style>
