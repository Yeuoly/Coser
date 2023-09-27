import { useUserStore } from '../../store/user'
import { apiUserCheck } from '../../interface/user'

import { $t } from '../../locale/index'

export const refreshLoginStatus = async () => {
    const userStore = useUserStore()
    const response = await apiUserCheck<{
        username: string,
        role: string,
        uid: number,
    }>()

    if(response.isSuccess()) {
        userStore.setUserName(response.data?.username || $t('auth.login.596980-0'))
        userStore.setUserRole(response.data?.role || $t('auth.login.596980-1'))
        userStore.setUserId(response.data?.uid || 0)
        userStore.setUserOnline(true)
    } else {
        userStore.setUserName($t('auth.login.596980-0'))
        userStore.setUserRole($t('auth.login.596980-1'))
        userStore.setUserId(0)
        userStore.setUserOnline(false)
    }
}
