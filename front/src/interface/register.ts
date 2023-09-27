import { RequestMethods, apiBase } from "./base"

export const apiRegister = (username: string, password: string, vercode_token: string, vercode: string, invite_code: string) => apiBase<{
    success: boolean,
    new_token: string,
}>('/v1/user/reg', RequestMethods.POST, {
    username, password, vercode_token, vercode, invite_code
})