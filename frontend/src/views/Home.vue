<template>
  <div class="home">
    <img class="logo" alt="Kubernetes Whack-A-Mole logo" src="../assets/logo.png">
    <HomeContent/>
    <router-link class="play-button" to="/game" tag="button">Play the Game!</router-link>
  </div>
</template>

<script>
// @ is an alias to /src
import HomeContent from '@/components/HomeContent.vue'

export default {
  name: 'home',
  components: {
    HomeContent
  }
}
</script>

<style lang="scss">
.logo {
  // Centre logo
  margin-left: -50px;
}
.play-button {
  padding: 20px 40px 20px;
  font-family : inherit;
  font-size: 1em;
  background-color: #42b983;
}
</style>
