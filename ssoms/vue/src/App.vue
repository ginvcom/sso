<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png" /> -->
  <!-- <HelloWorld msg="Hello Vue 3 + Vite" /> -->
  <a-layout>
    <a-layout-header class="header__bar">
      <div class="header__logo">
        <img alt="Vue logo" src="./assets/logo.png" />
        <h1>单点管理系统</h1>
      </div>
    </a-layout-header>
    <a-layout>
      <a-layout-sider theme="light" class="sider__bar">
        <a-menu
          v-model:selectedKeys="state.selectedKeys"
          v-model:openKeys="state.openKeys"
          mode="inline"
          :style="{ height: '100%', borderRight: 0 }"
        >
        <a-menu-item key="1">
          <bars-outlined />
          <router-link to="/menu">对象管理</router-link>
        </a-menu-item>
        <a-menu-item key="2">
          <audit-outlined />
          <router-link to="/user">用户管理</router-link>
        </a-menu-item>
        <a-menu-item key="5">
          <appstore-outlined />
          <router-link to="/user">角色管理</router-link>
        </a-menu-item>
        <a-menu-item key="6">
          <audit-outlined />
          <router-link to="/user">授权管理</router-link>
        </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <a-layout-content class="main__bar">
        <router-view />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script setup lang="ts">
// import HelloWorld from './components/HelloWorld.vue'
import { reactive } from 'vue'
import {
  HomeOutlined,
  BarsOutlined,
  AuditOutlined,
  AppstoreOutlined,
  LoadingOutlined,
} from '@ant-design/icons-vue'

const state = reactive({
  selectedKeys: [],
  openKeys: []
})
</script>
<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.header__bar.ant-layout-header{
  height: 48px;
  background-color: #fff;
  box-shadow: 0 0 8px 0 rgb(0 0 0 / 10%);
  display: flex;
  justify-content: space-between;
  z-index: 999;
}
.header__logo{
  display: flex;
  align-items: center;
}
.header__logo img{
  width: 24px;
  height: 24px;
}
.header__logo h1{
  font-size: 15px;
  margin: 0 0 0 8px;
  color: #262626;
}
.sider__bar.ant-layout-sider-light, .sider__bar .ant-menu{
  height: calc(100vh - 48px);
  background-color: #f8f8f8;
}

.main__bar{
  background-color: #fff;
  padding: 30px;
}
.content-header{
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 15px;
}
.content-header h1{
  font-size: 24px;
  margin: 0;
}
</style>
