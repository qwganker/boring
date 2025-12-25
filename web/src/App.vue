<template>
  <el-container style="height:100vh;">
    <el-aside :width="collapsed ? '60px' : '220px'" class="app-aside" style="background-color: #545c64">
      <Sidebar :collapsed="collapsed" />
    </el-aside>

    <el-container>
      <el-header class="app-header">
        <el-button @click="toggleCollapse" circle>
          <el-icon v-if="!collapsed">
            <Fold />
          </el-icon>
          <el-icon v-else>
            <Expand />
          </el-icon>
        </el-button>
        <Breadcrumbs />
      </el-header>

      <el-main class="app-main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import { ref, onMounted } from 'vue'
import Sidebar from './components/Sidebar.vue'
import Breadcrumbs from './components/Breadcrumbs.vue'
import { Expand, Fold } from '@element-plus/icons-vue'

export default {
  name: 'App',
  components: { Sidebar, Breadcrumbs, Expand, Fold },
  setup() {
    const collapsed = ref(false)

    function toggleCollapse() {
      collapsed.value = !collapsed.value
    }

    // make sidebar collapsed on small screens
    onMounted(() => {
      if (window.innerWidth < 700) collapsed.value = true
      window.addEventListener('resize', () => {
        if (window.innerWidth < 700) collapsed.value = true
      })
    })

    return { collapsed, toggleCollapse }
  }
}
</script>

<style scoped>
.app-header {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-title {
  font-weight: 600;
  font-size: 16px;
}

.app-main {
  padding: 14px;
  background: #f5f7fa;
}

.app-aside {
  transition: width 0.25s ease, background-color 0.3s ease;
  overflow: hidden;
}
</style>
