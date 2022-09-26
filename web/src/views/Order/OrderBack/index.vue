<template>
  <div>
    <div class="films">
      <ul class="film">
        <li v-for="(item,index) in movies" :key="item">
          <div class="plan">
          <div class="planmovie"><img :src="item.avatar" ></div>
          <span class="add">
            <el-row>
              <el-button type="primary" @click="add(index)">添加演出计划</el-button>
    
          </el-row>
          </span>
          </div>
        </li>
      </ul>
    </div>
    <el-dialog
      title="添加剧目"
      :visible.sync="dialogVisible"
      width="70%"
      :before-close="handleClose">

  <!-- 内容区域 -->
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm"   class="demo-ruleForm">



<el-form-item label="开始时间(例如:2022-06-23 12:00:00)" prop="start_at">
    <el-input v-model="ruleForm.start_at"></el-input>
  </el-form-item>
<el-form-item label="输入影院" prop="cinema_id">
    <el-input v-model="ruleForm.cinema_id"></el-input>
  </el-form-item><el-form-item label="输入版本" prop="version">
    <el-input v-model="ruleForm.version"></el-input>
  </el-form-item><el-form-item label="输入价格" prop="price">
    <el-input v-model="ruleForm.price"></el-input>
  </el-form-item>
    <el-button type="primary" @click="submitForm('ruleForm')">确定</el-button>
    <el-button @click="resetForm('ruleForm')">重置</el-button>
</el-form>
</el-dialog>
    <my-pagnation :total="total" :pageSize="pageSize" @changePage='changePage'></my-pagnation>  

  </div>

</template>

<script>
import MyPagnation from '../../../components/MyPagnation.vue'

export default {
  data () {
    return {
      movies:[],
      total:10,
      pageSize: 1,
      movie_id:'',
      dialogVisible: false,
        ruleForm: {
          start_at: '',
          cinema_id:'',
          version:'',
          price:'',
          type: [],
        },
        rules: {
          start_at: [
            { required: true, message: '请输入开始时间', trigger: 'blur' },
          ],
          cinema_id: [
            { required: true, message: '请输入影厅', trigger: 'blur' }
          ],
           version: [
            { required: true, message: '请输入版本', trigger: 'blur' },
          ],  
          price: [
            { required: true, message: '请输入价格', trigger: 'blur' },
          ],  
    },

    }
  },
    components: {
    MyPagnation
  },
  methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            console.log(this.ruleForm);
            this.$api.getPlan({
          start_at:new Date(this.ruleForm.start_at).getTime()/1000,
          cinema_id:parseInt(this.ruleForm.cinema_id),
          version:this.ruleForm.version,
          price:parseFloat(this.ruleForm.price),
          movie_id:this.movie_id
        })
        .then(res=>{
          console.log(res.data);
          if(res.data.status_code==0){
            this.dialogVisible = false
            this.$message({
                                message: '添加成功!!',
                                center:true,
                                type: 'success'
                            })
          }
          if (res.data.status_code == 3003) {
            this.dialogVisible = false
            this.$message({
                                message: '时间冲突添加失败!!',
                                center:true,
                                type: 'error'
                            })
          }
                    if (res.data.status_code == 1001) {
            this.dialogVisible = false
           this.$message({
                                message: '起始时间晚于当前时间添加失败!!',
                                center:true,
                                type: 'error'
                            })
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
      },
      add(index) {
        this.dialogVisible = true
        this.movie_id = this.movies[index].id
      },
      http(page,pageSize) {
        this.$api.getMovieList({
        page,
        pageSize : 10
    }).then(res=>{
      if(res.data.status_code==0) {
              console.log(res.data);
      this.movies = res.data.data.list
      this.total = res.data.data.list[0].total
      this.pageSize = 10
      }
      
    })
      },
      changePage(num) {
        this.http(num)
      }
  },
  created () {
this.http(1)
  }
}
</script>

<style> 
  .plan {
    position: relative;
  }
  .films {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    height: 1000px;
  }

  .film li {
    float: left;
    height: 270px;
    width: 250px;
    margin-left: 100px;
    margin-top: 60px;
  }
  .planmovie img {
    width: 218px;
    height: 300px;
  }
  .add {
    position: absolute;
    bottom: 4px;
  }
  .add .el-row {
    margin-left: 50px;
    /* margin-top: 320px; */
  }

</style>