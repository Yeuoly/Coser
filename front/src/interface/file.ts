import axios from "axios"
import { RequestMethods, Response, apiBase } from "./base"

export const apiUploadFile = (filename: string, file: File) => new Promise<Response<{
    url: string
}>>(async (resolve) => {
    const response = await apiBase<{
        url: string
    }>('/v1/file/upload', RequestMethods.POST, {
        content_type: file.type,
        filename
    })

    if (!response.isSuccess()) {
        resolve(response)
        return
    }

    const { url } = response.data as {
        url: string
    }

    var xhr = new XMLHttpRequest()
    xhr.open("PUT", url, true)
    xhr.setRequestHeader("Content-Type", file.type)
    xhr.onload = async function () {
        if (xhr.status === 200) {
            const readurl = url.split('?')[0]
            const result = new Response<{url: string}>()
            result.code = 0
            result.success = true
            result.data = {
                url: readurl
            }

            const finishResponse = await apiBase<{
                success: boolean
            }>('/v1/file/upload/finish', RequestMethods.POST, {
                url: readurl
            })
        
            if (!finishResponse.isSuccess()) {
                response.message = finishResponse.message
                response.code = finishResponse.code
                response.success = finishResponse.success
                resolve(response)
                return
            }

            resolve(result)
        } else {
            const response = new Response<{url: string}>()
            response.code = xhr.status
            response.message = xhr.statusText
            resolve(response)
        }
    }

    xhr.send(file)
})