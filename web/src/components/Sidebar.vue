<template>
  <div class="sidebar-wrap" :class="{ 'is-collapsed': collapsed }">
    <div class="brand">Boring</div>
    <el-menu :default-active="active" class="el-menu-vertical-demo" @select="handleSelect" :collapse="collapsed"
      background-color="#545c64" text-color="#fff" active-text-color="#ffd04b" :default-openeds="['alert']">
      <!-- <el-menu-item index="/portal">
        <el-icon>
          <Warning />
        </el-icon>
        <span>首页</span>
      </el-menu-item> -->
      <el-sub-menu index="job">
        <template #title>
          <el-icon>
            <Timer />
          </el-icon>
          <span>定时任务</span>
        </template>
        <el-menu-item index="/job/jobtask">任务列表</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="collector">
        <template #title>
          <el-icon>
            <Files />
          </el-icon>
          <span>收集器</span>
        </template>
        <el-menu-item index="/collector/sqltask">SQL任务</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="alert">
        <template #title>
          <el-icon>
            <Bell />
          </el-icon>
          <span>告警中心</span>
        </template>
        <el-menu-item index="/alert/rules">告警规则</el-menu-item>
        <el-menu-item index="/alert/type">告警类型</el-menu-item>
        <el-menu-item index="/alert/prometheus">普米管理</el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch, onMounted, PropType, Ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { RouteLocationNormalizedLoaded } from 'vue-router'
import { Warning, Bell, Timer, Files } from '@element-plus/icons-vue'

export default defineComponent({
  name: 'Sidebar',
  components: { Warning, Bell, Timer,Files },
  props: { collapsed: { type: Boolean as PropType<boolean>, default: false } },
  setup(props) {
    const route = useRoute() as RouteLocationNormalizedLoaded
    const router = useRouter()
    const active: Ref<string> = ref(route.path)

    onMounted(() => {
      active.value = route.path
    })

    watch(() => route.path, (p: string) => {
      active.value = p
    })

    function handleSelect(key: string): void {
      active.value = key
      router.push(key)
    }

    return { active, handleSelect }
  }
})
</script>

<style scoped>
.sidebar-wrap {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  background: #545c64;
  /* dark sidebar background */
  color: #fff;
}

.brand {
  padding: 20px;
  font-size: 18px;
  font-weight: 700;
  text-align: center;
  color: #fff;
}

.sidebar-wrap .brand {
  transition: opacity 0.18s ease, transform 0.18s ease;
}

.sidebar-wrap.is-collapsed .brand {
  opacity: 0;
  transform: translateX(-6px);
  pointer-events: none;
}

.el-menu-vertical-demo {
  background: transparent;
  /* menu already uses props for colors */
}

/* Animate menu label visibility when collapsed */
.el-menu-vertical-demo ::v-deep(.el-menu-item__content),
.el-menu-vertical-demo ::v-deep(.el-sub-menu__title) {
  display: flex;
  align-items: center;
  gap: 8px;
}

.el-menu-vertical-demo ::v-deep(.el-menu-item__text),
.el-menu-vertical-demo ::v-deep(.el-sub-menu__title .el-sub-menu__title-text),
.el-menu-vertical-demo ::v-deep(.el-sub-menu__title span) {
  transition: opacity 0.18s ease, transform 0.18s ease;
  white-space: nowrap;
}

.sidebar-wrap.is-collapsed ::v-deep(.el-menu-item__text),
.sidebar-wrap.is-collapsed ::v-deep(.el-sub-menu__title span) {
  opacity: 0;
  transform: translateX(-6px);
  pointer-events: none;
}

.sidebar-wrap ::v-deep(.el-menu-item__text),
.sidebar-wrap ::v-deep(.el-sub-menu__title span) {
  opacity: 1;
  transform: translateX(0);
}

/* footer removed */

/* Keep mobile behavior: hide aside on very small screens */
@media (max-width: 700px) {
  :host ::v-deep(.app-aside) {
    display: none;
  }
}
</style>
