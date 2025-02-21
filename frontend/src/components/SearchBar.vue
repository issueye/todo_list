<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  searchKeyword: {
    type: String,
    default: ''
  },
  selectedDate: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['search', 'dateChange'])

const handleSearch = (event) => {
  emit('search', event.target.value)
}

const handleDateChange = (event) => {
  emit('dateChange', event.target.value)
}
</script>

<template>
  <div class="mb-6 space-y-4">
    <input 
      type="text" 
      :value="searchKeyword"
      placeholder="搜索待办事项" 
      @change="handleSearch"
      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
    >
    <input 
      type="date" 
      :value="selectedDate"
      @change="handleDateChange"
      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
    >
  </div>
</template>