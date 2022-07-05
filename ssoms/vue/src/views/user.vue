<template>
  <div class="content-header">
    <div>
      <h1>用户</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary" @click="initAdd">添加用户</a-button>
    </div>
  </div>
  <a-table
  :dataSource="respState.list"
  :columns="columns"
  :pagination="{
    total: respState.total,
    current: state.params.page,
    pageSize: state.params.pageSize,
    showSizeChanger: true, showTotal: (total) => `共 ${total} 条`
  }"
  @change="onTableChange">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'name'">
        <a-avatar style="color: #f56a00; background-color: #fde3cf">
          <template #icon><UserOutlined /></template>
        </a-avatar>
        <span style="margin-left: 10px">{{record.name}}</span>
      </template>
      <template v-if="column.key === 'gender'">
        <span v-if="record.gender == 1">男</span>
        <span v-else-if="record.gender == 2">女</span>
        <span v-else-if="record.gender == 3">未知</span>
      </template>
      <template v-if="column.key === 'actions'">
        <a @click="initEdit(record.uuid)">编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要删除该用户吗?"
          ok-text="Yes"
          cancel-text="No"
          @confirm="onDelete(record.uuid)">
          <a>删除</a>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
  <a-modal
    v-model:visible="formState.visible"
    :title="formState.type == 'add' ? '新增用户' : '修改用户'"
    @ok="onSubmit">
    <a-form layout="vertical" :model="formState.form">
      <a-row :gutter="24">
      <a-col :span="12">
        <a-form-item label="姓名">
          <a-input v-model:value="formState.form.name" placeholder="姓名" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="性别">
          <a-radio-group v-model:value="formState.form.gender">
            <a-radio :value="1">男</a-radio>
            <a-radio :value="2">女</a-radio>
            <a-radio :value="3">未知</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="手机号">
          <a-input v-model:value="formState.form.mobile" placeholder="13位手机号" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="密码">
          <a-input v-model:value="formState.form.password" placeholder="6~8位密码" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="生日">
          <a-date-picker style="width: 100%" v-model:value="formState.form.birth" placeholder="出生日期"/>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item label="状态">
          <a-switch v-model:checked="formState.form.status" :checked-value="1" :unChecked-value="0" checked-children="启用" un-checked-children="停用" />
        </a-form-item>
      </a-col>
    </a-row>
    </a-form>
  </a-modal>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import {
  userList,
  userDetail,
  addUser,
  updateUser,
  deleteUser,
  UserListReply,
  UserForm
} from '../api/ssoms'
import type { FormInstance } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { UserOutlined } from '@ant-design/icons-vue'

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
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
    align: 'center',
    width: '120px'
  }
]


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive({
  loading: false, // 列表是否加载完成
  params: {
    name: undefined,
    mobile: undefined,
    page: 1,
    pageSize: 10
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<UserListReply>({
  total: 0,
  list: []
})

onMounted(() => {
  getList()
})

/**
 * 获取列表
 */
const getList = () => {
  state.loading = true
  userList(state.params).then((data: UserListReply) => {
    respState.total = data.total
    respState.list = data.list
  }).finally(() => {
    state.loading = false
  })
}

const onTableChange = ({ current, pageSize }) => {
  state.params.page = current
  state.params.pageSize = pageSize
  getList()
} 

/**
 * 删除用户
 */
const onDelete = (uuid: string) => {
  deleteUser({ uuid }).then(() => {
    message.success('删除用户操作成功')
    getList()
  })
}

const modalFormRef = ref<FormInstance>()

interface Form {
  type: string
  visible: boolean,
  loading: boolean,
  form: UserForm
}

const formState = reactive<Form>({
  visible: false,
  loading: false,
  type: 'add',
  form: {
    uuid: '',
    name: '',
    mobile: '',
    password: '',
    avatar: '',
    gender: 1,
    birth: '',
    status: 1,
  }
})

const initAdd = () => {
  modalFormRef.value?.resetFields()
  formState.visible = true
}


/**
 * 获取用户详情, 用于编辑
 * @param uuid
 */
const initEdit = (uuid: string) => {
  userDetail({ uuid }).then((data: UserForm) => {
    formState.form = data
    formState.visible = true
  })
}

/**
 * 新增或修改用户
 */
const onSubmit = () =>{
  modalFormRef.value?.validateFields().then(() => {
    formState.loading = true
    if (formState.type === 'add') {
      addUser(formState.form).then(() => {
        message.success('新增用户成功')
      }).finally(() => {
        formState.loading = false
      })
    } else {
      updateUser(formState.form).then(() => {
        message.success('修改用户成功')
      }).finally(() => {
        formState.loading = false
      })
    }
  })
}
</script>

<style scoped>
</style>
