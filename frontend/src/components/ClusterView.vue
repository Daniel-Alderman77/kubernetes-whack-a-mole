<template>
  <div class="cluste-view">
    <div v-if="pods" class="content">
      {{ pods }}
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
        .get('http://127.0.0.1:8080/getPods')
        .then(response => (
          this.pods = response
        ))
        .catch(error => {
          console.log(error)
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
p {
  width: 60%;
  margin: 0 auto;
  padding-bottom: 20px;
}
h3 {
  margin: 40px 0 0;
}
ul {
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
