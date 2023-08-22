import axios from 'axios'
import { resReject, resResolve, reqReject, reqResolve } from './interceptors'

export function createAxios(options = {}) {
  const defaultOptions = {
    timeout: 12000,
  }
  const service = axios.create({
    ...defaultOptions,
    ...options,
  })
  service.interceptors.request.use(reqResolve, reqReject) // 请求拦截
  service.interceptors.response.use(resResolve, resReject) // 响应拦截
  return service
}

export const request = createAxios({
  baseURL: import.meta.env.VITE_BASE_API,
})
