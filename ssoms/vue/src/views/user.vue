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
    width="720px"
    v-model:visible="formState.visible"
    :title="formState.type == 'add' ? '添加用户' : '修改用户'"
    :loading="formState.loading"
    @cancel="onCancel"
    @ok="onSubmit">
    <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
      <a-row :gutter="30">
        <a-col :span="5">
          <a-form-item label="头像">
            <a-upload
              v-model:file-list="uploadState.fileList"
              name="avatar"
              list-type="picture-card"
              class="avatar-uploader"
              :show-upload-list="false"
              action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
              :before-upload="beforeAvatarUpload"
              @change="onAvatarChange"
            >
              <img v-if="uploadState.imageUrl" :src="uploadState.imageUrl" alt="avatar" />
              <div v-else>
                <loading-outlined v-if="uploadState.loading"></loading-outlined>
                <plus-outlined v-else></plus-outlined>
                <div class="ant-upload-text">上传</div>
              </div>
            </a-upload>
            <p style="font-size: 12px; margin-top: 20px; color: #bcbcbc">提示: 上传png、jpg格式图片, 最大不能超过2M</p>
          </a-form-item>
        </a-col>
        <a-col :span="19">
          <a-row :gutter="24">
          <a-col :span="12">
            <a-form-item label="姓名" name="form.name">
              <a-input v-model:value="formState.form.name" placeholder="姓名" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="性别" name="form.gender">
              <a-radio-group v-model:value="formState.form.gender">
                <a-radio :value="1">男</a-radio>
                <a-radio :value="2">女</a-radio>
                <a-radio :value="3">未知</a-radio>
              </a-radio-group>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="手机号" name="form.mobile">
              <a-input v-model:value="formState.form.mobile" placeholder="13位手机号" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="密码" name="form.password">
              <a-input v-model:value="formState.form.password" placeholder="6~8位密码" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="生日" name="form.birth">
              <a-date-picker style="width: 100%" valueFormat="YYYY-MM-DD" v-model:value="formState.form.birth" placeholder="出生日期"/>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="状态" name="form.status">
              <a-switch v-model:checked="formState.form.status" :checked-value="1" :unChecked-value="0" checked-children="启用" un-checked-children="停用" />
            </a-form-item>
          </a-col>
        </a-row>
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
import type { FormInstance, UploadChangeParam, UploadProps } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { UserOutlined, PlusOutlined, LoadingOutlined } from '@ant-design/icons-vue'
import * as components from "../api/ssomsComponents"

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

interface State {
  loading: boolean
  params: components.UserListReqParams
}


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive<State>({
  loading: false, // 列表是否加载完成
  params: {
    name: '',
    mobile: '',
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
  userList(state.params).then((data) => {
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
  // resetFileds初始化无效，手动初始化
  formState.form = {
    uuid: '',
    name: '',
    mobile: '',
    password: '',
    avatar: '',
    gender: 1,
    birth: '',
    status: 1,
  }
  formState.type = 'add'
  formState.visible = true
}


/**
 * 获取用户详情, 用于编辑
 * @param uuid
 */
const initEdit = (uuid: string) => {
  userDetail({ uuid }).then((data) => {
    formState.form = data
    formState.type = 'edit'
    formState.visible = true
  })
}

/**
 * 新增或修改用户
 */
const onSubmit = () =>{
  modalFormRef.value?.validate().then(() => {
    formState.loading = true
    if (formState.type === 'add') {
      addUser(formState.form).then(() => {
        message.success('新增用户成功')
        formState.visible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.loading = false
      })
    } else {
      updateUser({ uuid: formState.form.uuid! }, formState.form).then(() => {
        message.success('修改用户成功')
        formState.visible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.loading = false
      })
    }
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

const uploadState = reactive({
  fileList: [],
  imageUrl: '',
  loading: false
})

const beforeAvatarUpload = (files: UploadProps['fileList']) => {
  if (!files) {
    return
  }
  for(let i = 0; i < files?.length; i ++) {
    const file = files[i]
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
    if (!isJpgOrPng) {
      message.error('You can only upload JPG file!')
    }
    const isLt2M = file.size! / 1024 / 1024 < 2
    if (!isLt2M) {
      message.error('Image must smaller than 2MB!')
    }
    return isJpgOrPng && isLt2M
  }
}

const onAvatarChange = (info: UploadChangeParam) => {
  // if (info.file.status === 'uploading') {
  //   uploadState.loading = true
  //   return
  // }
  // if (info.file.status === 'done') {
  //   // Get this url from response in real world.
  //   getBase64(info.file.originFileObj, (base64Url: string) => {
  //     imageUrl.value = base64Url
  //     uploadState.loading = false
  //   });
  // }
  // if (info.file.status === 'error') {
  //   uploadState.loading = false
  //   message.error('upload error')
  // }
}
</script>

<style scoped>
.avatar-uploader:deep() > .ant-upload {
  width: 120px;
  height: 120px;
}
</style>
