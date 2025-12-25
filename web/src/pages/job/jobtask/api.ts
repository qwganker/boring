import request from '../../../utils/request'

export class JobTaskAPI {
    static async page(params: any) {
        const url = '/api/v1/job/jobtask/page'
        return await request.post(url, params)
    }

    // static async add(params: any) {
    //     const url = '/api/v1/alert/sqltask/add'
    //     return await request.post(url, params)
    // }

    // static async modify(params: any) {
    //     const url = '/api/v1/alert/sqltask/modify'
    //     return await request.post(url, params)
    // }

    // static async delete(params: any) {
    //     const url = '/api/v1/alert/sqltask/delete'
    //     return await request.post(url, params)
    // }

    // static async copy(params: any) {
    //     const url = '/api/v1/alert/sqltask/copy'
    //     return await request.post(url, params)
    // }

    // static async test(params: any) {
    //     const url = '/api/v1/alert/sqltask/run_test'
    //     return await request.post(url, params)
    // }

}

