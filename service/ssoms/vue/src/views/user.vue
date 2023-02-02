<template>
  <layout>
    <div class="content-header is-sticky">
      <div>
        <h1>用户管理</h1>
      </div>
      <div class="content-header__actions">
        <a-button type="primary" @click="initAdd"><template #icon><user-add-outlined /></template>添加用户</a-button>
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
          <a-avatar :src="ossConfig.doc.domain + record.avatar" style="color: #f56a00; background-color: #fde3cf">
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
          <span v-if="record.status == 1"><a-badge status="success" /> 启用</span>
          <span v-else-if="record.status == 0"><a-badge status="error" />停用</span>
        </template>
        <template v-if="column.dataIndex === 'roles'">
          <span v-for="(role,i) in record.roles" :key="role.value">
            <template v-if="i > 0">、</template>
            <router-link :to="{ path: '/role/assignedUsers', query: { roleUUID: role.value } }">{{role.label}}</router-link>
          </span>
        </template>
        <template v-if="column.dataIndex === 'actions'">
          <a @click="initAssignRole(record.uuid)">分配角色</a>
          <a-divider type="vertical" />
          <a @click="initEdit(record.uuid)">编辑</a>
          <a-divider type="vertical" />
          <a-popconfirm
            title="确定要删除该用户吗?"
            placement="topRight"
            @confirm="onDelete(record.uuid)">
            <a>删除</a>
          </a-popconfirm>
        </template>
      </template>
    </a-table>
    <a-modal
      width="694px"
      v-model:visible="formState.visible"
      :title="formState.type == 'add' ? '添加用户' : '修改用户'"
      :loading="formState.loading"
      :maskClosable="false"
      @cancel="onCancel"
      @ok="onSubmit">
      <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
        <a-row :gutter="24">
          <a-col :span="8">
            <a-form-item label="头像">
              <template v-if="cropperState.visible">
                <vue-cropper
                  ref="cropper"
                  :aspect-ratio="1 / 1"
                  :src="cropperState.imageUrl"
                  @keyup.enter="onCrop"
                />

              <a-button block style="margin-top: 16px;width: 200px;" @click="onCrop">完成裁剪并上传</a-button>
              </template>
              <a-upload  v-else
                v-model:file-list="uploadState.fileList"
                list-type="picture-card"
                class="avatar-uploader"
                :show-upload-list="false"
                :before-upload="beforeAvatarUpload"
                @change="onAvatarChange"
              >
                <!-- <img class="avatar-uploader__img" v-if="uploadState.imageUrl" :src="uploadState.imageUrl" alt="avatar" /> -->
                <a-avatar
                  v-if="uploadState.fileList && uploadState.fileList.length > 0"
                  class="avatar-uploader__img" :size="200" shape="square"
                  :src="ossConfig.doc.domain + uploadState.fileList[0].name"
                  style="color: #f56a00; background-color: #fde3cf"
                >
                  <template #icon><UserOutlined /></template>
                </a-avatar>
                <div>
                  <loading-outlined v-if="uploadState.fileList && uploadState.fileList.length > 0 && uploadState.fileList[0].status === 'uploading'"></loading-outlined>
                  <plus-outlined v-else></plus-outlined>
                  <div class="ant-upload-text">上传</div>
                </div>
              </a-upload>
            </a-form-item>
          </a-col>
          <a-col :span="16">
            <a-row :gutter="24">
            <a-col :span="12">
              <a-form-item label="姓名" name="name" :rules="[{ required: true, message: '请输入用户姓名' }]">
                <a-input v-model:value="formState.form.name" placeholder="姓名" />
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="性别" name="gender" :rules="[{ required: true, message: '请选择性别' }]">
                <a-radio-group v-model:value="formState.form.gender">
                  <a-radio :value="1">男</a-radio>
                  <a-radio :value="2">女</a-radio>
                  <a-radio :value="3">未知</a-radio>
                </a-radio-group>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="手机号" name="mobile" :rules="[{ required: true, message: '请输入用户手机号' }]">
                <a-input v-model:value="formState.form.mobile" placeholder="13位手机号" />
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="密码" name="password">
                <a-input v-model:value="formState.form.password" placeholder="6~8位密码" />
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="生日" name="birth" :rules="[{ required: true, message: '请输入用户出生日期' }]">
                <a-date-picker style="width: 100%" valueFormat="YYYY-MM-DD" v-model:value="formState.form.birth" placeholder="出生日期"/>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="状态" name="status" :rules="[{ required: true, message: '请选择用户状态' }]">
                <a-switch v-model:checked="formState.form.status" :checked-value="1" :unChecked-value="0" checked-children="启用" un-checked-children="停用" />
              </a-form-item>
            </a-col>
          </a-row>
          </a-col>
        </a-row>
        <a-form-item label="用户简介" name="summary">
          <a-textarea v-model:value="formState.form.introduction" placeholder="填写职业技能、擅长的事情、喜欢的事情等" :rows="3" show-count :maxlength="100" />
        </a-form-item>
      </a-form>
    </a-modal>
    <a-modal
      width="800px"
      v-model:visible="assignFormState.visible"
      :title="`用户“${assignFormState.name}”角色分配`"
      :loading="assignFormState.loading"
      :maskClosable="false"
      @cancel="onCancel"
      @ok="onAssignRole">
      <a-form layout="vertical" ref="userAssignFormmRef" :model="formState.form">
        <h3>已分配角色</h3>
        <a-row :gutter="24">
          <a-col v-for="role in assignFormState.assigned" :key="role.value" :span="8">
            <div class="user-assign__item is-active" @click="removeAssignRole(role.value)">
              <div>
                <team-outlined />
                <span style="margin-left: 10px">{{role.label}}</span>
              </div>
              <p>{{role.extra}}</p>
            </div>
          </a-col>
          <a-col v-if="assignFormState.assigned.length == 0" :span="8" :offset="8">
            <a-empty :image="Empty.PRESENTED_IMAGE_SIMPLE" />
          </a-col>
        </a-row>
        <h3 style="margin-top: 20px">可分配角色</h3>
        <a-row :gutter="24">
          <a-col v-for="role in assignFormState.options" :key="role.value" :span="8">
            <div class="user-assign__item" @click="addAssignRole(role.value)">
              <div>
                <team-outlined />
                <span style="margin-left: 10px">{{role.label}}</span>
              </div>
              <p>{{role.extra}}</p>
            </div>
          </a-col>
          <a-col v-if="assignFormState.options.length == 0" :span="8" :offset="8">
            <a-empty :image="Empty.PRESENTED_IMAGE_SIMPLE" />
          </a-col>
        </a-row>
      </a-form>
    </a-modal>
  </layout>
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import VueCropper from 'vue-cropperjs'
import 'cropperjs/dist/cropper.css'
import {
  userList,
  userDetail,
  addUser,
  updateUser,
  deleteUser,
  UserListReqParams,
  UserListReply,
  UserForm,
  assignedRoles,
  Option,
assignRole
} from '@/api/ssoms'
import { sts } from '@/api/oss'
import type { FormInstance, UploadChangeParam, UploadProps, UploadFile } from 'ant-design-vue'
import { message, Empty } from 'ant-design-vue'
import { UserOutlined, PlusOutlined, LoadingOutlined, UserAddOutlined, TeamOutlined } from '@ant-design/icons-vue'
import AliyunOSS, { StsInfo } from '@/utils/aliyunOSS'
import { getFileName, getBase64 } from '@/utils/file'
import { ossConfig } from '@/config'

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
    title: '分配的角色',
    dataIndex: 'roles'
  },
  {
    title: '操作',
    dataIndex: 'actions',
    align: 'center',
    width: '180px'
  }
]

interface State {
  aliOss?: any
  loading: boolean
  params: UserListReqParams
}

const simpleImage = ref(Empty.PRESENTED_IMAGE_SIMPLE)


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

const refreshSTSToken = async (bucket: string) => {
  const res = await sts({ bucket })
  return Promise.resolve(res as StsInfo)
}

const uploadState = reactive<UploadProps>({ fileList: [] })

onMounted(() => {
  state.aliOss = new AliyunOSS(ossConfig.doc.bucket, () => refreshSTSToken(ossConfig.doc.bucket), uploadState)
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
  deleteUser({}, uuid).then(() => {
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
    introduction: ''
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
    introduction: ''
  }
  formState.type = 'add'
  formState.visible = true
}


/**
 * 获取用户详情, 用于编辑
 * @param uuid
 */
const initEdit = (uuid: string) => {
  userDetail({}, uuid).then((data) => {
    formState.form = data
    uploadState.fileList = [{ uid: '', name: data.avatar, url: ossConfig.doc.domain + data.avatar }]
    formState.type = 'edit'
    formState.visible = true
  })
}

/**
 * 新增或修改用户
 */
const onSubmit = () => {
  if (uploadState.fileList && uploadState.fileList.length > 0) {
    formState.form.avatar = uploadState.fileList[0].name
  }
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
      updateUser({}, formState.form, formState.form.uuid!).then(() => {
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

const cropper = ref()

const beforeAvatarUpload = (file: UploadFile) => {
  uploadState.fileList = []
  if (!file) {
    return true
  }
  const isPic = file.type === 'image/jpeg' || file.type === 'image/png'
  if (!isPic) {
    message.error('头像仅支持png、jpg格式的图片!')
    return true
  }
  const isLt2M = file.size! / 1024 / 1024 < 10
  if (!isLt2M) {
    message.error('图片最大不能超过10MB!')
    return true
  }
  if (isPic && isLt2M) {
    return false
  }
  return false
}

const onAvatarChange = (info: UploadChangeParam) => {
  // getBase64((info.file as any), (base64Url: string) => {
  //   doUpload(info.file, base64Url)
  // })
  onUpload(info.file)
}

const cropperState = reactive({
  visible: false, // 1 是否开启裁剪
  imageUrl: '',
  uid: ''
})

const onCrop = () => {
  const canvas = cropper.value.getCroppedCanvas()
  // const cropImg = canvas.toDataURL()
  // uploadState.imageUrl = cropImg
  cropperState.visible = false
  canvas.toBlob((blob: Blob) => {
    let key = Math.floor(Math.random() * 10).toFixed()
    key += Math.floor(new Date().getTime() / 1000).toFixed()
    for (let i = 0; i < 5; i++) {
      key += Math.floor(Math.random() * 10).toFixed()
    }
    const filename = key + '.png'
    const file = new File([blob], filename, { type: "image/png", lastModified: Date.now() })
    file['uid'] = cropperState.uid
    state.aliOss.sendRequest(file as unknown as UploadFile)
  })
}

const onUpload = async (file: UploadFile) => {
  const reader = new FileReader()
  reader.addEventListener('load', () => {
    cropperState.visible = true
    cropperState.imageUrl = reader.result as string
    cropperState.uid = file.uid
    return false
  })
  reader.readAsDataURL(file as unknown as File)
}

interface AssignFormState {
  visible: boolean
  loading: boolean
  name: string
  uuid: string
  assigned: Array<Option>
	options: Array<Option>
}

const assignFormState = reactive<AssignFormState>({
  visible: false,
  loading: false,
  name: '',
  uuid: '',
  assigned: [],
  options: []
})

// 用户分配角色弹窗
const initAssignRole = (uuid: string) => {
  assignedRoles({ uuid }).then((data) => {
    assignFormState.assigned = data.assigned
    assignFormState.options = data.options
    assignFormState.uuid = data.uuid
    assignFormState.name = data.name
    assignFormState.visible = true
  })
}

const onAssignRole = () => {
  assignFormState.loading = true
  const roleUUIDArray = assignFormState.assigned.map(item => item.value)
  assignRole({ uuid: assignFormState.uuid, roleUUIDArray }).then(() => {
    message.success('用户分配角色修改成功')
    assignFormState.visible = false
    getList()
  }).finally(() => {
    assignFormState.loading = false
  })
}

const addAssignRole = (roleUUID: string) => {
  const index = assignFormState.options.findIndex(item => item.value === roleUUID)
  if (index > -1) {
    const role = assignFormState.options[index]
    assignFormState.assigned.push(role)
    assignFormState.options.splice(index, 1)
  }
}

const removeAssignRole = (roleUUID: string) => {
  const index = assignFormState.assigned.findIndex(item => item.value === roleUUID)
  if (index > -1) {
    const role = assignFormState.assigned[index]
    assignFormState.options.unshift(role)
    assignFormState.assigned.splice(index, 1)
  }
}
</script>
