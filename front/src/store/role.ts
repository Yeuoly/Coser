export const getDefaultCoser = () => {
    return new Promise<string>((resolve) => {
        const coser = localStorage.getItem('default-coser')
        if (coser) {
            resolve(coser)
            return
        }
        
        resolve('')
    })
}

export const setDefaultCoser = (coser: string) => {
    localStorage.setItem('default-coser', coser)
}

export const getDefaultPhotographer = () => {
    return new Promise<string>((resolve) => {
        const photographer = localStorage.getItem('default-photographer')
        if (photographer) {
            resolve(photographer)
            return
        }
        
        resolve('')
    })
}

export const setDefaultPhotographer = (photographer: string) => {
    localStorage.setItem('default-photographer', photographer)
}
