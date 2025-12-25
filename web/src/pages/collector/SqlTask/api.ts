import request from '../../../utils/request'

export class SqlTaskAPI {
    static async page(params: any) {
        const url = '/api/v1/collector/sqltask/page'
        return await request.post(url, params)
    }

    static async add(params: any) {
        const url = '/api/v1/collector/sqltask/add'
        return await request.post(url, params)
    }

    static async modify(params: any) {
        const url = '/api/v1/collector/sqltask/modify'
        return await request.post(url, params)
    }

    static async delete(params: any) {
        const url = '/api/v1/collector/sqltask/delete'
        return await request.post(url, params)
    }

    static async copy(params: any) {
        const url = '/api/v1/collector/sqltask/copy'
        return await request.post(url, params)
    }

    static async test(params: any) {
        const url = '/api/v1/collector/sqltask/run_test'
        return await request.post(url, params)
    }

    static async run(params: any) {
        const url = '/api/v1/collector/sqltask/run'
        return await request.post(url, params)
    }

    static async stop(params: any) {
        const url = '/api/v1/collector/sqltask/stop'
        return await request.post(url, params)
    }

}

