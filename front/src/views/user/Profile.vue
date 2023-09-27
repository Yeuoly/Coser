<template>
    <NH2 prefix="bar">
        <NText type="primary">
            {{ UserName }} {{ $t('user.Profile.672155-0') }}
        </NText>
    </NH2>
    <NBlockquote>
        {{ $t('user.Profile.672155-1') }}{{ UserSign }}
    </NBlockquote>
    <NDivider></NDivider>
    <NH4 prefix="">
        <NText type="primary">{{ $t('user.Profile.157805-0') }}</NText>
        <NForm
            label-placement="left"
            label-width="100"
        >
            <NFormItem :label="$t('user.Profile.157805-1')">
                <NSpace justify="center" class="w100">
                    <NTooltip trigger="hover">
                        <template #trigger>
                            <NAvatar class="clickable" :size="120" :src="userProfileForm.avatar" round></NAvatar>
                        </template>
                        {{ $t('user.Profile.157805-2') }}
                    </NTooltip>
                </NSpace>
            </NFormItem>
            <NFormItem :label="$t('user.Profile.157805-3')">
                <NInput v-model:value="userProfileForm.username"></NInput>
            </NFormItem>
            <NFormItem :label="$t('user.Profile.157805-4')">
                <NInput v-model:value="userProfileForm.sign"></NInput>
            </NFormItem>
        </NForm>
    </NH4>
    <NDivider></NDivider>
</template>

<script setup lang="ts">
import { NAvatar, NBlockquote, NDivider, NForm, NFormItem, NH2, NH4, NInput, NLi, NSpace, NText, NTooltip, NUl, useLoadingBar, useMessage } from 'naive-ui'
import { useUserStore } from '../../store/user'
import { storeToRefs } from 'pinia'
import { onMounted, ref, watch } from 'vue'
import { apiUserProfile } from '../../interface/user'
import { copyToClipboard } from '../../utils/copy'
import { useI18n } from 'vue-i18n'

const { t: $t } = useI18n()

const message = useMessage()

const { UserName, UserSign, UserId, UserAvatar } = storeToRefs(useUserStore())

const userProfileForm = ref({
    username: UserName.value,
    sign: UserSign.value,
    avatar: UserAvatar.value,
})

watch(UserId, () => {
    loadProfile()
})


const loadProfile = async () => {
    const response = await apiUserProfile()
    if (!response.isSuccess()) {
        message.error(response.getError())
    } else {
    }
}

onMounted(async () => {
    const loadingBar = useLoadingBar()
    loadingBar.start()
    await loadProfile()
    loadingBar.finish()
})

</script>