<template>
  <div class="staff">
     <!--1 搜索区域 -->
    <div class="header">
      <el-input @change="searchInp" v-model="input" placeholder="请输入内容"></el-input>
      <el-button type="primary" round @click="search">查询</el-button>
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
      prop="id"
      label="用户ID"
      width="180">
    </el-table-column>
    <el-table-column
      prop="username"
      label="用户名称"
      width="180"
      show-overflow-tooltip>
    </el-table-column>
        <el-table-column
      prop="privilege"
      label="用户权限"
      width="120">
    </el-table-column>
    <el-table-column
      prop="email"
      label="用户邮箱">
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
    <div><my-pagnation :total="total" :pageSize="pageSize" @changePage='changePage'></my-pagnation></div>
  </div>
</template>

<script>
import MyPagnation from '../../components/MyPagnation.vue'
export default {
  name: "Staff",
  components: {
    MyPagnation,
  },
  data() {
    return {
      page_size:1,
      tableData: [],
      total:10,
      input:''
    }
  },
  methods: {
    changePage(num){
      this.http(num)
    },
    http(page) {
      this.$api.getUser({
        page,
      })
      .then(res=>{
        console.log(res.data);
        if(res.data.status_code == 0) {
          // console.log(res.data);
          this.tableData = res.data.data.user_infos 
          this.total = res.data.data.user_num;
          this.page_size = 10
        }
      })
    }, 
    handleDelete(index,rows) {
      console.log('删除',index,rows);
      this.$api.deleteUser({
        UserId: rows.id
      })
      .then(res => {
        // console.log(res.data);
        window.alert('删除成功!!')
        this.http(1)
      })
    },
    search() {
      this.$api.getUsersearch({
        Username: this.input,
      })
      .then(res=>{
        // console.log(res.data);
        this.tableData = []
        this.tableData = res.data.data.user_infos
      })
    }
  },
  created() {
      this.http(1)
  }
}
</script>

<style>
  .staff {
    margin: 20px;
  }
</style>