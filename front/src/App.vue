<script setup lang="ts">
import {
    NConfigProvider,
    NLayout,
    NSpace,
    NMessageProvider,
    NDialogProvider,
    NLoadingBarProvider,
    useMessage,
} from 'naive-ui'
import Header from './layout/Header.vue'
import Footer from './layout/Footer.vue'
import Body from './layout/Body.vue'

</script>

<template>
    <NConfigProvider>
        <NDialogProvider>
            <NMessageProvider>
                <NLoadingBarProvider>
                    <NSpace vertical>
                        <NLayout>
                            <Header />
                            <Body />
                            <Footer />
                        </NLayout>
                    </NSpace>
                </NLoadingBarProvider>
            </NMessageProvider>
        </NDialogProvider>
    </NConfigProvider>
</template>

<style>
</style>