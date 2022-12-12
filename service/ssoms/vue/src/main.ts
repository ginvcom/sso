import { createApp } from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue'

import Layout from '@/layout/Layout.vue'

import 'ant-design-vue/dist/antd.css'

import './assets/index.css'

import router from './router' 

createApp(App).use(Antd).component('Layout', Layout).use(router).mount('#app')
