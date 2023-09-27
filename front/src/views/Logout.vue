<template>
    <div></div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { refreshLoginStatus } from '../middleware/auth/login'
import { loginTokenSet } from '../store/login'
import { useRouter } from 'vue-router'

const router = useRouter()

onMounted(() => {
    loginTokenSet(0, '')
    refreshLoginStatus()
    router.push({ path: '/'})
})
</script>

<style scoped>

</style>