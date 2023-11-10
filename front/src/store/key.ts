export const getBrowserKey = () => {
    const key = localStorage.getItem('browser-key')
    if (key) {
        return key
    }

    const newKey = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15)
    localStorage.setItem('browser-key', newKey)
    return newKey
}
