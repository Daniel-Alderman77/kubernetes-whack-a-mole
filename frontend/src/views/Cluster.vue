<template>
  <div class="cluster">
    <h1>Cluster Info</h1>

    <ClusterView/>
  </div>
</template>

<script>
// @ is an alias to /src
import ClusterView from '@/components/ClusterView.vue'

export default {
  name: 'cluster',
  components: {
    ClusterView
  }
}
</script>
