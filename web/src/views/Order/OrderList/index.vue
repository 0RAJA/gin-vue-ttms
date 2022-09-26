
<template>
  <div class="invite">
    <!--1 搜索区域 -->
    <div class="header">
      
      <!-- <el-button type="primary" round>
        <router-link to="/add" style="color:#fff">页面添加</router-link>
      </el-button> -->
    </div>  
    <!-- 2 表格区域 -->
    <div class="wrapper">
    <el-table
    :data="tableData"
    stripe
    border
    style="width: 100%">
     <el-table-column
      type="selection"
      width="55">
    </el-table-column>
    <el-table-column
      prop="cinemaname"
      label="影厅"
      width="250">
    </el-table-column>
    <el-table-column
      prop="name"
      label="电影"
      width="350"
      show-overflow-tooltip>
    </el-table-column>
        <el-table-column
      prop="price"
      label="价格"
      width="120">
    </el-table-column>
    <el-table-column
      prop="start_at"
      label="开始时间"
      width="300">
    </el-table-column>
        <el-table-column
      prop="version"
      label="版本">
    </el-table-column>
    
      <el-table-column
      label="操作"
      width="180  ">
      <template slot-scope="scope">
        <el-button
          size="mini"
          type="danger"
          @click="handleDelete(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
    </div> 
        <!-- 3 分页展示 -->
    <my-pagnation :total="total" :pageSize="pageSize" @changePage='changePage'></my-pagnation>  
  </div>
</template>

<script>
import MyPagnation from '../../../components/MyPagnation.vue'

export default {
  data () {
    return {
      dialogVisible : false,
      total:10,
      tableData:[],
      pageSize: 1,
    }
  },
  components: {
    MyPagnation
  },
  methods: {
    add() {
      this.dialogVisible = true
    },
    changeDialog() {
      this.dialogVisible = false
    },
    changePage(num) {
    this.http(num)
  },
  http(page) {
    this.$api.getPlanpage({
      page,
    })
    .then(res=>{
      console.log(res.data);
      this.tableData = res.data.data.list
      this.total = res.data.data.list[0].total
      this.pageSize = 10
    })
  },
  handleDelete(index,rows){
    this.$api.deletePlan({
      planID: rows.id
    })
    .then(res => {
      // console.log(res.data);
      if(res.data.status_code==0){
        window.alert('删除成功!!!')
        this.http(1);
      }
    })
  },

  },
  created () {
    this.http(1)
  }
}
</script>

<style scoped>
  .invite {
    margin: 20px;
  }
</style>