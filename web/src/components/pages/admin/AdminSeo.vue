<template lang="pug">
  v-container(grid-list-lg)
    AdminNavigation()
    v-layout(row, wrap)
      v-flex(xs12)
        v-card()
          v-toolbar(color="primary").white--text
            v-toolbar-title SEO
          v-card-text
            v-layout(row, wrap)
              v-flex(xs12)
                v-data-table(:headers="headers", :items="content" v-if="content.length > 0").elevation-1
                  template(slot="items" slot-scope="el")
                    td {{el.item.site}}
                    td {{el.item.visits}}
            v-layout(row, justify-end)
              v-btn(color="error" @click="resetCounters") Сбросить счетчики
</template>

<script>
import AdminNavigation from '@/components/Navigation/AdminNavigation'

export default {
  name: 'AdminCategory',
  data: () => {
    return {
      headers: [{
        text: 'Сайт',
        value: 'site'
      }, {
        text: 'Переходы',
        value: 'visits'
      }],
      content: [
        {
          site: 'Vk',
          visits: '0'
        },
        {
          site: 'Instagram',
          visits: '0'
        }
      ]
    }
  },
  methods: {
    async getSeoLinks () {
      let vk = await this.$api.send('get', '/app/property', null, {
        key: 'vk'
      })
      let inst = await this.$api.send('get', '/app/property', null, {
        key: 'inst'
      })
      if (vk.data.value && inst.data.value) {
        this.content[0].visits = vk.data.value
        this.content[1].visits = inst.data.value
      }
    },
    async resetCounters () {
      await this.$api.send('put', '/app/property', {
        key: 'vk',
        value: '0'
      })
      await this.$api.send('put', '/app/property', {
        key: 'inst',
        value: '0'
      })
      await this.getSeoLinks()
    }
  },
  async mounted () {
    await this.getSeoLinks()
  },
  components: {
    AdminNavigation
  }
}
</script>
