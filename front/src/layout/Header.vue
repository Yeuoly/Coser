<template>
    <NLayoutHeader bordered class="m-l-header" id="intro-header">
        <NPageHeader>
            <template #title>
                <NMenu mode="horizontal" :options="mainMenu"></NMenu>
            </template>
        </NPageHeader>
    </NLayoutHeader>
</template>

<script setup lang="ts">
import { NLayoutHeader, NMenu, MenuOption, NPageHeader, NIcon, NAvatar, NText, NDropdown, FormInst, NH2, NInput, NInputGroup, NInputGroupLabel, NButton } from 'naive-ui'
import { computed, h, onMounted, watch, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { getAssetsFile } from '../utils/index'
import { setLocale } from '../locale/index'
import { BatteryChargingOutline, InformationCircleOutline, LanguageOutline, LogInOutline, LogInSharp, LogOutOutline, Map, Search, SearchCircleOutline } from '@vicons/ionicons5'
const { t: $t } = useI18n()
import { useRouter } from 'vue-router'
const router = useRouter()

import { useUserStore } from '../store/user'
import { storeToRefs } from 'pinia'

const { UserOnline } = storeToRefs(useUserStore())

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
            } }, { default: () => 'Coshub' })
        },
        {
            label: () => h('div', {
                onClick: () => {
                    router.push('/map')
                }
            }, '地图'),
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(Map) }),
        },
        {
            label: () => h('div', {
                onClick: () => {
                    router.push('/search')
                }
            }, '综合搜索'),
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(Search) }),
        }
    ]
    return menu
})

const userMenu = computed(() => {
    if(UserOnline.value) {
        return [{
            label: () => $t('layout.Header.111223-2'),
            key: 'profile',
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(InformationCircleOutline) }),
            path: '/user/profile',
        },{
            label: () => $t('layout.Header.776236-5'),
            key: 'logout',
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(LogOutOutline) }),
            path: '/logout',
        }]
    } else {
        return [{
            label: () => $t('layout.Header.776236-6'),
            icon: () => h(NIcon, { clsPrefix: 'ion', size: 20 }, { default: () => h(LogInOutline) }),
            key: 'login',
            path: '/login',
        }]
    }
})

const languages = computed(() => {
    return [{
        label: () => '简体中文',
        key: 'zh-CN',
    }, {
        label: () => 'English',
        key: 'en-US',
    }, {
        label: () => '日本語',
        key: 'ja-JP',
    }]
})

const handleSelectLanguage = (key: string) => {
    setLocale(key)
}

const handleMenuSelect = (key: string) => {
    if(key === 'logout') {
        router.push('/logout')
    } else if(key === 'login') {
        router.push('/login')
    } else if(key === 'profile') {
        router.push('/user/profile')
    }
}

</script>

<style scoped>
.m-l-header {
    padding-top: 15px;
    padding-bottom: 10px;
}
</style>