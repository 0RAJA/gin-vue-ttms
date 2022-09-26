<template>
  <div class="cardcup">
    <el-carousel :interval="4000" type="card" height="300px"  class="card">
        <el-carousel-item v-for="item,index in swiperImg" :key="index">
            <img :src='item.url' style="width:100%">
        </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
    name:'Swiper',
    data(){
        return {
            swiperImg:[
              {url: require('../assets/img/1.gif')},
              {url: require('../assets/img/2.gif')},
              {url: require('../assets/img/3.gif')},
              {url: require('../assets/img/4.gif')},
              {url: require('../assets/img/5.gif')}
            ],
        }
    },
}
</script>

<style scoped>
.cardcup{
    display:flex;
    padding-top: 30px;
    justify-content: center;
    background-color: #fff;
}
.card{
    width: 90%;
    margin: auto;
}
</style>