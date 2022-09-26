<template>
  <div>
    <div class="searchBox">
      <div class="search">
            <input class="kw" type="text"  
            @keydown.enter="changeKey"
            v-model="keyValue"
            maxlength="32"  placeholder="找影视剧、影人、影院" >
            <button class="submit" 
            @click="changeKey"
            type="submit"></button>
        </div>
    </div>
    <div class="result">
      <ul class="resultList">
        <li class="resultItem"
        @click="toFilm(item.id)"
        v-for="item in movieList" :key="item.id">
          <img :src="item.avatar" alt="电影海报">
          <div class="movieDetails">
            <div class="name">
              {{item.name}}
            </div>
            <div class="nickname">
              {{item.alias_name}}
            </div>
            
            <div class="score">
              {{item.score.toFixed(1)}}
            </div>
            <div class="label">
              奇幻,剧情
            </div>
            <div class="time">
              {{item.period.slice(0,10)}} {{item.period.slice(11,16)}}中国大陆上映
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import request from '../api/index'
export default {
  data() {
    return {
      keyValue: '',
      movieList: []
    }
  },
  methods: {
    changeKey(){
      if(this.keyValue!='')
      {
        request.getMovieBykey({
          key:this.keyValue,
          page:1,
          page_size:20,
        }).then(res=>{
          this.movieList= res.data.data.list;
          this.keyValue=''
        },
        err=>{
          console.log(err);
        })
      }
    },
    toFilm(movie_id){
      this.$router.push({name: 'thefilm',query:{movie_id,}})
    },
  },
  created(){
    request.getMovieBykey({
      key:this.$route.query.keyvalue,
      page:1,
      page_size:20,
    }).then(res=>{
      this.movieList= res.data.data.list;
    },
    err=>{
      console.log(err);
    })
  }
}
</script>

<style scoped>
.time{
    color: #999;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    margin-top:4px;
    border-bottom: 1px solid #e5e5e5;
    padding-bottom:52.4px;
}
.label{
    margin-top: 5px;
    font-size: 16px;
    color: #333;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}
.name{
    width: 160px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    color: #333;
    font-size: 26px;
    width: auto;
    margin-top: 27px;
}
.score{
    width: 160px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    margin-top: 10px;
    font-size: 26px;
    font-weight: 700;
    color: #ffb400;
}
.nickname{
  font-size: 14px;
    color: #999;
    margin-top: 3px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    padding-bottom: 1px;
    margin-bottom: -1px;
}
.result{
  width:980px;
  margin:0 auto;
}
.resultList{
  width:100%;
  list-style: none;
  overflow: hidden;
  margin:0;
}
.resultItem{
  width: 460px;
  height:220px;
  float: left;
  margin-top: 40px;
  margin-left: 10px;
  line-height: 1.2;
  position: relative;
}
.resultItem img{
  float: left;
    margin-right: 20px;
    width: 158.5px; 
    height: 220px;
}
.search{
  position: relative;
}
.searchBox {
    width: 100%;
    min-width: 1200px;
    background: #47464a;
    height: 156px;
    display: flex;
    justify-content: center;
    align-items: center;
}
.kw{
    box-sizing: border-box;
    line-height: 18px;
    width: 630px;
    padding: 15px 50px 15px 20px;
    border: 1px solid #999;
    border-radius: 50px;
    outline:none;
}
.submit{
    position: absolute;
    border: none;
    left: auto;
    top: 12px;
    right: 13px;
    width: 27px;
    height: 27px;
    cursor: pointer;
    background: url('../assets/img/searchlogo.png') no-repeat 50%;
    background-size: contain;
}
</style>