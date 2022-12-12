<template>
  <a-result
    status="error"
    title="Error Occurred"
    :sub-title="`the component file “@/views${route.path}.vue” not exits.`"
  >
    <template #extra>
      <a-button type="primary" @click="onBackToHome">返回首页</a-button>
    </template>
  </a-result>
</template>
<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const onBackToHome = () => {
  router.push('/')
}
</script>
