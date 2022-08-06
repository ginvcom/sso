import { createRouter, createWebHistory } from 'vue-router'
// import { createRouter, createWebHistory, RouteRecordRaw, RouteLocationNormalized, NavigationGuardNext } from 'vue-router'
// import User from '@/views/user.vue'
// import Role from '@/views/role.vue'
// import Menu from '@/views/menu.vue'

// 1. 定义路由组件.
// 也可以从其他文件导入
const Home = { template: '<div>Home</div>' }

// 2. 定义一些路由
// 每个路由都需要映射到一个组件。
// 我们后面再讨论嵌套路由。
const routes = [
  { path: '/', component: Home }
  // { path: '/sign-up', component: User },
  // { path: '/role', component: Role },
  // { path: '/forgot-password', component: Menu },
]

// 3. 创建路由实例并传递 `routes` 配置
// 你可以在这里输入更多的配置，但我们在这里
// 暂时保持简单
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
