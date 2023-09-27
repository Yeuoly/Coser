import {
    createRouter,
    createWebHistory,
    RouteRecordRaw
} from 'vue-router'

import userRoutes from './user'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Index',
        component: () => import('../views/Index.vue')
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/Login.vue')
    },
    {
        path: '/logout',
        name: 'Logout',
        component: () => import('../views/Logout.vue')
    },
    {
        path: '/redirect',
        name: 'Redirect',
        component: () => import('../views/Redirect.vue')
    },
    {
        path: '/user',
        name: 'User',
        component: () => import('../views/user/Index.vue'),
        children: userRoutes
    },
    {
        path: '/:catchAll(.*)',
        name: 'NotFound',
        redirect: '/'
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router