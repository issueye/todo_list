<script setup>
import { defineProps, defineEmits } from 'vue'
import { ExclamationTriangleIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    default: '确认'
  },
  message: {
    type: String,
    required: true
  },
  type: {
    type: String,
    default: 'warning',
    validator: (value) => ['warning', 'error', 'info'].includes(value)
  },
  buttons: {
    type: Array,
    default: () => [
      {
        label: '取消',
        role: 'cancel'
      },
      {
        label: '确认',
        role: 'confirm',
        style: 'danger'
      }
    ]
  }
})

const emit = defineEmits(['confirm', 'cancel'])

const handleButtonClick = (button) => {
  if (button.role === 'confirm') {
    emit('confirm')
  } else {
    emit('cancel')
  }
}
</script>

<template>
  <Transition
    enter-active-class="ease-out duration-300"
    enter-from-class="opacity-0"
    enter-to-class="opacity-100"
    leave-active-class="ease-in duration-200"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div v-if="isOpen" class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity z-50" />
  </Transition>

  <Transition
    enter-active-class="ease-out duration-300"
    enter-from-class="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
    enter-to-class="opacity-100 translate-y-0 sm:scale-100"
    leave-active-class="ease-in duration-200"
    leave-from-class="opacity-100 translate-y-0 sm:scale-100"
    leave-to-class="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
  >
    <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
        <div
          class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg"
        >
          <div class="bg-white px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div
                class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full sm:mx-0 sm:h-10 sm:w-10"
                :class="{
                  'bg-red-100': type === 'error',
                  'bg-yellow-100': type === 'warning',
                  'bg-blue-100': type === 'info'
                }"
              >
                <ExclamationTriangleIcon
                  class="h-6 w-6"
                  :class="{
                    'text-red-600': type === 'error',
                    'text-yellow-600': type === 'warning',
                    'text-blue-600': type === 'info'
                  }"
                  aria-hidden="true"
                />
              </div>
              <div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                <h3 class="text-base font-semibold leading-6 text-gray-900">{{ title }}</h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">{{ message }}</p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
            <template v-for="(button, index) in buttons" :key="index">
              <button
                type="button"
                :class="[
                  'inline-flex w-full justify-center rounded-md px-3 py-2 text-sm font-semibold sm:ml-3 sm:w-auto',
                  button.style === 'danger'
                    ? 'bg-red-600 text-white hover:bg-red-500 focus:ring-red-500'
                    : button.style === 'primary'
                    ? 'bg-indigo-600 text-white hover:bg-indigo-500 focus:ring-indigo-500'
                    : 'bg-white text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:ring-indigo-500'
                ]"
                @click="handleButtonClick(button)"
              >
                {{ button.label }}
              </button>
            </template>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>