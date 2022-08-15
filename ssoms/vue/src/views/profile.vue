<template>
  <div>
    <div class="content-header is-sticky">
      <div>
        <h1>个人中心</h1>
      </div>
    </div>
    <div class="profile__header">
      <div class="profile__bg"></div>
      <div class="profile__info" style="justify-content: space-between;">
        <div class="pofile__avatar">
          <a-tooltip placement="bottom">
            <template #title>点击修改头像</template>
            <div style="width:200px">
            <template v-if="cropperState.visible">
              <vue-cropper
                ref="cropper"
                :aspect-ratio="1 / 1"
                :src="cropperState.imageUrl"
                @keyup.enter="onCrop"
              />

            <a-button block style="margin-top: 16px;width: 200px;" @click="onCrop">完成裁剪并上传</a-button>
            </template>
            <a-upload v-else
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
                <div class="ant-upload-text">上传头像</div>
              </div>
            </a-upload>
            </div>
          </a-tooltip>
        </div>
        <div  class="profile__profile">
          <h1>{{profileData.info.name}}</h1>
          <p>个人说明：{{profileData.info.introduction}}</p>
        </div>
        <div class="profile__basic">
          <div>
            <b>uuid：</b>
            <span>{{profileData.info.uuid}}</span>
          </div>
          <div>
            <b>手机号：</b>
            <span>{{profileData.info.mobile}}</span>
          </div>
          <div>
            <b>生日：</b>
            <span>{{profileData.info.birth}}</span>
          </div>
          <div>
            <b>性别：</b>
            <span>{{profileData.info.mobile}}</span>
          </div>
        </div>
      </div>
    </div>
    <a-tabs v-model:activeKey="activeKey" :tabBarStyle="{padding: '0 10px'}">
      <a-tab-pane key="2" tab="最近我添加菜单" force-render>Content of Tab Pane 2</a-tab-pane>
      <a-tab-pane key="3" tab="批量导入菜单">Content of Tab Pane 3</a-tab-pane>
      <a-tab-pane key="1" tab="基础设置"><profile-change v-bind="profileData.info" @change="onProfileChange" /></a-tab-pane>
      <a-tab-pane key="4" tab="密码修改"><password-reset /></a-tab-pane>
    </a-tabs>
  </div>
</template>
<script setup lang="ts">
import { onBeforeMount, onMounted, reactive, ref } from 'vue'
import VueCropper from 'vue-cropperjs'
import { message } from 'ant-design-vue'
import { UserOutlined, PlusOutlined, LoadingOutlined } from '@ant-design/icons-vue'
import type { FormInstance, UploadFile, UploadChangeParam, UploadProps } from 'ant-design-vue'
import { getFileName } from '../utils/file'
import { ossConfig } from '@/config'
import PasswordReset from '@/components/PasswordReset.vue'
import ProfileChange from '@/components/ProfileChange.vue'
import {
  profile,
  UserForm
} from '../api/ssoms'

interface State {
  aliOss?: any
  loading: boolean
  uuid: string
}

/**
 * 这是列表的入参
 * 没有动态变化，所以不使用reactive
 */
const state = reactive<State>({
  loading: false, // 列表是否加载完成
  uuid: ''
})

interface ProfileInfo {
  loading: boolean
  info: UserForm
}

const profileData = reactive<ProfileInfo>({
  loading: false,
  info: {
    uuid: '',
    name: '',
    mobile: '',
    avatar: '',
    gender: 1,
    birth: '',
    introduction: '',
    status: 1
  }
})

onBeforeMount(() => {
  profile().then(data => {
    profileData.info = data
    uploadState.imageName = data.avatar
    uploadState.imageUrl = ossConfig.ginvdoc.domain + data.avatar
  })
})

const onProfileChange = ({ introduction }) => {
  profileData.info.introduction = introduction
}

const activeKey = ref('4')

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
      profileData.info.avatar = res.name
    } catch (e) {
      // onError()
      message.error((e as Error).message)
    }
  }
}
</script>
