<template>
  <div class="theater">
    <!--1 搜索区域 -->
    <div class="header">
      <el-input @change="searchInp" v-model="input" placeholder="请输入内容"></el-input>
      <el-button type="primary" round @click="moviesearch">查询</el-button>
        <el-button type="primary" round @click="add">添加</el-button>
<el-button
          size="mini"
          @click="edit()">编辑</el-button>
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
      label="ID"
      width="50"
      show-overflow-tooltip>
    </el-table-column>   
    <el-table-column    
      prop="name"
      label="剧目名称"
      width="180"
      show-overflow-tooltip>
    </el-table-column>
    <el-table-column
      prop="avatar"
      label="剧目封面"
      show-overflow-tooltip>
      <template scope="scope">
        <img :src="scope.row.avatar" width="100px" height="100px">
      </template>
    </el-table-column>
    <el-table-column
      prop="actors"
      label="演员列表"
      width="180">
    </el-table-column>

    <el-table-column
      prop="tags"
      label="剧目标签">
    </el-table-column>
    
        <el-table-column
      prop="duration"
      label="剧目时长">
    </el-table-column>
      <el-table-column
      prop="period"
      label="剧目上映时间">
    </el-table-column>
      <el-table-column
      prop="area"
      label="剧目地区">
    </el-table-column>
        <el-table-column
      prop="alias_name"
      label="剧目别名">
    </el-table-column>
      <el-table-column
      prop="content"
      label="剧目简介"
      width="250"
      show-overflow-tooltip>
    </el-table-column>
      <el-table-column
      prop="director"
      label="导演">
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
    <MyPagnation :total="total" :pageSize="pageSize" @changePage='changePage'></MyPagnation>
    <Dialog :dialogVisible='dialogVisible' @changeDialog='changeDialog'></Dialog>
    <Edit :dialogTableVisible="dialogTableVisible" @changeEdit='changeEdit' ></Edit>
  </div>
</template>

<script>
import Dialog from './dialog.vue'
import MyPagnation from '../../components/MyPagnation.vue'
import Edit from './Edit.vue'
export default {
    data () {
      return {
      dialogVisible : false,
      dialogTableVisible:false,
      tableData:[],
      total:10,
      pageSize: 10,
      input:''
      }
    },
    components: {
      MyPagnation,
      Dialog,
      Edit
    },
    methods: {
      add() {
      this.dialogVisible = true
      // this.$ref.dialog.dialogVisible = true;
      // this.$api.getCinema({
      //     cols:this.tableData.cols,
      //     name: this.tableData.name,
      //     rows:this.tableData.rows,
      //     avatar: this.tableData.avatar
      //   })
      //   .then(res=>{
      //     console.log(res);
      //   })

    },
    edit() {
      this.dialogTableVisible = true

    },
    changeDialog() {
      this.dialogVisible = false
    },
    changePage(num) {
    this.http(num);
      // else {
      //   //搜索分页
      //   console.log('搜索的分页处理--');
      // }
    },
    moviesearch() {
      this.$api.getMovieBykey({
        key: this.input,
        page:1,
      })
      .then(res=>{
        console.log(res.data.data);
        this.tableData = [];
        this.tableData = res.data.data.list
      })
    },
    handleDelete(index,rows) {
      console.log('删除',rows)
      this.$api.deleteMovie({
        movie_id: rows.id
      }) 
      .then(res=>{
        console.log(res.data);
        if (res.data.status_code==0){
          window.alert('删除成功!!!')
          this.http(1)
        }
        if(res.data.status_code==3013){
          window.alert('电影存在演出计划!!! 删除失败!!!')
          this.http(1)
        }
      })
    },
    http(page) {
      this.$api.getMovie({
        page,
      })
      .then(res=>{
        console.log(res.data);
        if (res.data.status_code==0){
          this.tableData = res.data.data.list;
          this.total = res.data.data.list[0].total;
          this.pageSize=10;
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