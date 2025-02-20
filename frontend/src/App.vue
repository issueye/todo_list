<script setup>
import { ref, computed, onMounted } from 'vue'
import { PlusIcon } from '@heroicons/vue/24/solid'
import TodoModal from './components/TodoModal.vue'
import { GetAllTodos, CreateTodo, UpdateTodo, DeleteTodo } from '../wailsjs/go/main/App'

const todos = ref([])
const isModalOpen = ref(false)
const editingTodo = ref(null)
const currentFilter = ref('all')
const selectedDate = ref(null)
const searchKeyword = ref('')  // 添加搜索关键字
const todoCount = computed(() => {
  return {
    all: todos.value.length,
    active: todos.value.filter(todo => !todo.completed).length,
    completed: todos.value.filter(todo => todo.completed).length
  }
})

// 在组件挂载时加载数据
onMounted(async () => {
  try {
    selectedDate.value = new Date().toISOString().split('T')[0];
    qryData()
  } catch (error) {
    console.error('加载待办事项失败:', error)
  }
})

const filteredTodos = computed(() => {
  return todos.value;
})

const openModal = (todo = null) => {
  editingTodo.value = todo
  isModalOpen.value = true
}

const handleClose = () => {
  isModalOpen.value = false
  qryData();
}

const handleSave = async ({ text, level, dueDate }) => {
  try {
    if (editingTodo.value) {
      await UpdateTodo(editingTodo.value.id, text, level, editingTodo.value.completed, dueDate)
      const todo = todos.value.find(t => t.id === editingTodo.value.id)
      if (todo) {
        todo.title = text
        todo.due_date = dueDate
      }
    } else {
      await CreateTodo(text, level, dueDate)
    }
    isModalOpen.value = false
    editingTodo.value = null
    qryData()
  } catch (error) {
    console.error('保存待办事项失败:', error)
  }
}

/**
 * 获取数据
 */
const getData = async () => {
  console.log('selectedDate.value', selectedDate.value);
  // 将时间转换为  YYYY-MM-DD 格式
  let qry_date = selectedDate.value.substr(0, 10);

  let state = () => {
    switch (currentFilter.value) {
      case 'active':
        return 0
      case 'completed':
        return 1
      default:
        return -1
    }
  }

  console.log('state', state());

  todos.value = await GetAllTodos(searchKeyword.value, qry_date, state())
}

const qryData = () => {
  getData();
}

const stateQryData = (state) => {
  currentFilter.value = state;
  getData();
}

const toggleTodo = async (todo) => {
  try {
    await UpdateTodo(todo.id, todo.title, todo.level, !todo.completed)
    todo.completed = !todo.completed
    qryData()
  } catch (error) {
    console.error('更新待办事项状态失败:', error)
  }
}

const deleteTodo = async (id) => {
  try {
    await DeleteTodo(id)
    todos.value = todos.value.filter(todo => todo.id !== id)
  } catch (error) {
    console.error('删除待办事项失败:', error)
  }
}
</script>

<template>
  <div class="h-full bg-gray-100 py-6 flex flex-col sm:py-12">
    <div class="relative flex-1 py-3 sm:max-w-xl sm:mx-auto w-full px-4 sm:px-0">
      <div class="relative h-full px-4 py-8 bg-white shadow-lg sm:rounded-3xl sm:p-12">
        <div class="h-full max-w-lg mx-auto">
          <div class="h-full flex flex-col">
            <div class="flex justify-between items-center mb-8">
              <h1 class="text-3xl font-bold text-indigo-600">待办事项</h1>
              <button @click="openModal()"
                class="inline-flex items-center px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                <PlusIcon class="-ml-1 mr-2 h-5 w-5" />
                新建</button>
            </div>

            <!-- 日期筛选 -->
            <div class="mb-6 space-y-4">
              <!-- 搜索框 -->
              <input type="text" v-model="searchKeyword" placeholder="搜索待办事项" @change="qryData()"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-indigo-500 focus:border-indigo-500">

              <input type="date" v-model="selectedDate" @change="qryData()"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-indigo-500 focus:border-indigo-500">
            </div>

            <!-- 任务分类标签栏 -->
            <div class="flex justify-center gap-4 mb-6">
              <button @click="stateQryData('all')" class="px-4 py-2 rounded-lg transition-colors text-sm font-medium flex items-center gap-2"
                :class="currentFilter === 'all' ? 'bg-indigo-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300'">
                <span>所有</span>
                <span class="bg-white bg-opacity-20 px-2 py-0.5 rounded-full text-xs">{{ todoCount.all }}</span>
              </button>
              <button @click="stateQryData('active')" class="px-4 py-2 rounded-lg transition-colors text-sm font-medium flex items-center gap-2"
                :class="currentFilter === 'active' ? 'bg-indigo-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300'">
                <span>未完成</span>
                <span class="bg-white bg-opacity-20 px-2 py-0.5 rounded-full text-xs">{{ todoCount.active }}</span>
              </button>
              <button @click="stateQryData('completed')" class="px-4 py-2 rounded-lg transition-colors text-sm font-medium flex items-center gap-2"
                :class="currentFilter === 'completed' ? 'bg-indigo-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300'">
                <span>已完成</span>
                <span class="bg-white bg-opacity-20 px-2 py-0.5 rounded-full text-xs">{{ todoCount.completed }}</span>
              </button>
            </div>
            <!-- 待办事项列表 -->
            <div class="space-y-4 h-[calc(100vh-25rem)] overflow-y-auto pr-2 custom-scrollbar">
              <div v-for="todo in filteredTodos" :key="todo.id"
                class="flex items-center justify-between p-4 bg-gray-50 rounded-lg" :class="{
                'opacity-75': todo.completed
              }">
                <div class="flex items-center space-x-3">
                  <input type="checkbox" :checked="todo.completed" @change="toggleTodo(todo)"
                    class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded">
                  <span class="text-gray-800" :class="{
                'line-through': todo.completed
              }">
                    {{ todo.title }}
                  </span>
                  <span class="px-2 py-1 text-xs font-medium rounded" :class="{
                'bg-red-100 text-red-800': todo.level === 'high',
                'bg-yellow-100 text-yellow-800': todo.level === 'medium',
                'bg-green-100 text-green-800': todo.level === 'low'
              }">
                    {{ todo.level === 'high' ? '高' : todo.level === 'medium' ? '中' : '低' }}
                  </span>
                </div>
                <div class="flex space-x-2">
                  <button @click="openModal(todo)" class="text-sm text-indigo-600 hover:text-indigo-900">
                    编辑</button>
                  <button @click="deleteTodo(todo.id)" class="text-sm text-red-600 hover:text-red-900">
                    删除</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <TodoModal :is-open="isModalOpen" :editing-todo="editingTodo" @close="handleClose" @save="handleSave" />
</template>
