/*
  请求的方法
*/
import axios from "axios"
import base from "./base"
const api = {
  //剧目列表
  getorderList(params) {//{page:xx}
    return axios.get(base.orderList,{
      params
    });
  },
  getLogin(params) {
    return axios.post(base.login,params)
  },
  getMail(params) {
    return axios.post(base.send,params)
  },
  getRegister(params) {
    return axios.post(base.register,params)
  },
  getCinema(params) {
    return axios.post(base.addcinema,params)  
  },
  deleteCinema(x) {
    return axios.delete(base.delete,{data:x})
  },
  isRepeat(params) {
    return axios.get(base.repeat,{
      params
    })
  },
  //获取榜单电影1
  getMovieList(params){
    return axios.get(base.movieList, {
      params
    })
  },
  //获取榜单电影2
  getMovieListByfollow(params){
    return axios.get(base.movieListByfollow, {
      params
    })
  },
  //获取电影详情
  getFilmDetail(params){
    return axios.get(base.theFilm, {
      params
    })
  },
  //根据标签获取电影
  getAllfilms(params){
    return axios.get(base.allFilm, {
      params
    })
  },
  getComment(params){
    return axios.get(base.comment, {
      params
    })
  },
  //关键字搜索
  getMovieBykey(params) {
    return axios.get(base.movieBykey, {
      params
    })
  },
  //获取演出计划
  getPlanlist(params) {
    return axios.get(base.planlist, {
      params
    })
  },
  getSearchcinema(params) {
    return axios.get(base.searchcinema, {
      params  
    })
  },
  getPlan(params) {
    return axios.post(base.createplan,params)
  },
  getCreatemovie(params) {
    return axios.post(base.createmovie,params)
  },
  deleteMovie(x) {
    return axios.delete(base.deletemovie,{data:x})
  },
  getMovie(params) {
    return axios.get(base.getmovie,{
      params
    })
  },
  getVisitmovie(params) {
    return axios.get(base.visitmovie,{
      params
    })
  },
  getUser(params) {
    return axios.get(base.getuser,{params})
  },
  //获取座位列表
  getSeatlist(params){
    return axios.post(base.seatlist, 
      params
    )
  },
  deleteUser(params) {
    return axios.post(base.deleteuser,params)
  },
  checkCinema(params) {
    return axios.get(base.checkcinema,{params})
  },
  //获取用户信息
  getuserInfo(params){
    return axios.get(base.userInfo, {
      params
    })
  },
  getPlanpage(params) {
    return axios.get(base.planpage,{params})
  },
  deletePlan(x) {
    return axios.delete(base.deleteplan,{data:x})
  },
  getTickets(params) {
    return axios.get(base.gettickets,{params})
  },
  updateCinema(params) {
    return axios.put(base.updatecinema,params)
  },
  //更新用户信息
  updateInfo(params){
    return axios.put(base.updateInfo,
      params
    )
  },
  updataMovie(params) {
    return axios.put(base.updatemovie,params)
  },
  //上传文件
  uploadFile(params){
    return axios.post(base.upFile, 
      params
    )
  },
  //更新头像
  updateAdatar(params){
    return axios.put(base.upAdatar,
      params
    )
  },
  //锁定票
  rockTicket(params){
    return axios.post(base.rock, 
      params
    )
  },
  //获取用户订单
   getOrder(params){
    return axios.get(base.order, {params})
  },
  //轮询是否支付
  IsPay(params){
    return axios.get(base.isPay, {
      params
    })
  },
  //支付操作
  toPay(params){
    return axios.post(base.Pay, 
      params
    )
  },
  getUsersearch(params) {
    return axios.get(base.usersearch,{
      params
    })
  },
  allOrderlist(params) {
    return axios.get(base.allorderlist,{
      params
    })
  },
  searchOrder(params) {
    return axios.get(base.searchorder,{params})
  },
  deleteComment(x){
    return axios.delete(base.delComment,{data:x})
  }
}
export default api