import OSS from 'ali-oss'
import { getKey } from './file'
import webapi from './webapi'
import type { UploadFile } from 'ant-design-vue'
import { UploadFileStatus } from 'ant-design-vue/lib/upload/interface'

class AliyunOSS implements UploadFile {
  public uid: string = ''
  public name: string = ''
  public file?: File
  public size?: number | undefined
  public percent?: number | undefined
  public status?: UploadFileStatus | undefined
  private accessKeyId: string
  private accessKeySecret: string
  private bucket: string
  private endpoint: string
  private store?: OSS
  private retryMax: number = 3
  private abortCheckpoint: any // 用于存放中断点

  constructor (bucket: string, accessKeyId: string, accessKeySecret: string, endpoint: string) {
    this.bucket = bucket
    this.accessKeyId = accessKeyId
    this.accessKeySecret = accessKeySecret
    this.endpoint = endpoint
    this.uid = ''
  }

  // Starts the upload process.
  xhr () {
    this.initRequest()
    return this.sendRequest()
  }

  // 中断分片上次
  abort () {
    if (this.store && this.abortCheckpoint) {
      this.store.abortMultipartUpload(this.abortCheckpoint.name, this.abortCheckpoint.uploadId)
    }
  }

  // Initializes the XMLHttpRequest object using the URL passed to the constructor.
  private initRequest () {
    this.store = new OSS({
      accessKeyId: this.accessKeyId,
      accessKeySecret: this.accessKeySecret,
      bucket: this.bucket,
      endpoint: this.endpoint,
      refreshSTSToken: async () => {
        const info = await webapi.get('aliyun/sts', { bucket: 'doc' }) as any
        return {
          accessKeyId: info.accessKeyId,
          accessKeySecret: info.accessKeySecret,
          stsToken: info.stsToken
        }
      },
      refreshSTSTokenInterval: 300000
    })
  }

  private async sendRequest () {
    // const data = new FormData()
    // data.append('upload', file)
    // this.xhr.send(data)
    console.log(this.store)
    if (!this.store || !this.file) {
      return
    }
    for (let i = 0; i <= this.retryMax; i++) {
      try {
        // object表示上传到OSS的文件名称。
        // file表示浏览器中需要上传的文件，支持HTML5 file和Blob类型。
        this.name = getKey(this.file)
        const result = await this.store.multipartUpload(this.name, this.file, {
          progress: (p: any, cpt: any, res: any) => {
            // 为中断点赋值。
            this.abortCheckpoint = cpt
            // 获取上传进度。
            console.log(cpt, res)
            this.status = 'uploading'
            // // this.loader.uploadedPercent = p
            this.percent = p
          }
        })
        this.percent = 100
        this.status = 'success'
        return Promise.resolve({ default: (result.res as any).requestUrls[0] })
      } catch (e) {
        this.status = 'error'
        return Promise.reject((e as Error).message)
      }
    }
  }
}
  