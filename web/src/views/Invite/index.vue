<template>
  <div class="list">
     <!--1 搜索区域 -->
    <div class="header">
      <el-input @change="searchInp" v-model="input" placeholder="请输入内容"></el-input>
      <el-button type="primary" round @click="moviesearch">查询</el-button>
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
      prop="order_id"
      label="订单ID"
      width="180">
    </el-table-column>
    <el-table-column
      prop="movie_name"
      label="影院名称"
      width="180"
      show-overflow-tooltip>
    </el-table-column>
    <el-table-column
      prop="movie_avatar"
      label="电影封面">
      <template scope="scope">
        <img :src="scope.row.movie_avatar" width="100px" height="100px">
      </template>
    </el-table-column>
    <el-table-column
      prop="cinema_name"
      label="影厅名称"
      show-overflow-tooltip>
    </el-table-column>
      <el-table-column
      prop="create_at"
      label="创建时间"
      width="400"
      show-overflow-tooltip>
    </el-table-column>
      <el-table-column
      prop="price"
      label="价格"
      show-overflow-tooltip>
    </el-table-column>
          <el-table-column
      prop="seats"
      label="座位"
      show-overflow-tooltip>
    </el-table-column>
        <el-table-column
      prop="status"
      label="状态"
      show-overflow-tooltip>
    </el-table-column>
  </el-table> 
    </div>
    <div>
    <my-pagnation :total="total" :pageSize="pageSize" @changePage='changePage'></my-pagnation>
    </div>
  </div>
</template>

<script>
import MyPagnation from '../../components/MyPagnation.vue'
export default {
    components: {
    MyPagnation
  },
  data () {
    return {
      tableData:[],
      total:10,
      pageSize: 1,
      input:''
    }
  },
  methods: {
    http(page){
      this.$api.allOrderlist({
        page
      })
      .then(res=>{
        console.log(res.data);
        if(res.data.status_code==0){
          this.tableData = res.data.data.orders
          this.total = res.data.data.total
          this.pageSize = 10
        }
      })
    },
    changePage(num) {
      this.http(num)
    },
    moviesearch() {
      this.$api.searchOrder({
        condition: this.input,
        page:0,
      })
      .then(res=>{
        console.log(res.data);
        this.tableData = [];
        this.tableData = res.data.data.orders
        
      })
    }
  },
  created () {
    this.http(1)
  }
}
</script>

<style>
  .list {
    margin: 20px;
    margin-top: 50px;
  }
</style>