import Vue from "vue"
import VueRouter from "vue-router"
import start from "../components/start.vue"
//解决冗余导航问题
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}
const originalreplace = VueRouter.prototype.replace
VueRouter.prototype.replace = function replace(location) {
  return originalreplace.call(this, location).catch(err => err)
}
const Home = () => import('../views/Home');
const Films = () => import ('../views/Films');
const List = () => import ('../views/List');
const TheFilm = () => import ('../views/TheFilm');
const Search = () => import ('../views/Search');
const Login = () => import ('../views/Register');
const Layout = () => import ('../views/Layout/index')
const BackHome = () => import ('../views/BackHome/index')
const Staff = () => import ('../views/Staff/index')
const Params = () => import('../views/Params/index')
const Invite = () => import('../views/Invite/index')
const Order = () => import('../views/Order/index')
const OrderList = () => import('../views/Order/OrderList/index')
const OrderBack = () => import('../views/Order/OrderBack/index')
const Theater = () => import('../views/Theater/index')

const Add = () => import ("../views/Theater/Add")
const Comment = () => import ('../views/comment');
const ChooseTicket = () => import ('../views/chooseTicket');
const ChooseSeat = () => import ('../views/chooseSeat');
const HomePage = () => import('../views/HomePage');
Vue.use(VueRouter);


export const routes = [
        {
            path:'/',
            redirect:'/home',
            name:'start',
            component:start,
            children: [
              {
                path: "/home",
                name: "home",
                component:Home,
                meta: { title: "首页" }
              },
              {
                  path: "/films",
                  name: "films",
                  component: Films,
                meta: { title: "电影" }

              },
              {
                path: "/list",
                name: "list",
                component: List,
                meta: { title: "榜单" }

              },
              {
                path: "/thefilm",
                name: "thefilm",
                redirect:'/comment',
                component: TheFilm,
                children:[
                  {
                    path: "/chooseticket",
                    name: "chooseticket",
                    component:ChooseTicket,
                  },
                  {
                    path: "/comment",
                    name: "comment",
                    component:Comment,
                  },
                ]
              },
              {
                path: "/search",
                name: "search",
                component: Search
              },
              {
                path: "/chooseSeat",
                name: "chooseSeat",
                component: ChooseSeat
              },
              {
                path: "/homepage",
                name: "homepage",
                component: HomePage
              },
            ]
        },
        {
          path: "/login",
          name:'login',
          component: Login,
          meta: { title: '登录'}
        },
        {
          path: '/layout',
          name: 'layout',
          redirect: '/backhome',
          component: Layout,
          meta: { isLogin:true},
          children : [
            {
              path: '/BackHome',
              name: 'BackHome',
              component: BackHome
            },
            {
              path: "/staff",
              name: "Staff",
              component: Staff
            },
            {
              path: "/params",
              name: "Params",
              component: Params
            },
            {
              path: "/invite",
              name: "Invite",
              component: Invite
            },
            {
              path: "/order",
              name: "Order",
              redirect: "/order/order-list",
              component: Order,
              children: [
                {
                  path: "order-list",
                  component: OrderList
                },
                {
                  path: 'order-back',
                  component: OrderBack
                }
              ]
            },
            {
              path: "/theater",
              name: "Theater",
              component: Theater
            },
            {
              path: "/add",
              name: "Add",
              component: Add
            }
          ]
        }
    ]
const router = new VueRouter({
    routes,
})

router.beforeEach((to, from, next) => {
  // if(to.path==='/chooseticket'&&!localStorage.token)
  // {
  //   return next('/login');
  // }
  // console.log('---to',to);
  //判断是否需要登录
  if(to.matched.some(ele=>ele.meta.isLogin)) {
    if(!sessionStorage.getItem('token') && !localStorage.getItem('token')){
      next({  
        path: '/login',
        query: {redirect: to.fullPath}  // 将跳转的路由path作为参数，登录成功后跳转到该路由 
      })
    }
    else {
      // console.log(next());
    	 next();
    }
  }
  else {
    next();
  }
});
export default router;

