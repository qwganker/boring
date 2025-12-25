<template>
    <div>
        <el-space style="margin-bottom:12px">
            <el-button type="primary" @click="onAddConfig">新增</el-button>
            <el-button type="primary" @click="onRefresh" v-loading="tableLoading">刷新</el-button>
        </el-space>

        <el-card>
            <el-table :data="configs" v-loading="tableLoading" size="small" border style="width:100%">
                <el-table-column prop="ID" label="ID" width="60" />
                <el-table-column prop="SchedID" label="调度ID" width="60" />
                <el-table-column label="存储普米" width="140">
                    <template #default="{ row }">
                        <div v-for="cfg in PrometheusConfigList">
                            <el-tag type="info" v-if="cfg.ID === row.PrometheusConfigID">{{ cfg.Remark }}</el-tag>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column prop="Remark" label="用途" min-width="140" />
                <el-table-column prop="DBType" label="数据库类型" min-width="60" />
                <el-table-column prop="SQL" label="SQL语句" min-width="180" />
                <el-table-column prop="Cron" label="Cron表达式" min-width="140" />
                <el-table-column prop="Enabled" label="启用" width="80">
                <template #default="{ row }">
                    <el-tag type="success" v-if="row.Enabled == '1'">是</el-tag>
                    <el-tag type="danger" v-else>否</el-tag>
                </template>
                </el-table-column>
                <el-table-column prop="UpdatedAt" label="更新时间" width="140" />
                <el-table-column label="操作" fixed="right" width="240">
                    <template #default="{ row }">
                        <!-- <el-button size="small" type="primary" @click="onRun(row)">运行</el-button>
                        <el-button size="small" type="primary" @click="onStop(row)">停止</el-button> -->
                        <el-button size="small" type="primary" @click="onEdit(row)">修改</el-button>
                        <el-button type="primary" size="small" @click="onCopy(row)"
                            style="margin-left:8px">复制</el-button>
                        <el-button size="small" type="warning" @click="onTest(row)">测试</el-button>    
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

        <el-drawer v-model="editVisible" title="修改SQL采集" size="50%" direction="rtl">
            <el-form ref="editFormRef" :model="editForm" label-width="auto">
                <el-form-item label="ID">
                    <el-input v-model="editForm.ID" disabled />
                </el-form-item>
                <el-form-item label="告警普米" prop="PrometheusConfigID"
                    :rules="[{ required: true, message: '请选择存储普米', trigger: 'blur' }]">
                    <el-select v-model="editForm.PrometheusConfigID" placeholder="请选择存储普米">
                        <el-option v-for="cfg in PrometheusConfigList" :key="cfg.ID" :label="cfg.Remark"
                            :value="cfg.ID" />
                    </el-select>
                </el-form-item>
                <el-form-item label="用途" prop="Remark" :rules="[{ required: true, message: '请输入用途', trigger: 'blur' }]">
                    <el-input v-model="editForm.Remark" />
                </el-form-item>
                <el-form-item label="数据库类型" prop="DBType"
                    :rules="[{ required: true, message: '选择数据库类型', trigger: 'blur' }]">
                    <el-select v-model="editForm.DBType" placeholder="选择数据库类型" style="width: 240px" clearable>
                        <el-option v-for="item in dbTypeOptions" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="DSN" prop="DSN" :rules="[{ required: true, message: '请输入DSN', trigger: 'blur' }]">
                    <el-input v-model="editForm.DSN" />
                </el-form-item>
                <el-form-item label="SQL语句" prop="SQL"
                    :rules="[{ required: true, message: '请输入SQL语句', trigger: 'blur' }]">
                    <el-input type="textarea" autosize v-model="editForm.SQL" />
                </el-form-item>
                <el-form-item label="指标配置" prop="MetricDefine"
                    :rules="[{ required: true, message: '请输入指标配置', trigger: 'blur' }]">
                    <el-input type="textarea" :rows="4" autosize v-model="editForm.MetricDefine" />
                </el-form-item>
                <el-form-item label="Cron表达式" prop="Cron"
                    :rules="[{ required: true, message: '请输入Cron表达式', trigger: 'blur' }]">
                    <el-input v-model="editForm.Cron" />
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

        <el-dialog v-model="addVisible" title="新增SQL采集">
            <el-form ref="addFormRef" :model="addForm" label-width="auto">
                <el-form-item label="告警普米" prop="PrometheusConfigID"
                    :rules="[{ required: true, message: '请选择存储普米', trigger: 'blur' }]">
                    <el-select v-model="addForm.PrometheusConfigID" placeholder="请选择存储普米">
                        <el-option v-for="cfg in PrometheusConfigList" :key="cfg.ID" :label="cfg.Remark"
                            :value="cfg.ID" />
                    </el-select>
                </el-form-item>
                <el-form-item label="用途" prop="Remark" :rules="[{ required: true, message: '请输入用途', trigger: 'blur' }]">
                    <el-input v-model="addForm.Remark" />
                </el-form-item>
                <el-form-item label="数据库类型" prop="DBType"
                    :rules="[{ required: true, message: '请输入数据库类型', trigger: 'blur' }]">
                    <el-select v-model="addForm.DBType" placeholder="选择数据库类型" style="width: 240px" clearable>
                        <el-option v-for="item in dbTypeOptions" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="DSN" prop="DSN" :rules="[{ required: true, message: '请输入DSN', trigger: 'blur' }]">
                    <el-input v-model="addForm.DSN" />
                </el-form-item>
                <el-form-item label="SQL语句" prop="SQL"
                    :rules="[{ required: true, message: '请输入SQL语句', trigger: 'blur' }]">
                    <el-input type="textarea" :rows="3" v-model="addForm.SQL" />
                </el-form-item>
                <el-form-item label="指标配置" prop="MetricDefine"
                    :rules="[{ required: true, message: '请输入指标配置', trigger: 'blur' }]">
                    <el-input type="textarea" :rows="4" v-model="addForm.MetricDefine" />
                </el-form-item>
                <el-form-item label="Cron表达式" prop="Cron"
                    :rules="[{ required: true, message: '请输入Cron表达式', trigger: 'blur' }]">
                    <el-input v-model="addForm.Cron" />
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
import { SqlTaskAPI } from './api'
import { PrometheusConfigAPI } from '../../PrometheusConfig/api'
import { ElMessage, ElMessageBox } from 'element-plus'

import type { FormInstance } from 'element-plus'
import { el } from 'element-plus/es/locale/index.mjs'

const configs = ref<any[]>([])
const tableLoading = ref(false)

const pageQuery = reactive({ page: 1, pageSize: 10, total: 0 })

const editVisible = ref(false)
const editFormRef = ref<FormInstance>()
const editForm = reactive<any>({ ID: 0, Remark: '', DBType: '', DSN: '', SQL: '', MetricDefine: '', Cron: '', PrometheusConfigID: 0, Enabled: '1'})

const addVisible = ref(false)
const addFormRef = ref<FormInstance>()
const addForm = reactive<any>({ Remark: '', DBType: '', DSN: '', SQL: '', MetricDefine: '', Cron: '', PrometheusConfigID: undefined, Enabled: '1' })
const dbTypeOptions = [
    {
        value: 'sqlite',
        label: 'sqlite',
    },
    {
        value: 'mysql',
        label: 'mysql',
    },
]

async function loadSqlTasks() {
    tableLoading.value = true
    try {
        const res = await SqlTaskAPI.page(pageQuery)
        configs.value = res.data || []
        pageQuery.total = res.total || 0
    } catch (err) {
        console.error(err)
    } finally {
        tableLoading.value = false
    }
}

function onRefresh() { loadSqlTasks() }

function onPageChange(p: number) { pageQuery.page = p; loadSqlTasks() }
function onSizeChange(s: number) { pageQuery.pageSize = s; loadSqlTasks() }

function onEdit(row: any) { Object.assign(editForm, row); editVisible.value = true }

async function onDelete(row: any) {
    try {
        await SqlTaskAPI.delete({ ID: row.ID })
        loadSqlTasks()
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
        await SqlTaskAPI.modify(editForm)
        editVisible.value = false
        loadSqlTasks()
    } catch (err) {
    }
}

function onAddConfig() {
    Object.assign(addForm, { Remark: '', DBType: '', SQL: '', Cron: '' })
    addVisible.value = true
}

async function saveAdd() {
    try {
        await (addFormRef.value as any).validate()
    } catch (err) {
        return
    }
    try {
        await SqlTaskAPI.add(addForm)
        addVisible.value = false
        loadSqlTasks()
    } catch (err) {
    }
}

async function onCopy(row: any) {
    await SqlTaskAPI.copy({ ID: row.ID })
    await loadSqlTasks()
}

const PrometheusConfigList = ref<any[]>([]);

async function loadAllPrometheusConfig() {
    const res = await PrometheusConfigAPI.listall()
    PrometheusConfigList.value = res || []
}

async function onTest(row: any) {
    await SqlTaskAPI.test({ ID: row.ID })
}

async function onRun(row: any) {
    await SqlTaskAPI.run({ ID: row.ID })
    loadSqlTasks()
}

async function onStop(row: any) {
    await SqlTaskAPI.stop({ ID: row.ID })
    loadSqlTasks()
}

onMounted(() => {
    loadSqlTasks()
    loadAllPrometheusConfig()
})

</script>

<style scoped></style>
