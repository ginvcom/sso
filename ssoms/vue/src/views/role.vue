<template>
  <div class="content-header">
    <div>
      <a-breadcrumb>
        <a-breadcrumb-item>首页</a-breadcrumb-item>
        <a-breadcrumb-item>角色</a-breadcrumb-item>
      </a-breadcrumb>
      <h1>角色</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary">添加角色</a-button>
    </div>
  </div>
  <a-table :dataSource="dataSource" :columns="columns">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'actions'">
        <a>编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="Are you sure delete this task?"
          ok-text="Yes"
          cancel-text="No"
          @confirm="confirm">
          <a>删除</a>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
</template>
<script setup lang="ts">
import { reactive, ref } from 'vue'
const columns = [
  {
    title: '角色',
    dataIndex: 'roleName',
    key: 'roleName'
  },
  {
    title: '概述',
    dataIndex: 'summary',
    key: 'summary'
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status'
  },
  {
    title: '操作',
    key: 'actions',
    width: '120px'
  }
]

const dataSource = reactive( [
  {
    key: '1',
    roleName: '胡彦斌',
    status: 1,
    summary: '西湖区湖底公园1号'
  },
  {
    key: '2',
    roleName: '胡彦祖',
    status: 2,
    summary: '西湖区湖底公园1号'
  }
])
const count = ref(0)

const confirm = () => {
  console.log('on delete')
}
</script>

<style scoped>
</style>
