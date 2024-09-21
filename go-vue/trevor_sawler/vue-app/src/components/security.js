import { store } from '@/components/store'
import router from '@/router/index'

let Security = {
  // make sure user is authenticated
  requireToken: function () {
    if (store.token === '') {
      router.push('/login')
      return false
    }
  },

  // create request options and send them back
  requestOptions: function (payload) {
    const headers = new Headers()
    headers.append('Content-Type', 'application/json')
    headers.append('Authorization', 'Bearer ' + store.token)

    return {
      method: 'POST',
      body: JSON.stringify(payload),
      headers
    }
  },

  // check token (navigation guard)
  checkToken: function () {
    if (store.token !== '') {
      const payload = {
        token: store.token
      }

      const headers = new Headers()
      headers.append('Content-Type', 'application/json')

      let requestOptions = {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: headers
      }

      fetch(`${import.meta.env.VITE_APP_API_URL}/validate-token`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          if (data.error) {
            console.log(data.error)
          } else {
            if (!data.data) {
              console.log('removing token...')
              store.token = ''
              store.user = {}
              document.cookie =
                '_site_data=; Path=/; ' +
                'SameSite=strict; Secure; ' +
                'Expires=Thu, 01 Jan 1970 00:00:01 GMT;'
            }
          }
        })
    }
  }
}

export default Security
