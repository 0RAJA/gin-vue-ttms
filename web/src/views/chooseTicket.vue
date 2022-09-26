<template>
  <div>
    <div class="detail">
          <h3 class="labelname">{{$route.query.movieInfo.name}}</h3>
          <hr/>
        <div>
          <div class="viewing">
            <span>观影时间：</span>
            <div @click="handletime" :class="{time:true,isSelected:currentTime==timelabel[0]}">{{timelabel[0]}}</div>
            <div @click="handletime" :class="{time:true,isSelected:currentTime==timelabel[1]}">{{timelabel[1]}}</div>
            <div @click="handletime" :class="{time:true,isSelected:currentTime==timelabel[2]}">{{timelabel[2]}}</div>
          </div>
        </div>
        <br>
        <br>
        <el-table
        :data="tableData"
        stripe
        style="width: 100%">
        <el-table-column
          prop="start_at"
          label="放映时间"
          width="180">
        </el-table-column>
        <el-table-column
          prop="version"
          label="语言版本"
          width="180">
        </el-table-column>
        <el-table-column
          prop="cinemaname"
          label="放映厅">
        </el-table-column>
        <el-table-column
          prop="price"
          label="售价（元）">
        </el-table-column>
        <el-table-column
          label="选座购票">
          <template slot-scope="scope">
          <el-button
          size="small "
          type="primary" round
          @click="chooseSeat(scope.$index)">选座购票</el-button>
        </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import request from '../api/index'
export default {
  data() {
    return {
       tableData: [],
        currentTime: '今天',
        timelabel:['今天','明天','后天'],
        period: {
          start: '',
          end: '',
        }
    }
  },
  methods:{
    handletime(e){
      this.currentTime=e.target.innerHTML
      switch(this.currentTime){
          case '今天':this.period.start=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()}`);
                           this.period.end=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()+1}`);
                           break;
          case '明天':this.period.start=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()+1}`);
                           this.period.end=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()+2}`);
                           break;
          case '后天':this.period.start=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()+2}`);
                           this.period.end=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()+3}`);
                           break;
      }
      request.getPlanlist({
        movie_id: this.$route.query.movie_id,
        start_time:`${this.period.start}`.slice(0,10),
        end_time:`${this.period.end}`.slice(0,10),
      }).then(
        res=>{
          for(let i of res.data.data.list)
          {
            i.start_at=i.start_at.slice(11,16);
          }
          this.tableData=res.data.data.list;
        },
        err=>{
          console.log(err);
        }
      )
    },
    chooseSeat(x){
      let token=localStorage.getItem("token");
      if(token)
          this.$router.push({name:'chooseSeat',query:{ticket:JSON.stringify(this.tableData[x]),movieInfo: JSON.stringify(this.$route.query.movieInfo)}})
      else{
        this.$router.push({name:'login'});
      }
    }
  },
  created(){
    this.period.start=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()}`)
    this.period.end=Date.parse(`${new Date().getFullYear()}-${new Date().getMonth()+1}-${new Date().getDate()+1}`);
    request.getPlanlist({
      movie_id: this.$route.query.movie_id,
      start_time:`${this.period.start}`.slice(0,10),
      end_time:`${this.period.end}`.slice(0,10),
    }).then(
      res=>{
        for(let i of res.data.data.list)
        {
          i.start_at=i.start_at.slice(11,16);
        }
        this.tableData=res.data.data.list;
      },
      err=>{
        console.log(err);
      }
    )
  },
  beforeRouteLeave(to,from,next){
        let token=localStorage.getItem('token');
        if(to.name=='chooseSeat')
        {
          if(token)
            {
                next()
            }
        }
        next()
    }
}
</script>

<style scoped>
.time {
  border-radius:14px;
  padding: 3px 9px;
  display: inline-block;
  margin-left: 12px;
  cursor: pointer;
}
.isSelected{
  background-color: #75c4ff;
  color: #fff;
}
.labelname::before{
    content: "";
    display: inline-block;
    width: 4px;
    height: 18px;
    margin-right: 6px;
    background-color: #11a8cd;
}
.detail{
    width:1000px;
    margin:50px auto;
    padding:10px;
}
</style>