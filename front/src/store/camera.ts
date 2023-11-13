const cameras = [] as string[]
const lens = [] as string[]

export const getCommonlyUsedCameras = () => {
    if (cameras.length === 0) {
        return localStorage.getItem('commonlyUsedCameras')?.split(',') || []
    }

    return cameras
}

export const addCommonlyUsedCamera = (camera: string) => {
    const cameras = getCommonlyUsedCameras()

    if (cameras.includes(camera)) {
        return
    }

    cameras.push(camera)

    localStorage.setItem('commonlyUsedCameras', cameras.join(','))

    return cameras
}

export const removeCommonlyUsedCamera = (camera: string) => {
    const cameras = getCommonlyUsedCameras()

    if (!cameras.includes(camera)) {
        return
    }

    const index = cameras.indexOf(camera)

    cameras.splice(index, 1)

    localStorage.setItem('commonlyUsedCameras', cameras.join(','))

    return cameras
}

export const getCommonlyUsedLens = () => {
    if (lens.length === 0) {
        return localStorage.getItem('commonlyUsedLens')?.split(',') || []
    }

    return lens
}

export const addCommonlyUsedLens = (lens: string) => {
    const lensList = getCommonlyUsedLens()

    if (lensList.includes(lens)) {
        return
    }

    lensList.push(lens)

    localStorage.setItem('commonlyUsedLens', lensList.join(','))

    return lensList
}

export const removeCommonlyUsedLens = (lens: string) => {
    const lensList = getCommonlyUsedLens()

    if (!lensList.includes(lens)) {
        return
    }

    const index = lensList.indexOf(lens)

    lensList.splice(index, 1)

    localStorage.setItem('commonlyUsedLens', lensList.join(','))

    return lensList
}