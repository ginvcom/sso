<template>
  <layout>
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
        <a-tab-pane key="1" tab="基础设置"><profile-change v-bind="profileData.info" @change="onProfileChange" /></a-tab-pane>
        <a-tab-pane key="2" tab="密码修改"><password-reset /></a-tab-pane>
      </a-tabs>
    </div>
  </layout>
</template>
<script setup lang="ts">
import { onBeforeMount, reactive, ref } from 'vue'
import VueCropper from 'vue-cropperjs'
import { sts } from '@/api/oss'
import { message, UploadChangeParam, UploadFile, UploadProps } from 'ant-design-vue'
import { UserOutlined, PlusOutlined, LoadingOutlined } from '@ant-design/icons-vue'
import AliyunOSS, { StsInfo } from '@/utils/aliyunOSS'
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

const refreshSTSToken = async (bucket: string) => {
  const res = await sts({ bucket })
  return Promise.resolve(res as StsInfo)
}

onBeforeMount(() => {
  state.aliOss = new AliyunOSS(ossConfig.doc.bucket, () => refreshSTSToken(ossConfig.doc.bucket), uploadState)
  profile().then(data => {
    profileData.info = data
    uploadState.fileList = [{ uid: '', name: data.avatar, url: ossConfig.doc.domain + data.avatar }]
  })
})

const onProfileChange = ({ introduction }) => {
  profileData.info.introduction = introduction
}

const activeKey = ref('1')

const cropper = ref()

const uploadState = reactive<UploadProps>({ fileList: [] })

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
</script>
