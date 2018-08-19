import axios from 'axios'
import moment from 'moment'

const Api = {}
const api = '/api'

Api.install = (Vue, {store, R}) => {
  Vue.prototype.moment = moment
  Vue.prototype.$api = {

    send: async (type, path, data, params) => {
      let headers = {
        'Authorization': 'Bearer ' + store.getters.getToken
      }
      if (R.isNil(data)) {
        if (params) {
          headers.params = params
        }
      }
      let result = {}
      try {
        if (data) {
          result = await axios[type](api + path, data, { headers })
        } else {
          console.log('here', headers)
          result = await axios[type](api + path, headers)
        }
        if (result.status !== 200) {
          result = {'error': true}
        }
      } catch (err) {
        result = {
          'error': true,
          'message': err
        }
      }

      return result
    }

  }
}

export default Api
