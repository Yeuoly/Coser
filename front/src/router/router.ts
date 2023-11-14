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
        path: '/search',
        name: 'Search',
        component: () => import('../views/Search.vue')
    },
    {
        path: '/detail/:gid',
        name: 'Detail',
        component: () => import('../views/Detail.vue')
    },
    {
        path: '/galleries/my',
        name: 'MyGalleries',
        component: () => import('../views/MyGalleries.vue')
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