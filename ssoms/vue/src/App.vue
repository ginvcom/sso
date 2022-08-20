<template>
  <a-config-provider :locale="zhCN">
    <a-layout>
      <a-layout-header class="header__bar" theme="dark">
        <div class="header__logo">
          <a class="header__logo-link" href="/">
            <img alt="Vue logo" src="./assets/logo.png" />
            <h1>单点管理系统</h1>
          </a>
        </div>
        <div class="header__actions">
          <ul>
            <li class="header__account-info">
              <a-dropdown>
                <a class="ant-dropdown-link" @click.prevent>
                  <a-avatar :size="24" :src="ossConfig.ginvdoc.domain + user.avatar">
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
          <a-menu-item key="1">
            <block-outlined />
            <router-link to="/object">对象管理</router-link>
          </a-menu-item>
          <a-menu-item key="2">
            <UserOutlined />
            <router-link to="/user">用户管理</router-link>
          </a-menu-item>
          <a-menu-item key="5">
            <appstore-outlined />
            <router-link to="/role">角色管理</router-link>
          </a-menu-item>
          </a-menu>
        </a-layout-sider>
        <a-layout-content class="main__bar">
          <router-view />
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-config-provider>
</template>
<script setup lang="ts">
import { onBeforeMount, reactive } from 'vue'
import zhCN from 'ant-design-vue/es/locale/zh_CN'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import {
  BlockOutlined,
  UserOutlined,
  AppstoreOutlined,
  SettingOutlined,
  LogoutOutlined,
} from '@ant-design/icons-vue'
import router from '@/router'
import { systemCode } from '@/config'
import { server } from '@/utils/ajax'
import { message } from 'ant-design-vue'
import { ossConfig } from '@/config'

dayjs.locale('zh-cn')

const user = reactive({ name: '', avatar: '' })

onBeforeMount(() => {
  getUserCookie()
})

const getUserCookie = () => {
  let arr = document.cookie.match(new RegExp('(^| )ssoUser=([^;]*)(;|$)'))
  if (arr !== null) {
    try {
      // 先uri_decode转义一下
      const u = JSON.parse(decodeURIComponent(arr[2]))
      user.name = u.name
      user.avatar = u.avatar
    } catch (err: Error) {
      message.error(err.message)
    }
  }
}

const state = reactive({
  selectedKeys: [],
  openKeys: []
})

const onHeaderClick = ({ key }) => {
  if (key === 'profile') {
    router.push('/profile')
  } else if (key === 'logout') {
    window.location.href = server + '/sign-out?system=' + systemCode
  }
}
</script>
