/*
  接口路径配置

*/
let baseURL = 'https://ttms.humraja.xyz/ttms';

const base = {
  orderList: baseURL +'/cinema/list',//剧目列表
  searchcinema:baseURL +'/cinema/get',//影厅搜索功能
  login: baseURL +'/user/login',//登录接口
  send: baseURL +'/email/send',//验证码
  register: baseURL +'/user/register',
  addcinema: baseURL +'/cinema/create',//添加影厅
  delete: baseURL +'/cinema/delete',
  repeat:baseURL +'/user/isRepeat',
  checkcinema:baseURL +'/cinema/check_name',//是否重复
  movieList: baseURL +'/movie/list/visit_count',//电影榜单1
  movieListByfollow: baseURL +'/movie/list/user_movie_count',//电影榜单2
  theFilm:baseURL +'/movie/get',//电影详情
  allFilm:baseURL +'/movie/list/tag_period_area',//根据标签获取电影
  comment:baseURL +'/comment/list/movie_id',//根据标签获取电影
  movieBykey:baseURL +'/movie/list/key',//关键字搜索电影
  planlist:baseURL +'/plan/list/movie_id',//根据电影获取演出计划
  createplan: baseURL +'/plan/create',   
  createmovie: baseURL +'/movie/create',
  deletemovie: baseURL +'/movie/delete',
  getmovie: baseURL +'/movie/list/info',//电影列表
  visitmovie: baseURL +'/movie/list/box_office',//首页票房
  getuser: baseURL +'/user/listInfo',//获取用户
  deleteuser: baseURL +'/user/delete',//删除用户
  seatlist:baseURL +'/ticket/list',//座位列表
  userInfo:baseURL +'/user/get',//用户信息
  planpage: baseURL +'/plan/list/info',//演出计划
  deleteplan: baseURL +'/plan/delete',
  gettickets: baseURL +'/ticket/listAll',
  updatecinema: baseURL +'/cinema/update',
  updateInfo:baseURL +'/user/info/modify',//更新用户信息
  updatemovie: baseURL + '/movie/update',
  usersearch: baseURL + '/user/search',//搜索用户
  allorderlist: baseURL + '/order/listAll',
  updateInfo:baseURL+'/user/info/modify',//更新用户信息
  upFile:baseURL+'/upload/file',//上传文件
  upAdatar:baseURL+'/user/updateAvatar',//更新头像
  rock:baseURL+'/ticket/check',//锁定票
  order:baseURL+'/order/list',//我的订单
  isPay:baseURL+'/ticket/qrresult',//是否支付
  Pay:baseURL+'/ticket/pay',//支付操作
  searchorder: baseURL+ '/order/listCondition',
  delComment: baseURL+'/comment/delete'//删除评论
}

export default base;
