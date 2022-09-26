<template>
  <div>
        <!-- 4弹框 -->
  <el-dialog
  title="添加影厅"
  :visible.sync="dialogVisible"
  width="70%"
  :before-close="handleClose">

  <!-- 内容区域 -->
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

  <el-form-item label="影厅名称" prop="name">
    <el-input v-model="ruleForm.name"></el-input>
  </el-form-item>
    <el-form-item label="座位列" prop="cols">
    <el-input v-model="ruleForm.cols"></el-input>
  </el-form-item>
  <el-form-item label="座位行" prop="rows">
    <el-input v-model="ruleForm.rows"></el-input>
  </el-form-item>
    <el-form-item label="影厅封面" prop="avatar">
    <el-input v-model="ruleForm.avatar"></el-input>
  </el-form-item>

    <el-button type="primary" @click="submitForm('ruleForm')">确定</el-button>
    <el-button @click="resetForm('ruleForm')">重置</el-button>
</el-form>

  <!-- 1.内弹框 剧目选择 -->
  <!-- <el-dialog
      width="30%"
      title="内层 Dialog"
      :visible.sync="innerVisible"
      append-to-body>
      <Treeadd></Treeadd>
      <span slot="footer" class="dialog-footer">
    <el-button @click="innerVisible = false" >取 消</el-button>
    <el-button type="primary" @click="innerVisible = false">确定</el-button>
  </span>
    </el-dialog> -->
</el-dialog>
  </div>
</template>

<script>
import Treeadd from './Treeadd.vue'
export default {
  props: ['dialogVisible'],
data() {
      return {
        // dialogVisible: false,
        innerVisible:false,
        ruleForm: 
          {
          name: '',
          cols: '',
          rows: '',
          avatar:'',
          delivery: false,
          type: [],
          },
        rules: {
          name: [
            { required: true, message: '请输入影厅名称', trigger: 'blur' },
          ],
          cols: [
            { required: true, message: '请输入影厅列', trigger: 'blur' }
          ],
           rows: [
            { required: true, message: '请选择影厅行', trigger: 'blur' },
          ],  
          type: [
            { type: 'array', required: true, message: '请至少选择一个剧目类型', trigger: 'change' }
          ],
        }
      };
    },
  methods: {
    close() {
      this.$emit('changeDialog')
    },
    submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            console.log(this.ruleForm);
            this.$api.getCinema({
              name:this.ruleForm.name,
              cols:parseInt(this.ruleForm.cols),
              rows:parseInt(this.ruleForm.rows),
              avatar:this.ruleForm.avatar
            })
            .then(res=>{
              console.log(res.data);
              if (res.data.status_code==0) {
                this.dialogVisible = false
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
  },
  components: {
    Treeadd
  }
}
</script>

<style>
  
</style>  