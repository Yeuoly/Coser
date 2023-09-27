import { getAssetsFile } from '../utils'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
    state : () => ({
        _Id : 0,
        _Username : '游客',
        _Avatar: getAssetsFile('long.gif'),
        _Online : false,
        _Role: '未登录',
        _Sign: '这个人很懒，什么都没有留下'
    }),
    getters : {
        UserId(state) {
            return state._Id
        },
        UserName(state) {
            return state._Username
        },
        UserAvatar(state) {
            return state._Avatar
        },
        UserOnline(state) {
            return state._Online
        },
        UserRole(state) {
            return state._Role
        },
        UserSign(state) {
            return state._Sign
        },
        IsAdmin(state) {
            return state._Role === 'admin'
        }
    },
    actions : {
        setUserId(id : number) {
            this._Id = id
        },
        setUserName(name : string) {
            this._Username = name
        },
        setUserAvatar(avatar : string) {
            this._Avatar = avatar
        },
        setUserOnline(online : boolean) {
            this._Online = online
        },
        setUserRole(role : string) {
            this._Role = role
        },
        setUserSign(sign : string) {
            this._Sign = sign
        }
    }
})