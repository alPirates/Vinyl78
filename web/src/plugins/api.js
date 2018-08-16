import axios from 'axios'
import moment from 'moment'

const Api = {}
const api = '/api'

Api.install = (Vue, {store}) => {
  Vue.prototype.moment = moment
  Vue.prototype.$api = {

    send: async (type, path, data) => {
      let headers = {
        // 'Authorization': 'bearer ' + store.getters.getToken
      }

      let result = {}
      try {
        if (data) {
          result = await axios[type](api + path, data, { headers })
        } else {
          result = await axios[type](api + path, { headers })
        }
        if (result.status !== 200) {
          result = {'error': true}
        }
      } catch (err) {
        result = {'error': true}
      }

      return result
    }

  }
}

export default Api
