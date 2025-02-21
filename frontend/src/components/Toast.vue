<script setup>
import { ref } from 'vue'

const visible = ref(false)
const message = ref('')
const type = ref('info')
let timer = null

const show = (msg, msgType = 'info', duration = 3000) => {
  message.value = msg
  type.value = msgType
  visible.value = true

  if (timer) clearTimeout(timer)
  timer = setTimeout(() => {
    visible.value = false
  }, duration)
}

defineExpose({
  show
})
</script>

<template>
  <Transition
    enter-active-class="transition duration-300 ease-out"
    enter-from-class="transform -translate-y-4 opacity-0"
    enter-to-class="transform translate-y-0 opacity-100"
    leave-active-class="transition duration-200 ease-in"
    leave-from-class="transform translate-y-0 opacity-100"
    leave-to-class="transform -translate-y-4 opacity-0"
  >
    <div
      v-if="visible"
      class="fixed top-4 left-1/2 transform -translate-x-1/2 z-50 px-4 py-2 rounded-lg shadow-lg"
      :class="{
        'bg-green-500 text-white': type === 'success',
        'bg-red-500 text-white': type === 'error',
        'bg-yellow-500 text-white': type === 'warning',
        'bg-blue-500 text-white': type === 'info'
      }"
    >
      {{ message }}
    </div>
  </Transition>
</template>