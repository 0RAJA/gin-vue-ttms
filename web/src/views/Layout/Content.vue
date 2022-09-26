<template>
  <div>
    <div class="header">
      <i v-if="!isCollapse" @click="changeMenu"><svg t="1654351883978" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6753" width="20" height="20"><path d="M853.333333 714.666667v64H170.666667v-64h682.666666zM320 405.333333v213.333334l-149.333333-106.666667 149.333333-106.666667z m533.333333 85.333334v64H405.333333v-64h448z m0-234.666667v64H170.666667v-64h682.666666z" p-id="6754"></path></svg></i>
      <i v-else @click="changeMenu"><svg t="1654348866778" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6965" width="20" height="20"><path d="M853.333333 725.034667v64H170.666667v-64h682.666666z m-149.333333-309.333334l149.333333 106.666667-149.333333 106.666667v-213.333334z m-85.333333 74.666667v64H170.666667v-64h448zM853.333333 256v64H170.666667v-64h682.666666z" p-id="6966"></path></svg></i>
      顶部区域</div>
    <!-- 路由出口 -->
    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: "Content",
  props: ['isCollapse'],
  methods: {
    changeMenu() {
      //点击切换按钮,修改父组件的数据  isCollapse
      this.$emit("changeCollapse");
    }
  }
}
</script>

<style scoped>
  .header   {
    height: 50px;
    line-height: 50px;
    background: rgb(84, 92, 100);
    color: #fff;
  }
  .header i {
    cursor: pointer;
  }
  .content {
    background-color: #eee;
  }
</style>