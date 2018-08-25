<template lang="pug">
  v-container(grid-list-sm)
    | edit here
    pre {{sticker}}
    FileUpload(
      v-model="files"
      :data="{linked_id: sticker.id}"
    )
</template>

<script>
import FileUpload from '@/components/Utils/FileUpload'
export default {
  name: '',
  data: () => {
    return {
      files: []
    }
  },
  methods: {
  },
  props: {
    sticker: {
      type: Object,
      default: () => {
        return {}
      }
    }
  },
  components: {
    FileUpload
  }
}
</script>
