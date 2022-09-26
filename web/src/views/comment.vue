<template>
  <div>   
      <div class="plot">
            <div class="topBox">
                <h1 class="labelname">热门评论</h1>
            </div>
             <hr/> 
            <div class="comment">
                <ul class="commentlist">
                    <li class="commentitem" v-for="item in comment" :key="item.commentid">
                        <div class="photo">
                            <img :src="item.avatar" alt="头像">
                            <i class="level"></i>
                        </div>
                        <div class="main">
                            <div class="main-header">
                                <div class="user">
                                    <span class="name">{{item.username}}</span>
                                    <span class="tag"></span>
                                    <span class="ip">来自:{{item.ip_address}}</span>
                                </div>
                                <div class="time">
                                    <span>{{item.created_at.slice(0,10)}}</span>
                                    <el-rate
                                        style="display:inline"
                                        :value="parseInt((item.score/2).toFixed(1))"
                                        disabled
                                        show-score
                                        text-color="#ff9900"
                                        score-template="{value}">
                                    </el-rate>
                                    <div class="operation">
                                    <i class="el-icon-thumb"
                                    v-show="!item.is_star"
                                    @click="givelike(item)"
                                    ></i>
                                    <img src="../assets/img/已点赞.png" 
                                    v-show="item.is_star"
                                    @click="givelike(item)"
                                    style="width:16px;height:16px">
                                    <span>{{item.star_num}}</span>
                                </div>
                                </div>
                                <div class="comment-content">{{item.content}}</div>
                            </div>
                        </div>
                    </li>
                </ul>  
            </div>
        <div class="addcomment" v-show="isComment">
            <p class="title">请点击星星进行评分</p>
            <div class="star">
                <el-rate show-text v-model="value1"></el-rate>
            </div>
            <div class="opinion">
                <textarea cols="60" rows="10" placeholder="快来说说你的看法吧" v-model="commentValue">
                </textarea>
            </div>
            <el-button 
            @click="addComment()"
            style="float: right"
            size="small"
            type="primary" plain>提交</el-button>
            <el-button 
            @click="isComment=false"
            style="float: left;margin-left:0px"
            size="small"
            type="primary" plain>退出</el-button>
        </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            value1:0,
            isComment:false,
            comment:[],
            commentValue:'',
        }
    },
    watch:{
        isComment(){
            this.value1=0;
        }
    },
    computed:{
    },
    methods: {
        givelike(item){
            let token=localStorage.getItem('token');
            if(!token)
            {
                this.$message({
                                message: '请登录',
                                type: 'warning'
                            });
                return;
            }
            let opt=item.is_star==true ?false :true
            this.$http.post('https://ttms.humraja.xyz/ttms/comment/star',{
                comment_id:item.commentid,
                opt,
            }).then(res=>{
                if(res.data.status_code==0)
                {
                    item.is_star==true?item.star_num-=1:item.star_num+=1;
                    item.is_star = !item.is_star;
                }
            },
            err=>{console.log(err)}
            )            
        },
        addComment(){
            this.$http(
                { 
                    method:'post',
                    url:'https://ttms.humraja.xyz/ttms/comment/create',
                    data: JSON.stringify({
                            content: this.commentValue,
                            movie_id: parseInt(this.$route.query.movie_id),
                            score: parseInt(this.value1)*2,
                        })
                }
            ).then((res)=>{
                this.isComment=false;
                if(res.data.status_code==1003){
                    this.$message({
                                message: '系统繁忙，稍后重试',
                                type: 'error'
                            });
                    return;
                }
                let user = JSON.parse(localStorage.getItem("user"));
                res.data.data.avatar=user.avatar;
                res.data.data.star_num=0
                res.data.data.commentid=res.data.data.id
                res.data.data.username=user.Username
                this.comment.unshift(
                    res.data.data
                )
            },err=>{
                console.log(err)
            })
        },

    },
    created(){
        this.$bus.$on('addComment',()=>{
            this.isComment=true;
        })
        this.$api.getComment({
            movie_id: this.$route.query.movie_id,
            page:1,
            page_size:15,
        }).then((res)=>{
            this.comment=res.data.data.list
        },err=>{
            console.log(err)
        })
    }
}
</script>

<style scoped>
.ip{
    margin-left: 20px;
    font-size:10px;
    color:rgb(175, 169, 169);
}
.addcomment /deep/ .el-rate__icon{
    font-size: 30px;
}
.star{
    text-align: center;
    margin-bottom:15px;

}
.title {
    color: #ffc600;
    text-align: center;
    font-size: 12px;
}
.addcomment{
    position:fixed;
    left:50%;
    top:50%;
    transform:translate(-50%,-50%);
    z-index: 100000;
    padding:30px;
    border:1px solid #11a8cd;
    background-color: #fff;
    border-radius: 5px;
}
.opinion textarea{
    resize: none;
    padding:5px;
    border-radius: 5px;
    outline-color:#11a8cd;
}
.topBox{
    display:flex;
    justify-content: space-between;
    align-items: center;
}
.labelname::before{
    content: "";
    display: inline-block;
    width: 8px;
    height: 28px;
    margin-right: 6px;
    background-color: #11a8cd;
}
.synopsis{
    margin-left:20px;
    text-indent: 2em;
}
.plot{
    width:1000px;
    margin:50px auto;
    padding:10px;
}
.operation{
    float:right;
    cursor: pointer;
}
.comment-content{
    margin-top:20px;
    padding-bottom: 30px;
    text-indent: 2em;
    border-bottom: 1px solid#e5e5e5;
}
.time{
    margin-top:4px;
    color: #999;
}
.time span{
    margin-right: 10px;
}
.el-icon-star-on{
    color:#ffc600;
}
.user{
    margin-top:2px;
}
.comment {
    display: flex;
    /* position: relative; */
}
.main{
    margin-left:70px;
}
.level{
    display: inline-block;
    width: 20px;
    height: 20px;
    position: absolute;
    bottom: -2px;
    right: -2px;
    background: url(../assets/img/icon.png);
}
.commentlist {
    list-style: none;
    width:100%;
}
.photo{
    width: 50px;
    height: 50px;
    position: relative;
    margin-right: 20px;
    float:left;
}
.photo img{
    width:100%;
    height: 100%;
    border-radius:50%;

}
</style>