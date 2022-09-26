<template>
  <div>
    <p style="color:#409EFF">我的订单</p>
        <hr>
        <ul class="orderlist">
            <li class="orderitem" v-for="item in orderlist" :key="item.order_id">
                <div class="order-header">
                    <span class="order-date">{{item.create_at.slice(0,10)}}</span>
                    <span>订单号:{{item.order_id}}</span>
                </div>
                <div class="order-body">
                    <div class="poster">
                        <img :src="item.movie_avatar" alt="movie">
                    </div>
                    <div class="content">
                        <div class="movie-name">《{{item.movie_name}}》</div>
                        <div class="hall">
                            <span>{{item.cinema_name}}</span>
                            <span v-for="x,index in item.seats" :key="index">{{x.slice(0,2)}}排{{x.slice(4)}}座</span>
                        </div>
                        <div class="showtime">{{item.create_at.slice(0,10)}} {{item.create_at.slice(11,16)}}</div>
                    </div>
                    <div class="order-price">￥{{item.price}}</div>
                    <div class="order-status">{{item.status}}</div>
                    <div class="action" v-show="item.status=='待支付'"><el-button type="primary" round @click="toPay(item)">去付款</el-button></div>
                </div>
            </li>
        </ul>
        <div class="pay" v-if="isPay">
            <div id="qrcode">
            </div>
            <el-button type="primary" @click="noPay">退出</el-button>
        </div>
  </div>
</template>

<script>
 import QRCode  from "qrcodejs2"
export default {
    data() {
        return{
            orderlist:[],
            isPay:false,
            poll:null,
        }
    },
    components:{
        QRCode,
    },
    methods: {
        qrcode (orderid) {
                let qrcode = new QRCode('qrcode', {
                    width: 124,
                    height: 124,        // 高度
                    text:  `https://ttms.humraja.xyz/ttms/ticket/payUrl?uuid=${orderid}`,   // 二维码内容
                })
            },
        noPay(){
                clearInterval(this.poll);
                this.isPay=false;
        },
        toPay(item){
                if(this.isPay==false)
                {

                    this.isPay=true;
                    this.$nextTick (function () {
                    this.qrcode(item.order_id);
                    })
                    this.poll=setInterval(this.payment.bind(null, item),1000)
                }
            },
            payment(item){
                this.$api.IsPay(
                {
                    uuid: item.order_id,
                }
                ).then(res=>{
                    if(res.data.data.is_pay==true){
                    clearInterval(this.poll);
                    this.payorder(item);
                }
                })
            },
            payorder(item){
                this.$api.toPay({
                    order_id:item.order_id,
                    plan_id:item.plan_id,
                    seats_id:item.seats_id,
                    user_id:item.user_id,
                }).then(res=>{
                     if(res.data.status_code==0)
                    {
                        this.isPay=false;
                        this.$message({
                            message: '支付成功',
                            type: 'success'
                            }); 
                        item.status='已支付'
                    }
                })
            }
    },
    created(){
        let user_id = JSON.parse(localStorage.getItem("user")).userid;
        this.$api.getOrder(
            { 
                user_id,
            }
        ).then(
            res=>{
                this.orderlist=res.data.data.list.orders
            }
        )
        
    }
}
</script>

<style scoped>
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
.content {
    width: 49%;
}
.action{
    line-height: 95px;
    font-size: 14px;
    color: #fff;
}
.order-status {
    width: 15%;
    font-size: 14px;
    color: #333;
    line-height: 95px;
}
.order-price {
    font-size: 14px;
    color: #333;
    width: 12%;
    line-height: 95px;
}
.orderlist{
    list-style: none;
    width: 880px;
    margin: 0 auto;
}
.orderitem{
    border: 1px solid #e5e5e5;
    margin: 0 40px 30px 0;
}
.order-header{
    padding: 16px 20px;
    background-color: #f7f7f7;
    font-size: 14px;
}
.order-date {
    color: #333;
    display: inline-block;
    margin-right: 30px;
}
.order-body{
    padding: 20px;
    padding-right: 0;
}
.poster{
    border: 2px solid #fff;
    box-shadow: 0 1px 2px 0 hsl(0deg 0% 53% / 50%);
    margin-right: 11px;
    font-size: 0;
}
.poster img{
    widows: 66px;
    height:91px;
}
 .order-body>div {
    display: inline-block;
    vertical-align: top;
}
.movie-name {
    font-size: 16px;
    font-weight: 700;
    color: #333;
    margin: 4px 0 7px -6px;
}
.hall{
    font-size: 12px;
    color: #999;
    margin-bottom: 4px;
}
 .showtime {
    font-size: 12px;
    color: #f03d37;
}
</style>