<template>
  <div class="content-header">
    <div>
      <h1>用户</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary">添加用户</a-button>
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
    title: '用户',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: '性别',
    dataIndex: 'gender',
    key: 'gender'
  },
  {
    title: '手机号',
    dataIndex: 'mobile',
    key: 'mobile'
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
    name: '胡彦斌',
    gender: 1,
    address: '西湖区湖底公园1号'
  },
  {
    key: '2',
    name: '胡彦祖',
    gender: 2,
    address: '西湖区湖底公园1号'
  }
])
const count = ref(0)

const confirm = () => {
  console.log('on delete')
}
</script>

<style scoped>
</style>
