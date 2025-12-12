import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

const service: AxiosInstance = axios.create({
  baseURL: (import.meta as any).env?.VITE_API_BASE || '',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor: can add auth token here
service.interceptors.request.use(
  (config) => {
    // example: attach token from localStorage
    // const token = localStorage.getItem('token')
    // if (token) config.headers.Authorization = `Bearer ${token}`
    return config
  },
  (error) => Promise.reject(error)
)

// Response interceptor — unwraps data and enforces business-code === 2000
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const res = response.data

    if (res.code !== 2000) {
        ElMessage.error(res.msg || `请求出错，错误码：${res.code}`)
        return Promise.reject()
    } else {
        if ((res.msg !== undefined || res.msg !== null || res.msg !== "") && res.msg.length > 0) {
            ElMessage.success(res.msg)
        }
    }

    return res.data
  },
  (error) => {
    return Promise.reject(error)
  }
)

export function request<T = any>(config: AxiosRequestConfig): Promise<T> {
  return service.request<T>(config) as unknown as Promise<T>
}

export default {
  request,
  get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return service.get<T>(url, config).then((r) => r as unknown as T)
  },
  post<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return service.post<T>(url, data, config).then((r) => r as unknown as T)
  },
}
