<template lang="pug">
  div
    v-layout(row, wrap)
      v-badge(color="red" right id="red-badge")
        span(slot="badge" v-if="reqCount > 0") {{reqCount}}
        v-btn(color="secondary", to="/admin") Заявки
      v-btn(color="secondary", to="/settings") Настройки
      v-btn(color="secondary", to="/admin/category") Категории
      v-btn(color="secondary", to="/admin/seo") SEO
    v-layout().mb-1
      v-divider
</template>

<script>
export default {
  name: 'AdminNavigation',
  data: () => {
    return {
    }
  },
  props: {
    reqCount: {
      type: Number,
      default: 0
    }
  }
}
</script>
