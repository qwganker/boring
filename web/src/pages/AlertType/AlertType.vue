<template>
  <div>
    <el-space style="margin-bottom:12px">
      <el-button type="primary" @click="onAddConfig">新增</el-button>
      <el-button type="primary" @click="onRefresh" v-loading="tableLoading">刷新</el-button>
    </el-space>

    <el-card>
      <el-table :data="configs" v-loading="tableLoading" size="small" border style="width:100%">
        <el-table-column prop="ID" label="ID" width="60" />
        <el-table-column prop="Name" label="名称" min-width="180" />
        <el-table-column prop="Code" label="Code" min-width="180" />
        <el-table-column prop="UpdatedAt" label="更新时间" width="140"/>

        <el-table-column label="操作" fixed="right" width="160">
          <template #default="{ row }">
            <el-button size="small" type="primary" @click="onEdit(row)">修改</el-button>
            <el-popconfirm title="确认删除？" confirm-button-text="确定" cancel-button-text="取消"
              @confirm="() => onDelete(row)">
              <template #reference>
                <el-button size="small" type="danger" style="margin-left:8px">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>

      <div style="margin-top:12px; display:flex; justify-content:flex-end;">
        <el-pagination background :page-size="pageQuery.pageSize" :page-sizes="[5, 10, 20, 50]"
          :current-page="pageQuery.page" :total="pageQuery.total" layout="prev, pager, next, sizes, ->, total"
          @current-change="onPageChange" @size-change="onSizeChange" />
      </div>
    </el-card>

    <!-- Edit Drawer -->
    <el-drawer v-model="editVisible" title="修改告警类型" size="50%" direction="rtl">
      <el-form ref="editFormRef" :model="editForm" label-width="auto">
        <el-form-item label="ID">
          <el-input v-model="editForm.ID" disabled />
        </el-form-item>
        <el-form-item label="名称" prop="Name" :rules="[{ required: true, message: '请输入名称', trigger: 'blur' }]">
          <el-input v-model="editForm.Name" />
        </el-form-item>
        <el-form-item label="代码" prop="Code" :rules="[{ required: true, message: '请输入代码', trigger: 'blur' }]">
          <el-input v-model="editForm.Code" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="editVisible = false">取消</el-button>
        <el-button type="primary" @click="saveEdit">保存</el-button>
      </template>
    </el-drawer>

    <!-- Add Dialog -->
    <el-dialog v-model="addVisible" title="新增告警类型">
      <el-form ref="addFormRef" :model="addForm" label-width="auto">
        <el-form-item label="名称" prop="Name" :rules="[{ required: true, message: '请输入名称', trigger: 'blur' }]">
          <el-input v-model="addForm.Name" />
        </el-form-item>
        <el-form-item label="代码" prop="Code" :rules="[{ required: true, message: '请输入代码', trigger: 'blur' }]">
          <el-input v-model="addForm.Code" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="addVisible = false">取消</el-button>
        <el-button type="primary" @click="saveAdd">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { AlertTypeAPI } from './api'
import type { FormInstance } from 'element-plus'

const configs = ref<any[]>([])
const tableLoading = ref(false)

const pageQuery = reactive({ page: 1, pageSize: 10, total: 0 })

const editVisible = ref(false)
const editFormRef = ref<FormInstance>()
const editForm = reactive<any>({ ID: 0,  Name: '', Code: ''})

const addVisible = ref(false)
const addFormRef = ref<FormInstance>()
const addForm = reactive<any>({ Name: '', Code: ''})

async function loadConfigs() {
  tableLoading.value = true
  try {
    const res = await AlertTypeAPI.pageAlertType(pageQuery)
    configs.value = res.data || []
    pageQuery.total = res.total || 0
  } catch (err) {
    console.error(err)
  } finally {
    tableLoading.value = false
  }
}

function onRefresh() { loadConfigs() }

function onPageChange(p: number) { pageQuery.page = p; loadConfigs() }
function onSizeChange(s: number) { pageQuery.pageSize = s; loadConfigs() }

function onEdit(row: any) { Object.assign(editForm, row); editVisible.value = true }

async function onDelete(row: any) {
  try {
    await AlertTypeAPI.deleteAlertType({ ID: row.ID })
    loadConfigs()
  } catch (err) {
  }
}

async function saveEdit() {
  try {
    await (editFormRef.value as any).validate()
  } catch (err) {
    return
  }
  try {
    await AlertTypeAPI.modifyAlertType(editForm)
    editVisible.value = false
    loadConfigs()
  } catch (err) {
  }
}

function onAddConfig() {
  Object.assign(addForm, { Remark: '', Address: '', Username: '', Password: '', CtrlAddress: '', Config: '', Rule: '', Enabled: '1' })
  addVisible.value = true
}

async function saveAdd() {
  try {
    await (addFormRef.value as any).validate()
  } catch (err) {
    return
  }
  try {
    await AlertTypeAPI.addAlertType(addForm)
    addVisible.value = false
    loadConfigs()
  } catch (err) {
  }
}

onMounted(() => loadConfigs())

</script>

<style scoped></style>
