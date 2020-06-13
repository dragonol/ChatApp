<template>
  <v-container fluid>
    <v-form>
      <v-text-field
        v-model="email"
        :rules="emailRules"
        label="Email"
        type="email"
        filled
        prepend-icon="mdi-mail"
      />
      <v-text-field
        v-model="password"
        :rules="passwordRules"
        label="Password"
        type="password"
        filled
        prepend-icon="mdi-lock"
      />
      <v-row justify="center" no-gutters>
        <v-btn min-width="10vw" color="primary" :loading="signingIn" @click="signIn">
          Sign In
        </v-btn>
      </v-row>
      <v-row justify="center" align-content="center" no-gutters>
        <v-list>
          <v-list-item>
            <v-btn color="blue" large icon rounded>
              <v-icon>
                mdi-twitter
              </v-icon>
            </v-btn>
          </v-list-item>
        </v-list>
      </v-row>
    </v-form>
  </v-container>
</template>

<script>
export default {
  name: 'LoginForm',
  data () {
    return {
      email: '',
      password: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
      ],
      passwordRules: [
        v => !!v || 'Password is required'
      ],
      signingIn: false
    }
  },
  methods: {
    signIn () {

    }
  }
}
</script>

<style scoped>

</style>
