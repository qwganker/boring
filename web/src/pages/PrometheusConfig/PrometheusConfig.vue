<template>
  <div>
    <el-space style="margin-bottom:12px">
      <el-button type="primary" @click="onAddConfig">新增</el-button>
      <el-button type="primary" @click="onRefresh" v-loading="tableLoading">刷新</el-button>
    </el-space>

    <el-card>
      <el-table :data="configs" v-loading="tableLoading" size="small" border style="width:100%">
        <el-table-column prop="ID" label="ID" width="60" />
        <el-table-column prop="Remark" label="备注" min-width="180" />
        <el-table-column prop="Address" label="访问地址" min-width="200" />
        <el-table-column prop="CtrlAddress" label="控制地址" min-width="200" />
        <el-table-column prop="Status" label="运行状态" width="100" />
        <el-table-column prop="Enabled" label="启用" width="100">
          <template #default="{ row }">
            <el-tag type="success" v-if="row.Enabled == '1'">是</el-tag>
            <el-tag type="danger" v-else>否</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="UpdatedAt" label="更新时间" width="140"/>
        <el-table-column label="操作" fixed="right" width="240">
          <template #default="{ row }">
            <el-button size="small" type="primary" @click="onEdit(row)">修改</el-button>
            <el-button size="small" type="primary" @click="onCopy(row)" style="margin-left:8px">复制</el-button>
            <el-button size="small" type="warning" @click="onSumbit(row)" style="margin-left:8px">提交</el-button>
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
    <el-drawer v-model="editVisible" title="修改普米配置" size="50%" direction="rtl">
      <el-form ref="editFormRef" :model="editForm" label-width="auto">
        <el-form-item label="ID">
          <el-input v-model="editForm.ID" disabled />
        </el-form-item>
        <el-form-item label="备注" prop="Remark" :rules="[{ required: true, message: '请输入备注', trigger: 'blur' }]">
          <el-input v-model="editForm.Remark" />
        </el-form-item>
        <el-form-item label="访问地址" prop="Address" :rules="[{ required: true, message: '请输入地址', trigger: 'blur' }]">
          <el-input v-model="editForm.Address" />
        </el-form-item>
        <el-form-item label="控制地址" prop="CtrlAddress"
          :rules="[{ required: true, message: '请输入控制地址', trigger: 'blur' }]">
          <el-input v-model="editForm.CtrlAddress" />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="editForm.Username" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="editForm.Password" type="password" show-password />
        </el-form-item>
        <el-form-item label="配置">
          <el-input type="textarea" autosize v-model="editForm.Config" />
        </el-form-item>
        <el-form-item label="规则">
          <el-input type="textarea" autosize v-model="editForm.Rule" disabled />
        </el-form-item>
        <el-form-item label="更新时间">
          <el-input v-model="editForm.UpdatedAt" disabled />
        </el-form-item>
        <el-form-item label="创建时间">
          <el-input v-model="editForm.CreatedAt" disabled />
        </el-form-item>
        <el-form-item label="启用">
          <el-radio-group v-model="editForm.Enabled">
            <el-radio label="1">是</el-radio>
            <el-radio label="0">否</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="editVisible = false">取消</el-button>
        <el-button type="primary" @click="saveEdit">保存</el-button>
      </template>
    </el-drawer>

    <!-- Add Dialog -->
    <el-dialog v-model="addVisible" title="新增普米配置">
      <el-form ref="addFormRef" :model="addForm" label-width="auto">
        <el-form-item label="备注" prop="Remark" :rules="[{ required: true, message: '请输入备注', trigger: 'blur' }]">
          <el-input v-model="addForm.Remark" />
        </el-form-item>
        <el-form-item label="访问地址" prop="Address" :rules="[{ required: true, message: '请输入地址', trigger: 'blur' }]">
          <el-input v-model="addForm.Address" />
        </el-form-item>
        <el-form-item label="控制地址" prop="CtrlAddress"
          :rules="[{ required: true, message: '请输入控制地址', trigger: 'blur' }]">
          <el-input v-model="addForm.CtrlAddress" />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="addForm.Username" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="addForm.Password" type="password" show-password />
        </el-form-item>
        <el-form-item label="配置">
          <el-input type="textarea" autosize v-model="addForm.Config" />
        </el-form-item>
        <el-form-item label="规则">
          <el-input type="textarea" autosize v-model="addForm.Rule" />
        </el-form-item>
        <el-form-item label="启用">
          <el-radio-group v-model="addForm.Enabled">
            <el-radio label="1">是</el-radio>
            <el-radio label="0">否</el-radio>
          </el-radio-group>
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
import { PrometheusConfigAPI } from './api'
import { ElMessage } from 'element-plus'
import type { FormInstance } from 'element-plus'

const configs = ref<any[]>([])
const tableLoading = ref(false)

const pageQuery = reactive({ page: 1, pageSize: 10, total: 0 })

const editVisible = ref(false)
const editFormRef = ref<FormInstance>()
const editForm = reactive<any>({ ID: 0, Remark: '', Address: '', Username: '', Password: '', CtrlAddress: '', Config: '', Rule: '', Enabled: '1', CreatedAt: '', UpdatedAt: '' })

const addVisible = ref(false)
const addFormRef = ref<FormInstance>()
const addForm = reactive<any>({ Remark: '', Address: '', Username: '', Password: '', CtrlAddress: '', Config: '', Rule: '', Enabled: '1' })

async function loadConfigs() {
  tableLoading.value = true
  try {
    const res = await PrometheusConfigAPI.pageConfigs(pageQuery)
    configs.value = res.data || []
    pageQuery.total = res.total || 0
  } catch (err) {
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
    await PrometheusConfigAPI.deleteConfig({ ID: row.ID })
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
    await PrometheusConfigAPI.modifyConfig(editForm)
    editVisible.value = false
    loadConfigs()
  } catch (err) {
  }
}

async function onCopy(row: any) {
  await PrometheusConfigAPI.copyConfig({ ID: row.ID })
  loadConfigs()
}

async function onSumbit(row: any) {
  try {
    await PrometheusConfigAPI.sumbitConfig({ ID: row.ID })
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
    await PrometheusConfigAPI.addConfig(addForm)
    addVisible.value = false
    loadConfigs()
  } catch (err) {
  }
}

onMounted(() => loadConfigs())

</script>

<style scoped></style>
