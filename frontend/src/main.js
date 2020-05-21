import Vue from 'vue'
import router from './router'
import loginPage from './views/Login.vue'
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(loginPage)
}).$mount('#app')
