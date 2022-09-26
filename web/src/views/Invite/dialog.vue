<template>
   <div>
        <!-- 4弹框 -->
    <el-dialog
  title="添加剧目"
  :visible.sync="dialogVisible"
  width="70%"
  :before-close="handleClose">

  <!-- 内容区域 -->
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
  <el-form-item label="影厅" prop="cinema_id">
    <el-input v-model="ruleForm.name"></el-input>
  </el-form-item>
    <el-form-item label="电影" prop="movie_id">
    <el-input v-model="ruleForm.cols"></el-input>
  </el-form-item>
  <el-form-item label="剧目行" prop="rows">
    <el-input v-model="ruleForm.rows"></el-input>
  </el-form-item>

    <el-form-item label="剧目描述" prop="avatar">
    <el-input v-model="ruleForm.avatar"></el-input>
  </el-form-item>


    <el-button type="primary" @click="submitForm('ruleForm')">确定</el-button>
    <el-button @click="resetForm('ruleForm')">重置</el-button>
</el-form>
</template>

<script>
export default {

}
</script>

<style>

</style>