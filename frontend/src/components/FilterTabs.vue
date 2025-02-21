<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  currentFilter: {
    type: String,
    required: true
  },
  todoCount: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['filterChange'])

const handleFilterChange = (filter) => {
  emit('filterChange', filter)
}
</script>

<template>
  <div class="flex justify-center gap-4 mb-6">
    <button 
      @click="handleFilterChange('all')" 
      class="px-4 py-2 rounded-lg transition-colors text-sm font-medium flex items-center gap-2"
      :class="currentFilter === 'all' ? 'bg-indigo-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300'"
    >
      <span>所有</span>
      <span class="bg-white bg-opacity-20 px-2 py-0.5 rounded-full text-xs">{{ todoCount.all }}</span>
    </button>
    <button 
      @click="handleFilterChange('active')" 
      class="px-4 py-2 rounded-lg transition-colors text-sm font-medium flex items-center gap-2"
      :class="currentFilter === 'active' ? 'bg-indigo-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300'"
    >
      <span>未完成</span>
      <span class="bg-white bg-opacity-20 px-2 py-0.5 rounded-full text-xs">{{ todoCount.active }}</span>
    </button>
    <button 
      @click="handleFilterChange('completed')" 
      class="px-4 py-2 rounded-lg transition-colors text-sm font-medium flex items-center gap-2"
      :class="currentFilter === 'completed' ? 'bg-indigo-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300'"
    >
      <span>已完成</span>
      <span class="bg-white bg-opacity-20 px-2 py-0.5 rounded-full text-xs">{{ todoCount.completed }}</span>
    </button>
  </div>
</template>