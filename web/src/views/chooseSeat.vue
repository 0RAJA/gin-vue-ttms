<template>
    <div>
        <div class="main">
            <div class="seat">
                <div class="flag">
                    <img src="../assets/img/seat/可选座位.png" alt="zuowei">
                    <span>可选座位</span>
                    <img src="../assets/img/seat/已售座位.png" alt="zuowei">
                    <span>已售座位</span>
                    <img src="../assets/img/seat/已选座位.png" alt="zuowei">
                    <span>已选座位</span>
                </div>
                <div class="seatlayout">
                    <div>   
                        <img src="../assets/img/yingmu.jpg" alt="">
                        <br>
                        银幕中央
                    </div>
                    <ul class="seatList">
                        <li class="seatRow" v-for="index in seat.length" :key="index">
                        {{index}}
                            <ul class="seatRowList">
                                <li class="seatItem" v-for="(item,x) in seat[index-1]" :key="item.id">
                                    <img v-show="item.status=='正常'&&item.seats_status=='正常'" @click="changeState(index-1,x)"  src="../assets/img/seat/可选座位.png" alt="zuowei">
                                    <img v-show="(item.status=='已售'||item.status=='锁定')&&item.seats_status=='正常'" src="../assets/img/seat/已售座位.png" alt="zuowei">
                                    <img v-show="item.status=='选中'&&item.seats_status=='正常'" @click="changeState(index-1,x)" src="../assets/img/seat/已选座位.png" alt="zuowei">
                                    <div v-show="item.seats_status=='走廊'" class="emptySeat"></div>
                                </li>
                            </ul> 
                        </li>
                    </ul>
                </div>
            </div>
            <div class="movie">
                <div class="header">
                    <img :src="movieInfo.cover" alt="movie">
                    <div class="title">
                        <h2>{{movieInfo.name}}</h2>
                        <p class="labelname">类型：<span>{{movieInfo.type}}</span></p>
                        <p class="labelname">时长：<span>{{movieInfo.duration}}分钟</span></p>
                    </div>
                    <p class="labelname">影厅：<span>{{JSON.parse($route.query.ticket).cinemaname}}</span></p>
                    <p class="labelname">版本：<span>{{JSON.parse($route.query.ticket).version}}</span></p>
                    <p class="labelname">场次：<span>{{JSON.parse($route.query.ticket).end_at.slice(0,10)}} {{JSON.parse($route.query.ticket).start_at}}</span></p>
                    <p class="labelname">票价：<span>￥{{JSON.parse($route.query.ticket).price}}/张</span></p>
                </div>
                <div class="price">
                    <p class="labelname">
                        座位：
                        <el-tag 
                        v-for="(item,index) in selectSeat" :key="index"
                        color="#f56c6c" style="color:#fff;">{{item[0]}}排{{item[1]}}座</el-tag>
                    </p>
                    <div class="allprice">
                        总价：<div>￥{{JSON.parse($route.query.ticket).price*selectSeat.length}}</div>
                    </div>
                </div>
                <div class="enter">
                    <el-button 
                    @click="toPay()"
                    @click.once="deadline()"
                    style="width: 300px"
                    type="danger" round>确认选座</el-button>
                </div>
            </div>
        </div>
        <div class="pay" v-if="isPay">
            <div class="time">
                {{timeRemaining.minutes}}:{{timeRemaining.seconds>=10?timeRemaining.seconds:'0'+timeRemaining.seconds}}
            </div>
            <div id="qrcode">
            </div>

        </div>
    </div>
</template>

<script>
 import QRCode  from "qrcodejs2"
export default {
    name:'chooseSeat',
    data() {
        return {
            seat:[
            ],
            timeRemaining:{
                minutes:15,
                seconds:0,
            },
            selectSeat:[],
            isPay:false,
            pay_url:'',
            order_id:'',
            poll:null,
        }
    },
    components:{
        QRCode,
    },
    computed:{
        movieInfo()
        {
            return JSON.parse(this.$route.query.movieInfo)
        }
    },
    methods: {
        changeState(row, column){
            let newState =this.seat[row][column]
            if(newState.status=='正常'&&this.selectSeat.length<5)
            {
                newState.status='选中'
                this.selectSeat.push([row+1,column+1]);
            }
            else{
                newState.status='正常' 
                for(let i in this.selectSeat)
                {
                    if(this.selectSeat[i][0]==row+1&&this.selectSeat[i][1]==column+1)
                        this.selectSeat.splice(i,1);
                }   
            }
            this.seat[row].splice(column,1,newState);
        },
        deadline(){
            this.poll=setInterval(this.payment ,1000)
            let timer=setInterval(()=>{
                if(this.timeRemaining.minutes>=0)
                {
                    if(this.timeRemaining.minutes==0&&this.timeRemaining.seconds==0)
                    {
                        this.isPay=false;
                        clearInterval(timer);
                        clearInterval(this.poll);
                        return;
                    }       
                    if(this.timeRemaining.seconds==0)
                    {
                        this.timeRemaining.seconds=60;
                        this.timeRemaining.minutes--;
                    }
                    this.timeRemaining.seconds--;
                }
            },1000)
        },
        payment(){
            this.$api.IsPay({
                uuid: this.order_id
            }).then(res=>{
                if(res.data.data.is_pay==true){
                    clearInterval(this.poll);
                    this.payorder();
                }
            })
        },
        payorder(){
            let user_id = JSON.parse(localStorage.getItem("user")).userid;
            let plan_id=this.seat[0][0].plan_id;
            let arr=[];
            for(let i of this.selectSeat)
            { 
                arr.push(this.seat[i[0]-1][i[1]-1].seat_id)
            }
            this.$api.toPay({
                order_id:this.order_id,
                plan_id,
                seats_id:arr,
                user_id,
            }).then(res => {
                console.log(res.data);
                if(res.data.status_code==0)
                {
                    this.isPay=false;
                    this.$message({
                        message: '支付成功',
                        type: 'success'
                        }); 
                    this.$router.push({name:'homepage'});
                }
            })
        },
        rocking(seats_id){
            let user_id = JSON.parse(localStorage.getItem("user")).userid;
            let plan_id=this.seat[0][0].plan_id;
            this.$api.rockTicket(
                {
                    plan_id,
                    seats_id,
                    user_id,
                }
            ).then(res=>{
                this.pay_url=res.data.data.pay_url,
                this.order_id=res.data.data.order_id,
                this.$nextTick (function () {
                    this.qrcode();
                    })
                if(res.data.status_code == 0)
                {
                    this.$message({
                        message: '请尽快支付',
                        type: 'warning'
                        }); 
                }
                else{
                    this.$message({
                        message: '请刷新重试',
                        type: 'warning'
                        }); 
                }
            })
        },
        toPay(){
            if(!confirm("确定购买吗？"))
            { 
                return
            }
            if(this.selectSeat.length>0)
            {
                let arr=[];
                for(let i of this.selectSeat)
                { 
                    arr.push(this.seat[i[0]-1][i[1]-1].seat_id)
                }
                this.rocking(arr)
                if(this.isPay==false)
                {
                    this.isPay=true;
                }
            }
            else{
                this.$message({
                        message: '请选座',
                        type: 'warning'
                        }); 
            }

        },
        qrcode () {
                let qrcode = new QRCode('qrcode', {
                    width: 124,
                    height: 124,        // 高度
                    text:  this.pay_url,   // 二维码内容
                    // render: 'canvas' ,   // 设置渲染方式（有两种方式 table和canvas，默认是canvas）
                    // background: '#f0f',   // 背景色
                    // foreground: '#ff0'    // 前景色
                })
            },
            
    },
    created(){
        this.$api.getSeatlist({
            PlanId:JSON.parse(this.$route.query.ticket).id
        }).then(res=>{
            this.seat=res.data.data.list.tickets
        },err=>{
            console.log(err)
        })
    },
    beforeRouteEnter(to,from,next){
        let token=localStorage.getItem('token');
            if(token)
            {
                next()
            }
    }
}
</script>

<style scoped>
.time{
    font-size: 35px;
    text-align: center;
}
#qrcode{
    display: flex;
    justify-content:center;
    margin:40px 0px;
}
.pay{
    padding:20px;
    position: fixed;
    width: 300px;
    border:1px solid #ccc;
    box-shadow: 3px 3px 3px #ccc;
    top:50%;
    left:50%;
    background:#fff;
    text-align: center;
    transform:translate(-50%,-50%);
}
.enter{
    text-align: center;
    margin: 40px 0;
}
.allprice{
    color:#000;
    padding: 30px 0px;
    border-bottom: 1px solid rgb(155, 150, 150);

}
.allprice >div{
    display: inline-block;
    color:#f56c6c;
    font-size:30px;
}
.labelname {
    color: rgb(155, 150, 150);
    font-size:14px;
}
.labelname span{
    color: #000;
}
.labelname:nth-child(6){
    padding-bottom: 10px;
    border-bottom: 1px solid rgb(155, 150, 150);
}
.title{
    display: inline-block;
    margin-left:20px;
    vertical-align:top;
}
.header img{
  width: 160px;
    height:222.62px;
}
.title h2{
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    width:230px;
}
.movie{
    width:30%;
    padding:54px;
    background-color:#f3f0f0;
}
.emptySeat {
    width: 48px;
    height: 48px;
}
.seatRowList{
    list-style: none;
    white-space: nowrap;
}
.seatItem{
    display: inline-block;
    margin-left: 6px;
    cursor: pointer;
}
.emptySeat{
    margin:0px;
}
.seatList {
    overflow-x: auto;
    list-style: none;
}
 /* 滚动条宽度 */
.seatList::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}
 /* 滚动条轨道 */
.seatList::-webkit-scrollbar-track {
  background: rgb(239, 239, 239);
  border-radius: 2px;
}
/*  小滑块 */
.seatList::-webkit-scrollbar-thumb {
  background: #40a0ff49;
  border-radius: 10px;
}
.seatList::-webkit-scrollbar-thumb:hover {
  background: #40a0ff;
}

.seatRow {
    margin-top:6px;
    display:flex;
    align-items: center;
    color:#999;
}
.flag{
    display: flex;
    align-items: center;
    margin-left: 80px;
}
.flag img{
    margin-left:40px;
    margin-right:5px;
}
.main{
    margin: 0 auto;
    margin-top: 100px;
    display:flex;
    border-top: 1px solid #ccc;
    border-bottom: 1px solid #ccc;
}
.seat{
    width:70%;
}
.seatlayout >div{
    text-align: center;
    color: #999;
    margin:15px 0;
}
</style>