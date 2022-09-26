<template>
  <div>
    <Myheader/>
    <router-view></router-view>
    <Footer></Footer>
  </div>
</template>

<script>
import Myheader from "./Myheader.vue";
import Footer from "./Footer.vue"
export default {
    components:{
      Myheader,
      Footer
    },
    // created () {
    //   if (window.localStorage) {
        
    //   }
    // }
}
</script>

<style scoped>

</style>