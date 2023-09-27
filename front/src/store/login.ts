export const loginTokenSet = (uid: number, token: string) => {
    localStorage.setItem('uid', uid.toString())
    localStorage.setItem('token', token)
}

export const loginTokenGet = () => {
    return {
        uid: parseInt(localStorage.getItem('uid') as string),
        token: localStorage.getItem('token')
    }
}