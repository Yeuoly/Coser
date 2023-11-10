<template>
    <NLayoutContent class="m-l-body">
        <RouterView v-slot="{ Component }">
            <Transition name="slide">
                <component :is="Component" />
            </Transition>
        </RouterView>
    </NLayoutContent>
</template>

<script setup lang="ts">
import { NLayoutContent } from 'naive-ui'
import { Transition } from 'vue'
</script>

<style scoped>
.m-l-body {
    min-height: 85vh;
}

.slide-enter-active {
  animation: slide-in 0.5s;
}
.slide-leave-active {
  animation: slide-out 0.5s;
}

@keyframes slide-in {
  0% {
    opacity: 0;
    transform: translateX(-100px);
  }
  60% {
    opacity: 0;
    transform: translateX(-50px);
  }
  100% {
    opacity: 1;
    transform: translateX(0);
  }
}
@keyframes slide-out {
    0% {
        transform: translateX(0);
        opacity: 1;
    }
    50% {
        transform: translateX(50vw);
        opacity: 0;
        position: absolute;
    }
    100% {
        transform: translateX(100vw);
        opacity: 0;
        position: absolute;
    }
}

</style>

