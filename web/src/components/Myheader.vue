<template>
    <div class="top">
        <div class="logo">
            <img src="../assets/img/logo.png" alt="logo">
            <span>海绵摆摆</span>
        </div>
        <div class="nav">
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" active-text-color="#ffd04b" text-color="#fff" style="background:url(../assets/img/background.jpg)">
        <el-menu-item index="1" @click="changeroute(1)">首页</el-menu-item>
        <el-menu-item index="2" @click="changeroute(2)">电影</el-menu-item>
        <el-menu-item index="3" @click="changeroute(3)">榜单</el-menu-item>
        </el-menu>
        </div>
        <div class="search">
            <el-input placeholder="请输入内容" v-model="keyvalue" class="input-with-select">
            <el-button 
            @click="goSearch"
            slot="append" icon="el-icon-search"></el-button>
        </el-input>
        </div>
        <div class="user" >
            <img @click="toHomePage" :src="photo" alt="user">
            <span @click="tologin">{{isLogin?'登录':'注销'}}</span>
        </div>
    </div>
</template>

<script>
import home from "../views/Home.vue";
import film from "../views/Films.vue";
import list from "../views/List";
import search from "../views/Search";
import homepage from "../views/HomePage.vue";
export default {
    data() {
      return {
        activeIndex: '1',
        keyvalue:'',
        isLogin: true,
      };
    },
    components:{
        home,
        film,
        list,
        search,
        homepage
    },
    computed:{
        photo()
        {
            let user= JSON.parse(localStorage.getItem('user'));
            if(user) 
            {
                this.isLogin=false
                return user.avatar
            }
            return require('../assets/img/user.png')
        }
    },
    methods: {   
        changeroute(T)
        {
            switch (T) {
            case 1:this.$router.replace({
                name:'home'
            }); break;
            case 2:this.$router.replace({
                name:'films'
            }); break;
            case 3:this.$router.replace({
                name:'list'
            }); break;
            }
        },
        tologin(){
            if(this.isLogin==false )
                localStorage.clear();
            this.$router.push({name:'login'});
        },
        goSearch(){
            if(this.keyvalue!='')
            {
                this.$router.push({name:'search',query:{keyvalue:this.keyvalue}});
                this.keyvalue='';
            }
            else{
                this.$message({
                                message: '输入一点东西',
                                type: 'warning'
                            });
            }
        },
        toHomePage(){
            this.$router.push({name:'homepage'});
        }
    }
}
</script>

<style scoped>
.el-menu-item:hover{
  color: #e6a23c !important;
}
.top{
    display:flex;
    justify-content:center;
    box-shadow: 3px 3px 1px #888888;
    background: url(../assets/img/background.jpg);
    background-size: 100%;
    font-size: 1.5em;
    font-weight: bold;
    align-items:center;
    color: white;
    min-width: 988px;
    margin-bottom: 3px;
}
.user{
        cursor: pointer;
        margin-left:50px;
        position: relative;
        font-size: 16px;
        display:flex;
        align-items: center;
}
.user img{
    background-color:#888888;
    border-radius:50%;
    width:30px;
    height:30px;
    margin-right: 8px;
}
.logo{
    height:30px;
    margin-right: 50px;
    display:flex;
}
.logo img{
    height:120%;
}
.logo span{
    align-self: center;
    margin-left: 8px;
    }
.nav{
    margin: 0 100px;
}

</style>