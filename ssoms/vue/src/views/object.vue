<template>
  <a-tabs
   v-model:activeKey="activeKey"
   :tabBarStyle="{
      position: 'fixed',
      zIndex: '100',
      width: '100%',
      margin: '0 -30px',
      padding: '0 30px',
      background: '#fff'
    }">
    <a-tab-pane key="system" tab="后台系统" style="margin-top:123px">
      <system-tab />
    </a-tab-pane>
    <a-tab-pane key="others" tab="菜单 & 操作" style="margin-top:123px">
      <menu-tab />
    </a-tab-pane>
  </a-tabs>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import SystemTab from '@/components/SystemTab.vue'
import MenuTab from '@/components/MenuTab.vue'
const activeKey = ref('system')
</script>
