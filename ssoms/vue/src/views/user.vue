<template>
  <div class="content-header is-sticky">
    <div>
      <h1>用户</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary" @click="initAdd">添加用户</a-button>
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
      <template v-if="column.key === 'name'">
        <a-avatar :src="ossConfig.ginvdoc.domain + record.avatar" style="color: #f56a00; background-color: #fde3cf">
          <template #icon><UserOutlined /></template>
        </a-avatar>
        <span style="margin-left: 10px">{{record.name}}</span>
      </template>
      <template v-if="column.key === 'gender'">
        <span v-if="record.gender == 1">男</span>
        <span v-else-if="record.gender == 2">女</span>
        <span v-else-if="record.gender == 3">未知</span>
      </template>
      <template v-if="column.key === 'status'">
        <span v-if="record.status == 1"><a-badge status="success" /> 启用</span>
        <span v-else-if="record.status == 0"><a-badge status="error" />停用</span>
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
    width="730px"
    v-model:visible="formState.visible"
    :title="formState.type == 'add' ? '添加用户' : '修改用户'"
    :loading="formState.loading"
    :maskClosable="false"
    @cancel="onCancel"
    @ok="onSubmit">
    <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
      <a-row :gutter="40">
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
              <a-avatar v-if="uploadState.imageUrl" class="avatar-uploader__img" :size="200" shape="square" :src="uploadState.imageUrl" style="color: #f56a00; background-color: #fde3cf">
                <template #icon><UserOutlined /></template>
              </a-avatar>
              <div>
                <loading-outlined v-if="uploadState.loading"></loading-outlined>
                <plus-outlined v-else></plus-outlined>
                <div class="ant-upload-text">上传</div>
              </div>
            </a-upload>
          </a-form-item>
        </a-col>
        <a-col :span="16">
          <a-row :gutter="40">
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
            <a-form-item label="密码" name="password" :rules="[{ required: true, message: '请输入6~8位密码用户登录' }]">
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
    </a-form>
  </a-modal>
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
  UserForm
} from '../api/ssoms'
import { sts } from '../api/aliOss'
import type { FormInstance, UploadFile, UploadChangeParam, UploadProps } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { UserOutlined, PlusOutlined, LoadingOutlined } from '@ant-design/icons-vue'
import OSS from 'ali-oss'
import { getFileName } from '../utils/file'
import { ossConfig } from '@/config'

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
    title: '状态',
    dataIndex: 'status',
    key: 'status'
  },
  {
    title: '操作',
    key: 'actions',
    align: 'center',
    width: '120px'
  }
]

interface State {
  aliOss?: any
  loading: boolean
  params: UserListReqParams
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
// //ginvdoc.oss-cn-shenzhen.aliyuncs.com/1165839836329658.png
const respState = reactive<UserListReply>({
  total: 0,
  list: []
})

onMounted(() => {
  state.aliOss = new OSS({
    ...ossConfig.ginvdoc,
    refreshSTSToken: async () => {
      const info = await sts({ bucket: 'doc' })
      return info
    },
    refreshSTSTokenInterval: 300000
  })
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
  uploadState.imageName = ''
  uploadState.imageUrl = ''
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
    uploadState.imageName = data.avatar
    uploadState.imageUrl = ossConfig.ginvdoc.domain + data.avatar
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

const cropper = ref()

const uploadState = reactive({
  fileList: [],
  imageUrl: '',
  imageName: '',
  fileType: '',
  loading: false
})

const beforeAvatarUpload = (files: UploadProps) => {
  console.log('beforeAvatarUpload', files)
  if (!files) {
    return
  }
  // console.log('beforeAvatarUpload', files.length)
  // if (files.length > 1) {
  //   message.error('只能上传一张图片!')
  //   return
  // }
  const file = files as File
  uploadState.fileType = file.type
  const isPic = file.type === 'image/jpeg' || file.type === 'image/png'
  if (!isPic) {
    message.error('You can only upload JPG file!')
  }
  console.log(44)
  const isLt2M = file.size! / 1024 / 1024 < 2
  if (!isLt2M) {
    message.error('Image must smaller than 2MB!')
  }
  if (isPic && isLt2M) {
    uploadState.imageName = getFileName(file)
    getBase64((file as any), (base64Url: string) => {
      cropperState.visible = true
      cropperState.imageUrl = base64Url
      uploadState.loading = false
      return false
    })
    return false
  }
  return false
}

const ossUpload = async (options: any) => {
  return
  const { onSuccess, onError, file, data, onProgress } = options
  const name = getFileName(file)
  options.filename = name
  console.log(state.aliOss)
  if (state.aliOss) {
    try {
      const res = await state.aliOss.multipartUpload(name, file, {
        progress: (progress: number, checkpoint: any) => {
          onProgress({ percent: progress * 100 })  // 执行onProgress 并传入当前进度，使得上传组件正确显示进度条
        },
      })
      onSuccess(res)
    } catch (e) {
      onError()
    }
  }
}

const getBase64 = (img: Blob, callback: (base64Url: string) => void) => {
  const reader = new FileReader()
  reader.addEventListener('load', () => callback(reader.result as string))
  reader.readAsDataURL(img)
}

const onAvatarChange = (info: UploadChangeParam) => {
  console.log(info)
  if (info.file.status === 'uploading') {
    uploadState.loading = true
    return
  }
  if (info.file.status === 'done') {
    uploadState.imageUrl = ossConfig.ginvdoc.domain + info.file.response.name
    // Get this url from response in real world.
    // getBase64((info.file as any).originFileObj, (base64Url: string) => {
    //   uploadState.imageUrl = base64Url
    //   uploadState.loading = false
    // })
  }
  if (info.file.status === 'error') {
    uploadState.loading = false
    message.error('upload error')
  }
}

const cropperState = reactive({
  visible: false, // 1 是否开启裁剪
  imageUrl: ''
})

const onCrop = () => {
  console.log('onCrop', cropper.value.getCroppedCanvas())
  const canvas = cropper.value.getCroppedCanvas()
  const cropImg = canvas.toDataURL()
  uploadState.imageUrl = cropImg
  cropperState.visible = false
  canvas.toBlob((blob: Blob) => {
    // send the blob to server etc.
    doUpload(blob)
  }, uploadState.fileType, 1)
}

const doUpload = async (blob: Blob) => {
  if (state.aliOss) {
    try {
      const res = await state.aliOss.multipartUpload(uploadState.imageName, blob, {
        progress: (progress: number, checkpoint: any) => {
          uploadState.loading = true
          // onProgress({ percent: progress * 100 })  // 执行onProgress 并传入当前进度，使得上传组件正确显示进度条
        },
      })
      console.log(res)
      uploadState.loading = false
      uploadState.imageUrl = ossConfig.ginvdoc.domain + res.name
      formState.form.avatar = res.name
    } catch (e) {
      // onError()
      message.error((e as Error).message)
    }
  }
}
</script>

<style scoped>
.avatar-uploader:deep() > .ant-upload {
  width: 200px;
  height: 200px;
}
.upload-btns{
  display: flex;
  justify-content:space-between;
}
.avatar-uploader__img{
  position: absolute;
  left: 0;
  top: 0;
  width: 200px;
  height: 200px;
}
</style>
