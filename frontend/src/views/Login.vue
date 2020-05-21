<template>
  <div class="login-page">
    <div class="panel">
      <transition
        name="login-register-transition"
        enter-active-class="animated bounceInDown"
        leave-active-class="animated bounceOutDown"
      >
        <div v-if="screenPhase == 0" class="banner animation">
          <h1>Welcome</h1>

          <p>Enter your information and join us on new journey!</p>

          <button @click="changeScreen">
            <p>Register</p>
          </button>
        </div>
      </transition>
      <transition
        name="login-register-transition"
        enter-active-class="animated bounceInDown"
        leave-active-class="animated bounceOutDown"
      >
        <div v-if="screenPhase == 2" class="banner animation">
          <h1>Already have an account?</h1>

          <p>Sign in and start using Jusay</p>

          <button @click="changeScreen">
            <p>Sign in</p>
          </button>
        </div>
      </transition>
    </div>

    <div class="page">
      <transition
        name="login-register-transition"
        enter-active-class="animated bounceInUp"
        leave-active-class="animated bounceOutUp"
      >
        <div v-if="screenPhase == 0" class="form animation">
          <i class="fas fa-comment-alt logo" />

          <h1>Sign in to Jusay</h1>
          <div class="form-message">
            <div class="message">
              <p v-if="signIn_emptyField">Email or password must not be empty</p>
            </div>
            <div class="message">
              <p v-if="signIn_invalidEmail">Email address is not valid</p>
            </div>
            <div class="message">
              <p v-if="signIn_incorrectField">Email or password is incorrect</p>
            </div>
          </div>
          <form action method="post">
            <div class="input">
              <i class="fas fa-envelope" />
              <input v-model="signIn_email" type="text" placeholder="Email" />
            </div>
            <div class="input">
              <i class="fas fa-lock" />
              <input v-model="signIn_password" type="password" placeholder="Password" />
            </div>
          </form>

          <button @click="login">
            <p>Sign in</p>
          </button>
        </div>
      </transition>
      <transition
        name="login-register-transition"
        enter-active-class="animated bounceInUp"
        leave-active-class="animated bounceOutUp"
      >
        <div v-if="screenPhase == 2" class="form animation">
          <i class="fas fa-comment-alt logo" />

          <h1>Join Jusay</h1>
          <div class="form-message">
            <div class="message">
              <p v-if="signUp_emptyField">Email or password must not be empty</p>
            </div>
            <div class="message">
              <p v-if="signUp_invalidEmail">Email address is not valid</p>
            </div>
            <div class="message">
              <p v-if="signUp_mismatchPassword">Confirm password does not match password</p>
            </div>
            <div class="message">
              <p v-if="signUp_usedEmail">This email is already taken</p>
            </div>
          </div>
          <form action method="post">
            <div class="input">
              <i class="fas fa-envelope" />
              <input v-model="signUp_email" type="text" placeholder="Email" />
            </div>
            <div class="input">
              <i class="fas fa-lock" />
              <input v-model="signUp_password" type="password" placeholder="Password" />
            </div>
            <div class="input">
              <i class="fas fa-lock" />
              <input
                v-model="signUp_confirmPassword"
                type="password"
                placeholder="Confirm Password"
              />
            </div>
          </form>

          <button @click="register">
            <p>Sign up</p>
          </button>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
const axios = require('axios').default
export default {
  name: 'LoginPage',
  data: function () {
    return {
      screenPhase: 0,
      signIn_email: '',
      signIn_password: '',
      signUp_email: '',
      signUp_password: '',
      signUp_confirmPassword: '',
      signIn_emptyField: false,
      signIn_incorrectField: false,
      signIn_invalidEmail: false,
      signUp_emptyField: false,
      signUp_invalidEmail: false,
      signUp_mismatchPassword: false,
      signUp_usedEmail: false
    }
  },
  probs: {
  },
  methods: {
    changeScreen: function () {
      this.screenPhase = (this.screenPhase + 1) % 4
      const self = this
      setTimeout(function () {
        self.screenPhase = (self.screenPhase + 1) % 4
        console.log('run')
      }, 500)
    },
    validateEmail: function (s) {
      if (/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(s)) {
        console.log('is email')
        return (true)
      }
      return (false)
    },
    login: function () {
      this.signIn_emptyField = false
      this.signIn_incorrectField = false
      this.signIn_invalidEmail = false

      if (this.signIn_email === '' || this.signIn_password === '') {
        this.signIn_emptyField = true
      } else if (!this.validateEmail(this.signIn_email)) {
        this.signIn_invalidEmail = true
      } else {
        const postData = new FormData()
        postData.append('email', this.signIn_email)
        postData.append('password', this.signIn_password)

        const self = this

        axios({
          method: 'post',
          url: '/api/login',
          data: postData,
          headers: {
            'content-type': 'multipart/form-data'
          }
        }).then(function (response) {
          if (response.data.status === 'fail') {
            self.signIn_incorrectField = true
          } else {

          }
        })
      }
    },
    register: function () {
      this.signUp_emptyField = false
      this.signUp_usedEmail = false
      this.signUp_invalidEmail = false
      this.signUp_mismatchPassword = false

      if (this.signUp_email === '' || this.signUp_password === '') {
        this.signUp_emptyField = true
      } else if (!this.validateEmail(this.signUp_email)) {
        this.signUp_invalidEmail = true
      } else if (this.signUp_password !== this.signUp_confirmPassword) {
        this.signUp_mismatchPassword = true
      } else {
        const postData = new FormData()
        postData.append('email', this.signUp_email)
        postData.append('password', this.signUp_password)

        const self = this

        axios({
          method: 'post',
          url: '/api/register',
          data: postData,
          headers: {
            'content-type': 'multipart/form-data'
          }
        }).then(function (response) {
          if (response.data.status === 'fail') {
            self.signUp_usedEmail = true
          } else {

          }
        })
      }
    }
  }
}
</script>

<style scoped lang="scss">
.login-page {
  display: flex;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  p,
  h1,
  button {
    font-family: quicksand;
  }
  h1 {
    font-size: 45px;
  }
  p {
    font-size: 18px;
  }
  button {
    padding: 10px 40px 10px 40px;
    margin: 10px;
    font-size: 15px;
    border-radius: 30px;
    font-weight: bold;
    cursor: pointer;
    p {
      margin: 0;
      padding: 0;
    }
  }
  .animation {
    animation-duration: 0.5s;
  }
  .panel {
    width: 30vw;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    background: #01509e;
    @keyframes slideOutUp {
      100% {
        margin-bottom: 500px;
      }
    }
    .banner {
      width: 20vw;
      color: white;
      font-family: quicksand;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      text-align: center;
      button {
        color: #01509e;
        background: white;
        border: #01509e 1.5px solid;
      }
      button:hover {
        color: white;
        background: #01509e;
        border: white 1.5px solid;
      }
    }
  }
  .page {
    width: 70vw;
    height: 100vh;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .form {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;

      .form-message {
        font-weight: 600;
        height: 40px;
        .message {
          color: white;
          background: #e94836;
          border-radius: 5px;
        }
        p {
          padding: 5px 15px 5px 15px;
          font-size: 15px;
          margin: 0;
        }
      }
      .input {
        display: flex;
        justify-content: center;
        align-items: center;
        i {
          font-size: 30px;
          margin: 15px;
          color: #01509e;
        }
      }
      .logo {
        font-size: 70px;
        color: #01509e;
      }
      form {
        display: flex;
        flex-direction: column;
        input {
          margin: 5px;
          padding: 15px;
          font-family: quicksand;
          font-size: 18px;
          font-weight: 600;
          width: 400px;
          border-radius: 5px;
          border: 0.5px rgba(128, 128, 128, 0.041) solid;
          background: #01509e0e;
        }
      }
      button {
        color: white;
        background: #01509e;
        border: #01509e 1.5px solid;
        margin: 25px;
      }
      button:hover {
        color: #01509e;
        background: white;
        border: #01509e 1.5px solid;
      }
      h1 {
        color: #01509e;
        margin-bottom: 20px;
      }
    }
  }
}
</style>
