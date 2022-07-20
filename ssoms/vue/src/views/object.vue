<template>
  <a-tabs v-model:activeKey="activeKey">
    <a-tab-pane key="system" tab="后台系统">
      <system-tab />
    </a-tab-pane>
    <a-tab-pane key="others" tab="菜单 & 操作">
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
