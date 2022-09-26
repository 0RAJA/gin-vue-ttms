<template>
  <div>
      <p style="color:#409EFF">基本信息</p>
            <hr>
            <div class="photo">
                <img :src="adatar?adatar:userInfo.avatar" alt="头像">
                <br/><br/>
                <el-button @click="changeAim()" type="primary">更换</el-button>
                <el-button @click="submit()" type="primary">确认</el-button>
                <input type="file" ref="upload" enctype="multipart/form-data" accept="image/gif,image/jpeg,image/jpg,image/png" @change="fileChange" style="display:none;">
                <p>支持JPG,JPEG,PNG格式，且文件要小于1M</p>
            </div>
            <div class="info">
                昵称：<input type="text" v-model="userInfo.username"><br><br>
                性别：<input type="radio" name="gender" value="男" v-model="userInfo.gender"><label for="male">男</label> 
                <input type="radio" name="gender" value="女" v-model="userInfo.gender"><label for="female">女</label><br><br>
                生日：<input type="date" v-model="birthday"><br><br>
                生活状态：<input type="radio" value="单身" v-model="userInfo.lifestate" name="status" > 单身
                <input type="radio" value="热恋中" v-model="userInfo.lifestate" name="status"> 热恋中
                <input type="radio" value="已婚" v-model="userInfo.lifestate" name="status"> 已婚
                <input type="radio" value="为人父母" v-model="userInfo.lifestate" name="status"> 为人父母<br><br>
                兴趣：
                <div class="hobby">
                    <input type="checkbox" value="美食" v-model="userInfo.hobby"> 美食
                    <input type="checkbox" value="动漫" v-model="userInfo.hobby"> 动漫
                    <input type="checkbox" value="摄影" v-model="userInfo.hobby"> 摄影
                    <input type="checkbox" value="电影" v-model="userInfo.hobby"> 电影
                    <input type="checkbox" value="体育" v-model="userInfo.hobby"> 体育<br>
                    <input type="checkbox" value="财经" v-model="userInfo.hobby"> 财经
                    <input type="checkbox" value="音乐" v-model="userInfo.hobby"> 音乐
                    <input type="checkbox" value="游戏" v-model="userInfo.hobby"> 游戏
                    <input type="checkbox" value="科技" v-model="userInfo.hobby"> 科技
                    <input type="checkbox" value="旅游" v-model="userInfo.hobby"> 旅游<br>
                    <input type="checkbox" value="文学" v-model="userInfo.hobby"> 文学
                    <input type="checkbox" value="公益" v-model="userInfo.hobby"> 公益
                    <input type="checkbox" value="汽车" v-model="userInfo.hobby"> 汽车
                    <input type="checkbox" value="时尚" v-model="userInfo.hobby"> 时尚
                    <input type="checkbox" value="宠物" v-model="userInfo.hobby"> 宠物<br><br>
                    </div>
                个性签名：<input type="text" placeholder="20字以内" v-model="userInfo.signature" maxlength="50"><br><br>
                <el-button type="primary"
                @click="update()"
                >确定</el-button>
            </div>
  </div>
</template>

<script>
export default {
    data(){
        return {
            userInfo: {
                birthday:'',
            },
            adatar:''
        }
    },
    methods: {
        update(){
            let config={};
            config.Birthday=Date.parse(this.birthday)/1000;
            config.Email=this.userInfo.email;
            config.Gender=this.userInfo.gender;
            config.Hobbys=this.userInfo.hobby;
            config.LifeState=this.userInfo.lifestate;
            config.Signature=this.userInfo.signature;
            config.UserId=this.userInfo.id;
            config.Username=this.userInfo.username;
            this.$api.updateInfo(  
                config,
            ).then(res=>{
                if(res.data.status_code==0)
                 {
                    this.$message({
                        message: '更新成功',
                        type: 'success'
                        });
                    }
            })
        },
        changeAim(){
            const inputFile=document.querySelector("input[type='file']");
            inputFile.click();
        },
        submit(){
            var formdata = new FormData();
            formdata.append('file',this.$refs.upload.files[0], "file");
            this.$api.uploadFile(formdata).then(res=>{
                let user = JSON.parse(localStorage.getItem('user'))
                user.avatar = res.data.data;
                localStorage.setItem('user',JSON.stringify(user));
                this.updateAdatar(res.data.data)
            })
        },
        updateAdatar(address){
            this.$api.updateAdatar(
                {
                    NewAvatar:address,
                    UserId:this.userInfo.id,
                }
            ).then(res=>{
                if(res.data.status_code==0)
                {
                    this.$message({
                        message: '修改成功',
                        type: 'success'
                        });        
                }
            })
        },
        fileChange(e) {
            var that = this;
            var file = e.target.files[0];
            var reader = new FileReader();
            reader.onload = function(e){
               that.adatar = e.target.result;
            }
            reader.readAsDataURL(file);
        },
    },
    computed:{
        birthday:{
            get(){
                return this.userInfo.birthday.slice(0,10)
            },
            set(value){
                this.userInfo.birthday=value
            }
        }
    },
    created () {
        let UserId=JSON.parse(localStorage.getItem("user")).userid;
        this.$api.getuserInfo({
            UserId,
        }).then(res=>{
            this.userInfo=res.data.data.user;
        })
    }
}
</script>

<style scoped>
.hobby{
    margin-left: 48px;
}
input{
    outline: none;
}
textarea{
    outline: none;
    resize:none;
}
.info{
    margin-left: 300px;
}
.photo{
    text-align: center;
    font-size: 12px;
    color:#ccc;
    float: left;
}
.photo img{
    width:160px;
    height:160px;
    border-radius:3px;
    border: 1px solid black;
}
</style>