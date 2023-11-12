
<template>
    <div class="c-v-map">
        <div class="map-container">
            <BMap @onClick="handleMapClick" @onLocate="handleMapLocate" :points="points" :enable-modify="false"
                :enable-add-center-marker="false" enable-search mode="multi"></BMap>
        </div>
    </div>
    <NDrawer v-model:show="showPlace" :width="500">
        <NDrawerContent>
            <div v-if="!currentPlace">
                <NH3 prefix="bar">
                    <NText type="primary">这里似乎什么都没有，来创建一个吧！</NText>
                </NH3>
                <NResult status="warning">
                    请注意，当您创建完地点后，有且只有您当前使用的浏览器可以修改、删除该地点，其他人只拥有查看权限，当然了，其他人也可以在这个地点创建正片集。
                </NResult>
                <NH6 prefix="bar" class="w100">
                    <div class="text-13 w100">
                        这里是哪
                    </div>
                    <NInput v-model:value="newPlaceForm.name" maxlength="64" show-count />
                </NH6>
                <NH6 prefix="bar" class="w100">
                    <div class="text-13 w100">
                        简单描述一下
                    </div>
                    <NInput v-model:value="newPlaceForm.description" maxlength="1024" type="textarea" show-count />
                </NH6>
                <NH6 prefix="bar" class="w100" v-if="!newPlaceForm.avatar">
                    <div class="text-13 w100">
                        上传一张封面吧
                    </div>
                    <NUpload directory-dnd :max="5" :custom-request="uploadAvatar">
                        <NUploadDragger>
                            <div style="margin-bottom: 12px">
                                <NIcon size="48" :depth="3">
                                    <Image />
                                </NIcon>
                            </div>
                            <NText style="font-size: 16px">
                                点击或者拖动文件到该区域来上传
                            </NText>
                            <NP depth="3" style="margin: 8px 0 0 0">
                                请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码
                            </NP>
                        </NUploadDragger>
                    </NUpload>
                </NH6>
                <NH6 prefix="bar" class="w100" v-else>
                    <div class="text-13 w100">
                        封面预览
                    </div>
                    <div class="w100" style="height: 300px;">
                        <NImage style="border-radius: 12px;" :src="newPlaceForm.avatar" object-fit="cover" height="300"
                            width="450"></NImage>
                    </div>
                </NH6>
                <NH6 prefix="bar" class="w100">
                    <NButton block type="primary" @click="createPlace">创建</NButton>
                </NH6>
            </div>
            <div v-else>
                <NThing>
                    <template #header>
                        <NText type="primari">
                            {{ currentPlace.name }}
                        </NText>
                    </template>
                    <img :src="currentPlace.avatar" style="width: 100%; height: 300px; object-fit: contain;" />
                    <template #footer>
                        <div class="text-13">
                            {{ currentPlace.description }}
                        </div>
                    </template>
                    <template #header-extra>
                        <NPopover trigger="hover" style="padding: 0;">
                            <template #trigger>
                                <NButton text>
                                    <template #icon>
                                        <NIcon :component="Menu"></NIcon>
                                    </template>
                                </NButton>
                            </template>
                            <div>
                                <NButtonGroup vertical>
                                    <NButton type="primary" @click="updatePlace">
                                        <template #icon>
                                            <NIcon :component="CreateOutline"></NIcon>
                                        </template>
                                        编辑地点
                                    </NButton>
                                    <NButton type="error" @click="deletePlace(currentPlace.ID)">
                                        <template #icon>
                                            <NIcon :component="TrashBinOutline"></NIcon>
                                        </template>
                                        删除地点
                                    </NButton>
                                </NButtonGroup>
                            </div>
                        </NPopover>
                    </template>
                </NThing>
                <NDivider></NDivider>
                <NThing>
                    <template #header>
                        正片列表
                    </template>
                    <template #header-extra>
                        <NButton type="primary" @click="createGallery">
                            <template #icon>
                                <NIcon size="20" :component="Image"></NIcon>
                            </template>
                            <span>创建正片</span>
                        </NButton>
                    </template>
                    <NThing v-for="gallery in currentPlace.galleries" :key="gallery.ID">
                        <template #header>
                            {{ gallery.name }}
                        </template>
                        <template #header-extra>
                            <NButtonGroup>
                                <NButton size="small" type="info">
                                    让我康康！
                                </NButton>
                                <NTooltip trigger="hover">
                                    <template #trigger>
                                        <NButton size="small" type="primary" @click="updateGallery(gallery)">
                                            <template #icon>
                                                <NIcon :component="CreateOutline"></NIcon>
                                            </template>
                                            编辑
                                        </NButton>
                                    </template>
                                    只有您创建的才可以编辑哦
                                </NTooltip>
                                <NTooltip trigger="hover">
                                    <template #trigger>
                                        <NButton size="small" type="error" @click="deleteGallery(gallery)">
                                            <template #icon>
                                                <NIcon :component="Trash"></NIcon>
                                            </template>
                                            删除
                                        </NButton>
                                    </template>
                                    只有您创建的才可以删除哦
                                </NTooltip>
                            </NButtonGroup>
                        </template>
                        <NImage style="border-radius: 12px;" width="450" object-fit="contain" height="300"
                            :src="(gallery.images && gallery.images.length > 0) ? gallery.images[0].url + '?x-oss-process=image/resize,mfit,h_400,w_800' : ''">
                        </NImage>
                        <NDivider></NDivider>
                    </NThing>
                </NThing>
            </div>
        </NDrawerContent>
    </NDrawer>
</template>

<script setup lang="ts">
import { computed, h, ref } from 'vue'
import BMap from '../components/BMap.vue'
import { Gallery, Place } from '../interface/types'
import {
    NButton, NButtonGroup, NDivider, NDrawer, NDrawerContent, NDropdown, NH3, NH6,
    NIcon, NImage, NInput, NP, NPopover, NResult, NText,
    NThing,
    NTooltip,
    NUpload, NUploadDragger, UploadCustomRequestOptions,
    useDialog, useMessage
} from 'naive-ui'
import { apiUploadFile } from '../interface/file'
import { CreateOutline, Image, Menu, Trash, TrashBinOutline } from '@vicons/ionicons5'
import { apiCreatePlace, apiPlaceInfo, apiPlaceNearby, apiGalleryDelete, apiDeletPlace, apiUpdatePlace } from '../interface/cos'
import { getBrowserKey } from '../store/key'
import GalleryEditor from './GalleryEditor.vue'

const dialog = useDialog()
const message = useMessage()

const showPlace = ref<boolean>(false)
const currentPlace = ref<Place | null>(null)

const newPlaceForm = ref({
    name: '',
    description: '',
    avatar: '',
    lat: 0,
    lng: 0,
})

const uploadAvatar = async ({
    file,
    onFinish,
    onError,
    onProgress
}: UploadCustomRequestOptions) => {
    if (!file.file) {
        return
    }

    let index = 0
    let timer = setInterval(() => {
        onProgress({
            percent: index
        })
        index++
        if (index == 99) {
            clearInterval(timer)
            return
        }
    }, 100)
    const response = await apiUploadFile(file.name, file.file as File)
    clearInterval(timer)
    if (!response.isSuccess()) {
        onError()
        message.error(response.getError())
    } else {
        onFinish()
        newPlaceForm.value.avatar = response.data?.url || ''
    }
}

const creating = ref<boolean>(false)
const createPlace = async () => {
    if (creating.value) {
        message.warning('创建中')
        return
    }
    creating.value = true
    const response = await apiCreatePlace(
        newPlaceForm.value.name, newPlaceForm.value.description, newPlaceForm.value.lat, newPlaceForm.value.lng, getBrowserKey(), newPlaceForm.value.avatar
    )
    creating.value = false

    if (!response.isSuccess()) {
        message.error(response.getError())
    } else {
        message.success('创建成功')
        places.value.push({
            ID: response.data?.id || 0,
            name: newPlaceForm.value.name,
            description: newPlaceForm.value.description,
            avatar: newPlaceForm.value.avatar,
            point: {
                x: newPlaceForm.value.lng,
                y: newPlaceForm.value.lat,
            },
            CreatedAt: new Date().toLocaleString(),
            UpdatedAt: new Date().toLocaleString(),
            DeletedAt: '',
        })
        currentPlace.value = places.value[places.value.length - 1]
    }
}

const updatePlace = () => {
    dialog.create({
        style: {
            width: '800px'
        },
        title: '编辑',
        closable: true,
        content: () => [
            h(NH6, { prefix: 'bar' }, {
                default: () => [
                    h('div', { class: 'text-13 w100' }, '这里是哪'),
                    h(NInput, {
                        value: currentPlace.value?.name,
                        maxlength: 64,
                        showCount: true,
                        onUpdateValue: (v: string) => {
                            currentPlace.value!.name = v
                        }
                    }),
                ]
            }),
            h(NH6, { prefix: 'bar' }, {
                default: () => [
                    h('div', { class: 'text-13 w100' }, '简单描述一下'),
                    h(NInput, {
                        value: currentPlace.value?.description,
                        maxlength: 1024,
                        type: 'textarea',
                        showCount: true,
                        onUpdateValue: (v: string) => {
                            currentPlace.value!.description = v
                        }
                    }),
                ]
            }),
            h(NH6, { prefix: 'bar' }, {
                default: () => [
                    h('div', { class: 'text-13 w100' }, '上传一张封面吧'),
                    h(NUpload, {
                        directoryDnd: true,
                        max: 5,
                        customRequest: uploadAvatar,
                    }, {
                        default: () => [
                            h(NUploadDragger, {
                                style: {
                                    height: '200px',
                                }
                            }, {
                                default: () => [
                                    h(NIcon, {
                                        size: 48,
                                        depth: 3,
                                    }, {
                                        default: () => h(Image)
                                    }),
                                    h(NText, {
                                        style: {
                                            fontSize: '16px',
                                        }
                                    }, {
                                        default: () => '点击或者拖动文件到该区域来上传'
                                    }),
                                    h(NP, {
                                        depth: 3,
                                        style: {
                                            margin: '8px 0 0 0',
                                        }
                                    }, {
                                        default: () => '请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码'
                                    }),
                                ]
                            })
                        ]
                    }),
                ]
            }),
        ],
        positiveText: '更新',
        onPositiveClick: async () => {
            currentPlace.value!.avatar = newPlaceForm.value.avatar
            const response = await apiUpdatePlace(
                currentPlace.value!.ID, currentPlace.value!.name, currentPlace.value!.description, 
                currentPlace.value?.point.y || 0, currentPlace.value?.point.x || 0, 
                getBrowserKey(), newPlaceForm.value.avatar
            )

            if (!response.isSuccess()) {
                message.error(response.getError())
            } else {
                message.success('更新成功')
                const index = places.value.findIndex(v => v.ID === currentPlace.value?.ID)
                if (index !== undefined && index !== -1) {
                    places.value.splice(index, 1, currentPlace.value!)
                }
                currentPlace.value = places.value[index]
            }
        }
    })
}

const deletePlace = async (id: number) => {
    const response = await apiDeletPlace(id, getBrowserKey())
    if (!response.isSuccess()) {
        message.error(response.getError())
    } else {
        message.success('删除成功')
        const index = places.value.findIndex(v => v.ID === id)
        if (index !== undefined && index !== -1) {
            places.value.splice(index, 1)
        }
        showPlace.value = false
    }
}

const fetchNearbyPlaces = async (lat: number, lng: number) => {
    const response = await apiPlaceNearby(lat, lng)
    if (!response.isSuccess()) {
        message.error(response.getError())
    } else {
        for (const place of response.data?.places || []) {
            // check if place already exists
            if (places.value.find(v => v.ID === place.ID)) {
                continue
            }
            places.value.push(place)
        }
    }
}

const places = ref<Place[]>([])

const isClickPlace = ref<boolean>(false)
const points = computed(() => places.value.map((v, k) => {
    return {
        lng: v.point.x,
        lat: v.point.y,
        click: () => {
            isClickPlace.value = true
            showPlace.value = true
            setTimeout(async () => {
                currentPlace.value = v
                // check if place is already fetched
                if (!v.galleries) {
                    const response = await apiPlaceInfo(v.ID)
                    if (!response.isSuccess()) {
                        message.error(response.getError())
                    } else {
                        places.value[k] = response.data?.place || v
                        currentPlace.value = places.value[k]
                    }
                }
                isClickPlace.value = false
            }, 100)
        },
    }
}))

const handleMapClick = (e: {
    lng: number, lat: number
}) => {
    if (isClickPlace.value) {
        isClickPlace.value = false
        return
    }
    newPlaceForm.value.lat = e.lat
    newPlaceForm.value.lng = e.lng
    currentPlace.value = null
    showPlace.value = true
}

const handleMapLocate = (e: {
    lng: number, lat: number
}) => {
    fetchNearbyPlaces(e.lat, e.lng)
}

const createGallery = () => {
    dialog.create({
        style: {
            width: '1000px'
        },
        title: '编辑',
        closable: true,
        content: () => h(GalleryEditor, {
            galleryId: 0,
            placeId: currentPlace.value?.ID || 0,
            onCreated: (gallery: Gallery) => {
                // check if galleries exists
                if (!currentPlace.value?.galleries) {
                    currentPlace.value!.galleries = []
                }
                currentPlace.value?.galleries?.push(gallery)
            },
            onUpdated: (gallery: Gallery) => {
                const index = currentPlace.value?.galleries?.findIndex(v => v.ID === gallery.ID)
                if (index !== undefined && index !== -1) {
                    currentPlace.value?.galleries?.splice(index, 1, gallery)
                }
            },
        }),
    })
}

const updateGallery = async (gallery: Gallery) => {
    dialog.create({
        style: {
            width: '1000px'
        },
        title: '编辑',
        closable: true,
        content: () => h(GalleryEditor, {
            galleryId: gallery.ID,
            placeId: currentPlace.value?.ID || 0,
            gallery: gallery,
            onUpdated: (gallery: Gallery) => {
                const index = currentPlace.value?.galleries?.findIndex(v => v.ID === gallery.ID)
                if (index !== undefined && index !== -1) {
                    currentPlace.value?.galleries?.splice(index, 1, gallery)
                }
            },
        }),
    })
}

const deleteGallery = async (gallery: Gallery) => {
    const response = await apiGalleryDelete(gallery.ID, getBrowserKey())
    if (!response.isSuccess()) {
        message.error(response.getError())
    } else {
        message.success('删除成功')
        const index = currentPlace.value?.galleries?.findIndex(v => v.ID === gallery.ID)
        if (index !== undefined && index !== -1) {
            currentPlace.value?.galleries?.splice(index, 1)
        }
    }
}

</script>

<style scoped>
.c-v-map {
    min-height: calc(100vh - 128px);
    width: 100%;
}

.map-container {
    height: calc(100vh - 128px);
}
</style>