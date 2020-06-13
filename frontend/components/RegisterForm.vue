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
      <v-text-field
        v-model="confirmedPassword"
        :rules="confirmedPasswordRules"
        label="Re-enter Password"
        type="password"
        filled
        prepend-icon="mdi-lock"
      />
      <v-row justify="center" no-gutters>
        <v-btn min-width="10vw" color="primary" :loading="signingUp" @click="signUp">
          Sign Up
        </v-btn>
      </v-row>
    </v-form>
  </v-container>
</template>

<script>
export default {
  name: 'RegisterForm',
  data () {
    return {
      email: '',
      password: '',
      confirmedPassword: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
      ],
      passwordRules: [
        v => !!v || 'Password is required'
      ],
      confirmedPasswordRules: [
        v => !!v || 'Must confirm password',
        v => this.password === this.confirmedPassword || 'Password is not match'
      ],
      signingUp: false
    }
  },
  methods: {
    signUp () {

    }
  }
}
</script>

<style scoped>

</style>
