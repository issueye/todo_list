<script setup>
import { ref, defineProps, defineEmits, watch } from 'vue'

const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true
  },
  editingTodo: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['close', 'save'])

const todoText = ref('')
const priority = ref('medium')
const dueDate = ref('')

// 监听编辑状态变化，更新表单数据
watch(() => props.editingTodo, (newTodo) => {
  if (newTodo) {
    todoText.value = newTodo.title || ''
    priority.value = newTodo.priority || 'medium'
    dueDate.value = newTodo.due_date ? new Date(newTodo.due_date).toISOString().slice(0, 16) : ''
  } else {
    // 新建时设置默认值
    todoText.value = ''
    priority.value = 'medium'
    dueDate.value = new Date().toISOString().slice(0, 16)
  }
}, { immediate: true })

const handleSubmit = () => {
  if (todoText.value.trim()) {
    emit('save', {
      text: todoText.value.trim(),
      priority: priority.value,
      dueDate: dueDate.value
    })
    todoText.value = ''
    priority.value = 'medium'
    dueDate.value = ''
    emit('close')
  }
}

const handleClose = () => {
  todoText.value = ''
  priority.value = 'medium'
  dueDate.value = ''
  emit('close')
}
</script>

<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto bg-black bg-opacity-50 backdrop-blur-sm">
    <div class="flex items-center justify-center min-h-screen p-4">
      <!-- 模态框 -->
      <div class="relative w-full max-w-2xl bg-white rounded-2xl shadow-2xl transform transition-all">
        <form @submit.prevent="handleSubmit" class="p-8">
          <h3 class="text-2xl font-bold text-gray-900 mb-6">
            {{ editingTodo ? '编辑待办事项' : '新建待办事项' }}
          </h3>
          
          <div class="mb-6">
            <label for="todo-text" class="block text-sm font-medium text-gray-700 mb-2">
              内容
            </label>
            <textarea
              id="todo-text"
              v-model="todoText"
              class="w-full px-3 py-2 text-gray-700 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 placeholder-gray-400 transition duration-150 ease-in-out"
              rows="8"
              placeholder="请输入待办事项内容"
            ></textarea>
          </div>

          <div class="mb-6">
            <label for="due-date" class="block text-sm font-medium text-gray-700 mb-2">
              截止日期
            </label>
            <input
              id="due-date"
              type="datetime-local"
              v-model="dueDate"
              class="w-full px-3 py-2 text-gray-700 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition duration-150 ease-in-out"
            >
          </div>

          <div class="mb-8">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              优先级
            </label>
            <div class="flex gap-6 justify-start">
              <label class="inline-flex items-center hover:opacity-80 cursor-pointer">
                <input
                  type="radio"
                  v-model="priority"
                  value="high"
                  class="form-radio h-5 w-5 text-red-600 focus:ring-red-500"
                >
                <span class="ml-2 text-sm text-gray-700">高</span>
              </label>
              <label class="inline-flex items-center hover:opacity-80 cursor-pointer">
                <input
                  type="radio"
                  v-model="priority"
                  value="medium"
                  class="form-radio h-5 w-5 text-yellow-600 focus:ring-yellow-500"
                >
                <span class="ml-2 text-sm text-gray-700">中</span>
              </label>
              <label class="inline-flex items-center hover:opacity-80 cursor-pointer">
                <input
                  type="radio"
                  v-model="priority"
                  value="low"
                  class="form-radio h-5 w-5 text-green-600 focus:ring-green-500"
                >
                <span class="ml-2 text-sm text-gray-700">低</span>
              </label>
            </div>
          </div>

          <div class="flex justify-end gap-3">
            <button
              type="button"
              @click="handleClose"
              class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-400 transition duration-150 ease-in-out"
            >
              取消
            </button>
            <button
              type="submit"
              class="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-150 ease-in-out"
            >
              保存
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>