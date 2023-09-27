<template>
    <div class="m-v-direct mx-auto">
        <NSpin
            v-if="state === State.Processing"
            :size="90"
        >
            <template #description>
                {{ $t('views.Redirect.169953-0') }}
            </template>
        </NSpin>
        <NResult
            v-else-if="state === State.Success"
            status="success"
            :title="$t('views.Redirect.872758-0')"
            :description="$t('views.Redirect.169953-2')"
        >
            <template #footer>
                {{ message }}
            </template>
        </NResult>
        <NResult
            v-else-if="state === State.Fail"
            status="error"
            :title="$t('views.Redirect.169953-3')"
            :description="$t('views.Redirect.169953-4')"
        >
            <template #footer>
                {{ message }}
            </template>
        </NResult>
    </div>
</template>

<script setup lang="ts">
import { NResult, NSpin, useDialog } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { apiAcceptKasumiLogin, apiGithubLogin } from '../interface/login'
import { loginTokenSet } from '../store/login'
import { refreshLoginStatus } from '../middleware/auth/login'
import { useI18n } from 'vue-i18n'
const { t: $t } = useI18n()

const router = useRouter()
const route = useRoute()
const dialog = useDialog()

enum State {
    Processing,
    Success,
    Fail,
}

const state = ref(State.Processing)
const message = ref($t('views.Redirect.169953-5'))

const handleGithubLogin = (code: string) => new Promise<{
    login_token: string,
    uid: number,
}>(async (resolve) => {
    const response = await apiGithubLogin<{
        success: boolean,
        login_token: string,
        uid: number,
    }>(code)
    if(!response.isSuccess()) {
        message.value = response.getError()
        state.value = State.Fail
        throw new Error(response.getError())
    }
    
    if(!response.data?.success) {
        message.value = $t('views.Redirect.169953-6')
        state.value = State.Fail
        throw new Error($t('views.Redirect.169953-6'))
    }

    state.value = State.Success
    resolve(response.data)
})

const handleKasumiLogin = (token: string) => {
    dialog.info({
        title: $t('views.Redirect.381743-0'),
        content: $t('views.Redirect.381743-1'),
        positiveText: $t('views.Redirect.381743-2'),
        negativeText: $t('views.Redirect.381743-3'),
        onPositiveClick: async () => {
            const response = await apiAcceptKasumiLogin(token)
            if(!response.isSuccess()) {
                message.value = response.getError()
                if (response.getError() === 'Unauthorized') {
                    message.value = $t('views.Redirect.222093-0')
                    setTimeout(() => {
                        router.push('/login')
                    }, 500)
                    return
                }
                state.value = State.Fail
                return
            }
            state.value = State.Success
            message.value = $t('views.Redirect.381743-4')
        },
        onNegativeClick: () => {
            state.value = State.Fail
            message.value = $t('views.Redirect.381743-4')
        }
    })
}

const login = async () => {
    const type = route.query.type as string
    let token = {
        login_token: '',
        uid: 0,
    }
    switch(type) {
    case 'github':
        const code = route.query.code as string
        token = await handleGithubLogin(code)
        break
    case 'kasumi':
        const kasumiToken = route.query.token as string
        handleKasumiLogin(kasumiToken)
        return
    default:
        router.push('/')
    }
    if (token.uid != 0) {
        loginTokenSet(token.uid, token.login_token)
        refreshLoginStatus()
        router.push('/')
    }
}

const handlePayStripe = () => {
    const success = route.query.success as string
    if(success === 'true') {
        state.value = State.Success
        message.value = $t('views.Redirect.742147-0')
    } else {
        state.value = State.Fail
        message.value = $t('views.Redirect.742147-1')
    }
}

const pay = async () => {
    const type = route.query.type as string
    switch(type) {
    case 'stripe':
        handlePayStripe()
        break
    default:
        router.push('/')
    }
}

onMounted(() => {
    const method = route.query.method as string
    switch(method) {
    case 'login':
        login()
        break
    case 'pay':
        pay()
        break
    default:
        router.push('/')
    }
})

</script>

<style scoped>
.m-v-direct {
    width: 600px;
    height: 80vh;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
}
</style>