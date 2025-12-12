import request from '../../utils/request'

export class AlertRuleAPI {

  static async pageAlertRules(params: any) {
    const url = '/api/v1/alert/rule/page'
    return await request.post(url, params)
  }

  static async deleteAlertRule(params: any) {
    const url = '/api/v1/alert/rule/delete'
    return await request.post(url, params)
  }

  static async modifyAlertRule(params: any) {
    const url = '/api/v1/alert/rule/modify'
    return await request.post(url, params)
  }

  static async copyAlertRule(params: any) {
    const url = '/api/v1/alert/rule/copy'
    return await request.post(url, params)
  }

  static async addAlertRule(params: any) {
    const url = '/api/v1/alert/rule/add'
    return await request.post(url, params)
  }

  static async submitAlertRule(params: any) {
    const url = '/api/v1/alert/rule/submit'
    return await request.post(url, params)
  }

}

