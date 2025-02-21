<script setup>
import { defineProps, defineEmits } from 'vue'
import { WindowMinimise, WindowFullscreen, WindowIsFullscreen, WindowUnfullscreen, Quit } from '@/wailsjs/runtime'

const minimizeWindow = async () => {
  await WindowMinimise()
}

const maximizeWindow = async () => {
    let isFull = await WindowIsFullscreen()
    if (isFull) {
        WindowUnfullscreen()
    } else {
        WindowFullscreen()
    }
}

const closeWindow = async () => {
  await Quit()
}
</script>

<template>
  <div class="w-full h-8 bg-white flex justify-end items-center px-2 select-none">
    <div class="flex space-x-2">
      <button @click="minimizeWindow" class="p-1 hover:bg-gray-200 rounded">
        <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4"></path>
        </svg>
      </button>
      <button @click="maximizeWindow" class="p-1 hover:bg-gray-200 rounded">
        <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <rect width="18" height="18" x="3" y="3" stroke-width="2" rx="2"></rect>
        </svg>
      </button>
      <button @click="closeWindow" class="p-1 hover:bg-red-500 hover:text-white rounded">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    </div>
  </div>
</template>