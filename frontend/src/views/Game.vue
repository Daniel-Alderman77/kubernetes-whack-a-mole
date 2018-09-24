<template>
  <div class="game">
    <h1>Game</h1>

    <button class="game-button" alt="Pink Gopher" :style="{'background-image': 'url(' + require('../assets/pink-gopher.png') + ')'}" v-on:click="click(1)"></button>
    <button class="game-button" alt="Hipster Gopher" :style="{'background-image': 'url(' + require('../assets/hipster-gopher.png') + ')'}" v-on:click="click(2)"></button>
    <button class="game-button" alt="Tophat Gopher" :style="{'background-image': 'url(' + require('../assets/tophat-gopher.png') + ')'}" v-on:click="click(3)"></button>
    <br>
    <button class="game-button" alt="Buffalo Gopher" :style="{'background-image': 'url(' + require('../assets/buffalo-gopher.png') + ')'}" v-on:click="click(4)"></button>
    <button class="game-button" alt="Hawaiian Gopher" :style="{'background-image': 'url(' + require('../assets/hawaiian-gopher.png') + ')'}" v-on:click="click(5)"></button>
    <button class="game-button" alt="Bat Gopher" :style="{'background-image': 'url(' + require('../assets/bat-gopher.png') + ')'}" v-on:click="click(6)"></button>
    <br>
    <span>Gopher's clicked: {{clickCount}}</span>
  </div>
</template>

<script>
export default {
  data () {
    return {
      clickCount: 0
    }
  },
  methods: {
    click: function (gopherClicked, event) {
      alert(gopherClicked)
      this.clickCount += 1
    }
  }
}
</script>

<style lang="scss">
.game-button {
  width: 100px;
  height: 110px;
  background-size: cover;
  border: none;
  margin: 20px;
}
.menu-button {
  padding: 20px 40px 20px;
  font-family : inherit;
  font-size: 1em;
  background-color: #42b983;
  color: #fff;
  font-weight: bold;
}
span {
  font-family : inherit;
  font-size: 1em;
  font-weight: bold;
}
</style>
