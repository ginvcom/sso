<template>
  <div class="content-header is-sticky">
    <div>
      <h1>授权管理</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary"><save-outlined />修改授权</a-button>
    </div>
  </div>
  <a-row type="flex" :gutter="32">
  <a-col flex="260px" class="permission__left">
    <div v-for="obj in state.systems" :key="obj.uuid" class="object__item">
      <div
        class="object__current-value object__option"
        :class="{ 'is-active': state.params.topKey == obj.key }"
        @click="onChangeSystem(obj)"
      >
        <a-image
          :width="48"
          :height="48"
          :src="ossConfig.ginvdoc.domain + obj.icon"
          :preview="false"
          fallback="data:image/png;base64,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"
        />
        <p class="object__current-system-name">{{obj.objectName}}</p>
      </div>
    </div>
  </a-col>
  <a-col flex="1">
  <a-checkbox-group  v-model:value="formState.ActionUUIDArray" style="width: 100%">
  <a-table
  :loading="state.loading"
  class="apis-check__table"
  bordered
  rowKey="value"
  size="small"
  :row-selection="{ checkStrictly: false, selectedRowKeys: formState.MenuUUIDArray, onChange }"
  :dataSource="respState.list"
  :columns="columns"
  :pagination="false"
  @change="onTableChange">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'apis'">
        <div class="apis-check__col">
          <div class="apis-check__group">
            <a-checkbox v-for="item in record.apis" :key="item.value" :value="item.value">
              <!-- <span style="color: red">{{ item.label }}</span> -->
              <span v-if="item.subType == 1" class="method-get">{{ item.label }}</span>
              <span v-else-if="item.subType == 2" class="method-post">{{ item.label }}</span>
              <span v-else-if="item.subType == 3" class="method-put">{{ item.label }}</span>
              <span v-else-if="item.subType == 4" class="method-patch">{{ item.label }}</span>
              <span v-else-if="item.subType == 5" class="method-delete">{{ item.label }}</span>
            </a-checkbox>
          </div>
          <div class="apis-check__all" v-if="record.apis.length > 0 && !isAllApiChecked(record.apis)" @click="checkAllApis(record.apis)">全选</div>
        </div>
      </template>
    </template>
  </a-table>
  </a-checkbox-group>
  </a-col>
  </a-row>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import 'cropperjs/dist/cropper.css'
// Object 是js的关键字, 别名处理一下
import {
  roleOperations,
  objectList,
  ObjectListReply,
  ObjectOption,
  RoleOperationsReply,
  Object as Obj
} from '../api/ssoms'
import type { FormInstance } from 'ant-design-vue'
import { SaveOutlined } from '@ant-design/icons-vue'
import { ossConfig } from '@/config'

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
const columns = [
  {
    title: '菜单',
    dataIndex: 'label',
    key: 'value',
    width: '200px'
  },
  {
    title: '操作',
    key: 'apis'
  }
]

interface System {
  name: string
  icon: string
}

interface State {
  loading: boolean
  currentSystem: System
  systems: Array<Obj>
  params: any
}


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive<State>({
  loading: false, // 列表是否加载完成
  currentSystem: {
    icon: '',
    name: ''
  },
  systems: [],
  params: {
    pUUID: '',
    mobile: '',
    page: 1,
    pageSize: 10
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<RoleOperationsReply>({
  list: []
})

onMounted(() => {
  getSystemOptions()
})

/**
 * 获取列表
 */
const getList = () => {
  state.loading = true
  roleOperations(state.params).then((data) => {
    respState.list = data.list
    formState.MenuUUIDArray = ['k7vcb745vcsb', '8yzo4kzfucsb']
  }).finally(() => {
    state.loading = false
  })
}

const onTableChange = ({ current, pageSize }) => {
  state.params.page = current
  state.params.pageSize = pageSize
  getList()
}

const onChangeSystem = (obj: Obj) => {
  console.log(obj)
  if (state.params.topKey === obj.key) {
    return
  }
  state.currentSystem = { icon: obj.icon, name: obj.objectName }
  state.params = { pUUID: obj.uuid, topKey: obj.key, objectName: obj.objectName }
  // getList()
}

const modalFormRef = ref<FormInstance>()

interface Form {
  loading: boolean
  MenuUUIDArray: Array<string>
  ActionUUIDArray: Array<string>
}

const formState = reactive<Form>({
  loading: false,
  MenuUUIDArray: [],
  ActionUUIDArray:[]
})

const getSystemOptions = () => {
  objectList({ topKey: '', key: '', objectName: '' }).then((data: ObjectListReply) => {
    state.systems = data.list
    if (state.systems.length > 0) {
      const firstSystem = state.systems[0]
      state.currentSystem = { name: firstSystem.objectName, icon: firstSystem.icon }
      state.params.topKey = firstSystem.key
      getList()
    }
  }).finally(() => {
    state.loading = false
  })
}


const onChange = (selectedRowKeys: string[]) => {
  console.log(selectedRowKeys)
  formState.MenuUUIDArray = selectedRowKeys
}

const checkAllApis = (record: ObjectOption[]) => {
  console.log('checkAllApis', record)
  record.forEach(obj => {
    const index = formState.ActionUUIDArray.indexOf(obj.value)
    if (index < 0) {
      formState.ActionUUIDArray.push(obj.value)
    }
  })
  console.log('checkAllApis', formState.ActionUUIDArray)
}

const isAllApiChecked = (record: ObjectOption[]) => {
  for (const obj of record) {
    const index = formState.ActionUUIDArray.indexOf(obj.value)
    if (index < 0) {
      return false
    }
  }
  return true
}

/**
 * 获取用户详情, 用于编辑
 * @param uuid
 */
// const initEdit = (uuid: string) => {
//   userDetail({ uuid }).then((data) => {
//     formState.form = data
//   })
// }

/**
 * 新增或修改用户
 */
const onSubmit = () =>{
  modalFormRef.value?.validate().then(() => {
    formState.loading = true
    // updateUser({ uuid: formState.form.uuid! }, formState.form).then(() => {
    //   message.success('修改用户成功')
    //   modalFormRef.value?.resetFields()
    //   getList()
    // }).finally(() => {
    //   formState.loading = false
    // })
  }).catch((err) => {
    console.log(err)
  })
}

/**
 * 关闭modelForm，重置表单
 */
const onCancel = () => {
  modalFormRef.value?.resetFields()
}
</script>
