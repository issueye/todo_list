<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  todos: {
    type: Array,
    required: true
  }
})

const emit = defineEmits(['edit', 'delete', 'toggle'])

const handleEdit = (todo) => {
  emit('edit', todo)
}

const handleDelete = (id) => {
  emit('delete', id)
}

const handleToggle = (todo) => {
  emit('toggle', todo)
}
</script>

<template>
  <div class="space-y-4 h-[calc(100vh-30rem)] overflow-y-auto pr-2 custom-scrollbar">
    <div v-for="todo in todos" :key="todo.id"
      class="flex items-center justify-between p-4 bg-gray-50 rounded-lg transition-all duration-200 hover:bg-white hover:shadow-md hover:border hover:border-indigo-200"
      :class="{
        'opacity-75': todo.completed
      }">
      <div class="flex items-center space-x-3">
        <input type="checkbox" :checked="todo.completed" @change="handleToggle(todo)"
          class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded">
        <span
          class="max-w-[300px] truncate inline-block align-middle"
          :class="{
            'line-through text-gray-500': todo.completed
          }"
        >
          {{ todo.title }}
        </span>
        <div class="flex items-center space-x-2 text-sm text-gray-500">
          <span>{{ todo.estimated_hours }}h</span>
          <span>{{ todo.progress }}%</span>
          <span
            class="px-2 py-1 text-xs font-medium rounded"
            :class="{
              'bg-red-100 text-red-800': todo.level === 'high',
              'bg-yellow-100 text-yellow-800': todo.level === 'medium',
              'bg-green-100 text-green-800': todo.level === 'low'
            }"
          >
            {{ 
              todo.level === 'high' ? '高' :
              todo.level === 'medium' ? '中' : '低'
            }}
          </span>
        </div>
      </div>
      <div class="flex space-x-2">
        <button @click="handleEdit(todo)" class="text-sm text-indigo-600 hover:text-indigo-900">
          编辑</button>
        <button @click="handleDelete(todo.id)" class="text-sm text-red-600 hover:text-red-900">
          删除</button>
      </div>
    </div>
  </div>
</template>