import request from '../../utils/request'

export class PrometheusConfigAPI {
  static async pageConfigs(params: any) {
    const url = '/api/v1/alert/prometheus/page'
    return await request.post(url, params)
  }

  static async addConfig(params: any) {
    const url = '/api/v1/alert/prometheus/add'
    return await request.post(url, params)
  }

  static async modifyConfig(params: any) {
    const url = '/api/v1/alert/prometheus/modify'
    return await request.post(url, params)
  }

  static async deleteConfig(params: any) {
    const url = '/api/v1/alert/prometheus/delete'
    return await request.post(url, params)
  }

  static async copyConfig(params: any) {
    const url = '/api/v1/alert/prometheus/copy'
    return await request.post(url, params)
  }

  static async listall() {
    const url = '/api/v1/alert/prometheus/listall'
    return await request.post(url)
  }

  static async sumbitConfig(params: any) {
    const url = '/api/v1/alert/prometheus/submit'
    return await request.post(url, params)
  }

    static async checkStatus(params: any) {
    const url = '/api/v1/alert/prometheus/status'
    return await request.post(url, params)
  }

}

