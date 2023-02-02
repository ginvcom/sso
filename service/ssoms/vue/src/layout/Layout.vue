<template>
  <a-layout>
    <a-layout-header class="header__bar" theme="dark">
      <div class="header__logo">
        <a class="header__logo-link" href="/">
          <img alt="Vue logo" src="../assets/logo.png" />
          <h1>单点管理系统</h1>
        </a>
      </div>
      <div class="header__actions">
        <ul>
          <li class="header__account-info">
            <a-dropdown>
              <a class="ant-dropdown-link" @click.prevent>
                <a-avatar :size="24" :src="ossConfig.doc.domain + user.avatar">
                  <template #icon><UserOutlined /></template>
                </a-avatar>
                <span class="header__account-content">hi! {{ user.name }}</span>
              </a>
              <template #overlay>
                <a-menu @click="onHeaderClick">
                  <a-menu-item key="profile">
                    <setting-outlined />
                    个人中心
                  </a-menu-item>
                  <a-menu-item key="logout">
                    <logout-outlined />
                    退出登录
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </li>
        </ul>
      </div>
    </a-layout-header>
    <a-layout>
      <a-layout-sider theme="light" class="sider__bar">
        <a-menu
          v-model:selectedKeys="state.selectedKeys"
          v-model:openKeys="state.openKeys"
          mode="inline"
          :style="{ borderRight: 0 }"
        >
        <menu-item v-for="menu in state.menus" :key="menu.n" v-bind="menu" />
        </a-menu>
      </a-layout-sider>
      <a-layout-content class="main__bar">
        <slot></slot>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script setup lang="ts">
import { onBeforeMount, onMounted, reactive } from 'vue'
import {
  UserOutlined,
  SettingOutlined,
  LogoutOutlined,
} from '@ant-design/icons-vue'
import router from '@/router'
import { systemCode } from '@/config'
import { server } from '@/utils/ajax'
import MenuItem from '@/components/MenuItem.vue'
import { message } from 'ant-design-vue'
import { ossConfig, SERVER_ROUTER_MENU_KEY } from '@/config'
import { Menu } from '@/api/auth'
import { useRoute } from 'vue-router'

const route = useRoute()
const user = reactive({ name: '', avatar: '' })

onBeforeMount(() => {
  getUserCookie()
})

onMounted(() => {
  getMenus(1)
})

const getUserCookie = () => {
  let arr = document.cookie.match(new RegExp('(^| )ssoUser=([^;]*)(;|$)'))
  if (arr !== null) {
    try {
      // 先uri_decode转义一下
      const u = JSON.parse(decodeURIComponent(arr[2]))
      user.name = u.name
      user.avatar = u.avatar
    } catch (err: any) {
      message.error(err.message)
    }
  }
}



interface State {
  menus: Array<Menu>
  selectedKeys: Array<string>
  openKeys: Array<string>
}

const state = reactive<State>({
  menus: [],
  selectedKeys: [],
  openKeys: []
})

const getMenus = (times: number) => {
  try {
    const serverRouter = sessionStorage.getItem(SERVER_ROUTER_MENU_KEY)
    if (serverRouter !== null) {
      state.menus = JSON.parse(serverRouter)
      getselectedAndOpendKeys(state.menus, [])
    }
  } catch (err) {
    console.log(err)
  }
}

const getselectedAndOpendKeys = (menus: Array<Menu>,parents: Array<string>) => {
  console.log(parents)
  for (const menu of menus) {
    console.log(route.path, menu.p)
    if (menu.p === route.path) {
      state.openKeys = parents
      state.selectedKeys = [menu.n]
      
      return
    }
    if (menu.c && menu.c.length > 0) {
      const newParent = [...parents, menu.n]
      getselectedAndOpendKeys(menu.c, newParent)
    }
  }
}

const onHeaderClick = ({ key }) => {
  if (key === 'profile') {
    router.push('/profile')
  } else if (key === 'logout') {
    window.location.href = server + '/sign-out?system=' + systemCode
  }
}
</script>
