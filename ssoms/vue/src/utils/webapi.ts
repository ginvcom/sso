import axios, { AxiosRequestConfig } from 'axios'
import qs from 'qs'

const instance = axios.create()
instance.defaults.timeout = 2500

instance.interceptors.request.use((config: AxiosRequestConfig) => {
  if (config.headers) {
    config.headers['Content-Type'] = 'application/json'
  }
})


