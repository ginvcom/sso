import { createRouter, createWebHistory, RouteRecordRaw, RouteLocationNormalized } from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import Home from '@/views/home.vue'
import NotFound from '@/NotFound.vue'
import { SERVER_ROUTER_MENU_KEY } from '@/config'
import { menus, Menu } from '@/api/auth'

// 1. 定义一些路由
const routes: RouteRecordRaw[] = [
  { path: '/', component: Home },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
]

const modules = import.meta.glob("@/views/**/**.vue")

let asyncRouter: Array<Menu> | null

// 2. 创建路由实例并传递 `routes` 配置
const router = createRouter({
  history: createWebHistory(),
  routes
})

// https://router.vuejs.org/zh/guide/advanced/navigation-guards.html#%E5%85%A8%E5%B1%80%E5%89%8D%E7%BD%AE%E5%AE%88%E5%8D%AB
router.beforeEach(async to => {
  NProgress.start()
  if (!asyncRouter) {
    try {
      const serverRouter = sessionStorage.getItem(SERVER_ROUTER_MENU_KEY)
      if (serverRouter !== null) {
        asyncRouter = JSON.parse(serverRouter)
      } else {
        const { list: routers } = await menus()
        sessionStorage.setItem(SERVER_ROUTER_MENU_KEY, JSON.stringify(routers))
        asyncRouter = routers
      }
    } catch (err) {
      // 意料之外的错误，取消导航并把错误传给全局处理器
      // https://router.vuejs.org/zh/guide/advanced/navigation-guards.html#%E5%85%A8%E5%B1%80%E8%A7%A3%E6%9E%90%E5%AE%88%E5%8D%AB
      throw err
    }
    const formatedRouters = routerFormat(asyncRouter)
    if (formatedRouters) {
      formatedRouters.forEach(item => {
        router.addRoute(item)
      })
    }
    return to.fullPath
  }
})

router.afterEach(() => {
  NProgress.done(asyncRouter)
})

const routerFormat = (routers: Array<Menu> | null) => {
  const records:RouteRecordRaw[] = []
  if (routers === null) {
    return
  }
  routers.forEach(item => {
    const { p: path, n: name } = item
    const meta = { ...item.m }
    const children = routerFormat(item.c)
    const record: RouteRecordRaw = { path, name, meta, component: modules[`../views${path}.vue`], children }
    if (item.m.t == 3) {
      router.addRoute(record)
    } else {
      records.push(record)
    }
  })
  return records
}

// https://router.vuejs.org/zh/api/#onerror
router.onError((err: any, to: RouteLocationNormalized, from: RouteLocationNormalized) => {
  console.log(err)
  // router.push({ path: '/error', query: { msg: error }})
})

export default router
