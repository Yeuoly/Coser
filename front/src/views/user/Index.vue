<template>
    <div class="k-v-admin mx-auto">
        <NLayout has-sider>
            <NLayoutSider
                bordered
                :collapsed-width="64"
                :width="240"
                :native-scrollbar="false"
            >
                <NMenu
                    :collapsed-width="64"
                    :collapsed-icon-size="22"
                    :options="menuOptions"
                    @update:value="handleSelect"
                />
            </NLayoutSider>
            <NLayoutContent class="px5">
                <router-view />
            </NLayoutContent>
        </NLayout>
    </div>
</template>

<script setup lang="ts">
import { AtOutline, MaleFemaleOutline } from '@vicons/ionicons5'
import { NIcon, NLayout, NLayoutContent, NLayoutSider, NMenu } from 'naive-ui'
import { Component, h } from 'vue'
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router'

const router = useRouter()
const { t: $t } = useI18n()

const renderIcon = (icon: Component) => {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions = [{
    label: $t('user.Index.201369-0'),
    key: 'user-center',
    icon: renderIcon(MaleFemaleOutline),
    children: [{
        label: $t('user.Index.201369-1'),
        key: 'user-profile',
        icon: renderIcon(AtOutline),
    }]
}]

const handleSelect = (key: string) => {
    switch (key) {
        case 'user-profile':
            router.push('/user/profile')
            break
    }
}

</script>

<style scoped>
.k-v-admin {
    padding: 20px;
    max-width: 1200px;
    min-height: 80vh;
}
</style>