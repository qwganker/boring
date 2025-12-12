import request from '../../utils/request'

export class AlertTypeAPI {

  static async pageAlertType(params: any) {
    const url = '/api/v1/alert/alert_type/page'
    return await request.post(url, params)
  }

  static async listAllAlertType() {
    const url = '/api/v1/alert/alert_type/listall'
    return await request.post(url)
  }

  static async deleteAlertType(params: any) {
    const url = '/api/v1/alert/alert_type/delete'
    return await request.post(url, params)
  }

  static async modifyAlertType(params: any) {
    const url = '/api/v1/alert/alert_type/modify'
    return await request.post(url, params)
  }

  static async addAlertType(params: any) {
    const url = '/api/v1/alert/alert_type/add'
    return await request.post(url, params)
  }

}

