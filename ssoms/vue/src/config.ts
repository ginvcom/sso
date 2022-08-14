import Ajax from "./utils/ajax"

export const systemCode = <string>import.meta.env.VITE_SYSTEM_NAME || ''

export const ossConfig = {
  ginv: {
    bucket: 'ginv',
    accessKeyId: '',
    accessKeySecret: '',
    endpoint: ''
  },
  ginvdoc: {
    bucket: 'ginvdoc',
    domain: '//ginvdoc.oss-cn-shenzhen.aliyuncs.com/',
    accessKeyId: 'LTAI5t7n3KoUNHQgrUep8auP',
    accessKeySecret: 'K8AaLsyaJXksMSsuEviu4DqSJEU5OJ',
    endpoint: 'oss-cn-shenzhen.aliyuncs.com'
  }
}

export const ssoms = new Ajax(systemCode, 'ssoms')
