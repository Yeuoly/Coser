import { RequestMethods, apiBase } from "./base"
import { Place } from "./types"

export const apiCreatePlace = (name: string, description: string, lat: number, lng: number, key: string, avatar: string) => apiBase<{
    id: number
}>('/v1/place/create', RequestMethods.POST, {
    name, description, lat, lng, key, avatar
})

export const apiUpdatePlace = (id: number, name: string, description: string, lat: number, lng: number, key: string, avatar: string) => apiBase<{

}>('/v1/place/update', RequestMethods.POST, {
    id, name, description, lat, lng, key, avatar
})

export const apiPlaceInfo = (id: number) => apiBase<{

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

export const apiCreateTag = (name: string) => apiBase<{

}>('/v1/tag/create', RequestMethods.POST, {
    name
})

export const apiSearchTag = (keyword: string) => apiBase<{

}>('/v1/tag/search', RequestMethods.GET, {
    keyword
})

export const apiCreateGallery = (
    place_id: number, name: string, coser: string, description: string, photographers: string, character: string,
    series: string, tags: string[], key: string
) => apiBase<{
        
}>('/v1/gallery/create', RequestMethods.POST, {
    place_id, name, coser, description, photographers, character, series, tags, key
})

export const apiUpdateGallery = (
    id: number, place_id: number, name: string, coser: string, description: string, photographers: string, character: string,
    series: string, tags: string[], key: string
) => apiBase<{
        
}>('/v1/gallery/update', RequestMethods.POST, {
    id, place_id, name, coser, description, photographers, character, series, tags, key
})

export const apiGalleryInfo = (id: number) => apiBase<{
        
}>('/v1/gallery/info', RequestMethods.GET, {
    id
})

export const apiGallerySearch = (keyword: string) => apiBase<{
        
}>('/v1/gallery/search', RequestMethods.GET, {
    keyword
})

export const apiGalleryUploadImage = (gallery_id: number, key: string) => apiBase<{
        
}>('/v1/gallery/upload', RequestMethods.POST, {
    gallery_id, key
})

export const apiGalleryDeleteImage = (gallery_id: number, key: string) => apiBase<{
        
}>('/v1/gallery/delete', RequestMethods.POST, {
    gallery_id, key
})
