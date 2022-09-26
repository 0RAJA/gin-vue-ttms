<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import start from "./components/start.vue"
import register from "./views/Register.vue"
export default {
  name: 'App',
  data(){
    return {
      isLogin: 'start',
    }
  },
  components: {
    start,
    register
  },
  methods:{

  },
}
</script>
  
<style>
@import "./assets/css/normalize.css";
</style>
