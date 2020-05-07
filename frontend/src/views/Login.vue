<template>
  <div class="login-page">
    <div class="panel">
      <div class="welcome">
        <h1>Welcome</h1>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Eos perferendis voluptates ab natus non
          consequuntur corrupti quis
        </p>
      </div>
    </div>

    <div class="login">
      <h1>Login</h1>
      <form action method="post">
        <input type="text" />
        <input type="text" />
      </form>
    </div>
  </div>
</template>

<style scoped lang="scss">

.login-page {
  display: flex;
  height: 100vh;
  width: 100vw;

  .panel {
    width: 400px;
    height: 100vh;
    display: flex;
    align-items: flex-end;
    background: #01509e;

    .welcome {
      margin: 30px;
      color: white;
      font-family: quicksand;
    }
  }
}

</style>
