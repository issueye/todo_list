<script setup>
import { ref, computed, onMounted, reactive } from 'vue'
import { PlusIcon } from '@heroicons/vue/24/solid'
import TodoModal from '@/components/TodoModal.vue'
import TodoGroupList from '@/components/TodoGroupList.vue'
import SearchBar from '@/components/SearchBar.vue'
import FilterTabs from '@/components/FilterTabs.vue'
import TodoList from '@/components/TodoList.vue'
import Header from '@/components/Header.vue'
import { GetTodoList, CreateTodo, UpdateTodo, DeleteTodo, UpdateTodoState } from '@/wailsjs/go/server/App'

const todos = ref([])
const isModalOpen = ref(false)
const editingTodo = ref(null)
const currentFilter = ref('all')
const selectedDate = ref(null)
const searchKeyword = ref('')  // 添加搜索关键字
const selectedGroup = ref(null)  // 添加选中的事项组

const todoCount = reactive({
  all: 0,
  active: 0,
  completed: 0,
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
const handleSave = async ({ text, level, dueDate, estimatedHours, progress }) => {
  try {
    if (editingTodo.value) {
      let data = {
        id: editingTodo.value.id,
        title: text,
        level: level,
        computed: editingTodo.value.completed,
        due_date: `${dueDate}:00Z`,
        estimated_hours: estimatedHours,
        progress: progress
      }
      await UpdateTodo(data)
    } else {
      let data = {
        title: text,
        level: level,
        due_date: `${dueDate}:00Z`,
        estimated_hours: estimatedHours,
        progress: progress
      }
      await CreateTodo(data)
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
  let params = {
    filter: searchKeyword.value,
    date: qry_date,
    state: state(),
  }
  let data = await GetTodoList(params)
  todos.value = data.list;
  todoCount.all = data.all_count;
  todoCount.active = data.uncompleted_count;
  todoCount.completed = data.completed_count;
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
    await UpdateTodoState(todo.id)
    qryData()
  } catch (error) {
    console.error('更新待办事项状态失败:', error)
  }
}
const handleGroupSelected = (group) => {
  selectedGroup.value = group
  qryData()
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
  <div class="h-full bg-gray-100 flex flex-col border border-solid border-gray-400" style="--wails-draggable: drag">
    <Header />
    <div class="flex-1 py-3 sm:py-12">
      <div class="relative flex-1 sm:max-w-5xl sm:mx-auto w-full sm:px-0">
        <div class="relative h-full px-4 py-8 bg-white shadow-lg sm:rounded-3xl sm:p-12">
          <div class="h-full mx-auto flex gap-8">
            <!-- 左侧事项组列表 -->
            <div class="w-64 border-r pr-6">
              <TodoGroupList @group-selected="handleGroupSelected" />
            </div>
            <!-- 右侧待办事项列表 -->
            <div class="flex-1">
              <div class="h-full flex flex-col">
                <div class="flex justify-between items-center mb-8">
                  <h1 class="text-3xl font-bold text-indigo-600">待办事项</h1>
                  <button @click="openModal()"
                    class="inline-flex items-center px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                    <PlusIcon class="-ml-1 mr-2 h-5 w-5" />
                    新建
                  </button>
                </div>
                <SearchBar :search-keyword="searchKeyword" :selected-date="selectedDate"
                  @search="(value) => { searchKeyword = value; qryData(); }"
                  @date-change="(value) => { selectedDate = value; qryData(); }" />
                <FilterTabs :current-filter="currentFilter" :todo-count="todoCount" @filter-change="stateQryData" />
                <TodoList :todos="filteredTodos" @edit="openModal" @delete="deleteTodo" @toggle="toggleTodo" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <TodoModal :is-open="isModalOpen" :editing-todo="editingTodo" @close="handleClose" @save="handleSave" />
</template>
