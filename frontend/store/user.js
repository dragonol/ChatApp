import axios from '~/.nuxt/axios'

export const state = () => ({
  user: '',
  token: ''
})

export const mutations = {
  setUserData (state, userData) {
    axios.defaults.headers.common.Authorization = localStorage.getItem('token')
    state.user = userData
  }
}

export const getters = {

}

export const actions = {
  login ({ commit }, credentials) {
    return new Promise((resolve, reject) => {
      const postData = new FormData()
      postData.append('email', credentials.email)
      postData.append('password', credentials.password)
      axios({
        method: 'post',
        data: postData,
        headers: {
          'content-type': 'multipart/form-data'
        }
      }).then((response) => {
        commit('setUserData', response.data)
        resolve()
      }).catch((err) => {
        reject(err)
      })
    })
  }
}
