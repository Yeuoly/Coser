
<template>
    <div class="c-v-map">
        <div class="map-container">
            <BMap @onClick="handleMapClick" @onLocate="handleMapLocate" :points="points" :enable-modify="false" :enable-add-center-marker="false"
                enable-search mode="multi"></BMap>
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
                        <NImage :src="newPlaceForm.avatar" fit="cover" style="width: 100%; height: 100%;"></NImage>
                    </div>
                </NH6>
                <NH6 prefix="bar" class="w100">
                    <NButton block type="primary" @click="createPlace">创建</NButton>
                </NH6>
            </div>
            <div v-else>
                <NH3 prefix="bar">
                    <NText type="primary">{{ currentPlace.name }}</NText>
                </NH3>
                <NThing>
                    <img :src="currentPlace.avatar" style="width: 100%; height: 300px; object-fit: contain;" />
                    <template #footer>
                        <div class="text-13">
                            {{ currentPlace.description }}
                        </div>
                    </template>
                    <template #header-extra>
                        创建于{{ new Date(currentPlace.CreatedAt).toLocaleString() }}
                    </template>
                </NThing>
                <NDivider></NDivider>
                <NThing>
                    <template #header>
                        正片列表
                    </template>
                    <template #header-extra>
                        <NButton type="primary">
                            <template #icon>
                                <NIcon size="20" :component="Image"></NIcon>
                            </template>
                            <span>创建正片</span>
                        </NButton>
                    </template>
                </NThing>
            </div>
        </NDrawerContent>
    </NDrawer>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import BMap from '../components/BMap.vue'
import { Place } from '../interface/types'
import {
    NButton, NDivider, NDrawer, NDrawerContent, NH3, NH6,
    NIcon, NImage, NInput, NP, NResult, NText,
    NThing,
    NUpload, NUploadDragger, UploadCustomRequestOptions,
    useDialog, useMessage
} from 'naive-ui'
import { apiUploadFile } from '../interface/file'
import { Image } from '@vicons/ionicons5'
import { apiCreatePlace, apiPlaceNearby } from '../interface/cos'
import { getBrowserKey } from '../store/key'

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

    onProgress({ percent: 0 })
    const response = await apiUploadFile(file.name, file.file as File)
    onProgress({ percent: 100 })
    if (!response.isSuccess()) {
        onError()
        message.error(response.getError())
    } else {
        onFinish()
        newPlaceForm.value.avatar = response.data?.url || ''
    }
}

const createPlace = async () => {
    const response = await apiCreatePlace(
        newPlaceForm.value.name, newPlaceForm.value.description, newPlaceForm.value.lat, newPlaceForm.value.lng, getBrowserKey(), newPlaceForm.value.avatar
    )

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
const points = computed(() => places.value.map(v => {
    return {
        lng: v.point.x,
        lat: v.point.y,
        blue: false,
        click: () => {
            isClickPlace.value = true
            setTimeout(() => {
                currentPlace.value = v
                showPlace.value = true
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