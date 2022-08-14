import axios, { AxiosError, AxiosRequestConfig } from 'axios'
import { message } from 'ant-design-vue'
import { getCookie } from './cookie'

// import qs from 'qs'
let baseURL = <string>import.meta.env.VITE_BASEURL || '/'
const system = <string>import.meta.env.VITE_SYSTEM_NAME || ''
// 线上灰度不同环境支持
const env = window.location.origin.search('gray.') > -1 ? 'gray' : 'online'
// test不同环境支持
const baseMatch = baseURL.match(/\.(test[1-9]{0,1})\./)
if (baseMatch !== null && baseMatch.length > 1) {
  const urlMatch = window.location.origin.match(/\.(test[1-9]{0,1})\./)
  if (urlMatch !== null && urlMatch.length > 1) {
    if (baseMatch[1] !== urlMatch[1]) {
      baseURL = baseURL.replace(baseMatch[1], urlMatch[1])
    }
  }
}

export const server = baseURL

export const instance = axios.create({
  baseURL,
  withCredentials: true
})

function signOutForward () {
  localStorage.removeItem('router')
  // Cookies.remove('ticket')
  // Cookies.remove('user')
  // Cookies.remove('app')
  // TODO 跳转到登出页，清cookie
  const beforeTemp = localStorage.getItem('beforeTemp')
  // window.location.href = server + '/redirect/sso?service=' + appName + '&env=' + env
}

// 请求拦截器
instance.interceptors.request.use((config: AxiosRequestConfig) =>{
  if (config.headers !== undefined) {
    const userStr = getCookie('user')
    if (userStr) {
      const user = JSON.parse(userStr)
      if (config.url !== '/sign-in' && user) {
        config.headers.Authorization = user.accessToken
      }
    }
  }
  if (config.method === 'post' || config.method === 'put') {
    if (config.headers !== undefined) {
      config.headers['Content-Type'] = 'application/json;charset=utf-8'
    }
  }
//   config.data = qs.stringify(config.data)
  return config
}, (err: Error) => {
  return Promise.reject(err)
})

// 响应拦截器
instance.interceptors.response.use((res) => {
  if (res.data.code >= 0) {
    return Promise.resolve(res.data.data)
  } else {
    if (res.data.code === -99999) {
      signOutForward()
      window.location.href = server + '/sign-out?system=' + system
    } else {
      message.error({ content: res.data.msg, key: 'golbal' })
      return Promise.reject(new Error(res.data.msg))
    }
  }
}, (error: AxiosError) => {
  if (error.response) {
    if (error.response.status == 401) {
      signOutForward()
      window.location.href = server + '/sign-out?system=' + system
      return Promise.reject(error)
    } else {
      if (error.response.data) {
        message.error({ content: <string>error.response.data, key: 'golbal' })
        return Promise.reject(error)
      }
    }
  }
    message.error({ content: error.message, key: 'golbal' })
    return Promise.reject(error)
})

export default class Ajax {
  private systemCode: string
  private serviceCode: string
  constructor (systemCode: string, serviceCode: string) {
    this.systemCode = systemCode
    this.serviceCode = serviceCode
  }
  private setPath (url: string, params: any) {
    const matches = url.match(/\/:[a-zA-z0-9]+/mg)
    if (matches) {
      for (const match of matches) {
        const key = match.slice(2)
        if (params && params[key]) {
          url = url.replace(match, '/' + params[key])
          delete params[key]
        }
      }
    }
    return { realpath: url, realParams: params }
  }
  private getHeaders (url: string) {
    return {
      'X-Origin-Uri' : url,
      'X-Origin-System': this.systemCode,
      'X-Origin-Service': this.serviceCode
    }
  }
  public get<T = any, R = T, D = any>(url: string, params?: D): Promise<R> {
    const { realpath, realParams } = this.setPath(url, params)
    const headers = this.getHeaders(url)
    return instance.get(realpath, { params: realParams, headers })
  }

  public delete<T = any, R = T, D = any>(url: string, params?: D): Promise<R> {
    const { realpath, realParams } = this.setPath(url, params)
    const headers = this.getHeaders(url)
    return instance.delete(realpath, { params: realParams, headers })
  }

  public post<T = any, R = T, D = any>(url: string, data?: D, params?: D): Promise<R> {
    const { realpath, realParams } = this.setPath(url, params)
    const headers = this.getHeaders(url)
    return instance.post(realpath, data, { params: realParams, headers })
  }

  public put<T = any, R = T, D = any>(url: string, params: D, data?: D): Promise<R> {
    const { realpath, realParams } = this.setPath(url, params)
    const headers = this.getHeaders(url)
    return instance.put(realpath, data, { params: realParams, headers })
  }

  public patch<T = any, R = T, D = any>(url: string, params: D, data?: D): Promise<R> {
    const { realpath, realParams } = this.setPath(url, params)
    const headers = this.getHeaders(url)
    return instance.patch(realpath, data, { params: realParams, headers })
  }
}
