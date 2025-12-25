<template>
    <div>
        <el-space style="margin-bottom:12px">
            <el-button type="primary" @click="onRefresh" v-loading="tableLoading">刷新</el-button>
        </el-space>

        <el-card>
            <el-table :data="configs" v-loading="tableLoading" size="small" border style="width:100%">
                <!-- <el-table-column prop="ID" label="ID" width="60" /> -->
                <el-table-column prop="SchedID" label="调度ID" width="60" />
                <el-table-column prop="Cron" label="Cron" min-width="60" />
                <el-table-column prop="State" label="状态" min-width="20" />
                <el-table-column prop="Type" label="类型" width="80"/>
                <el-table-column prop="UpdatedAt" label="更新时间" width="140" />
            </el-table>

            <div style="margin-top:12px; display:flex; justify-content:flex-end;">
                <el-pagination background :page-size="pageQuery.pageSize" :page-sizes="[5, 10, 20, 50]"
                    :current-page="pageQuery.page" :total="pageQuery.total" layout="prev, pager, next, sizes, ->, total"
                    @current-change="onPageChange" @size-change="onSizeChange" />
            </div>
        </el-card>

    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { JobTaskAPI } from './api'

const configs = ref<any[]>([])
const tableLoading = ref(false)

const pageQuery = reactive({ page: 1, pageSize: 10, total: 0 })

async function loadSqlTasks() {
    tableLoading.value = true
    try {
        const res = await JobTaskAPI.page(pageQuery)
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

onMounted(() => {
    loadSqlTasks()
})

</script>

<style scoped></style>
