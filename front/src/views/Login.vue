<template>
    <div class="login mx-auto">
        <NCard v-if="!isClient">
            <NTabs
                size="large"
                animated
                pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
                pane-wrapper-style="margin: 0 -4px"
                @update:value="handleTabUpdate"
                ref="tabsInstRef"
            >
                <NTabPane :name="$t('views.LogIn.169953-0')">
                    <NForm :model="loginForm" :rules="rules" ref="formRefLogin">
                        <NFormItemRow :label="$t('views.LogIn.169953-1')" path="email">
                            <NInput v-model:value="loginForm.email" :placeholder="$t('views.LogIn.169953-2')" />
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.169953-3')" path="password">
                            <NInput 
                                show-password-on="mousedown" 
                                type="password" 
                                v-model:value="loginForm.password" 
                                :placeholder="$t('views.LogIn.169953-4')" 
                            />
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.621117-0')">
                            <img :src="loginVercode.image" class="clickable" @click="getLoginVercode" style="height: 40px;" />
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.621117-1')">
                            <NInput v-model:value="loginVercode.result" />
                        </NFormItemRow>
                        <NFormItemRow>
                            <NButton block type="primary" @click="login">{{ $t('views.LogIn.169953-0') }}</NButton>
                        </NFormItemRow>
                    </NForm>
                </NTabPane>
                <NTabPane :name="$t('views.LogIn.169953-5')">
                    <NForm :model="regiterForm" :rules="rules" ref="formRefRegister">
                        <NFormItemRow :label="$t('views.LogIn.169953-6')" path="username">
                            <NInput v-model:value="regiterForm.username" :placeholder="$t('views.LogIn.169953-7')" />
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.169953-1')" path="email">
                            <NInput v-model:value="regiterForm.email" :placeholder="$t('views.LogIn.169953-2')" />
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.169953-3')" path="password">
                            <NInput 
                                v-model:value="regiterForm.password" 
                                :placeholder="$t('views.LogIn.169953-4')" 
                                show-password-on="mousedown"
                                type="password"
                            />
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.621117-2')">
                            <img :src="registerVercode.pre_image" class="clickable" @click="getRegisterVercodePre" style="height: 40px;" />
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.621117-3')">
                            <NInputGroup>
                                <NInput v-model:value="registerVercode.pre_result"></NInput>
                                <NButton type="primary" @click="getRegisterVercode">{{ $t('views.LogIn.621117-4') }}</NButton>
                            </NInputGroup>
                        </NFormItemRow>
                        <NFormItemRow :label="$t('views.LogIn.621117-5')">
                            <NInput v-model:value="registerVercode.result" />
                        </NFormItemRow>
                        <NFormItemRow>
                            <NButton block type="primary" @click="register">{{ $t('views.LogIn.169953-5') }}</NButton>
                        </NFormItemRow>
                    </NForm>
                </NTabPane>
            </NTabs>
            <NDivider></NDivider>
            <div class="mb2">{{ $t('views.LogIn.169953-8') }}</div>
            <div>
                <NAvatar class="clickable" :src="getAssetsFile('github.png')" :size="36" @click="githubLogin"></NAvatar>
            </div>
        </NCard>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { FormInst, NAvatar, NBlockquote, NButton, NCard, NDivider, NForm, NFormItemRow, NH2, NIcon, NInput, NInputGroup, NTabPane, NTabs, NText, NTooltip } from 'naive-ui'
import { useRoute, useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { getAssetsFile } from '../utils'
import { useI18n } from 'vue-i18n'
import { apiLogin, apiCheckKasumiLogin } from '../interface/login'
import { apiRegister } from '../interface/register'
import { apiLoginVercode, apiRegisterVercode, apiRegisterVercodePre } from '../interface/vercode'
import { loginTokenSet } from '../store/login'
import { refreshLoginStatus } from '../middleware/auth/login'
import { hashPassword } from '../utils/crypto'

const { t: $t } = useI18n()
const route = useRoute()
const router = useRouter()
const message = useMessage()

const isClient = ref(navigator.userAgent.toLocaleLowerCase().includes('electron'))

const loginForm = ref({
    email: '',
    password: '',
})

const regiterForm = ref({
    email: '',
    username: '',
    password: '',
    vercode: '',
    vercodeToken: '',
})

const loginVercode = ref({
    result: '',
    token: '',
    image: ''
})

const registerVercode = ref({
    pre_token: '',
    pre_result: '',
    pre_image: '',
    result: '',
    token: '',
    invite_code: ''
})

const rules = ref({
    username: {
        validator: (rule: any, value: string) => {
            if (!value) {
                return Error($t('views.LogIn.169953-7'))
            }
            if (value.length < 4) {
                return Error($t('views.LogIn.169953-9'))
            }
            return true
        },
        trigger: 'input'
    },
    email: {
        validator: (rule: any, value: string) => {
            if (!value) {
                return Error($t('views.LogIn.169953-2'))
            }
            if (!/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(value)) {
                return Error($t('views.LogIn.169953-10'))
            }
            return true
        },
        trigger: 'input'
    },
    password: {
        validator: (rule: any, value: string) => {
            if (!value) {
                return Error($t('views.LogIn.169953-4'))
            }
            if (value.length < 6) {
                return Error($t('views.LogIn.169953-11'))
            }
            return true
        },
        trigger: 'input'
    }
})

const handleTabUpdate = (name: string) => {
    if (name === $t('views.LogIn.169953-0')) {
        if (loginVercode.value.token === '') {
            getLoginVercode()
        }
    } else if (name === $t('views.LogIn.169953-5')) {
        if (registerVercode.value.pre_token === '') {
            getRegisterVercodePre()
        }
    }
}

const getFormValidate = (formRef: FormInst) => new Promise((resolve) => {
    try {
        formRef.validate((errros) => {
            if(!errros) {
                resolve(true)
            } else {
                resolve(false)
            }
        }).catch(() => {
            resolve(false)
        })
    } catch (err) {
        resolve(false)
    }
})

const getLoginVercode = async () => {
    const response = await apiLoginVercode()
    if (!response.isSuccess()) {
        message.error(response.getError())
        return
    }
    loginVercode.value.token = response.data?.token as string
    loginVercode.value.image = response.data?.image as string
}

const getRegisterVercodePre = async () => {
    const response = await apiRegisterVercodePre()
    if (!response.isSuccess()) {
        message.error(response.getError())
        return
    }
    registerVercode.value.pre_token = response.data?.token as string
    registerVercode.value.pre_image = response.data?.image as string
}

const getRegisterVercode = async () => {
    const response = await apiRegisterVercode(
        registerVercode.value.pre_token,
        registerVercode.value.pre_result,
        regiterForm.value.email
    )
    if (!response.isSuccess()) {
        message.error(response.getError())
        return
    }
    registerVercode.value.token = response.data?.token as string
    message.success($t('views.LogIn.621117-7'))
}

const githubLogin = () => {
    const redirect_uri = encodeURIComponent(`${(import.meta as any).env.VITE_APP_FRONTEND_BASE_URL}/redirect?method=login&type=github`)
    window.location.href = `https://github.com/login/oauth/authorize?client_id=${(import.meta as any).env.VITE_APP_GITHUB_CLIENT_ID}&redirect_uri=` + redirect_uri
}

const formRefLogin = ref<FormInst | null>(null)
const login = async () => {
    const isValid = await getFormValidate(formRefLogin.value as FormInst)
    if (!isValid) {
        message.error($t('views.LogIn.169953-12'))
    }

    const response = await apiLogin(
        loginForm.value.email,
        hashPassword(loginForm.value.password),
        loginVercode.value.token,
        loginVercode.value.result
    )

    if (!response.isSuccess()) {
        message.error(response.getError())
        return
    }

    if (response.data?.success) {
        message.success($t('views.LogIn.914867-0'))
        loginTokenSet(response.data?.uid as number, response.data?.login_token as string)
        refreshLoginStatus()

        setTimeout(() => {
            router.push('/')
        }, 500)
    } else {
        message.error($t('views.LogIn.621117-8'))
        loginVercode.value.token = response.data?.new_token as string
    }
}

const formRefRegister = ref<FormInst | null>(null)
const register = async () => {
    const isValid = await getFormValidate(formRefRegister.value as FormInst)
    if (!isValid) {
        message.error($t('views.LogIn.169953-12'))
    }

    const response = await apiRegister(
        regiterForm.value.username,
        hashPassword(regiterForm.value.password),
        registerVercode.value.token,
        registerVercode.value.result,
        registerVercode.value.invite_code
    )

    if (!response.isSuccess()) {
        message.error(response.getError())
        return
    }

    if (response.data?.success) {
        message.success($t('views.LogIn.621117-9'))

        setTimeout(() => {
            window.location.href = '/login'
        }, 500)
    } else {
        message.error($t('views.LogIn.621117-10'))
        registerVercode.value.token = response.data?.new_token as string
    }
}

onMounted(() => {
    getLoginVercode()
})

</script>

<style scoped>
.login {
    width: 600px;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>