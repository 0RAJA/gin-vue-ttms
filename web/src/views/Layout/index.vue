<template>
  <div class="layout">
    <!-- 左侧导航 -->
    <Mymenu class="menu" :isCollapse="isCollapse"></Mymenu>
    <!-- 右侧内容 -->
    <Content class="content" :class="{isActive:isCollapse}" @changeCollapse='changeCollapse' :isCollapse="isCollapse"></Content>
  </div>
</template>

<script>
import Content from "./Content"
import Mymenu from "./Mymenu.vue"
export default {
  name: "index",
  components: {
  Mymenu,
  Content,
  },
  data () {
    return {
      isCollapse: false,
    }
  },
  methods: {
    changeCollapse() {
      this.isCollapse = !this.isCollapse;
    }
  }
}

</script>

<style scoped>
  .layout .menu {
    /* width: 200px; */
    min-height: 500px;
    background-color: rgb(84, 92, 100);
    position: fixed;
    top: 0;
    bottom: 0;
  }
  .layout .content {
    margin-left: 200px;
  }
  .layout .isActive{
    margin-left: 64px;
  }
</style>