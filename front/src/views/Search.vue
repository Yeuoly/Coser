<template>
    <div class="c-v-search mt5">
        <div class="max-container default-container">
            <NInputGroup>
                <NInput v-model:value="keyword" size="large" placeholder="关键词" ref="input" @keypress.enter="search">
                </NInput>
                <NButton size="large" type="primary" @click="search">
                    搜索
                </NButton>
            </NInputGroup>
            <div class="result mt5">
                <div v-if="searching">
                    <NSpin class="w100" style="height: 200px;"></NSpin>
                </div>
                <div v-else-if="!galleries.length">
                    <NEmpty description="这里空空的"></NEmpty>
                </div>
                <div v-else>
                    <NCard size="small" v-for="gallery in galleries" class="mb2">
                        <NThing content-indented>
                            <template #header>
                                {{ gallery.name }}
                            </template>
                            <template #header-extra>
                                <NButton type="primary">仔细康康</NButton>
                            </template>
                            <template #avatar v-if="gallery.images && gallery.images.length > 0">
                                <NImage style="border-radius: 12px;" :src="gallery.images[0].url" height="200" width="200" object-fit="cover"></NImage>
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
import { NButton, NCard, NEmpty, NIcon, NImage, NInput, NInputGroup, NSpin, NTag, NThing, useMessage } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { Gallery } from '../interface/types'
import { apiGallerySearch } from '../interface/cos'
import { Camera, List, Man } from '@vicons/ionicons5';

const message = useMessage()

const input = ref<HTMLElement | null>(null)
const keyword = ref('')

const galleries = ref<Gallery[]>([])
const searching = ref(false)
onMounted(() => {
    input.value?.focus()
})

const search = async () => {
    if (keyword.value && !searching.value) {
        searching.value = true
        const response = await apiGallerySearch(keyword.value)
        searching.value = false
        if (!response.isSuccess()) {
            message.error(response.getError())
            return
        }

        galleries.value = response.data?.galleries || []
    }
}
</script>

<style scoped></style>