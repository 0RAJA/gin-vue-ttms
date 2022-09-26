<template>


    <section>
        <!-- 背景颜色 -->
        <div class="color"></div>
        <div class="color"></div>
        <div class="color"></div>
            <div class="logo2">
                <img src="../assets/img/logo2.png" alt="">
            </div>
            <div class="logo3">
                <img src="../assets/img/logo3.png" alt="">
            </div>
            <div class="logo4">
                <img src="../assets/img/logo4.png" alt="">
            </div>

        <div class="box">
            <!-- 背景圆 -->
            <div class="circle" style="--x:0"></div>
            <div class="circle" style="--x:1"></div>
            <div class="circle" style="--x:2"></div>
            <div class="circle" style="--x:3"></div>
            <div class="circle" style="--x:4"></div>
            <!-- 登录框 -->
            <div class="container">
                <div class="form"  v-if="isUser">
                <svg t="1653991850738" @click="confirm2" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4545" width="30" height="30"><path d="M512.105912 80.83812c-237.365082 0-429.779394 192.56883-429.779394 430.134481 0 237.564627 192.414311 430.132434 429.779394 430.132434 237.387595 0 429.801906-192.567807 429.801906-430.132434C941.907818 273.40695 749.494531 80.83812 512.105912 80.83812zM605.570607 696.710862c10.212596 10.166547 11.227716 25.591865 2.321882 34.477232-8.92937 8.885368-24.465205 7.844665-34.675754-2.297322l-199.88547-197.651592c-5.79089-5.7694-8.53028-13.19451-8.266267-20.180622-0.309038-7.048533 2.432399-14.541181 8.245801-20.35356l199.905936-197.651592c10.210549-10.166547 25.724895-11.183713 34.675754-2.298346 8.905834 8.884344 7.891737 24.311709-2.321882 34.477232L417.64554 510.972601 605.570607 696.710862z" p-id="4546" fill="#cdcdcd"></path>啊啊啊</svg>

                    <h2>用户登录</h2>
                    <div class="info">{{info}}</div>
                    <form v-on:submit.prevent>
                        <div class="inputBox" >
                            <input type="text" placeholder="用户名" v-model="user.Username">
                        </div>
                        <div class="inputBox">
                            <input type="password" placeholder="密码"  v-model="user.Password">
 
                        </div>
                        <div class="inputBox">
                            <input type="submit" value="登录" @click="blank">
                        </div>
                        <p class="forget"><button @click="change">没有账户?点击注册</button></p>
                    </form>
                </div>
                <div class="form" v-else>
                    <h2>用户注册</h2>
                    <form v-on:submit.prevent>
                        <div class="inputBox" >
                            <input type="text" placeholder="用户名" v-model="user.Username">
                        </div>
                        <div class="inputBox">
                            <input type="password" placeholder="密码" v-model="user.Password">

                        </div>
                        <div class="inputBox">
                            <input type="password" placeholder="确认密码" v-model="user.Password2">
                        </div>
                        <div class="inputBox">
                            <input type="mail" placeholder="邮箱验证" v-model="user.email">
                        </div>
                        <button class="aaa" @click="mail">点击发送验证码</button>
                        <div class="inputBox">
                            <input type="comfirm" placeholder="验证码" v-model="user.comfirm">
                        </div>
                        <div class="inputBox">
                            <input type="invite" placeholder="邀请码" v-model="user.invite">
                        </div>
                        <!-- <p class="forget"><button @click="backk">返回</button></p> -->
                        <p class="forget2"><button @click="blankk">确定</button></p>
                        <p class="forget1"><button @click="change()">返回</button></p>
                    </form>
                </div>
            </div>
        </div>
    </section>

 

</template>

<script>
  export default {
    name: "Register",
    data() {
      return {
          info: '',
          isUser: true,
          user: {
              Password:'',
              Username:'',
              comfirm:'',
              email:'',
              Password2:'',
              invite:''
          }
      }
    },
    methods: {
        change() {
            this.isUser =!this.isUser;
        },
        confirm2() {
        // alert("注册成功!!!")
        this.$router.replace({name: "start"});
        },
        blank() {
            if (this.user.Password==''&& this.user.Username!='') {
                this.$message({
                                message: '密码不能为空!!',
                                center:true,
                                type: 'error'
                            })
            }
            if(this.user.Username==''&& this.user.Password!='') {
                               this.$message({
                                message: '账号不能为空!!',
                                center:true,
                                type: 'error'
                            })
            }
            if(this.user.Password==''&& this.user.Username == '')
            {
                               this.$message({
                                message: '账号密码不能为空!!',
                                center:true,
                                type: 'error'
                            })
            }
            else {
                let {Username,Password} = this.user
                this.$api.getLogin({    
                    Username,Password
                })
                .then(res => {
                    if(res.data.status_code == 0) {
                        //1.存储登录信息 2.跳转页面 3.顶部区域显示用户
                        this.info=''
                        let obj = {
                            user:{
                                Username:res.data.data.username,
                                avatar:res.data.data.avatar,
                                userid:res.data.data.user_id,
                                privilege:res.data.data.privilege
                            },
                            token:res.data.data.refresh_token
                        }
                        // console.log(obj);
                        localStorage.setItem('token',obj.token)
                        localStorage.setItem('user',JSON.stringify(obj.user))
                        // //跳转
                        if (res.data.data.privilege=='用户') {
                            this.$router.push({path:'/'})     
                           this.$message({
                                    message: '登录成功',
                                    center: true,
                                    type: 'success'
                                    });
                        }
                        if(res.data.data.privilege=='管理员') {
                            this.$router.push({ path: "/layout" })
                            this.$message({
                                    message: '登录成功',
                                    center: true,
                                    type: 'success'
                                    });
                        }
                    }
                    else{
                        //账号和密码错误
                        this.info='账号或密码错误'
                    }
                })
                .catch (err=>{
                    console.log(err);
                })
            }
        },
        mail() {
            this.$api.getMail({
                email:this.user.email
            })
            .then(res=>{
                // console.log(res.data);
                if(res.data.status_code==0) {
                    this.$message({
                                    message: '验证码发送成功',
                                    center: true,
                                    type: 'success'
                    });
                }
                if(res.data.status_code==3004) {
                     this.$message({
                                    message: '验证码发送过于频繁!!!',
                                    center: true,
                                    type: 'error'
                    });
                }
            })
            .catch (err => {
                console.log(err);
            })
        },
        blankk() {
            if(this.user.Password==''|| this.user.Username == '' || this.user.email=='' || this.user.comfirm=='')
            {
                this.$message({
                                message: '账号密码邮箱验证码不能为空!!!',
                                center:true,
                                type: 'error'
                            })
            }
            else {
                if (this.user.Password !== this.user.Password2) {
                    wthis.$message({
                                message: '两次输入的密码不一致!!!',
                                center:true,
                                type: 'error'
                            })
                }
                else {
                    this.$api.isRepeat({
                        Username:this.user.Username
                    })
                    .then(res=>{
                        console.log(res.data);
                        if(res.data.status_code==3001)
                        this.$message({
                                message: '用户名已存在!!',
                                center:true,
                                type: 'error'
                            })
                    })
                    this.$api.getRegister({
                        Email: this.user.email,
                        Password:this.user.Password,
                        Username:this.user.Username,
                        VerifyCode:this.user.comfirm,
                        InviteCode:this.user.invite
                    })
                    .then(res=>{
                        console.log(res);
                        if(res.data.status_code==0){
                            this.$message({
                                    message: '注册成功',
                                    center: true,
                                    type: 'success'
                                    });
                            this.change();
                        }
                        if(res.data.status_code==3007){
                            this.$message({
                                message: '验证码错误或过期!!',
                                center:true,
                                type: 'error'
                            })
                        }
                    })
                    .catch (err => {
                    console.log(err);
            })
                }
            }
        }
    },

  }

</script>

<style scoped>
    @import "../assets/css/register.css";
    .aaa {
        margin-top: 15px;
    }
    .info {
        margin-left: 10px;
        color: tomato;
    }
</style>

