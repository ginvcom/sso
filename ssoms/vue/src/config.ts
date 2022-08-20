import Ajax from "@/utils/ajax"

export const systemCode = <string>import.meta.env.VITE_SYSTEM_NAME || ''

export const SERVER_ROUTER_MENU_KEY = systemCode + ':' + 'server_router_menu'

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
export const auth = new Ajax(systemCode, 'auth')
