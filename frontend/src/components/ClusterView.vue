<template>
  <div class="cluste-view">
    <div v-if="pods" class="content">
      <ul>
        <li v-for="value in pods.data" :key="value.name">
          <p>Pod name: {{ value.name }},  Pod IP: {{ value.podIP }}, Status: {{ value.phase }}, Age: {{ convertTime(value.startTime) }}</p>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'clusterView',
  data () {
    return {
      pods: null
    }
  },
  created () {
    this.fetchData()
  },
  watch: {
    // call again the method if the route changes
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      this.axios
        .get('/api/getPods')
        .then(response => (
          this.pods = response
        ))
        .catch(error => {
          console.log(error)
        })
    },
    convertTime (timeStamp) {
      // FIXME: fix timestamp calculation
      var time = new Date()
      time = Date.now() - timeStamp

      var result = new Date(time)

      var days = result.getDay()
      // Hours part from the timestamp
      var hours = result.getHours()
      // Minutes part from the timestamp
      var minutes = '0' + result.getMinutes()
      // Seconds part from the timestamp
      var seconds = '0' + result.getSeconds()

      var formattedTime = days + ' day(s) ' + hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2)

      return formattedTime
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  width: 60%;
  margin: 0 auto;
  padding-bottom: 20px;
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
