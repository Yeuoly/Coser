import { CheckmarkCircleOutline, LogoAndroid, LogoApple, LogoTux, LogoWindows } from "@vicons/ionicons5"

export const getPlatformIcon = (platform: string) => {
    switch(platform.toLocaleLowerCase()) {
        case 'all':
            return CheckmarkCircleOutline
        case 'osx':
        case 'ios':
            return LogoApple
        case 'android':
            return LogoAndroid
        case 'windows':
            return LogoWindows
        case 'linux':
            return LogoTux
        default:
            return CheckmarkCircleOutline
    }
}