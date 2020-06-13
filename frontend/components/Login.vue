<template>
  <v-row no-gutters class="fill-height transition-ease-in-out" justify="center" align-content="center">
    <v-col class="primary fill-height" cols="4">
      <v-row class="fill-height" justify="center" align-content="center" no-gutters>
        <create-account-banner />
      </v-row>
    </v-col>
    <v-col cols="8">
      <v-row no-gutters class="fill-height" justify="center" align-content="center">
        <v-col cols="6">
          <v-row justify="center" align-content="center">
            <v-icon class="logo" color="primary">
              mdi-comment
            </v-icon>
          </v-row>
          <v-row justify="center" align-content="center">
            <p class="display-2">
              Sign In
            </p>
          </v-row>
          <login-form />
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import LoginForm from '~/components/LoginForm'
import CreateAccountBanner from '~/components/CreateAccountBanner'

export default {
  name: 'Login',
  components: {
    'login-form': LoginForm,
    'create-account-banner': CreateAccountBanner
  }
}
</script>

<style lang="scss" scoped>
  .logo {
    font-size: 100px;
  }
</style>
