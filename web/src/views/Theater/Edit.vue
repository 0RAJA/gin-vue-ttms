<template>
    <div>
        <!-- 4弹框 -->
  <el-dialog
  title="编辑影厅"
  :visible.sync="dialogTableVisible"
  width="70%"
  :before-close="handleClose">

  <!-- 内容区域 -->
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

  <el-form-item label="影厅id" prop="id">
    <el-input v-model="ruleForm.id"></el-input>
  </el-form-item>
  <el-form-item label="影厅名称" prop="name">
    <el-input v-model="ruleForm.name"></el-input>
  </el-form-item>

    <el-form-item label="影厅封面" prop="avatar">
    <el-input v-model="ruleForm.avatar"></el-input>
  </el-form-item>

    <el-button type="primary" @click="submitForm('ruleForm')">确定</el-button>
    <el-button @click="resetForm('ruleForm')">重置</el-button>
</el-form>
</el-dialog>
  </div>
</template>

<script>
export default {
  props: ['dialogTableVisible'],
  data () {
    return {
      ruleForm:{
        name:'',
        avatar:'',
        id:''
      }
    }
  },
  methods: {
    close(){
      this.$emit('changeEdit')
    },
    submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            console.log(this.ruleForm);
            this.$api.updateCinema({
              cinema_id: parseInt(this.ruleForm.id),
              new_avatar: this.ruleForm.avatar,
              new_name: this.ruleForm.name
            })
            .then(res=>{
              console.log(res.data);
              if(res.data.status_code==0){
                this.dialogTableVisible = false
                this.http(1)
              }
            })
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
    resetForm(formName) {
        this.$refs[formName].resetFields();
      }
  }
}
</script>

<style>

</style>