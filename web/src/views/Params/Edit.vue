<template>
  <div>
    <el-dialog
  title="添加剧目"
  :visible.sync="dialogTableVisible"
  width="70%"
  :before-close="handleClose">

  <!-- 内容区域 -->
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

        <el-form-item label="剧目id" prop="id">
    <el-input v-model="ruleForm.id"></el-input>
  </el-form-item>
  <el-form-item label="演员列表" prop="actors">
    <el-input v-model="ruleForm.actors"></el-input>
  </el-form-item>
  <el-form-item label="电影名称" prop="name">
    <el-input v-model="ruleForm.name"></el-input>
  </el-form-item>
  <el-form-item label="别名" prop="alias_name">
    <el-input v-model="ruleForm.alias_name"></el-input>
  </el-form-item>
    <el-form-item label="电影简介" prop="content">
    <el-input v-model="ruleForm.content"></el-input>
  </el-form-item>

  <el-form-item label="电影地点" prop="area">
    <el-input v-model="ruleForm.area"></el-input>
  </el-form-item>
    <el-form-item label="电影导演" prop="director">
    <el-input v-model="ruleForm.director"></el-input>
  </el-form-item>
      <el-form-item label="剧目封面" prop="avatar">
    <el-input v-model="ruleForm.avatar"></el-input>
  </el-form-item>

  <el-form-item label="电影上映时间" prop="period">
    <el-input v-model="ruleForm.period"></el-input>
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
          movie_id:'',
      ruleForm: 
          {
          name: '',
          actors: [ 
          ],
          alias_name: '',
          avatar:'',
          content:'',
          area: '',
          period:'',
          director:'',
          },
        rules: {
          name: [
            { required: false, message: '请输入电影名称', trigger: 'blur' },
          ],
          actors: [
            { required: false, message: '请输入演员', trigger: 'blur' }
          ],
          director: [
            { required: false, message: '请输入导演', trigger: 'blur' }
          ],
           alias_name: [
            { required: false, message: '请输入电影别名', trigger: 'blur' },
          ],
          avatar: [
            { required: false, message: '图像', trigger: 'blur' }
          ],
          area: [
            { required: false, message: '请输地点', trigger: 'blur' }
          ],
          content: [
            { required: false, message: '请输入电影内容', trigger: 'blur' },
          ],
          period: [
            { required: false, message: '请输入电影上映时间', trigger: 'blur' },
          ],
          id: [
            { required: false, message: '电影id', trigger: 'blur'}
          ]
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
            this.$api.updataMovie({   
                actors: this.ruleForm.actors.split(','),
                alias_name: this.ruleForm.alias_name,
                area: this.ruleForm.area,
                avatar: this.ruleForm.avatar,
                content: this.ruleForm.content,
                director: this.ruleForm.director,
                id: parseInt(this.ruleForm.id),
                name: this.ruleForm.name,
                period: new Date(this.ruleForm.period).getTime()/1000,
            })
            .then(res=>{
              // console.log(res.data); 
              if(res.data.status_code==0) {
                this.dialogTableVisible = false
                window.alert('编辑成功!!!');
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