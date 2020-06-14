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
  login ({ commit }) {

  }
}
