<template>
<div>
    <div class="main">
        <el-tabs type="border-card" tab-position="left">
        <el-tab-pane label="我的订单">
            <MyOrder/>
        </el-tab-pane>
        <el-tab-pane label="基本信息">
            <MyInfo/>
        </el-tab-pane>
        <el-tab-pane label="我的评论">
           <MyComment/>
        </el-tab-pane>
        </el-tabs>
    </div>
</div>
</template>

<script>
import MyInfo from '../components/MyInfo.vue'
import MyComment from '../components/MyComment.vue'
import MyOrder from '../components/MyOrder.vue'
export default {
    name:'homepage',
    components:{
        MyInfo,
        MyComment,
        MyOrder
    },
    beforeRouteEnter(to,from,next){
        let token=localStorage.getItem('token');
        if(token)
        {
            next()
        }
    }
}
</script>

<style scoped>

.main{
    width:1200px;
    margin:0 auto;
    margin-bottom: 300px;
}
</style>