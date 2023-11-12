import { RequestMethods, Response, apiBase } from "./base"
import { Gallery, GalleryTag, Place } from "./types"

export const apiCreatePlace = (name: string, description: string, lat: number, lng: number, key: string, avatar: string) => apiBase<{
    id: number
}>('/v1/place/create', RequestMethods.POST, {
    name, description, lat, lng, key, avatar
})

export const apiUpdatePlace = (id: number, name: string, description: string, lat: number, lng: number, key: string, avatar: string) => apiBase<{

}>('/v1/place/update', RequestMethods.POST, {
    id, name, description, lat, lng, key, avatar
})

export const apiDeletPlace = (id: number, key: string) => apiBase<{
    success: boolean
}>('/v1/place/delete', RequestMethods.POST, {
    id, key
})

export const apiPlaceInfo = (id: number) => apiBase<{
    place: Place
}>('/v1/place/info', RequestMethods.GET, {
    id
})

export const apiPlaceList = (keyword: string) => apiBase<{
    
}>('/v1/place/list', RequestMethods.GET, {
    keyword
})

export const apiPlaceNearby = (lat: number, lng: number) => apiBase<{
    places: Place[]
}>('/v1/place/nearby', RequestMethods.GET, {
    lat, lng
})

const tagMap: {[name: string]: number} = {}
export const apiCreateTag = async (name: string) => {
    const response = await apiBase<{
        id: number
    }>('/v1/tag/create', RequestMethods.POST, {
        name
    })

    if (!response.isSuccess()) {
        return response
    }

    tagMap[name] = response.data?.id as number
    return response
}

export const apiSearchTag = (keyword: string) => apiBase<{
    tags: GalleryTag[]
}>('/v1/tag/search', RequestMethods.GET, {
    keyword
})

export const apiCreateGallery = (
    place_id: number, name: string, coser: string, description: string, photographers: string, character: string,
    series: string, tags: number[], key: string
) => apiBase<{
    id: number
}>('/v1/gallery/create', RequestMethods.POST, {
    place_id, name, coser, description, photographers, character, series, tags, key
})

export const apiUpdateGallery = (
    id: number, place_id: number, name: string, coser: string, description: string, photographers: string, character: string,
    series: string, tags: number[], key: string
) => apiBase<{
        
}>('/v1/gallery/update', RequestMethods.POST, {
    id, place_id, name, coser, description, photographers, character, series, tags, key
})

export const apiGalleryInfo = (id: number) => apiBase<{
        
}>('/v1/gallery/info', RequestMethods.GET, {
    id
})

export const apiGallerySearch = (keyword: string) => apiBase<{
    galleries: Gallery[]
}>('/v1/gallery/search', RequestMethods.GET, {
    keyword
})

export const apiGalleryUploadImage = (
    gallery_id: number, filename: string, key: string, content_type: string, camera: string, lens: string, focal_length: string,
    file: File
) => {
    return new Promise<Response<{
        url: string,
        id: number
    }>>(async (resolve) => {
        const response = await apiBase<{
            url: string,
            id: number
        }>('/v1/gallery/upload', RequestMethods.POST, {
            gallery_id, filename, key, content_type, camera, lens, focal_length
        })

        if (!response.isSuccess()) {
            resolve(response)
            return
        }
    
        const url = response.data?.url as string

        var xhr = new XMLHttpRequest()
        xhr.open("PUT", url, true)
        xhr.setRequestHeader("Content-Type", file.type)
        xhr.onload = async function () {
            if (xhr.status === 200) {
                const readurl = url.split('?')[0]
                const result = new Response<{url: string, id: number}>()
                result.code = 0
                result.success = true
                result.data = {
                    url: readurl,
                    id: response.data?.id as number
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
                const response = new Response<{url: string, id: number}>()
                response.code = xhr.status
                response.message = xhr.statusText
                resolve(response)
            }
        }
        xhr.send(file)
    })
}

export const apiGalleryDelete = (id: number, key: string) => apiBase<{
    success: boolean
}>('/v1/gallery/delete', RequestMethods.POST, {
    id, key
})

export const apiGalleryRemoveImage = (gallery_id: number, id: number, key: string) => apiBase<{
    success: boolean
}>('/v1/gallery/image/remove', RequestMethods.POST, {
    id, key, gallery_id
})