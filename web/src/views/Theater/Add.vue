<template>
  <div class="add">
    <div class="title">
      <el-breadcrumb separator="/">
  <el-breadcrumb-item :to="{ path: '/layout'}">首页</el-breadcrumb-item>
  <el-breadcrumb-item :to="{ path: '/theater'}">剧目管理</el-breadcrumb-item>
  <el-breadcrumb-item>添加剧目</el-breadcrumb-item>
</el-breadcrumb>
    </div>
    <div class="myform">
      <!-- 表单数据 -->
    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
  <el-form-item label="剧目选择">
    <el-button type='primary'>剧目选择</el-button>
  </el-form-item>
  <el-form-item label="剧目名称" prop="name">
    <el-input v-model="ruleForm.name"></el-input>
  </el-form-item>
    <el-form-item label="剧目价格" prop="price">
    <el-input v-model="ruleForm.price"></el-input>
  </el-form-item>
  <el-form-item label="剧目数量" prop="num">
    <el-input v-model="ruleForm.num"></el-input>
  </el-form-item>
    <el-form-item label="剧目时间" prop="time">
       <div class="block">
       <el-date-picker
      v-model="value2"
      type="datetimerange"
      :picker-options="pickerOptions"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
      align="right">
    </el-date-picker>
  </div>
  </el-form-item>
  
    <el-form-item label="剧目类型" prop="name">
      <el-button type='primary'>剧目类型</el-button>
  </el-form-item>
  <el-form-item label="剧目描述" prop="name">
    <textarea name="" id="" cols="30" rows="10"></textarea>
  </el-form-item>
    <el-button type="primary" @click="submitForm('ruleForm')">确定</el-button>
    <el-button @click="resetForm('ruleForm')">重置</el-button>
</el-form>
    </div>
  </div>
</template>

<script>
 export default {
    data() {
      return {
        ruleForm: {
          name: '',
          num:'',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: '',
          price:''
        },
        rules: {
          name: [
            { required: true, message: '请输入剧目名称', trigger: 'blur' },
          ],
          price: [
            { required: true, message: '请输入剧目价格', trigger: 'blur' }
          ],
           time: [
            { required: true, message: '请选择剧目时间', trigger: 'blur' },
          ],
          num: [
            { required: true, message: '请输入剧目数量', trigger: 'change' }
          ],
          date1: [
            { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
          ],
          date2: [
            { type: 'date', required: true, message: '请选择时间', trigger: 'change' }
          ],
          type: [
            { type: 'array', required: true, message: '请至少选择一个剧目类型', trigger: 'change' }
          ],
          resource: [
            { required: true, message: '请选择活动资源', trigger: 'change' }
          ],
          desc: [
            { required: true, message: '请填写活动形式', trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
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
  .add {
    margin: 20px;
  }
  .title {
    padding: 10px;
    background: #fff;
    border: 1px solid #eee;
    margin-bottom: 10px;
  }
  .myform {
    background: #fff;
    padding: 10px;
    padding-right: 30px;
  }
</style>