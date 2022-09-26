<template>
  <div>
      <div class="back">
          <div class="main">
              <div class="cover">
                  <img :src="moviedetail.avatar" alt="封面">
              </div>
              <div class="content">
                  <div class="info">
                      <h1 class="name">{{moviedetail.name}}</h1>
                      <div class="ename">{{moviedetail.alias_name}}</div>
                      <ul class="label">
                          <li>{{moviedetail.Tags.join(' ')}}</li>
                          <li>
                            评分：
                            <el-rate
                                style="display:inline"
                                :value="(moviedetail.score/2)"
                                disabled
                                text-color="#ff9900">
                            </el-rate>
                            <span style="color:#ffc600">{{moviedetail.score.toFixed(1)}}</span>
                          </li>
                          <li>{{moviedetail.area}} / {{moviedetail.duration}}分钟</li>
                          <li>{{moviedetail.period.slice(0,10)}} {{moviedetail.period.slice(11,16)}}中国大陆上映</li>
                      </ul>
                  </div>
                  <div style="margin-top:30px">
                      <el-button type="warning" icon="el-icon-video-play" @click="givelike()">{{moviedetail.IsFollow?'取消':'想看'}}</el-button>
                    <el-button type="warning"
                    @click="addComment"
                     icon="el-icon-star-off" >评分</el-button>
                    <br>
                    <el-button type="warning" 
                    @click="toChoose"
                    style="width: 188px; margin-top:10px" >特惠购票</el-button>
                  </div>
                  <div class="evaluate">
                      <div class="movieindex">
                          <div class="title">想看数</div>
                            <div class="number">{{moviedetail.follow_count}}</div>
                            <div class="title">累计票房</div>
                            <div class="number">{{moviedetail.box_office}}</div>
                      </div>
                  </div>
              </div>
          </div>
            <div class="plot">
            <h1 class="labelname">剧情简介</h1>
            <hr/>
            <p class="synopsis"><span style="color: red;font-weight: bold;">主要演员：</span> {{moviedetail.actors.join(' ')}}<br/></p>
            <p class="synopsis">{{moviedetail.content}}</p>
            </div>
      </div>
    <router-view></router-view>
  </div>
</template>

<script>
import request from '../api/index'
import comment from '../views/comment.vue'
import chooseticket from '../views/chooseTicket.vue'
export default {
    data() {
        return{
            moviedetail:{
            },
        }
    },
    components:{
        comment,
        chooseticket,
    },
    methods:{
        givelike(){
            let token=localStorage.getItem('token');
            if(this.$route.name=='chooseticket')
            { 
                return;
            }
            if(!token)
            {
                this.$message({
                                message: '请登录',
                                type: 'warning'
                            });
                return;
            }
            this.$http.post('https://ttms.humraja.xyz/ttms/user_movie/opt',{
                movie_id:this.moviedetail.id,
                opt:!this.moviedetail.IsFollow,
            }).then(res=>{
                if(res.data.status_code==0)
                {
                    this.moviedetail.IsFollow?this.moviedetail.follow_count-=1:this.moviedetail.follow_count+=1;
                    this.moviedetail.IsFollow=!this.moviedetail.IsFollow
                }
            },
            err=>{console.log(err)}
            )            
        },
        toChoose(){
            let obj={
                name: this.moviedetail.name,
                type: this.moviedetail.Tags.slice(0,2).join(' '),
                duration: this.moviedetail.duration,
                cover: this.moviedetail.avatar
            }
            this.$router.push({name:'chooseticket',query:{movie_id:this.$route.query.movie_id,movieInfo:obj}});
        },
        addComment(){
            if(this.$route.name=='chooseticket')
            { 
                return;
            }
            let token=localStorage.getItem('token');
            if(!token)
            {
                this.$message({
                                message: '请登录',
                                type: 'warning'
                            });
                return;
            }
             if(this.moviedetail.IsComment==true)
            {
                 this.$message({
                                message: '已经评价过了',
                                type: 'warning'
                            });
                
            }
            else
                {
                    this.$bus.$emit('addComment');
                    this.moviedetail.IsComment=true
                }
        }
    },
    created(){
        request.getFilmDetail({
            movie_id: this.$route.query.movie_id,
        }
        ).then(res=>{
            this.moviedetail=res.data.data;
        },
        err=>{
            console.log(err);
        }
        )
    }
}
</script>

<style scoped>
.synopsis{
    margin-left:20px;
    text-indent: 2em;
}
.plot{
    width:1000px;
    margin:50px auto;
    padding:10px;
}
.labelname::before{
    content: "";
    display: inline-block;
    width: 8px;
    height: 28px;
    margin-right: 6px;
    background-color: #11a8cd;
}
.title{
    font-size: 12px;
    margin-bottom: 8px;
}
.number{
    font-size: 30px;
    color: #ffc600;
    height: 30px;
    line-height: 30px;
    margin-bottom:16px;
}
.evaluate{
    position: absolute;
    width: 116.7px;
    height: 136.8px;
    top:158px;
    left:342px;
    z-index: 100;
}
.label {
    width: 250px;
    list-style: none;
    padding-left: 0;
    margin-bottom: 20px;
}
.label li{
    margin: 12px 0;
    line-height: 100%;
}
.name{  
    white-space:nowrap;
    width: 400px;
    margin-top: 0;
    font-size: 26px;
    line-height: 32px;
    font-weight: 700;
    margin-bottom: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    max-height: 64px;
}
.ename:hover{
    width: 700px;
}
.name:hover{
    width: 700px;
}
.ename{
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    width: 340px;
    font-size: 18px;
    line-height: 1.3;
    margin-bottom: 14px;
}
.content {
    color:#fff;
    position: relative;
    margin:70px 30px 0 300px;
    width: 870px;
    height:300px;
}
.back{
    height: 376px;
    width: 100%;
    display: inline-block;
    background: #392f59 url(../assets/img/theFilmBack.png) no-repeat 50%;
}
.cover{
    float: left;
    overflow: hidden;
    z-index: 9;
    width: 240px;
    height: 330px;
    margin: 0 30px;
}
.cover img{
    border: 4px solid #fff;
    height: 322px;
    width: 232px;
}
.main{
    width:1200px;
    margin: 0 auto;
}
</style>