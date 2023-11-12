<template>
    <div class="c-c-camera-editor">
        <NInputGroup>
            <NAutoComplete v-model:value="camera.camera" placeholder="相机"></NAutoComplete>
            <NAutoComplete v-model:value="camera.lens" placeholder="镜头"></NAutoComplete>
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

import { EXIF } from '../utils/camera'

const camera = defineModel<EXIF>('camera', {
    required: true
})


</script>

<style scoped>

</style>