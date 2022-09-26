<template>
  <div class="theater">
    <!--1 搜索区域 -->
    <div class="header">
      <el-input v-model="input" placeholder="请输入内容"></el-input>
      <el-button type="primary" round @click="getsearch">查询</el-button>
      <!-- <el-button type="primary" round>
        <router-link to="/add" style="color:#fff">页面添加</router-link>
      </el-button> -->
      <el-button type="primary" @click="add">添加</el-button>
              <el-button
          size="mini"
          @click="edit">编辑</el-button>
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
      label="影厅ID"
      width="180">
    </el-table-column>
    <el-table-column
      prop="name"
      label="影厅名称"
      width="300"
      show-overflow-tooltip>
    </el-table-column>
        <el-table-column
      prop="cols"
      label="行"
      width="120">
    </el-table-column>
    <el-table-column
      prop="rows"
      label="列"
      width="120">
    </el-table-column>
    <el-table-column
      prop="avatar"
      label="封面链接"
      show-overflow-tooltip>
      <template slot-scope="scope">
        <img :src="scope.row.avatar" width="100px" height="100px"> 
      </template>
    </el-table-column>

      <el-table-column
      label="操作"
      width="180 ">
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
    <!-- 4 弹框组件 1.父传子  2.ref-->
    <Add :dialogVisible='dialogVisible' @changeDialog='changeDialog'></Add>
    <Edit :dialogTableVisible="dialogTableVisible" @changeEdit='changeEdit' ></Edit>
    <!-- <Add ref="dialog"></Add> -->
  </div>
</template>

<script>    
import MyPagnation from '../../components/MyPagnation.vue'
import Add from './Dialog.vue'
import Edit from './Edit.vue'
// import request from '../../api/index'
export default {
  components: {
    MyPagnation,
    Add,
    Edit
  },
  data () {
    return {
      dialogVisible : false,
      dialogTableVisible:false,
      input: '',
      tableData: [
      ],
      total:10,
      pageSize: 1,
    }
  },
  methods: {
    //添加商品出现弹框
    add() {
      this.dialogVisible = true
    },
    edit() {
      this.dialogTableVisible = true
    },
    changeDialog() {
      this.dialogVisible = false
    },
    changeEdit() {
      this.dialogTableVisible = false
    },
    //分页的页码
    changePage(num) {
        this.http(num);
      // else {
      //   //搜索分页
      //   console.log('搜索的分页处理--');
      // }
    },
    getsearch(){
      this.$api.getSearchcinema({
        cinema_id: this.input
      })
      .then(res=>{
        console.log(res.data);
        this.tableData = [];
        this.tableData[0] = res.data.data
      })
    },
      handleDelete(index,rows) {
        console.log('删除',index,rows);
        this.$api.deleteCinema({
          cinema_id: rows.id
        })
        .then(res=>{
          console.log(res.data);
            if(res.data.status_code === 0){
              window.alert('删除成功!!!')
              this.http(1)
            }
            if(res.data.status_code == 3012){
              window.alert('删除失败!!! 影厅存在演出计划!!')
              this.http(1)
            }
        })
      },
    //剧目列表获取
    http(page) {     
      this.$api.getorderList({
      page,
      // page_size:8 
    })
    .then(res => {
      console.log(res.data);
      if (res.data.status_code === 0){
        this.tableData = res.data.data.list//数据列表
        this.total = res.data.data.list[0].total;
        this.pageSize = 10
      }
    })}
  },  
  //生命周期函数
  mounted () {
    this.http(1)
  } 
}
</script>

<style>
  .theater {
    margin: 20px;
  }
  .header {
    display: flex;
  }
  .header button {
    margin-left: 20px;
  }
  .wrapper {
    margin: 20px 0;
  }
</style>