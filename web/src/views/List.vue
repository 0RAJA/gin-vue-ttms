<template>
  <div>
    <div class="header">
        <h3>
          <span @click="list='listByvisit'" :class="{textcolor:list=='listByvisit'}">热映口碑榜&nbsp;&nbsp;&nbsp;</span>
          <span @click="list='listByfollow'" :class="{textcolor:list=='listByfollow'}">最受期待榜</span>
        </h3>
    </div>
      <component :is="list"></component>
  </div>
</template>

<script>
import listByvisit from '../components/listByvisit.vue'
import listByfollow from "../components/listByfollow.vue"
export default {
    data() {
      return {
        list:'listByvisit'
      }
    },
    components: {
      listByvisit,
      listByfollow
    }
}
</script>

<style scoped>
.textcolor{
  color: #eae852;
}
.header {
  background: url(../assets/img/background2.jpg);
  display: flex;
  align-items: center;
  justify-content: center;
  color:#fff;
  height:60px;
  width:950px;
  margin:0 auto;
  cursor: pointer;
}

</style>