import Vue from 'vue'
import App from './App.vue'
import 'element-ui/lib/theme-chalk/index.css'
// import store from './store'
import router from './router'
import api from './api/index'
import axios from 'axios'
import { Button,Tree,Message,Alert, TabPane,Tag,Select,Menu,MenuItem,Input,Dialog ,RadioGroup,Radio,Carousel,carouselItem,Pagination,Rate,Submenu,Table,Row,Col,Tabs,TableColumn,MenuItemGroup,Checkbox,BreadcrumbItem,Breadcrumb,Form,FormItem,TimeSelect,DatePicker,  } from 'element-ui';
import * as echarts from 'echarts'
Vue.prototype.$echarts = echarts
Vue.use(Menu);
Vue.use(MenuItem);
Vue.use(Button);
Vue.use(Select);
Vue.use(Input);
Vue.use(Radio);
Vue.use(RadioGroup);
Vue.use(Carousel);
Vue.use(carouselItem);
Vue.use(Pagination);
Vue.use(Rate);
Vue.use(Submenu);
Vue.use(Table);
Vue.use(Row);
Vue.use(Col);
Vue.use(Tabs);  
Vue.use(TabPane);
Vue.use(TableColumn);
Vue.use(MenuItemGroup);
Vue.use(Checkbox);
Vue.use(BreadcrumbItem);
Vue.use(Breadcrumb);
Vue.use(Form)
Vue.use(FormItem)
Vue.use(TimeSelect)
Vue.use(DatePicker)
Vue.use(Dialog)
Vue.use(Tree)
Vue.use(Tag);
Vue.use(Alert)
Message.install = function(Vue, options) {
  Vue.prototype.$message = Message
}
Vue.use(Message)
Vue.config.productionTip = false
Vue.prototype.$api = api
Vue.prototype.$http = axios
Vue.config.silent = true;
new Vue({
  render: h => h(App),
  router,
  beforeCreate(){
      Vue.prototype.$bus= this
  }
  // store
}).$mount('#app')
axios.interceptors.request.use(
  function (config) {
  	const Token = window.localStorage.getItem('token')
    config.headers["Content-type"] = "application/json;";
  if (Token) {
    // 在请求头中添加Token
    // config.data = { unused: 0 };
    config.headers.Authorization = "Bearer "+Token;
    // config.headers.x = "application/json"
  }
  return config
}, function (error) {
  // 请求发送失败
  return Promise.reject(error)
})

axios.interceptors.response.use(function (response) {
  // 请求成功，处理响应数据
  return response
}, function (error) {
  // 请求失败，处理错误响应数据
  if (error.response && error.response.status === 401) {
    window.localStorage.removeItem('token')
    router.push('/login')
  }
  return Promise.reject(error)
})