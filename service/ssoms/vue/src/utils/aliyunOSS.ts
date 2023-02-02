import OSS from 'ali-oss'
import { UnwrapNestedRefs, toRefs } from 'vue'
import type { UploadProps, UploadFile } from 'ant-design-vue'
import { UploadFileStatus } from 'ant-design-vue/lib/upload/interface'

export interface StsInfo {
  bucket: string
  accessKeyId: string
  accessKeySecret: string
  endpoint: string
  stsToken: string
}
export default class AliyunOSS implements UploadFile {
  public uid: string = ''
  public name: string = ''
  public file?: File
  public size?: number | undefined
  public percent?: number | undefined
  public status?: UploadFileStatus | undefined
  public store?: OSS
  private bucket: string
  private stsInfo: StsInfo
  private retryMax: number = 3
  private abortCheckpoint: any // 用于存放中断点
  private storagePrefix = 'aliyun.oss.'
  private refreshSTSToken: () => Promise<StsInfo>
  private uploadState: UnwrapNestedRefs<UploadProps>

  constructor (bucket: string, refreshSTSToken: () => Promise<StsInfo>, uploadState: UnwrapNestedRefs<UploadProps>) {
    this.bucket = bucket
    this.refreshSTSToken = refreshSTSToken
    this.uploadState = uploadState
    const key = this.storagePrefix + bucket
    const basicInfo = localStorage.getItem(key)
    this.stsInfo = { bucket: '', accessKeyId: '', accessKeySecret: '', endpoint: '', stsToken: '' }
    if (basicInfo === null) {
      this.refreshSTSToken().then(stsRes => {
        this.stsInfo = stsRes
        this.initRequest()
      })
    } else {
      const { bucket: realBucket, accessKeyId, accessKeySecret, endpoint, stsToken } = JSON.parse(basicInfo)
      this.stsInfo = { bucket: realBucket, accessKeyId, accessKeySecret, endpoint, stsToken }
      this.initRequest()
    }
  }

  // Starts the upload process.
  // xhr () {
  //   return this.sendRequest()
  // }

  // 中断分片上次
  abort () {
    if (this.store && this.abortCheckpoint) {
      this.store.abortMultipartUpload(this.abortCheckpoint.name, this.abortCheckpoint.uploadId)
    }
  }

  private initRequest () {
    this.store = new OSS({
      accessKeyId: this.stsInfo.accessKeyId || 'yourAccessKeyId',
      accessKeySecret: this.stsInfo.accessKeySecret || 'yourAccessKeySecret',
      bucket: this.stsInfo.bucket,
      endpoint: this.stsInfo.endpoint,
      stsToken: this.stsInfo.stsToken,
      refreshSTSToken: async () => {
        const stsRes = await this.refreshSTSToken()
        const key = this.storagePrefix + this.bucket
        localStorage.setItem(key, JSON.stringify(stsRes))
        return { accessKeyId: stsRes.accessKeyId, accessKeySecret: stsRes.accessKeySecret, stsToken: stsRes.stsToken }
      },
      refreshSTSTokenInterval: 300000
    })
  }

  public async sendRequest (file: UploadFile) {
    if (this.store && this.uploadState.fileList) {
      const index = this.uploadState.fileList.findIndex(f => f.uid === file.uid)
      if (index < 0) {
        return
      }
      try {
        await this.store.multipartUpload(file.name, file, {
          progress: (progress: number, checkpoint: any) => {
            this.uploadState.fileList![index].status = 'uploading'
            this.uploadState.fileList![index].percent = progress * 100
          },
        })
        this.uploadState.fileList[index].status = 'success'
        this.uploadState.fileList[index].name = file.name
      } catch (e) {
        this.uploadState.fileList[index].status = 'error'
      }
    }
  }
}
  