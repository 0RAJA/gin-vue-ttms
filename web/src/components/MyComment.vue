<template>
  <div>
       <ul class="commentList">
                <li class="commentItem" @click="moviedetail(item.movie_id)" @contextmenu.prevent="deletecomment(item.commentid)" v-for="item in commentList" :key="item.commentid">
                    <img :src="item.movieavatar" alt="film">
                    <div class="commentinfo">
                        <h4>{{item.moviename}}</h4>
                        <p class="content">
                            {{item.content}}
                        </p>
                        <div class="likes" style="float: left">
                            {{item.created_at.slice(0,10)}}
                        </div>
                        <div class="likes" style="margin-left:640px">
                            <i class="el-icon-thumb"></i>
                            <span>{{item.star_num}}</span>
                        </div>
                    </div>
                </li>
        </ul>
  </div>
</template>

<script>
export default {
    data(){
        return{
            commentList:[]
        }
    },
    methods: {
        moviedetail(id){
             this.$router.push({name: 'thefilm',query:{movie_id:id}})
        },
        deletecomment(id){
            if(confirm("Are you sure you want to delete"))
            { 
                this.$api.deleteComment({comment_id:id}).then(res => {
                if(res.data.status_code==0)
                    this.commentList=this.commentList.filter(item => item.commentid !== id)
                else {
                    this.$message({
                            message: '删除失败，稍后重试',
                            type: 'error'
                            }); 
                }
            })
            }
        }
    },
    created(){
        this.$http.get('https://ttms.humraja.xyz/ttms/comment/list/user_id'
        ).then(res => {
            this.commentList=res.data.data.list
        })
    }
}
</script>

<style scoped>
.likes{
    font-size: 12px;
    color:#a29d9d;
}
.content{
    overflow: hidden;
    text-overflow: ellipsis;
    font-size:12px;
    color:rgb(154, 148, 148);
    text-indent: 24px;
    width:100%;
    height:29px;
}
.commentinfo{
    padding:10px;
    padding-bottom: 0px;
}
.commentinfo h4{
    font-size:16px;
    font-weight: normal;
    margin: 0;
}
.commentList{
    list-style: none;
    width:800px;
    margin:0px auto;
    padding: 0;
    border:1 solid #e5e5e5;
    box-shadow: 3px 3px 3px 0 #e5e5e5;
}
.commentItem{
    display:flex;
    height: 93.8px;
    border-bottom:1px solid #e5e5e5;

}
.commentItem img{
    width:67.575px;
    height:93.8px;
}
</style>