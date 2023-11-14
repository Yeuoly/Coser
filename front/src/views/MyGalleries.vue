<template>
    <div class="c-v-my-galleries">
        <div class="max-container default-container">
            <div class="result mt5">
                <div v-if="!galleries.length">
                    <NEmpty description="这里空空的"></NEmpty>
                </div>
                <div v-else>
                    <NCard size="small" v-for="gallery in galleries" class="mb2">
                        <NThing content-indented>
                            <template #header>
                                {{ gallery.name }}
                            </template>
                            <template #header-extra>
                                <NButton type="primary" @click="toDetail(gallery)">仔细康康</NButton>
                            </template>
                            <template #avatar v-if="gallery.images && gallery.images.length > 0">
                                <NImage style="border-radius: 12px;" :src="gallery.images[0].url + '?x-oss-process=image/resize,mfit,h_1600'" height="200" width="200"
                                    object-fit="cover"></NImage>
                            </template>
                            <div>
                                {{ gallery.description }}
                            </div>
                            <div class="mt3">
                                <NTag v-if="gallery.photographers.trim()" type="primary" :bordered="false">
                                    <template #icon>
                                        <NIcon :component="Camera"></NIcon>
                                    </template>
                                    摄影师：{{ gallery.photographers }}
                                </NTag>
                                <NTag v-if="gallery.cosers.trim()" type="primary" :bordered="false">
                                    <template #icon>
                                        <NIcon :component="Man"></NIcon>
                                    </template>
                                    Coser：{{ gallery.cosers }}
                                </NTag>
                                <NTag v-if="gallery.character.trim()" type="primary" :bordered="false">
                                    <template #icon>
                                        <NIcon :component="Camera"></NIcon>
                                    </template>
                                    角色{{ gallery.character }}
                                </NTag>
                                <NTag v-if="gallery.place.name" class="clickable" type="warning" :bordered="false"
                                    @click="toPlace(gallery)">
                                    <template #icon>
                                        <NIcon :component="Locate"></NIcon>
                                    </template>
                                    {{ gallery.place.name }}
                                </NTag>
                                <NTag type="info" :bordered="false">
                                    <template #icon>
                                        <NIcon :component="List"></NIcon>
                                    </template>
                                    共{{ gallery.images.length }}张
                                </NTag>
                            </div>
                        </NThing>
                    </NCard>
                </div>
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { apiGalleryGetMy } from '../interface/cos'
import { Gallery } from '../interface/types'
import { getBrowserKey } from '../store/key'
import { NButton, NCard, NEmpty, NIcon, NImage, NTag, NThing, useMessage } from 'naive-ui'
import { useRouter } from 'vue-router'
import { Camera, List, Locate, Man } from '@vicons/ionicons5'

const message = useMessage()
const router = useRouter()

const galleries = ref<Gallery[]>([])

const loadGalleries = async () => {
    const response = await apiGalleryGetMy(getBrowserKey())

    if (!response.isSuccess()) {
        message.error(response.getError())
        return
    }

    galleries.value = response.data?.galleries || []
}

onMounted(() => {
    loadGalleries()
})

const toDetail = (gallery: Gallery) => {
    router.push(`/detail/${gallery.ID}`)
}

const toPlace = (gallery: Gallery) => {
    router.push(`/map?lat=${gallery.place.point.y}&lng=${gallery.place.point.x}`)
}
</script>

<style scoped></style>