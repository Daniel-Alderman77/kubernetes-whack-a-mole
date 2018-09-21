<template>
  <div class="content">
    <h1>Kubernetes Whack-A-Mole</h1>
    <p>
      A simple Whack-A-Mole game demonstrating the self-healing properties of Kubernetes applications by utilizing
      <a href="https://github.com/kubernetes/client-go" target="_blank" rel="noopener">Kubernetes Go Client Library</a>.
    </p>
  </div>
</template>

<script>
export default {
  name: 'HomeContent',
  props: {
    msg: String
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
