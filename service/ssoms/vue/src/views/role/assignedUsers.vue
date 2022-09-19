<template>
  <div class="content-header is-sticky">
    <div>
      <h1>角色“{{respState.roleName}}”分配用户列表</h1>
    </div>
    <div class="content-header__actions">
      <a-button  @click="onBack"><template #icon><arrow-left-outlined /></template>返回列表</a-button>
      <a-button type="primary" @click="initAssignUser"><template #icon><plus-outlined /></template>分配用户</a-button>
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
      <template v-if="column.dataIndex === 'name'">
        <a-avatar :src="ossConfig.ginvdoc.domain + record.avatar" style="color: #f56a00; background-color: #fde3cf">
          <template #icon><UserOutlined /></template>
        </a-avatar>
        <span style="margin-left: 10px">{{record.name}}</span>
      </template>
      <template v-if="column.dataIndex === 'gender'">
        <span v-if="record.gender == 1">男</span>
        <span v-else-if="record.gender == 2">女</span>
        <span v-else-if="record.gender == 3">未知</span>
      </template>
      <template v-if="column.dataIndex === 'status'">
        <span v-if="record.isDelete == 1" class="text-red"><a-badge status="success" /> 已删除</span>
        <template v-else>
          <span v-if="record.status == 1"><a-badge status="success" /> 启用</span>
          <span v-else-if="record.status == 0"><a-badge status="error" />停用</span>
        </template>
      </template>
      <template v-if="column.dataIndex === 'actions'">
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要将该用户移出当前角色吗?"
          placement="topRight"
          @confirm="onDeassignUser(record.uuid)">
          <a>移出角色</a>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
  <a-modal
    v-model:visible="formState.visible"
    :title="`选择用户分配到角色“${respState.roleName}”`"
    :maskClosable="false"
    @cancel="onCancel"
    @ok="onSubmit">
    <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
      <a-form-item name="userUUID" :rules="[{ required: true, message: '请选择用户' }]">
        <a-select
          v-model:value="formState.form.userUUID"
          show-search
          placeholder="输入用户姓名搜索用户"
          optionLabelProp="label"
          :default-active-first-option="false"
          :show-arrow="false"
          :filter-option="false"
          @search="onUserOptionSearch"
        >
          <a-select-option v-for="option in userOptionState.options" :key="option.value" :value="option.value" :label="option.label">
            <div>
              <a-avatar :src="ossConfig.ginvdoc.domain + option.extra" style="color: #f56a00; background-color: #fde3cf">
                <template #icon><UserOutlined /></template>
              </a-avatar>
              <span style="margin-left: 10px">{{option.label}}</span>
            </div>
          </a-select-option>
          <template v-if="userOptionState.loading" #notFoundContent>
            <div style="height: 78px;display:flex;align-items: center;justify-content: center;">
              <a-spin size="small" />
            </div>
          </template>
        </a-select>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref, watch } from 'vue'
import {
  assignedUsers,
  AssignedUsersReqParams,
  AssignedUsersReply,
  roleDetail,
  addRole,
  updateRole,
  Option,
  AssignUserReq,
RoleListReqParams,
userFilterOptions,
assignUser,
deassignUser
} from '@/api/ssoms'
import type { FormInstance } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { ArrowLeftOutlined, UserOutlined, PlusOutlined } from '@ant-design/icons-vue'
import { ossConfig } from '@/config'
import { useRoute, useRouter } from 'vue-router'
import Ajax from '@/utils/ajax'

const route = useRoute()

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
const columns = [
  {
    title: '用户',
    dataIndex: 'name'
  },
  {
    title: '性别',
    dataIndex: 'gender'
  },
  {
    title: '手机号',
    dataIndex: 'mobile'
  },
  {
    title: '状态',
    dataIndex: 'status'
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
  params: AssignedUsersReqParams
}


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive<State>({
  loading: false, // 列表是否加载完成
  params: {
    roleUUID: route.query.roleUUID as string,
    page: 1,
    pageSize: 10
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<AssignedUsersReply>({
  total: 0,
  list: [],
  roleName: ''
})

onMounted(() => {
  getList()
})

watch(
  () => route.query.roleUUID,
  (roleUUID) => {
    if (route.path == '/role/assignedUsers' && roleUUID) {
      state.params = { roleUUID: roleUUID as string, page: 1, pageSize: 10 }
      getList()
    }
  }
)

/**
 * 获取列表
 */
const getList = () => {
  state.loading = true
  assignedUsers(state.params).then(data => {
    respState.total = data.total
    respState.list = data.list
    respState.roleName = data.roleName
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
 * 将用户移出角色
 */
const onDeassignUser = (userUUID: string) => {
  deassignUser({ roleUUID: state.params.roleUUID }, { userUUID }).then(() => {
    message.success('用户移出角色操作成功')
    getList()
  })
}

const modalFormRef = ref<FormInstance>()

interface Form {
  visible: boolean,
  loading: boolean,
  form: AssignUserReq
}

const formState = reactive<Form>({
  visible: false,
  loading: false,
  form: {
    userUUID: ''
  }
})

interface UserOptionState {
  loading: boolean
  options: Array<Option>
}

const userOptionState = reactive<UserOptionState>({
  loading: false,
  options: []
})

const onUserOptionSearch = (name: string) => {
  userOptionState.loading = true
  userFilterOptions({ name }).then(data => {
    userOptionState.options = data.options
  }).finally(() => {
    userOptionState.loading = false
  })
}

const initAssignUser = () => {
  formState.form = {
    userUUID: ''
  }
  formState.visible = true
}



/**
 * 选择用户分配到角色
 */
const onSubmit = () =>{
  modalFormRef.value?.validate().then(() => {
    formState.loading = true
    assignUser({ roleUUID: state.params.roleUUID }, formState.form).then(() => {
      message.success('选择用户分配到角色操作成功')
      formState.visible = false
      modalFormRef.value?.resetFields()
      getList()
    }).finally(() => {
      formState.loading = false
    })
  })
}

/**
 * 关闭modelForm，重置表单
 */
const onCancel = () => {
  modalFormRef.value?.resetFields()
}

const router = useRouter()

const onBack = () => {
  router.push({ path: '/role', replace: true })
}
</script>

<style scoped>
</style>
