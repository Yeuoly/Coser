<template>
    <div class="c-v-detail max-container default-container">
        <div class="mt5">
            <NH3 prefix="bar">
                <NText type="primary">
                    {{ gallery?.name }}
                </NText>
            </NH3>
            <NTag type="primary" :bordered="false">
                <template #icon>
                    <NIcon :component="Time"></NIcon>
                </template>
                {{ new Date(gallery?.CreatedAt || 0).toLocaleString() }}
            </NTag>
            <NTag type="info" :bordered="false" v-if="gallery?.photographers">
                <template #icon>
                    <NIcon :component="Camera"></NIcon>
                </template>
                摄影师: {{ gallery.photographers }}
            </NTag>
            <NTag type="info" :bordered="false" v-if="gallery?.cosers">
                <template #icon>
                    <NIcon :component="Woman"></NIcon>
                </template>
                Coser: {{ gallery.cosers }}
            </NTag>
            <NTag type="info" :bordered="false" v-if="gallery?.character">
                <template #icon>
                    <NIcon :component="Man"></NIcon>
                </template>
                角色: {{ gallery.character }}
            </NTag>
            <NTag type="info" :bordered="false">
                <template #icon>
                    <NIcon :component="List"></NIcon>
                </template>
                共{{ gallery?.images.length }}张
            </NTag>
            <NTag class="clickable" type="warning" :bordered="false" @click="toPlace(gallery!)">
                <template #icon>
                    <NIcon :component="Locate"></NIcon>
                </template>
                {{ gallery?.place.name }}
            </NTag>
            <NTag v-for="camera in uniqueCameras" type="info" :bordered="false">
                <template #icon>
                    <NIcon :component="Camera"></NIcon>
                </template>
                {{ camera }}
            </NTag>
            <div class="mt5">
                关于拍摄地点：{{ gallery?.place.description }} <br>
                <span v-if="gallery?.description">
                    关于这套片：{{ gallery?.description }} <br>
                </span>
                坐标：{{ gallery?.place.point.x }}, {{ gallery?.place.point.y }}
            </div>
            <NH3 prefix="bar">
                <NText type="primary">
                    正片
                </NText>
            </NH3>
            <div>
                {{ gallery?.description }}
            </div>
            <NButton :type="showRaw ? 'error' : 'primary'" @click="showRaw = !showRaw" class="mt5 mb5">
                {{ showRaw ? '显示缩略图' : '显示原图' }}
            </NButton>
            <NDivider></NDivider>
            <div v-if="verticals.length > 0">
                <NH4 prefix="bar">
                    <NText type="primary">
                        竖构图(横向滚动)
                        {{ verticals.length }} 张
                    </NText>
                </NH4>
                <div class="verticals">
                    <NTooltip v-for="image in verticals">
                        <template #trigger>
                            <NImage class="mr5 mb5" lazy  :src="getImageUrl('vertical', image.url)" width="400" object-fit="cover"></NImage>
                        </template>
                        <template #default>
                            <span v-if="image.camera">
                                相机：{{ image.camera }} <br>
                            </span>
                            <span v-if="image.lens">
                                镜头：{{ image.lens }} <br>
                            </span>
                            尺寸：{{ image.width }} x {{ image.height }} <br>
                            <span v-if="parseInt(image.exposure_time)">
                                曝光时间：1/{{ image.exposure_time }}s <br>
                            </span>
                            <span v-if="parseInt(image.aperature)">
                                光圈：f{{ image.aperature }} <br>
                            </span>
                            <span v-if="parseInt(image.iso)">
                                ISO：{{ image.iso }} <br>
                            </span>
                            <span v-if="parseInt(image.focal_length)">
                                焦距：{{ image.focal_length }}mm <br>
                            </span>
                        </template>
                    </NTooltip>
                </div>
            </div>
            <div v-if="horizontals.length > 0">
                <NH4 prefix="bar">
                    <NText type="primary">
                        横构图
                        {{ horizontals.length }} 张
                    </NText>
                </NH4>
                <div class="horizontals">
                    <NTooltip v-for="image in horizontals">
                        <template #trigger>
                            <NImage class="mr5 mb5" lazy :src="getImageUrl('horizontal', image.url)" height="300" object-fit="contain"></NImage>
                        </template>
                        <template #default>
                            <span v-if="image.camera">
                                相机：{{ image.camera }} <br>
                            </span>
                            <span v-if="image.lens">
                                镜头：{{ image.lens }} <br>
                            </span>
                            尺寸：{{ image.width }} x {{ image.height }} <br>
                            <span v-if="parseInt(image.exposure_time)">
                                曝光时间：1/{{ image.exposure_time }}s <br>
                            </span>
                            <span v-if="parseInt(image.aperature)">
                                光圈：f{{ image.aperature }} <br>
                            </span>
                            <span v-if="parseInt(image.iso)">
                                ISO：{{ image.iso }} <br>
                            </span>
                            <span v-if="parseInt(image.focal_length)">
                                焦距：{{ image.focal_length }}mm <br>
                            </span>
                        </template>
                    </NTooltip>
                </div>
            </div>
            <div v-if="diagonals.length > 0">
                <NH4 prefix="bar">
                    <NText type="primary">
                        方构图（横向滚动）
                        {{ diagonals.length }} 张
                    </NText>
                </NH4>
                <div class="diagonals">
                    <NTooltip v-for="image in diagonals">
                        <template #trigger>
                            <NImage class="mr5 mb5" lazy :src="getImageUrl('diagonals', image.url)" width="400" object-fit="cover"></NImage>
                        </template>
                        <template #default>
                            <span v-if="image.camera">
                                相机：{{ image.camera }} <br>
                            </span>
                            <span v-if="image.lens">
                                镜头：{{ image.lens }} <br>
                            </span>
                            尺寸：{{ image.width }} x {{ image.height }} <br>
                            <span v-if="parseInt(image.exposure_time)">
                                曝光时间：1/{{ image.exposure_time }}s <br>
                            </span>
                            <span v-if="parseInt(image.aperature)">
                                光圈：f{{ image.aperature }} <br>
                            </span>
                            <span v-if="parseInt(image.iso)">
                                ISO：{{ image.iso }} <br>
                            </span>
                            <span v-if="parseInt(image.focal_length)">
                                焦距：{{ image.focal_length }}mm <br>
                            </span>
                        </template>
                    </NTooltip>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { NButton, NDivider, NH3, NH4, NIcon, NImage, NTag, NText, NTooltip, useMessage } from 'naive-ui'
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Gallery } from '../interface/types'
import { apiGalleryInfo } from '../interface/cos'
import { Camera, List, Locate, Man, Time, Woman } from '@vicons/ionicons5'

const route = useRoute()
const router = useRouter()
const message = useMessage()

const gid = ref(parseInt(route.params.gid as string))

const showRaw = ref(false)

const getImageUrl = computed(() => (mod: 'vertical' | 'horizontal' | 'diagonals', url: string) => {
    if (showRaw.value) {
        return url
    }
    if (mod === 'vertical') {
        return url + '?x-oss-process=image/resize,mfit,h_1600'
    }
    if (mod === 'horizontal') {
        return url + '?x-oss-process=image/resize,mfit,w_1600'
    }
    if (mod === 'diagonals') {
        return url + '?x-oss-process=image/resize,mfit,h_1600'
    }
    return url
})

const gallery = ref<Gallery>()
const uniqueCameras = computed(() => {
    if (!gallery.value) {
        return []
    }
    const cameras = gallery.value.images.map(v => v.camera)
    return [...new Set(cameras)]
})
const verticals = computed(() => {
    if (!gallery.value) {
        return []
    }
    // if w : h < 0.8, then it's a vertical
    return gallery.value.images.filter(v => v.width / v.height < 0.8)
})
const horizontals = computed(() => {
    if (!gallery.value) {
        return []
    }
    // if w : h > 1.2, then it's a horizontal
    return gallery.value.images.filter(v => v.width / v.height > 1.2 || v.height === 0)
})
const diagonals = computed(() => {
    if (!gallery.value) {
        return []
    }
    // if 0.8 < w : h < 1.2, then it's a diagonal
    return gallery.value.images.filter(v => (v.width / v.height) >= 0.8 && (v.width / v.height) <= 1.2)
})

const loading = ref(false)
const loadGallery = async () => {
    if (loading.value) {
        return
    }
    loading.value = true
    const response = await apiGalleryInfo(gid.value)
    loading.value = false
    if (!response.isSuccess()) {
        message.error(response.getError())
        return
    }

    gallery.value = response.data?.gallery
}

onMounted(() => {
    if (!gid.value) {
        message.error('参数错误')
        router.back()
        return
    }

    loadGallery()
})

const toPlace = (gallery: Gallery) => {
    router.push(`/map?lat=${gallery.place.point.y}&lng=${gallery.place.point.x}`)
}

</script>

<style scoped>
.c-v-detail {
    min-height: 100%;
}

.verticals {
    /** allow horizontal scroll */
    overflow-x: auto;
    white-space: nowrap;
}

.diagonals {
    /** allow horizontal scroll */
    overflow-x: auto;
    white-space: nowrap;
}

</style>