import { createI18n } from 'vue-i18n'

// @ts-ignore
const modules = import.meta.glob('../i18n/temp/*.json', { eager: true, import: 'default' }) as Record<string, { [key: string]: any }>

/**
 * modules: {
 *    'xxx1': { 'cn': { ... }, 'en': { ... } },
 *   'xxijl2': { 'cn': { ... }, 'en': { ... } },
 *  ...
 * }
 * 
 * remove top key and merge all locales into one object like
 * {
 *  'cn': { ... },
 * 'en': { ... },
 * ...
 * }
 */
const messages = Object.keys(modules).reduce((messages, key) => {
    const locales = modules[key]
    Object.keys(locales).forEach(locale => {
        if (!messages[locale]) {
            messages[locale] = {}
        }
        Object.assign(messages[locale], locales[locale])
    })
    return messages
}, {} as Record<string, { [key: string]: any }>)

// try to get browser language
const browserLang = localStorage.getItem('locale') || navigator.language
let userLang = ''
if (browserLang) {
    for (const lang in messages) {
        if (browserLang.includes(lang)) {
            userLang = lang
        }
    }
}

// return embeded translated text
export const $t = (key: string) => {
    // @ts-ignore
    return i18n.global.t(key)
}

// select translated text
export const $td = (values: {
    [lang: string]: string
}) => {
    const text = values[userLang || 'en'] as string
    if (!text) {
        return values['en']
    }
    return text
}

export const setLocale = (locale: string) => {
    localStorage.setItem('locale', locale)
    window.location.reload()
}

export const i18n = createI18n({
    legacy: false,
    locale: userLang || 'en',
    messages
})