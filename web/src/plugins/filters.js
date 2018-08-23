
const filters = {
  install (Vue, { moment }) {
    Vue.mixin({
      filters: {
        date (value) {
          return moment(value).format('DD.MM.YY hh:mm')
        },
        status (value) {
          switch (value) {
            case 0:
              return 'На проверке'
            case 1:
              return 'Проверен'
            case 2: 
              return 'Отказан'
          }
        }
      }
    })
  }
}

export default filters