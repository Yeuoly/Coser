import axios from 'axios'
import { stringify } from 'querystring'
import { loginTokenGet } from '../store/login'

export class Response<T = any> {
    code: number | undefined
    data: T | undefined
    message: string | undefined
    success: boolean | undefined

    isSuccess() {
        return this.success && this.code === 0
    }

    getError() {
        if (this.isSuccess()) {
            return ''
        }
        if (this.message) {
            return this.message
        }
        return '未知错误'
    }
}

export enum RequestMethods {
    GET,
    POST
}

export const apiBase = <T>(url: string, method: RequestMethods, args: any) => new Promise<Response<T>>(async (resolve) => {
    args = args || ''
    const { uid, token } = loginTokenGet()
    switch (method) {
        case RequestMethods.POST:
            try {
                const data = await axios.post<Response>(url, args, {
                    headers: {
                        'Content-Type': 'application/json',
                        'Billboards-User-Id': uid,
                        'Billboards-User-Token': token
                    }
                })
                const response = new Response<T>()
                if(typeof data === 'string') {
                    response.success = false
                    resolve(response)
                } else {
                    response.success = true
                    Object.assign(response, data.data)
                    resolve(response)
                }
            } catch (e: any) {
                const response = new Response<T>()
                try {
                    Object.assign(response, e.response.data)
                } catch (e) {
                    response.message = '未知错误'
                }
                response.success = false
                resolve(response)
            }
            break
        case RequestMethods.GET:
            if(typeof args === 'object') {
                args = stringify(args)
            }
            try {
                const data = await axios.get<Response>(url + '?' + args, {
                    headers: {
                        'Billboards-User-Id': uid,
                        'Billboards-User-Token': token
                    }
                })
                const response = new Response<T>()
                if(typeof data === 'string') {
                    response.success = false
                    resolve(response)
                } else {
                    response.success = true
                    Object.assign(response, data.data)
                    resolve(response)
                }
            } catch (e: any) {
                const response = new Response<T>()
                try {
                    Object.assign(response, e.response.data)
                } catch (e) {
                    response.message = '未知错误'
                }
                response.success = false
                resolve(response)
            }
            break
    }
})