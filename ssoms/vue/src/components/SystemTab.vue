<template>
  <div class="content-header">
    <div>
      <h1>后台系统</h1>
    </div>
    <div class="content-header__actions">
      <a-button type="primary" @click="initAdd">添加系统</a-button>
    </div>
  </div>
  <a-table
  :dataSource="respState.list"
  :columns="columns"
  :pagination="false">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'objectName'">
        <span style="margin-left: 10px">{{record.objectName}}</span>
      </template>
      <template v-if="column.key === 'icon'">
        <a-image
          :width="60"
          :height="60"
          :src="record.icon"
          fallback="data:image/png;base64,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"
        />
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
          title="确定要删除该对象吗?"
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
    :title="formState.type == 'add' ? '添加系统' : '修改系统'"
    :loading="formState.loading"
    @cancel="onCancel"
    @ok="onSubmit">
    <a-form layout="vertical" ref="modalFormRef" :model="formState.form">
      <a-row :gutter="30">
        <a-col :span="5">
          <a-form-item label="logo">
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
            <a-form-item label="系统代号" name="form.key">
              <a-input v-model:value="formState.form.key" placeholder="系统代号" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="系统名" name="form.objectName">
              <a-input v-model:value="formState.form.objectName" placeholder="系统中文名" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="域名" name="form.domain">
              <a-input v-model:value="formState.form.domain" placeholder="域名" />
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
  objectList,
  objectDetail,
  addObject,
  updateObject,
  deleteObject,
  ObjectListReply,
  ObjectForm
} from '../api/ssoms'
import type { FormInstance, UploadChangeParam, UploadProps } from 'ant-design-vue'
import { message } from 'ant-design-vue'
import { PlusOutlined, LoadingOutlined } from '@ant-design/icons-vue'

/**
 * 这是表格的列定义
 * 没有动态变化，所以不使用reactive
 */
const columns = [
  {
    title: '系统代号',
    dataIndex: 'key',
    key: 'key'
  },
  {
    title: '系统名',
    dataIndex: 'objectName',
    key: 'objectName'
  },
  {
    title: 'logo',
    dataIndex: 'icon',
    key: 'icon'
  },
  {
    title: '排序',
    dataIndex: 'sort',
    key: 'sort'
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
    topKey: undefined,
    objectName: undefined,
    key: undefined
  }
})

/**
 * 这是列表响应
 */
const respState = reactive<ObjectListReply>({
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
    domain: '',
    key: '',
    icon: '',
    type: 1,
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
    domain: '',
    key: '',
    icon: '',
    type: 1,
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
      updateObject(formState.form).then(() => {
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
