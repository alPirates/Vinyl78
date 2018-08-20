<template lang="pug">
.file-loader
  v-layout(row, wrap, align-end, justify-center).mt-2.mb-2
    v-flex(sm12).text-sm-center
      v-btn(:color="color")
        file-upload(
          ref="upload"
          :post-action="url"
          v-model="files"
          :multiple="true"
          :headers="headers"
          :extensions="getExtensions"
          :drop-directory="true"
          :drop="true"
        ) Загрузить
        v-icon(right) cloud_upload

    v-flex.text-sm-center(v-if="!($refs.upload && $refs.upload.uploaded)")
      ul.filenames(v-if="files.length")
        li(v-for="file in files") Файл: {{file.name}}
          img(v-if="file.thumb", :src="file.thumb" width="40" height="auto")
      v-btn(
        color="blue lighten-2"
        @click.prevent="$refs.upload.active = true"
        v-if="files.length"
      ) Начать загрузку

    v-flex.text-sm-center(v-else-if="showSuccess")
      h3 Успешно
      ul.filenames
        li(v-for="file in files")
         span(v-if="file.success") {{file.name}}
</template>

<script>
import FileUpload from 'vue-upload-component'

export default {
  name: 'FileLoader',
  data: () => {
    return {
      files: [],
      url: '/api/app/image',
      headers: {
        Authorization: 'Bearer' + this.getToken
      }
    }
  },
  watch: {
    files: {
      handler: function (oldVal, newVal) {
        let uploadedFiles = []
        uploadedFiles = this.files.reduce((acc, file) => {
          let len = file.response.data.length
          if (file.success && len > 0) {
            acc.push({
              filename: file.name,
              path: file.response.data[len - 1]
            })
          }
          return acc
        }, [])
        uploadedFiles = uploadedFiles || []
        this.$emit('input', uploadedFiles)
      },
      deep: true
    }
  },
  computed: {
    getExtensions () {
      if (this.images) {
        return ['jpg', 'png', 'jpeg']
      }
      return undefined
    }
  },
  mounted () {
    console.log('token is ', this.$store.getters.getToken)
  },
  props: {
    images: {
      type: Boolean,
      default: false
    },
    color: {
      type: String,
      default: 'success'
    },
    showSuccess: {
      type: Boolean,
      default: true
    },
    value: {
      type: Array,
      default: () => {
        return []
      }
    }
  },
  components: {
    FileUpload
  }
}
</script>

<style>
  .file-loader {
    position: relative;
    border: 2px dashed gray;
    border-radius: 3px;
  }

  .filenames {
    list-style: none;
  }
</style>
