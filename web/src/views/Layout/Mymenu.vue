<template>
<div>
<el-row class="tac">
  <el-col :span="12">
    <el-menu
      :default-active="$route.path"
      class="el-menu-vertical-demo"
      router
      :collapse='isCollapse'
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b">
      <el-menu-item>
        <span slot="title">海绵摆摆后台管理系统</span>
      </el-menu-item>
      <el-menu-item index="/backhome">
          <i class="el-icon-house"></i>
        <span slot="title">
          <span>首页</span>
        </span>
      </el-menu-item>
      <el-menu-item index="/theater">
        <i class="el-icon-document"></i>
        <span slot="title">影厅管理</span>
      </el-menu-item>
      <el-menu-item index="/params">
        <i class="el-icon-document"></i>
        <span slot="title">剧目管理</span>
      </el-menu-item>

      <el-submenu index="/order">
        <template slot="title">
          <i class="el-icon-s-order"></i>
          <span>演出计划</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="/order/order-list">演出计划列表</el-menu-item>
          <el-menu-item index="/order/order-back">新建演出计划</el-menu-item>
        </el-menu-item-group>
      </el-submenu>            
      <el-menu-item index="/invite">
        <i class="el-icon-document"></i>
        <span slot="title">订单管理</span>
      </el-menu-item>
      <el-menu-item index="/staff">
        <i class="el-icon-s-custom"></i>
        <span slot="title">用户管理</span>
      </el-menu-item>
      <el-menu-item index="/">
        <i class="el-icon-back"></i>
        <span>返回</span>
      </el-menu-item>
    </el-menu>
  </el-col>
</el-row>
</div>
</template>

<script>
export default {
  props: ['isCollapse'],
  name: "Mymenu",
  data () {
    return {
      // isCollapse: false
    }
}
}
</script>

<style scoped>
  .Htitle {
    text-align: center;
    color: #fff;
    height: 50px;
    line-height: 50px;
  }
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }
</style>