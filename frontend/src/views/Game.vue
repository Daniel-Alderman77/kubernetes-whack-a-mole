<template>
  <div class="game">
    <h1>Game</h1>

    <button class="game-button" alt="Pink Gopher" :style="{'background-image': 'url(' + require('../assets/pink-gopher.png') + ')'}" v-on:click="click(0)"></button>
    <span v-if="pods"> {{ this.pods[0].phase }} </span>
    <button class="game-button" alt="Hipster Gopher" :style="{'background-image': 'url(' + require('../assets/hipster-gopher.png') + ')'}" v-on:click="click(1)"></button>
    <span v-if="pods"> {{ this.pods[1].phase }} </span>
    <button class="game-button" alt="Tophat Gopher" :style="{'background-image': 'url(' + require('../assets/tophat-gopher.png') + ')'}" v-on:click="click(2)"></button>
    <span v-if="pods"> {{ this.pods[2].phase }} </span>
    <br>
    <button class="game-button" alt="Buffalo Gopher" :style="{'background-image': 'url(' + require('../assets/buffalo-gopher.png') + ')'}" v-on:click="click(3)"></button>
    <span v-if="pods"> {{ this.pods[3].phase }} </span>
    <button class="game-button" alt="Hawaiian Gopher" :style="{'background-image': 'url(' + require('../assets/hawaiian-gopher.png') + ')'}" v-on:click="click(4)"></button>
    <span v-if="pods"> {{ this.pods[4].phase }} </span>
    <button class="game-button" alt="Bat Gopher" :style="{'background-image': 'url(' + require('../assets/bat-gopher.png') + ')'}" v-on:click="click(5)"></button>
    <span v-if="pods"> {{ this.pods[5].phase }} </span>
    <br>
    <span>Gopher's clicked: {{clickCount}}</span>
  </div>
</template>

<script>
export default {
  data () {
    return {
      clickCount: 0,
      pods: null
    }
  },
  created () {
    this.fetchData()
  },
  computed: {
    podStatus: function (podID) {
      return this.pods[podID].phase
    }
  },
  watch: {
    // call again the method if the route changes
    '$route': 'fetchData'
  },
  methods: {
    click: function (gopherClicked, event) {
      this.deletePod(this.pods[gopherClicked].name, gopherClicked)
      this.clickCount += 1
    },
    sleep (ms) {
      return new Promise(resolve => setTimeout(resolve, ms))
    },
    fetchData () {
      this.axios
        .get('/api/getMolePods')
        .then(response => (
          this.pods = response.data
        ))
        .catch(error => {
          console.log(error)
        })
    },
    updateData (podID, gopherClicked) {
      this.axios
        .get('/api/getMolePods')
        .then(response => {
          this.pods.splice(gopherClicked, 1, response.data[6])
          console.log(this.pods[gopherClicked].phase)
        })
        .catch(error => {
          console.log(error)
        })
    },
    deletePod (podID, gopherClicked) {
      this.axios
        .get('/api/deletePod/' + podID)
        .then(response => {
          alert(podID + ' was deleted')
        })
        .catch(error => {
          console.log(error)
          this.errored = true
        })
        .finally(async () => {
          this.pods[gopherClicked].phase = 'Deleted'
          this.updateData(podID, gopherClicked)
          console.log(this.pods[gopherClicked].phase)
          this.sleep(2000)
          console.log(this.pods[gopherClicked].phase)
        })
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
