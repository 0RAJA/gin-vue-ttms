<template>
  <div>
    <el-dialog
  title="添加剧目"
  :visible.sync="dialogVisible"
  width="70%"
  :before-close="handleClose">

  <!-- 内容区域 -->
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm"   class="demo-ruleForm">

  <el-form-item label="演员列表" prop="actors">
    <el-input v-model="ruleForm.actors"></el-input>
  </el-form-item>
    <el-form-item label="电影名字" prop="name">
    <el-input v-model="ruleForm.name"></el-input>
  </el-form-item>
  <el-form-item label="别名" prop="alias_name">
    <el-input v-model="ruleForm.alias_name"></el-input>
  </el-form-item>
    <el-form-item label="电影简介" prop="content">
    <el-input v-model="ruleForm.content"></el-input>
  </el-form-item>
<el-form-item label="电影时长" prop="duration">
    <el-input v-model="ruleForm.duration"></el-input>
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
  </el-form-item><el-form-item label="电影标签" prop="tags">
    <el-input v-model="ruleForm.tags"></el-input>
  </el-form-item>
    <el-button type="primary" @click="submitForm('ruleForm')">确定</el-button>
    <el-button @click="resetForm('ruleForm')">重置</el-button>
</el-form>
</el-dialog>
  </div>

</template>

<script>
export default {
  props: ['dialogVisible'],
  data() {
      return {
        // dialogVisible: false,
        innerVisible:false,
        ruleForm: 
          {
          name: '',
          actors: [ 
          ],
          alias_name: '',
          avatar:'',
          content:'',
          area: '',
          duration:'',
          period:'',
          director:'',
          tags:[
          ],
          },
        rules: {
          name: [
            { required: true, message: '请输入电影名称', trigger: 'blur' },
          ],
          actors: [
            { required: true, message: '请输入演员', trigger: 'blur' }
          ],
          director: [
            { required: true, message: '请输入导演', trigger: 'blur' }
          ],
           alias_name: [
            { required: true, message: '请输入电影别名', trigger: 'blur' },
          ],
          avatar: [
            { required: false, message: '图像', trigger: 'blur' }
          ],
          area: [
            { required: true, message: '请输地点', trigger: 'blur' }
          ],
          content: [
            { required: true, message: '请输入电影内容', trigger: 'blur' },
          ],
          duration: [
            { required: true, message: '请输入电影时长', trigger: 'blur' },
          ],
          period: [
            { required: true, message: '请输入电影上映时间', trigger: 'blur' },
          ],
          tags: [
            { required: true, message: '请输入电影标签', trigger: 'blur' },
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
            this.$api.getCreatemovie({
              name:this.ruleForm.name,
              actors: this.ruleForm.actors.split(','),
              alias_name: this.ruleForm.alias_name,
              avatar: this.ruleForm.avatar,
              content: this.ruleForm.content,
              area: this.ruleForm.area,
              director: this.ruleForm.director,
              duration: parseInt(this.ruleForm.duration),
              period: new Date(this.ruleForm.period).getTime()/1000,
              tags:this.ruleForm.tags.split(',')
            }) 
            .then(res=>{
              console.log(this.ruleForm);
              console.log(res.data); 
              if(res.data.status_code==0){
                this.dialogVisible = false
                window.alert('添加成功!!!')
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