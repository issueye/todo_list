<script setup>
import { ref, onMounted } from 'vue'
import { PlusIcon } from '@heroicons/vue/24/solid'
import { GetAllTodoGroups, CreateTodoGroup, UpdateTodoGroup, DeleteTodoGroup } from '@/wailsjs/go/server/App'
import ConfirmDialog from './ConfirmDialog.vue'

const groups = ref([])
const selectedGroup = ref(null)
const isEditing = ref(false)
const editingName = ref('')
const showConfirmDialog = ref(false)
const confirmDialogProps = ref({})

const emit = defineEmits(['groupSelected'])

onMounted(async () => {
    await loadGroups()
})

async function loadGroups() {
    try {
        groups.value = await GetAllTodoGroups('')
        if (groups.value.length > 0) {
            if (!selectedGroup.value) {
                selectedGroup.value = groups.value[0]
                emit('groupSelected', groups.value[0])
            }
        }
    } catch (error) {
        console.error('加载事项组失败:', error)
    }
}

async function createGroup() {
    isEditing.value = true
    editingName.value = ''
    selectedGroup.value = null
}

async function editGroup(group) {
    isEditing.value = true
    editingName.value = group.name
    selectedGroup.value = group
}

async function saveGroup() {
    try {
        if (selectedGroup.value) {
            await UpdateTodoGroup({
                id: selectedGroup.value.id,
                name: editingName.value
            })
        } else {
            await CreateTodoGroup({
                name: editingName.value
            })
        }
        isEditing.value = false
        editingName.value = ''
        selectedGroup.value = null
        await loadGroups()
    } catch (error) {
        console.error('保存事项组失败:', error)
    }
}

async function cancelEdit() {
    isEditing.value = false
    editingName.value = ''
    selectedGroup.value = null
}

async function deleteGroup(group) {
    selectedGroup.value = group
    showConfirmDialog.value = true
    confirmDialogProps.value = {
        title: '删除确认',
        message: `确定要删除事项组「${group.name}」吗？\n删除后将无法恢复。`,
        type: 'warning',
        buttons: [
            {
                label: '取消',
                role: 'cancel'
            },
            {
                label: '删除',
                role: 'confirm',
                style: 'danger'
            }
        ]
    }
}

async function handleConfirmDelete() {
    try {
        await DeleteTodoGroup(selectedGroup.value.id)
        await loadGroups()

        selectedGroup.value = null
        emit('groupSelected', null)
        showConfirmDialog.value = false
    } catch (error) {
        console.error('删除事项组失败:', error)
    }
}

function selectGroup(group) {
    selectedGroup.value = group
    emit('groupSelected', group)
}
</script>

<template>
    <div class="h-full flex flex-col">
        <div class="flex justify-between items-center mb-4">
            <h2 class="text-lg font-semibold text-gray-700">事项组</h2>
            <button @click="createGroup"
                class="inline-flex items-center px-2 py-1 text-sm bg-indigo-600 text-white font-medium rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                <PlusIcon class="h-4 w-4" />
            </button>
        </div>

        <div v-if="isEditing" class="mb-4">
            <input v-model="editingName" type="text" placeholder="输入事项组名称"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                @keyup.enter="saveGroup">
            <div class="flex justify-end space-x-2 mt-2">
                <button @click="cancelEdit" class="px-2 py-1 text-sm text-gray-600 hover:text-gray-900">取消</button>
                <button @click="saveGroup"
                    class="px-2 py-1 text-sm bg-indigo-600 text-white rounded-md hover:bg-indigo-700">保存</button>
            </div>
        </div>

        <div class="flex-1 overflow-y-auto space-y-2 custom-scrollbar">
            <div v-for="group in groups" :key="group.id"
                class="flex items-center justify-between p-2 rounded-md hover:bg-gray-100 cursor-pointer"
                :class="{ 'bg-gray-100': selectedGroup?.id === group.id }" @click="selectGroup(group)">
                <span class="truncate flex-1">{{ group.name }}</span>
                <div class="flex space-x-2">
                    <button @click.stop="editGroup(group)"
                        class="text-xs text-indigo-600 hover:text-indigo-900">编辑</button>
                    <button @click.stop="deleteGroup(group)" class="text-xs text-red-600 hover:text-red-900">删除</button>
                </div>
            </div>
        </div>
    </div>
    <ConfirmDialog v-model:isOpen="showConfirmDialog" v-bind="confirmDialogProps" @confirm="handleConfirmDelete" @cancel="showConfirmDialog = false" />
</template>

<style scoped>
.custom-scrollbar {
    scrollbar-width: thin;
    scrollbar-color: #cbd5e0 #f7fafc;
}

.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
    background: #f7fafc;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: #cbd5e0;
    border-radius: 3px;
}
</style>