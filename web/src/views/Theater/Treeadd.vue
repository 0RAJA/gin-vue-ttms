<template>
  <div> 
    <el-tree
  :data="data"
  :props="defaultProps"
  accordion
  @node-click="handleNodeClick">
</el-tree>
  </div>
</template>

<script>
export default {
      data() {
      return {
        data: [{
          label: '国产',
           children: [{
            label: '喜剧',
          }, {
            label: '悲剧',
          },
          {
            label: '恐怖',
          },
          {
            label: '文艺',
          },
          {
            label: '爱情',
          },
          {
            label: '武侠',
          }]
        }, {
          label: '欧美',
          children: [{
            label: '喜剧',
          }, {
            label: '悲剧',
          },
          {
            label: '恐怖',
          },
          {
            label: '文艺',
          },
          {
            label: '爱情',
          },
          {
            label: '武侠',
          }]
        }, {
          label: '日韩',
          children: [{
            label: '喜剧',
          }, {
            label: '悲剧',
          },
          {
            label: '恐怖',
          },
          {
            label: '文艺',
          },
          {
            label: '爱情',
          },
          {
            label: '武侠',
          }]
        }],
        defaultProps: {
          children: 'children',
          label: 'label'
        }
      };
    },

}
</script>

<style>

</style>  