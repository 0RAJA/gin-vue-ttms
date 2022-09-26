<template>
  <div>
    <div class="sort">
      <ul class="selectbox">
        <li class="type">
          <div class="tagbox">类型：</div>
          <ul class="tags">
            <li @click="currentlabel.type='全部'" :class="{isSelected:currentlabel.type=='全部'}"><a href="#">全部</a></li>
            <li @click="currentlabel.type='爱情'" :class="{isSelected:currentlabel.type=='爱情'}"><a href="#">爱情</a></li>
            <li @click="currentlabel.type='喜剧'" :class="{isSelected:currentlabel.type=='喜剧'}"><a href="#">喜剧</a></li>
            <li @click="currentlabel.type='动画'" :class="{isSelected:currentlabel.type=='动画'}"><a href="#">动画</a></li>
            <li @click="currentlabel.type='剧情'" :class="{isSelected:currentlabel.type=='剧情'}"><a href="#">剧情</a></li>
            <li @click="currentlabel.type='恐怖'" :class="{isSelected:currentlabel.type=='恐怖'}"><a href="#">恐怖</a></li>
            <li @click="currentlabel.type='惊悚'" :class="{isSelected:currentlabel.type=='惊悚'}"><a href="#">惊悚</a></li>
            <li @click="currentlabel.type='科幻'" :class="{isSelected:currentlabel.type=='科幻'}"><a href="#">科幻</a></li>
            <li @click="currentlabel.type='动作'" :class="{isSelected:currentlabel.type=='动作'}"><a href="#">动作</a></li>
            <li @click="currentlabel.type='悬疑'" :class="{isSelected:currentlabel.type=='悬疑'}"><a href="#">悬疑</a></li>
            <li @click="currentlabel.type='犯罪'" :class="{isSelected:currentlabel.type=='犯罪'}"><a href="#">犯罪</a></li>
            <li @click="currentlabel.type='冒险'" :class="{isSelected:currentlabel.type=='冒险'}"><a href="#">冒险</a></li>
            <li @click="currentlabel.type='战争'" :class="{isSelected:currentlabel.type=='战争'}"><a href="#">战争</a></li>
            <li @click="currentlabel.type='奇幻'" :class="{isSelected:currentlabel.type=='奇幻'}"><a href="#">奇幻</a></li>
            <li @click="currentlabel.type='运动'" :class="{isSelected:currentlabel.type=='运动'}"><a href="#">运动</a></li>
            <li @click="currentlabel.type='家庭'" :class="{isSelected:currentlabel.type=='家庭'}"><a href="#">家庭</a></li>
            <li @click="currentlabel.type='古装'" :class="{isSelected:currentlabel.type=='古装'}"><a href="#">古装</a></li>
            <li @click="currentlabel.type='武侠'" :class="{isSelected:currentlabel.type=='武侠'}"><a href="#">武侠</a></li>
            <li @click="currentlabel.type='西部'" :class="{isSelected:currentlabel.type=='西部'}"><a href="#">西部</a></li>
            <li @click="currentlabel.type='历史'" :class="{isSelected:currentlabel.type=='历史'}"><a href="#">历史</a></li>
            <li @click="currentlabel.type='传记'" :class="{isSelected:currentlabel.type=='传记'}"><a href="#">传记</a></li>
            <li @click="currentlabel.type='歌舞'" :class="{isSelected:currentlabel.type=='歌舞'}"><a href="#">歌舞</a></li>
            <li @click="currentlabel.type='黑色电影'" :class="{isSelected:currentlabel.type=='黑色电影'}"><a href="#">黑色电影</a></li>
            <li @click="currentlabel.type='短片'" :class="{isSelected:currentlabel.type=='短片'}"><a href="#">短片</a></li>
            <li @click="currentlabel.type='纪录片'" :class="{isSelected:currentlabel.type=='纪录片'}"><a href="#">纪录片</a></li>
            <li @click="currentlabel.type='戏曲'" :class="{isSelected:currentlabel.type=='戏曲'}"><a href="#">戏曲</a></li>
            <li @click="currentlabel.type='音乐'" :class="{isSelected:currentlabel.type=='音乐'}"><a href="#">音乐</a></li>
            <li @click="currentlabel.type='灾难'" :class="{isSelected:currentlabel.type=='灾难'}"><a href="#">灾难</a></li>
            <li @click="currentlabel.type='青春'" :class="{isSelected:currentlabel.type=='青春'}"><a href="#">青春</a></li>
            <li @click="currentlabel.type='儿童'" :class="{isSelected:currentlabel.type=='儿童'}"><a href="#">儿童</a></li>
            <li @click="currentlabel.type='其他'" :class="{isSelected:currentlabel.type=='其他'}"><a href="#">其他</a></li>

          </ul>
        </li>
        <li class="region">
          <div class="tagbox">区域：</div>
          <ul class="tags">
            <li @click="currentlabel.region='全部'" :class="{isSelected:currentlabel.region=='全部'}"><a href="#">全部</a></li>
            <li @click="currentlabel.region='大陆'" :class="{isSelected:currentlabel.region=='大陆'}"><a href="#">大陆</a></li>
            <li @click="currentlabel.region='美国'" :class="{isSelected:currentlabel.region=='美国'}"><a href="#">美国</a></li>
            <li @click="currentlabel.region='韩国'" :class="{isSelected:currentlabel.region=='韩国'}"><a href="#">韩国</a></li>
            <li @click="currentlabel.region='日本'" :class="{isSelected:currentlabel.region=='日本'}"><a href="#">日本</a></li>
            <li @click="currentlabel.region='中国香港'" :class="{isSelected:currentlabel.region=='中国香港'}"><a href="#">中国香港</a></li>
            <li @click="currentlabel.region='中国台湾'" :class="{isSelected:currentlabel.region=='中国台湾'}"><a href="#">中国台湾</a></li>
            <li @click="currentlabel.region='泰国'" :class="{isSelected:currentlabel.region=='泰国'}"><a href="#">泰国</a></li>
            <li @click="currentlabel.region='印度'" :class="{isSelected:currentlabel.region=='印度'}"><a href="#">印度</a></li>
            <li @click="currentlabel.region='法国'" :class="{isSelected:currentlabel.region=='法国'}"><a href="#">法国</a></li>
            <li @click="currentlabel.region='英国'" :class="{isSelected:currentlabel.region=='英国'}"><a href="#">英国</a></li>
            <li @click="currentlabel.region='俄罗斯'" :class="{isSelected:currentlabel.region=='俄罗斯'}"><a href="#">俄罗斯</a></li>
            <li @click="currentlabel.region='意大利'" :class="{isSelected:currentlabel.region=='意大利'}"><a href="#">意大利</a></li>
            <li @click="currentlabel.region='西班牙'" :class="{isSelected:currentlabel.region=='西班牙'}"><a href="#">西班牙</a></li>
            <li @click="currentlabel.region='德国'" :class="{isSelected:currentlabel.region=='德国'}"><a href="#">德国</a></li>
            <li @click="currentlabel.region='波兰'" :class="{isSelected:currentlabel.region=='波兰'}"><a href="#">波兰</a></li>
            <li @click="currentlabel.region='澳大利亚'" :class="{isSelected:currentlabel.region=='澳大利亚'}"><a href="#">澳大利亚</a></li>
            <li @click="currentlabel.region='伊朗'" :class="{isSelected:currentlabel.region=='伊朗'}"><a href="#">伊朗</a></li>
            <li @click="currentlabel.region='其他'" :class="{isSelected:currentlabel.region=='其他'}"><a href="#">其他</a></li>
          </ul>
        </li>
        <li class="time">
          <div class="tagbox">年代：</div>
          <ul class="tags">
            <li @click="currentlabel.time='全部'" :class="{isSelected:currentlabel.time=='全部'}"><a href="#">全部</a></li>
            <li @click="currentlabel.time='2022'" :class="{isSelected:currentlabel.time=='2022'}"><a href="#">2022</a></li>
            <li @click="currentlabel.time='2021'" :class="{isSelected:currentlabel.time=='2021'}"><a href="#">2021</a></li>
            <li @click="currentlabel.time='2020'" :class="{isSelected:currentlabel.time=='2020'}"><a href="#">2020</a></li>
            <li @click="currentlabel.time='2019'" :class="{isSelected:currentlabel.time=='2019'}"><a href="#">2019</a></li>
            <li @click="currentlabel.time='2018'" :class="{isSelected:currentlabel.time=='2018'}"><a href="#">2018</a></li>
            <li @click="currentlabel.time='2017'" :class="{isSelected:currentlabel.time=='2017'}"><a href="#">2017</a></li>
            <li @click="currentlabel.time='2016'" :class="{isSelected:currentlabel.time=='2016'}"><a href="#">2016</a></li>
            <li @click="currentlabel.time='2015'" :class="{isSelected:currentlabel.time=='2015'}"><a href="#">2015</a></li>
            <li @click="currentlabel.time='2014'" :class="{isSelected:currentlabel.time=='2014'}"><a href="#">2014</a></li>
            <li @click="currentlabel.time='2013'" :class="{isSelected:currentlabel.time=='2013'}"><a href="#">2013</a></li>
            <li @click="currentlabel.time='2012'" :class="{isSelected:currentlabel.time=='2012'}"><a href="#">2012</a></li>
            <li @click="currentlabel.time='2011'" :class="{isSelected:currentlabel.time=='2011'}"><a href="#">2011</a></li>
            <li @click="currentlabel.time='2000-2010'" :class="{isSelected:currentlabel.time=='2000-2010'}"><a href="#">2000-2010</a></li>
            <li @click="currentlabel.time='90年代'" :class="{isSelected:currentlabel.time=='90年代'}"><a href="#">90年代</a></li>
            <li @click="currentlabel.time='80年代'" :class="{isSelected:currentlabel.time=='80年代'}"><a href="#">80年代</a></li>
            <li @click="currentlabel.time='70年代'" :class="{isSelected:currentlabel.time=='70年代'}"><a href="#">70年代</a></li>
            <li @click="currentlabel.time='更早'" :class="{isSelected:currentlabel.time=='更早'}"><a href="#">更早</a></li>
          </ul>
        </li>
      </ul>
    </div>
    <div class="movie">
      <div class="sortBy">
        <el-radio-group 
        v-model="changedata.radio">
          <el-radio :label="0">按热门排序</el-radio>
          <el-radio :label="1">按时间排序</el-radio>
          <el-radio :label="2">按评价排序</el-radio>
        </el-radio-group>
      </div>
      <div class="moviebox">
        <ul class="movieList">
          <li class="movieItem" v-for="item in movieList" :key="item.id" @click="toFilm(item.id)">
              <img :src="item.avatar" alt="cover">
              <p>{{item.name}}</p>
              <div class="score">
                <i style="font-size:24px">{{parseInt(item.score)}}.</i>
                <i style="font-size:22px">{{parseInt((item.score-parseInt(item.score))*10)}}</i>
              </div>
          </li>
        </ul>
      </div>
      <div class="page">
          <el-pagination
          background
          layout="prev, pager, next"
          :current-page="changedata.currentPage"
          @current-change="handlePage"
          :page-size="18"
          :total="total">
          </el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import request from '../api/index';
export default {
  name: "Films",
  data() {
    return {
        changedata: {
          radio:0,
          orderby:['visit_count','period', 'score'],
          currentPage:1,
        },        
        currentlabel:{
          type:'全部',
          region:'全部',
          time:'全部',
        },
      movieList:[],
    }
  },
  computed:{
    total(){
      if(this.movieList.length==0)
        return 18;
      else 
        return this.movieList[0].total
    }
  },
  methods: {
    handlePage(val){
      this.changedata.currentPage=val;
    },
    toFilm(movie_id){
      this.$router.push({name: 'thefilm',query:{movie_id,}})
    },
     deleteInvalid (obj){
        Object.keys(obj).forEach(item => {
            if (!obj[item] && obj[item] != 0) {
                delete obj[item]
            }
            })
          return obj;
        },
  },
  watch:{
        changedata:{
          handler:function(newval){ 
          let start_time=this.currentlabel.time=='全部'?'1':(Date.parse(`${this.currentlabel.time}-1-1`)+'').slice(0,10);
          let end_time=this.currentlabel.time=='全部'?'1938434400':(Date.parse(`${parseInt(this.currentlabel.time)+1}-1-1`)+'').slice(0,10);
          let area=this.currentlabel.region=='全部'?undefined:this.currentlabel.region;
          let tag_name=this.currentlabel.type=='全部'?undefined:this.currentlabel.type;
          let obj={
            area ,
            start_time,
            end_time,
            order_by:newval.orderby[newval.radio],
            page:newval.currentPage+'',
            page_size:'18',
            tag_name,
          }
          let config=this.deleteInvalid(obj);
          request.getAllfilms(
            config,      
          ).then((res)=>{
            this.movieList= res.data.data.list;
          },err=>{
            console.error(err)
          })
        },
        deep:true,
    },
     currentlabel:{
          handler:function(newval){ 
            this.changedata.currentPage=1;
            let start_time;
            let end_time;
            if(newval.time=='2000-2010')
            {
              start_time='946656000'
              end_time='1262275200'
            }
            else if(newval.time=='90年代')
            { 
              start_time='631123200'
              end_time='946656000'
            }
            else if(newval.time=='80年代')
            { 
              start_time='315504000'
              end_time='631123200'              
            }
            else if(newval.time=='70年代'||newval.time=='更早')
            { 
              start_time='1'
              end_time='315504000'             
            }
            else
            {
              start_time=newval.time=='全部'?'1':(Date.parse(`${newval.time}-1-1`)+'').slice(0,10);
              end_time=newval.time=='全部'?'1938434400':(Date.parse(`${parseInt(newval.time)+1}-1-1`)+'').slice(0,10);
            }
          let area=newval.region=='全部'?undefined:newval.region;
          let tag_name=newval.type=='全部'?undefined:newval.type;
          let obj={
            area ,
            start_time,
            end_time,
            order_by:this.changedata.orderby[this.changedata.radio],
            page:this.changedata.currentPage+'',
            page_size:'18',
            tag_name,
          }
          let config=this.deleteInvalid(obj);
          request.getAllfilms(
            config,      
          ).then((res)=>{
            this.movieList= res.data.data.list;
          },err=>{
            console.error(err)
          })
        },
        deep:true,
    }
  },
  created(){
    let start_time=this.currentlabel.time=='全部'?'1':`${Date.parse(`${parseInt(this.currentlabel.time)}-1-1`)}`.slice(0,10);
    let end_time=this.currentlabel.time=='全部'?'1938434400':(Date.parse(`${parseInt(this.currentlabel.time)+1}-1-1`)+'').slice(0,10);
    let area=this.currentlabel.region=='全部'?undefined:this.currentlabel.region;
    let tag_name=this.currentlabel.type=='全部'?undefined:this.currentlabel.type;
    let obj={
      area ,
      start_time,
      end_time,
      order_by:this.changedata.orderby[this.changedata.radio],
      page:this.changedata.currentPage,
      page_size:18,
      tag_name,
    }
    let config=this.deleteInvalid(obj);
    request.getAllfilms(
        config   
    ).then((res)=>{
      this.movieList= res.data.data.list;
    },err=>{
      console.error(err)
    })
  },
}
</script>

<style scoped>
.score{
  text-align: center;
    font-size: 16px;
    color: #ffb400;
    margin-top: 10px;
    font-style: italic;
}
.movieList{
  display:flex;
  flex-wrap: wrap;
  padding: 0;
}
.movieList li:nth-child(6n){
  margin-right: 0!important;
}
.movieItem img{
  width:160px;
  height:220.175px;
}
.movieItem p{
  width: 160px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: center;
    font-size: 16px;
    color: #333;
    margin-top: 10px;
}
.movieItem{
    list-style: none;
    width:160px;
    height:292.8px;
    margin: 30px 30px 0 0;
    background-color:#fff;
}
.movie{
  width: 1120px;
  margin: 0 auto;
  margin-top:40px;
}
.sort{
  width:1118px;
  border:1px solid #e5e5e5;
  margin:0px auto;
  box-sizing: content;
}
.sort ul{
  margin:0px;
}
.selectbox{
  list-style:none;
}
.selectbox >li{
  padding:10px 0px;
  display:flex;
  border-bottom:1px solid #e5e5e5;
}
.time{
  border: 0px !important; 
}
.tagbox{
    height: 24px;
    line-height: 24px;
    color: #999;
    font-size: 14px;
    position:absolute;
}
.tags{
  margin-left:40px;
  list-style: none;
  display: flex;
  flex-wrap: wrap;
}
.tags >li{
  margin-left:12px;
  margin-top:5px;
  padding: 3px 9px;
}
.tags a{
  text-decoration: none;
  font-size: 14px;
  color: black;
}
.isSelected{
  background-color: #75c4ff;
  border-radius:14px;
}
.isSelected a{
    color:#fff;
}
a:hover{
    color:#139bb1;
}
</style>