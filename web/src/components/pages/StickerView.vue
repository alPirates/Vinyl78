<template lang="pug">
  v-container(gridlist-sm)
    v-layout
      v-flex(xs12)
        h2 {{sticker.description}}
        p Hello something ehre
        .fotorama(data-nav="thumbs")
          img(:src="getMedia(el.name)" v-for="(el, index) in sticker.images")

</template>

<script>
import 'fotorama/fotorama.css'
import 'fotorama/fotorama.js'

export default {
  data () {
    return {
      sticker: {}
    }
  },
  methods: {
    async update () {
      let result = await this.$api.send('get', `/sticker/${this.$route.params.id}`)
      if (result) {
        console.log(result.data.resultd)
        this.sticker = result.data.result
      }
    }
  },
  async mounted () {
    await this.update()
    console.log('loaded')

    $('.fotorama').fotorama()
  }

}
</script>

<style>
  .fotorama__nav__shaft {
    float: left;
  }
</style>
