<template>
  <div>
    <el-space style="margin-bottom:12px">
      <el-button type="primary" @click="onAddAlertRule">新增</el-button>
      <el-button type="primary" @click="onRefresh" v-loading="tableLoading">刷新</el-button>
    </el-space>
    <el-card >
      <el-table border size="small" :data="alertRuleTableData" v-loading="tableLoading" style="width:100%">
        <el-table-column prop="ID" label="ID" width="60" />
        <el-table-column label="告警普米" width="140">
          <template #default="{ row }">
            <el-tag type="info" v-if="row.PrometheusConfig != undefined">{{ row.PrometheusConfig ? row.PrometheusConfig.Remark : '已删除' }}</el-tag>
            <el-tag type="danger" v-else>已删除普米</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="Type" label="告警类型" width="120" />
        <el-table-column prop="Title" label="告警标题" min-width="180" />
        <el-table-column prop="Level" label="告警等级" width="80">
          <template #default="{ row }">
            <el-tag type="danger" v-if="row.Level == 'emergency'">紧急</el-tag>
            <el-tag type="warning" v-else-if="row.Level == 'critical'">严重</el-tag>
            <el-tag type="info" v-else>一般</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="For" label="持续时间(秒)" width="100" />
        <el-table-column prop="Enabled" label="启用" width="80">
          <template #default="{ row }">
            <el-tag type="success" v-if="row.Enabled == '1'">是</el-tag>
            <el-tag type="danger" v-else>否</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="UpdatedAt" label="更新时间" width="140"/>
        <el-table-column label="操作" fixed="right" width="260">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="onModify(row)">修改</el-button>
            <el-button type="primary" size="small" @click="onCopy(row)" style="margin-left:8px">复制</el-button>
            <el-button type="warning" size="small" @click="onSbumit(row)" style="margin-left:8px"
              :loading="submitLoading">提交</el-button>
            <el-popconfirm title="确认删除？" confirm-button-text="确定" cancel-button-text="取消"
              @confirm="() => onDelete(row)">
              <template #reference>
                <el-button type="danger" size="small" style="margin-left:8px">删除</el-button>
              </template>
            </el-popconfirm>  
          </template>
        </el-table-column>
      </el-table>

      <div style="margin-top:12px; display:flex; justify-content:flex-end; align-items:center">
        <el-pagination size="small" background :page-size="pageQuery.pageSize" :page-sizes="[5, 10, 20, 50]"
          :current-page="pageQuery.page" :total="pageQuery.total" layout="prev, pager, next, jumper, sizes, ->, total"
          prev-text="上一页" next-text="下一页" @current-change="onPageChange" @size-change="onSizeChange">
          <template #total>
            共 {{ pageQuery.total }} 条
          </template>
        </el-pagination>
      </div>
    </el-card>

    <el-drawer v-model="editDialogVisible" title="修改告警规则" direction="rtl" size="50%">
      <el-card shadow="hover">
        <el-form ref="modifyFormRef" :model="modifyForm" label-width="auto">
          <el-form-item label="ID">
            <el-input v-model="modifyForm.ID" disabled />
          </el-form-item>
          <el-form-item label="告警普米" prop="PrometheusConfigID"
            :rules="[{ required: true, message: '请选择告警普米', trigger: 'blur' }]">
            <el-select v-model="modifyForm.PrometheusConfigID" placeholder="请选择告警普米">
              <el-option v-for="cfg in PrometheusConfigList" :key="cfg.ID" :label="cfg.Remark" :value="cfg.ID" />
            </el-select>
          </el-form-item>
          <el-form-item label="告警类型" prop="Type" :rules="[{ required: true, message: '请选择告警类型', trigger: 'blur' }]">
            <el-select v-model="modifyForm.Type" placeholder="请选择告警类型">
              <el-option v-for="type in AlertTypeList" :key="type.Code" :label="type.Name" :value="type.Name" />
            </el-select>
          </el-form-item>
          <el-form-item label="告警标题" prop="Title" :rules="[{ required: true, message: '请输入告警标题', trigger: 'blur' }]">
            <el-input v-model="modifyForm.Title" />
          </el-form-item>
          <el-form-item label="告警等级" prop="Level" :rules="[{ required: true, message: '请选择告警等级', trigger: 'blur' }]">
            <el-radio-group v-model="modifyForm.Level">
              <el-radio label="emergency">紧急</el-radio>
              <el-radio label="critical">严重</el-radio>
              <el-radio label="warning">一般</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="持续时间" prop="For">
            <el-input-number v-model="modifyForm.For" :min="0" />（秒）
          </el-form-item>

          <el-form-item label="告警规则" prop="PromQLRule"
            :rules="[{ required: true, message: '请输入 PromQL', trigger: 'blur' }]">
            <el-input type="textarea" :rows="3" v-model="modifyForm.PromQLRule" />
          </el-form-item>
          <el-form-item label="告警内容" prop="Content" :rules="[{ required: true, message: '请输入告警内容', trigger: 'blur' }]">
            <el-input type="textarea" :rows="6" v-model="modifyForm.Content" />
          </el-form-item>
          <el-form-item label="启用">
            <el-radio-group v-model="modifyForm.Enabled">
              <el-radio label="1">是</el-radio>
              <el-radio label="0">否</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="更新时间">
            <el-input v-model="modifyForm.UpdatedAt" disabled />
          </el-form-item>
          <el-form-item label="创建时间">
            <el-input v-model="modifyForm.CreatedAt" disabled />
          </el-form-item>
        </el-form>
      </el-card>
      <template #footer>
        <el-button @click="closeEditDialog">取消</el-button>
        <el-button type="primary" @click="saveModify">保存</el-button>
      </template>
    </el-drawer>

    <el-dialog v-model="addDialogVisible" title="新增告警规则">
      <el-card shadow="hover">
        <el-form ref="addFormRef" :model="addForm" label-width="auto">
          <el-form-item label="告警普米" prop="PrometheusConfigID"
            :rules="[{ required: true, message: '请选择告警普米', trigger: 'blur' }]">
            <el-select v-model="addForm.PrometheusConfigID" placeholder="请选择告警普米" clearable>
              <el-option v-for="cfg in PrometheusConfigList" :key="cfg.ID" :label="cfg.Remark" :value="cfg.ID" />
            </el-select>
          </el-form-item>
          <el-form-item label="告警类型" prop="Type" :rules="[{ required: true, message: '请选择告警类型', trigger: 'blur' }]">
            <el-select v-model="addForm.Type" placeholder="请选择告警类型">
              <el-option v-for="type in AlertTypeList" :key="type.Code" :label="type.Name" :value="type.Name" />
            </el-select>
          </el-form-item>
          <el-form-item label="告警标题" prop="Title" :rules="[{ required: true, message: '请输入告警标题', trigger: 'blur' }]">
            <el-input v-model="addForm.Title" />
          </el-form-item>
          <el-form-item label="告警等级" prop="Level" :rules="[{ required: true, message: '请选择告警等级', trigger: 'blur' }]">
            <el-radio-group v-model="addForm.Level">
              <el-radio label="emergency">紧急</el-radio>
              <el-radio label="critical">严重</el-radio>
              <el-radio label="warning">一般</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="持续时间" prop="For">
            <el-input-number v-model="addForm.For" :min="0" />（秒）
          </el-form-item>

          <el-form-item label="告警规则" prop="PromQLRule"
            :rules="[{ required: true, message: '请输入 PromQL', trigger: 'blur' }]">
            <el-input type="textarea" :rows="3" v-model="addForm.PromQLRule" />
          </el-form-item>
          <el-form-item label="告警内容" prop="Content" :rules="[{ required: true, message: '请输入告警内容', trigger: 'blur' }]">
            <el-input type="textarea" :rows="4" v-model="addForm.Content" />
          </el-form-item>
          <el-form-item label="启用">
            <el-radio-group v-model="addForm.Enabled">
              <el-radio label="1">是</el-radio>
              <el-radio label="0">否</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-form>
      </el-card>
      <template #footer>
        <el-button @click="addDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveAdd">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, reactive } from 'vue'
import { AlertRuleAPI } from './api'
import { PrometheusConfigAPI } from '../PrometheusConfig/api'
import { AlertTypeAPI } from '../AlertType/api'
import type { FormInstance } from 'element-plus'
import { ElMessage } from 'element-plus'

const alertRuleTableData = ref<any[]>([]);
const tableLoading = ref(false)

let pageQuery = reactive({
  page: 1,
  pageSize: 10,
  total: 0,
})

const addDialogVisible = ref(false)

interface IAddForm {
  Title: string
  Level: string
  For: number
  Type: string
  PromQLRule: string
  Content: string
  Enabled: string
  PrometheusConfigID: number | null
}
const addFormRef = ref<FormInstance>()
const addForm = reactive<IAddForm>({
  Title: '',
  Level: '',
  For: 0,
  Type: '',
  PromQLRule: '',
  Content: '',
  Enabled: '1',
  PrometheusConfigID: null,
})

const editDialogVisible = ref(false)
interface IModifyForm {
  ID: number
  Title: string
  Level: string
  For: number
  Type: string
  PromQLRule: string
  Content: string
  Enabled: string
  UpdatedAt: string
  CreatedAt: string
  PrometheusConfigID: number | null
}
const modifyFormRef = ref<FormInstance>()
const modifyForm = reactive<IModifyForm>({
  ID: 0,
  Title: '',
  Level: '',
  For: 0,
  Type: '',
  PromQLRule: '',
  Content: '',
  Enabled: '1',
  UpdatedAt: '',
  CreatedAt: '',
  PrometheusConfigID: null,
})

async function onRefresh() {
  await loadAlertRuleTableData()
}

async function loadAlertRuleTableData() {
  tableLoading.value = true
  const res = await AlertRuleAPI.pageAlertRules(pageQuery)

  alertRuleTableData.value = res.data || []
  pageQuery.total = res.total || 0

  tableLoading.value = false
}

function onPageChange(page: number) {
  pageQuery.page = page
  loadAlertRuleTableData()
}

function onSizeChange(size: number) {
  pageQuery.pageSize = size
  loadAlertRuleTableData()
}

function onModify(row: any) {
  Object.assign(modifyForm, row)
  editDialogVisible.value = true
}

function onAddAlertRule() {
  addForm.Title = ''
  addForm.Level = ''
  addForm.Type = ''
  addForm.PromQLRule = ''
  addForm.Content = ''
  addForm.Enabled = '1'
  addForm.PrometheusConfigID = null
  addDialogVisible.value = true
}

async function onDelete(row: any) {
  await AlertRuleAPI.deleteAlertRule({ ID: row.ID })
  await loadAlertRuleTableData()
}

async function saveModify() {
  try {
    await (modifyFormRef.value as any).validate()
  } catch (err) {
    ElMessage.error('请检查必填项')
    return
  }

  try {
    await AlertRuleAPI.modifyAlertRule(modifyForm)
    editDialogVisible.value = false
    await loadAlertRuleTableData()
  } catch (err) {

  }
}

function closeEditDialog() {
  editDialogVisible.value = false
}

async function saveAdd() {
  try {
    await (addFormRef.value as any).validate()
  } catch (err) {
    return
  }

  await AlertRuleAPI.addAlertRule(addForm)
  addDialogVisible.value = false
  await loadAlertRuleTableData()
}

async function onCopy(row: any) {
  await AlertRuleAPI.copyAlertRule({ ID: row.ID })
  await loadAlertRuleTableData()
}

const submitLoading = ref(false)

async function onSbumit(row: any) {
  submitLoading.value = true
  try {
    await AlertRuleAPI.submitAlertRule({ ID: row.ID })
  } finally {
    submitLoading.value = false
  }
}

const PrometheusConfigList = ref<any[]>([]);

async function loadAllPrometheusConfig() {
  const res = await PrometheusConfigAPI.listall()
  PrometheusConfigList.value = res || []
}

const AlertTypeList = ref<any[]>([]);

async function loadAllAlertTypes() {
  const res = await AlertTypeAPI.listAllAlertType()
  AlertTypeList.value = res || []
}

onMounted(() => {
  loadAlertRuleTableData()
  loadAllPrometheusConfig()
  loadAllAlertTypes()
})

</script>

<style scoped></style>
