<template>
  <div class="main">
        <ul class="list">
          <div class="statement">
            <p style="color: #ffb400">2022-05-30已更新</p>
            <p style="margin-top:6px">榜单规则: 将昨日国内热映的影片, 按照评分从高到低排列取前10名, 每天上午10点更新。相关数据来源于“王伟统计”及“赵锦阳统计”。</p>
          </div>
          <li 
          @click="toMovie(item.id)"
          class="item" v-for="(item,index) in movie" :key="item.id">
            <div class="aside">
              <img v-if="index==0" src="../assets/img/1.png" alt="" class="ranking">
              <div v-else-if="index!==0" class="ranking rankingimg">
                <p>{{index+1}}</p>
              </div>
              <img :src="item.avatar" alt="cover" class="cover">
            </div>
            <div class="describe">  
                <a>{{item.name}}</a>
                <p>主演：{{item.actors.slice(0,3).join(',')}}</p>
                <p style="margin-top:8px;  opacity:0.7 ;">上映时间: {{item.period.slice(0,10)}}</p>
            </div>
            <div class="score">
                  <i class="integer">{{parseInt(item.score)}}.</i>
                  <i class="fraction">{{parseInt((item.score-parseInt(item.score))*10)}}</i>
            </div>
          </li>
        </ul>
        <div class="page">
          <el-pagination
          background
          layout="prev, pager, next"
          @current-change="handlePage"
          :page-size="10"
          :total="50">
          </el-pagination>
      </div>
      </div>
</template>

<script>
export default {
    data() {
      return{
        movie:[],
      }
    },
    methods: {
      toMovie(movie_id){
        console.log(movie_id);
        this.$router.push({name:'thefilm',query:{movie_id}})
      },
      handlePage(val){
        this.$api.getMovieListByfollow({
          page:val,
        }).then(res=>{
          this.movie=res.data.data.list
        },error=>{
          console.log(error)
        })
      }
    },
    created() {
      this.$api.getMovieListByfollow({
        page:1,
      }).then(res=>{
        this.movie=res.data.data.list
      },error=>{
        console.log(error)
      })
    }
}
</script>

<style scoped>
.page {
  flex-basis: 100%;
  text-align: center;
}
.main{
  width: 100%;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}
.statement{
  text-align: center;
  font-size: 12px;
  color: #999;
  margin:70px 0 40px 0;
}
.item{
  list-style: none;
  width:950px;
  height:220px;
  background-color:#fff;
  display:flex;
  align-items: center;
  border-bottom: 0.5px solid rgb(181, 176, 176);
  cursor: pointer;
}
.list{
  margin-bottom:100px;
}
.aside{
  display:flex;
  align-items: center;
  position: relative;
}
.ranking{
  display: inline-block;
    width: 50px;
    height: 50px;
    background: #f7f7f7;
    font-size: 18px;
    color: #999;
    font-weight: 700;
    position: absolute;
    left: 0;
    top: 85px;
    display:flex;
    align-items: center;
    justify-content: center;
}
.cover{
  width:160px;
  margin-left:80px;
  height:220px;
}
.describe a{
    font-size: 26px;
    color: #333;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    width: 420px;
    display: block;
    text-decoration: none;
    margin-bottom:18px;
}
.describe a:hover{
      white-space: normal;
}
.describe{
  margin:0 30px;
}
.score{
  margin-left:60px
}
.integer{
  color: #ffb400;
  font-size: 56px;
  font-weight: 700;
  font-style: italic;
}
.fraction{
  color: #ffb400;
  font-size: 26px;
    font-weight: 700;
  font-style: italic;
}

</style>