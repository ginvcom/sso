<template>
  <div class="content-header">
    <div>
      <h1>后台系统</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary" @click="initAdd"><template #icon><appstore-add-outlined /></template>添加系统</a-button>
    </div>
  </div>
  <a-table
  :dataSource="respState.list"
  :columns="columns"
  :pagination="false">
    <template #bodyCell="{ column, record }">
      <template v-if="column.dataIndex === 'objectName'">
        <span style="margin-left: 10px">{{record.objectName}}</span>
      </template>
      <template v-if="column.dataIndex === 'icon'">
        <a-image
          :width="60"
          :height="60"
          :src="ossConfig.ginvdoc.domain + record.icon"
          fallback="data:image/png;base64,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"
        />
      </template>
      <template v-if="column.dataIndex === 'status'">
        <span v-if="record.status == 1"><a-badge status="success" /> 启用</span>
        <span v-else-if="record.status == 0"><a-badge status="error" />停用</span>
      </template>
      <template v-if="column.dataIndex === 'actions'">
        <a @click="initEdit(record.uuid)">编辑</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="确定要删除该对象吗?"
          placement="topRight"
          @confirm="onDelete(record.uuid)">
          <a>删除</a>
        </a-popconfirm>
      </template>
    </template>
  </a-table>
  <a-modal
    width="710px"
    v-model:visible="formState.visible"
    :title="formState.type == 'add' ? '添加系统' : '修改系统'"
    :loading="formState.loading"
    :maskClosable="false"
    @cancel="onCancel"
    @ok="onSubmit">
    <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
      <a-row :gutter="24">
        <a-col :span="8">
          <a-form-item label="logo">
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
              <a-image
                v-if="uploadState.imageUrl" class="avatar-uploader__img"
                :width="200"
                :height="200"
                :src="uploadState.imageUrl"
                :preview="false"
                fallback="data:image/png;base64,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"
              />
              <div>
                <loading-outlined v-if="uploadState.loading"></loading-outlined>
                <plus-outlined v-else></plus-outlined>
                <div class="ant-upload-text">上传</div>
              </div>
            </a-upload>
          </a-form-item>
        </a-col>
        <a-col :span="16">
          <a-row :gutter="24">
          <a-col :span="12">
            <a-form-item label="系统代号" name="key" :rules="[{ required: true, message: '请输入系统代号' }]">
              <a-input v-model:value="formState.form.key" placeholder="系统代号" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="系统名" name="objectName" :rules="[{ required: true, message: '请输入系统中文名' }]">
              <a-input v-model:value="formState.form.objectName" placeholder="系统中文名" />
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item label="域名" name="identifier" :rules="[{ required: true, message: '请输入域名' }]">
              <a-input v-model:value="formState.form.identifier" placeholder="域名" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item name="subType" :rules="[{ required: true, message: '请选择分类' }]">
              <template #label>
                <span>分类</span>
                （<a @click="onCategoryManage">分类管理</a>）
              </template>
              <a-input-number v-model:value="formState.form.subType" placeholder="排序值小的靠前" style="width: 100%" />
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-form-item label="排序值" name="sort">
              <a-input-number v-model:value="formState.form.sort" placeholder="排序值小的靠前" style="width: 100%" />
            </a-form-item>
          </a-col>
          <a-col :span="4">
            <a-form-item label="状态" name="status" :rules="[{ required: true, message: '请选择系统状态' }]">
              <a-switch v-model:checked="formState.form.status" :checked-value="1" :unChecked-value="0" checked-children="启用" un-checked-children="停用" />
            </a-form-item>
          </a-col>
        </a-row>
        </a-col>
      </a-row>
    </a-form>
  </a-modal>
  <category-modal ref="category" />
</template>
<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import CategoryModal from './CategoryModal.vue'
import VueCropper from 'vue-cropperjs'
import 'cropperjs/dist/cropper.css'
import {
  objectList,
  objectDetail,
  addObject,
  updateObject,
  deleteObject,
  ObjectListReply,
  ObjectForm,
  ObjectListReqParams
} from '../api/ssoms'
import type { FormInstance, UploadChangeParam, UploadProps } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { PlusOutlined, LoadingOutlined, AppstoreAddOutlined } from '@ant-design/icons-vue'
import { getFileName } from '../utils/file'
import { sts } from '../api/aliOss'
import OSS from 'ali-oss'
import { ossConfig } from '@/config'

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
const columns = [
  {
    title: '系统代号',
    dataIndex: 'key'
  },
  {
    title: '系统名',
    dataIndex: 'objectName'
  },
  {
    title: 'logo',
    dataIndex: 'icon'
  },
  {
    title: '状态',
    dataIndex: 'status'
  },
  {
    title: '排序',
    dataIndex: 'sort'
  },
  {
    title: '操作',
    dataIndex: 'actions',
    align: 'center',
    width: '120px'
  }
]

interface State {
  aliOss?: any
  loading: boolean
  params: ObjectListReqParams
}


/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive<State>({
  loading: false, // 列表是否加载完成
  params: {
    topKey: '',
    objectName: '',
    key: ''
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<ObjectListReply>({
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
  objectList(state.params).then((data: ObjectListReply) => {
    respState.list = data.list
  }).finally(() => {
    state.loading = false
  })
}

/**
 * 删除对象
 */
const onDelete = (uuid: string) => {
  deleteObject({ uuid }).then(() => {
    message.success('删除对象操作成功')
    getList()
  })
}

const modalFormRef = ref<FormInstance>()

interface Form {
  type: string
  visible: boolean,
  loading: boolean,
  form: ObjectForm
}

const formState = reactive<Form>({
  visible: false,
  loading: false,
  type: 'add',
  form: {
    uuid: '',
    objectName: '',
    identifier: '',
    key: '',
    icon: '',
    type: 1,
    subType: 1,
    pUUID: '',
    status: 1,
    sort: 0,
    topKey: ''
  }
})

const initAdd = () => {
  // resetFileds初始化无效，手动初始化
  formState.form = {
    uuid: '',
    objectName: '',
    identifier: '',
    key: '',
    icon: '',
    type: 1,
    subType: 1,
    pUUID: '',
    status: 1,
    sort: 0,
    topKey: ''
  }
  formState.type = 'add'
  formState.visible = true
}


/**
 * 获取对象详情, 用于编辑
 * @param uuid
 */
const initEdit = (uuid: string) => {
  objectDetail({ uuid }).then((data: ObjectForm) => {
    formState.form = data
    formState.type = 'edit'
    formState.visible = true
  })
}

/**
 * 新增或修改对象
 */
const onSubmit = () =>{
  modalFormRef.value?.validate().then(() => {
    formState.loading = true
    formState.form.topKey = formState.form.key
    if (formState.type === 'add') {
      addObject(formState.form).then(() => {
        message.success('新增对象成功')
        formState.visible = false
        modalFormRef.value?.resetFields()
        getList()
      }).finally(() => {
        formState.loading = false
      })
    } else {
      updateObject({ uuid: formState.form.uuid! }, formState.form).then(() => {
        message.success('修改对象成功')
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
    const getBase64 = (img: Blob, callback: (base64Url: string) => void) => {
      const reader = new FileReader()
      reader.addEventListener('load', () => callback(reader.result as string))
      reader.readAsDataURL(img)
    }

    const reader = new FileReader()
    reader.addEventListener('load', () => {
      cropperState.visible = true
      cropperState.imageUrl = reader.result as string
      uploadState.loading = false
      return false
    })
    reader.readAsDataURL(file)
    return false
  }
  return false
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
  console.log(state.aliOss)
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
      formState.form.icon = res.name
      console.log(formState.form)
    } catch (e) {
      // onError()
      message.error((e as Error).message)
    }
  }
}

const category = ref()

const onCategoryManage = () => {
  category.value.state.visible = true
}
</script>
