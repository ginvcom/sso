<template>
  <a-modal
    width="710px"
    v-model:visible="state.visible"
    title="分类管理"
    :loading="state.loading"
    :maskClosable="false"
    :footer="false">
    <a-button class="editable-add-btn" style="margin-bottom: 15px" @click="handleAdd">添加分类</a-button>
    <a-table bordered :data-source="dataSource" :columns="columns" :pagination="false">
    <template #bodyCell="{ column, text, record }">
      <template v-if="column.dataIndex === 'name'">
        <div class="editable-cell">
          <div v-if="editableData[record.key]" class="editable-cell-input-wrapper">
            <a-input v-model:value="editableData[record.key].name" @pressEnter="save(record.key)" />
            <check-outlined class="editable-cell-icon-check" @click="save(record.key)" />
          </div>
          <div v-else class="editable-cell-text-wrapper">
            {{ text || ' ' }}
            <edit-outlined class="editable-cell-icon" @click="edit(record.key)" />
          </div>
        </div>
      </template>
      <template v-else-if="column.dataIndex === 'operation'">
        <a-popconfirm
          v-if="dataSource.length"
          title="确定要删除该分类吗?"
          placement="topRight"
          @confirm="onDelete(record.key)"
        >
          <a>删除</a>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
  </a-modal>
</template>
<script setup lang="ts">
import { computed, reactive, ref } from 'vue'
import type { Ref, UnwrapRef } from 'vue'
import { CheckOutlined, EditOutlined } from '@ant-design/icons-vue'

interface State {
  visible: boolean
  loading: boolean
}

const state = reactive<State>({
  visible: false,
  loading: false
})

interface DataItem {
  key: string
  name: string
  age: number
  address: string
}

const columns = [
  {
    title: 'id',
    dataIndex: 'age',
    width: '100px'
  },
  {
    title: '分类名',
    dataIndex: 'name',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    align: 'center',
    width: '80px'
  }
]

const dataSource: Ref<DataItem[]> = ref([
  {
    key: '0',
    name: 'Edward King 0',
    age: 32,
    address: 'London, Park Lane no. 0',
  },
  {
    key: '1',
    name: 'Edward King 1',
    age: 32,
    address: 'London, Park Lane no. 1',
  },
])

const count = computed(() => dataSource.value.length + 1)
    
const editableData: UnwrapRef<Record<string, DataItem>> = reactive({})

const edit = (key: string) => {
  const item = dataSource.value.filter(item => key === item.key)
  editableData[key] = { ...item[0] }
}

const save = (key: string) => {
  Object.assign(dataSource.value.filter(item => key === item.key)[0], editableData[key])
  delete editableData[key]
}

const onDelete = (key: string) => {
  dataSource.value = dataSource.value.filter(item => item.key !== key)
}

const handleAdd = () => {
  const newData = {
    key: `${count.value}`,
    name: `Edward King ${count.value}`,
    age: 32,
    address: `London, Park Lane no. ${count.value}`,
  }
  dataSource.value.push(newData)
}

defineExpose({
  state
})
</script>
<style scoped>
.editable-cell {
  position: relative;
}
.editable-cell-input-wrapper,
.editable-cell-text-wrapper {
  padding-right: 24px;
}

.editable-cell-text-wrapper {
  padding: 5px 24px 5px 5px;
}

.editable-cell-icon,
.editable-cell-icon-check {
  position: absolute;
  right: 0;
  width: 20px;
  cursor: pointer;
}

.editable-cell-icon {
  margin-top: 4px;
  display: none;
}

.editable-cell-icon-check {
  line-height: 28px;
}

.editable-cell-icon:hover,
.editable-cell-icon-check:hover {
  color: #108ee9;
}

.editable-add-btn {
  margin-bottom: 8px;
}
.editable-cell:hover .editable-cell-icon {
  display: inline-block;
}
</style>
