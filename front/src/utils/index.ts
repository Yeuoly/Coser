export const getAssetsFile = (url: string) => {
    return new URL(`../assets/${url}`, import.meta.url).href
}

let queue: any[] = []
/**
 * request idle callback, every function will be executed after last timeout
 */
export const requestIdleCallback = async (fn: any, timeout = 1000) => {
    const sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))
    const handler = async () => {
        while(queue.length) {
            const { callback, timeout } = queue[0]
            callback()
            await sleep(timeout)
            queue.shift()
        }
    }
    queue.push({
        callback: fn,
        timeout
    })
    if (queue.length === 1) {
        handler()
    }
}