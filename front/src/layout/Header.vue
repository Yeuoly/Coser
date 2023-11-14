<template>
    <NLayoutHeader bordered class="m-l-header" id="intro-header">
        <NPageHeader>
            <template #title>
                <NMenu @update:value="handleSelectMenu" mode="horizontal" :options="mainMenu"></NMenu>
            </template>
        </NPageHeader>
    </NLayoutHeader>
</template>

<script setup lang="ts">
import { NLayoutHeader, NMenu, MenuOption, NPageHeader, NIcon, NH2 } from 'naive-ui'
import { computed, h } from 'vue'
import { getAssetsFile } from '../utils/index'
import { List, Locate, Map, Search } from '@vicons/ionicons5'
import { useRouter } from 'vue-router'
const router = useRouter()

const mainMenu = computed((): MenuOption[] => {
    const menu: MenuOption[] = [
        {
            label: () => h('img', {
                src: getAssetsFile('favicon.svg'),
                alt: 'logo',
                style: 'height: 80px;'
            }),
            key: 'logo',
            onClick: () => {
                router.push('/')
            }
        },
        {
            label: () => h(NH2, { style: {
                marginTop: '20px',
            } }, { default: () => 'Coshub' }),
            key: 'title',
        },
        {
            label: () => h('div', '地图'),
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(Map) }),
            key: 'map',
        },
        {
            label: () => h('div', '我标记的'),
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(Locate) }),
            key: 'my-galleries',
        },
        {
            label: () => h('div', '综合搜索'),
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(Search) }),
            key: 'search',
        }
    ]
    return menu
})

const handleSelectMenu = (key: string) => {
    switch (key) {
        case 'map':
            router.push('/map')
            break
        case 'my-galleries':
            router.push('/galleries/my')
            break
        case 'search':
            router.push('/search')
            break
    }
}



</script>

<style scoped>
.m-l-header {
    padding-top: 15px;
    padding-bottom: 10px;
}
</style>