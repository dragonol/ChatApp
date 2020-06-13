<template>
  <v-container fluid class="pa-0 fill-height">
    <login />
  </v-container>
</template>

<script>
import Login from '~/components/Login'

export default {
  name: 'Index',
  components: {
    login: Login
  },
  data () {
    return {
      isShow: false
    }
  }
}
</script>

<style lang="scss" scoped>
  .logo {
    font-size: 100px;
  }
</style>
