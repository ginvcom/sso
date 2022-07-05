import axios, { AxiosRequestConfig } from 'axios'
// import qs from 'qs'

let baseURL = 'http://localhost:8888'
// 线上灰度不同环境支持
const env = window.location.origin.search('gray.') > -1 ? 'gray' : 'online'
// test不同环境支持
const baseMatch = baseURL.match(/\.(test[1-5]{0,1})\./)
if (baseMatch !== null && baseMatch.length > 1) {
  const urlMatch = window.location.origin.match(/\.(test[1-5]{0,1})\./)
  if (urlMatch !== null && urlMatch.length > 1) {
    if (baseMatch[1] !== urlMatch[1]) {
      baseURL = baseURL.replace(baseMatch[1], urlMatch[1])
    }
  }
}
const instance = axios.create()

export interface CustomSuccessData<T> {
result: T;
success: boolean;
msg: string;
}
export interface SuccessData {
  <T>(url: string, params?: Record<string, unknown>, config?: AxiosRequestConfig): Promise<CustomSuccessData<T>>;
}

// 请求拦截器
instance.interceptors.request.use(function (config: AxiosRequestConfig) {
  if (config.method === 'post' || config.method === 'put') {
    if (config.headers !== undefined) {
      config.headers['Content-Type'] = 'application/json;charset=utf-8'
    }
    // config.withCredentials = true
  }
//   config.data = qs.stringify(config.data)
  return config
}, function (err: Error) {
  return Promise.reject(err)
})

// 响应拦截器
// api.interceptors.response.use(function (res) {
//   if (res.data.success) {
//     return res.data
//   } else {
//     // if (res.data.error.code >= 100000 && res.data.error.code < 100005) {
//     //   localStorage.removeItem('token')
//     //   localStorage.removeItem('user')
//     //   sessionStorage.removeItem('menu')
//     //   window.location = res.data.error.data
//     // }
//     message.error({ content: res.data.msg, key: 'golbal' })
//   }
// }, function (error) {
//   message.error({ content: error.message, key: 'golbal' })
//   // return Promise.reject(error)
// })

// const ajax: SuccessData = api.post

export interface HcResponse<T> {
  code: string;
  data: T;
  success: boolean;
  msg: string;
}

interface PostParams {
//   distinctRequestId: string;
  env: string;
  systemCode: string;
}

interface PostData {
  ticket: string | null;
  params: string;
}


function getUrl (url: string): string {
  const now = new Date()
  const dataWrapper: PostParams = {
    env: 'test',
    systemCode: 'test'
  }
  return `${url}?env=${env}&systemCode=test`
}

// function signOutForward () {
//   localStorage.removeItem('router')
//   Cookies.remove('ticket')
//   Cookies.remove('user')
//   Cookies.remove('app')
//   const beforeTemp = localStorage.getItem('beforeTemp')
//   window.location.href = server + '/redirect/sso?service=' + appName + '&env=' + env
// }

// eslint-disable-next-line
export async function request<T = any, _R = CustomSuccessData<T>> (url: string, data?: any) {
  const postURL = getUrl(url)
  const res = await instance.post(postURL, data).then(res => {
    if (res.data.code >= 0) {
      return Promise.resolve(res.data.data)
    } else {
      if (res.data.code === -99999) {
        // signOutForward()
      } else {
        return Promise.reject(new Error(res.data.msg))
      }
    }
  }).catch((err: Error) => {
    return Promise.reject(err)
  })
  return res
}

// eslint-disable-next-line
export default async function ajax<T = any, _R = CustomSuccessData<T>> (url: string, data?: any) {
  const res = request(url ,data).then(data => {
    return Promise.resolve(data)
  }).catch((err: Error) => {
    // alert(err.message)
    // message.error({ content: err.message, key: 'golbal' })
    return Promise.reject(err)
  })
  return res
}
