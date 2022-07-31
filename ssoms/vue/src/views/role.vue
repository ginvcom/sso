<template>
  <div class="content-header is-sticky">
    <div>
      <h1>角色</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary" @click="initAdd">添加角色</a-button>
    </div>
  </div>
  <a-table
  :loading="state.loading"
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
      <template v-if="column.dataIndex === 'roleName'">
        <team-outlined />
        <span style="margin-left: 10px">{{record.roleName}}</span>
      </template>
      <template v-if="column.dataIndex === 'summary'">
        <div>{{record.summary}}</div>
      </template>
      <template v-if="column.dataIndex === 'actions'">
        <a @click="initEdit(record.roleUUID)">编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要删除该角色吗?"
          placement="topRight"
          @confirm="onDelete(record.roleUUID)">
          <a>删除</a>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
  <a-modal
    v-model:visible="formState.visible"
    :title="formState.type == 'add' ? '添加角色' : '修改角色'"
    :maskClosable="false"
    @cancel="onCancel"
    @ok="onSubmit">
    <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
      <a-row :gutter="24">
      <a-col :span="24">
        <a-form-item label="角色名" name="roleName" :rules="[{ required: true, message: '请输入角色名称' }]">
          <a-input v-model:value="formState.form.roleName" placeholder="角色名称" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-form-item label="概述" name="summary" :rules="[{ required: true, message: '请简单描述一下角色' }]">
          <a-textarea
            v-model:value="formState.form.summary"
            placeholder="简单描述一下角色"
            :auto-size="{ minRows: 2, maxRows: 5 }"
          />
        </a-form-item>
      </a-col>
    </a-row>
    </a-form>
  </a-modal>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import {
  roleList,
  roleDetail,
  addRole,
  updateRole,
  deleteRole,
  RoleListReply,
  RoleForm,
RoleListReqParams
} from '../api/ssoms'
import type { FormInstance } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { TeamOutlined } from '@ant-design/icons-vue'

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
const columns = [
  {
    title: '角色',
    dataIndex: 'roleName'
  },
  {
    title: '概述',
    dataIndex: 'summary'
  },
  {
    title: '操作',
    dataIndex: 'actions',
    align: 'center',
    width: '120px'
  }
]

interface State {
  loading: boolean
  params: RoleListReqParams
}


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive({
  loading: false, // 列表是否加载完成
  params: {
    roleName: '',
    page: 1,
    pageSize: 10
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<RoleListReply>({
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
  roleList(state.params).then((data: RoleListReply) => {
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
 * 删除角色
 */
const onDelete = (roleUUID: string) => {
  deleteRole({ roleUUID }).then(() => {
    message.success('删除角色操作成功')
    getList()
  })
}

const modalFormRef = ref<FormInstance>()

interface Form {
  type: string
  visible: boolean,
  loading: boolean,
  form: RoleForm
}

const formState = reactive<Form>({
  visible: false,
  loading: false,
  type: 'add',
  form: {
    roleUUID: '',
    roleName: '',
    summary: ''
  }
})

const initAdd = () => {
  formState.form = {
    roleUUID: '',
    roleName: '',
    summary: ''
  }
  formState.type = 'add'
  formState.visible = true
}


/**
 * 获取角色详情, 用于编辑
 * @param roleUUID
 */
const initEdit = (roleUUID: string) => {
  roleDetail({ roleUUID }).then((data: RoleForm) => {
    formState.type = 'edit'
    formState.form = data
    formState.visible = true
  })
}

/**
 * 新增或修改角色
 */
const onSubmit = () =>{
  modalFormRef.value?.validate().then(() => {
    formState.loading = true
    if (formState.type === 'add') {
      addRole(formState.form).then(() => {
        message.success('新增角色成功')
        formState.visible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.loading = false
      })
    } else {
      updateRole({ roleUUID: formState.form.roleUUID! }, formState.form).then(() => {
        message.success('修改角色成功')
        formState.visible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.loading = false
      })
    }
  })
}

/**
 * 关闭modelForm，重置表单
 */
const onCancel = () => {
  modalFormRef.value?.resetFields()
}
</script>

<style scoped>
</style>
