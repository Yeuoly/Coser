<template>
    <div class="c-c-camera-editor">
        <NInputGroup>
            <NAutoComplete :options="cameraOptions" v-model:value="camera.camera" placeholder="相机（使用英文会有自动补全）"></NAutoComplete>
            <NAutoComplete :options="lensOptions" v-model:value="camera.lens" placeholder="镜头（使用英文会有自动补全）"></NAutoComplete>
        </NInputGroup>
        <NInputGroup>
            <NInputNumber v-model:value="camera.focal" placeholder="焦段">
                <template #suffix>
                    mm
                </template>
            </NInputNumber>
            <NInputNumber v-model:value="camera.apertureValue" placeholder="光圈">
                <template #prefix>
                    f
                </template>
            </NInputNumber>
            <NInputNumber v-model:value="camera.exposureTime" placeholder="曝光时间" :show-button="false">
                <template #prefix>
                    1/
                </template>
                <template #suffix>
                    秒
                </template>
            </NInputNumber>
            <NInputNumber v-model:value="camera.iso" placeholder="曝光时间" :show-button="false">
                <template #prefix>
                    ISO
                </template>
            </NInputNumber>
        </NInputGroup>
    </div>
</template>

<script setup lang="ts">
import { NAutoComplete, NInput, NInputGroup, NInputNumber, NSelect } from 'naive-ui'

import { EXIF, cameras, lens } from '../utils/camera'
import { getCommonlyUsedCameras, getCommonlyUsedLens } from '../store/camera'
import { computed } from 'vue'

const cameraOptions = computed(() => {
    // check if camera is in commonly used cameras
    const commonlyUsedCameras = getCommonlyUsedCameras()
    const cameraInCommonlyUsedCameras = commonlyUsedCameras.filter(
        v => v.toLowerCase().includes(camera.value.camera.toLowerCase())
    )
    if (cameraInCommonlyUsedCameras.length > 0) {
        return cameraInCommonlyUsedCameras.map(v => ({ label: v, value: v }))
    }
    return cameras.filter(
        v => v.toLowerCase().includes(camera.value.camera.toLowerCase())
    ).map(v => ({ label: v, value: v }))
})

const lensOptions = computed(() => {
    // check if lens is in commonly used lens
    const commonlyUsedLens = getCommonlyUsedLens()
    const lensInCommonlyUsedLens = commonlyUsedLens.filter(
        v => v.toLowerCase().includes(camera.value.lens.toLowerCase())
    )
    if (lensInCommonlyUsedLens.length > 0) {
        return lensInCommonlyUsedLens.map(v => ({ label: v, value: v }))
    }
    return lens.filter(
        v => v.toLowerCase().includes(camera.value.lens.toLowerCase())
    ).map(v => ({ label: v, value: v }))
})

const camera = defineModel<EXIF>('camera', {
    required: true
})


</script>

<style scoped>

</style>