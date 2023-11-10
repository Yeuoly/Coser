import {
    createRouter,
    createWebHistory,
    RouteRecordRaw
} from 'vue-router'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/map',
        name: 'Map',
        component: () => import('../views/Map.vue')
    },
    {
        path: '/:catchAll(.*)',
        name: 'NotFound',
        redirect: '/map'
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router