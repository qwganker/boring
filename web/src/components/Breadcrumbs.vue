<template>
  <el-breadcrumb class="app-breadcrumb" separator="/">
    <el-breadcrumb-item v-for="(item, idx) in crumbs" :key="idx">{{ item }}</el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

export default {
  name: 'Breadcrumbs',
  setup() {
    const route = useRoute()

    const crumbs = computed(() => {
      const matched = route.matched || []
      if (!matched.length) return ['首页']
      return matched
        .filter(r => r.meta && r.meta.title)
        .map(r => r.meta.title)
    })

    return { crumbs }
  }
}
</script>

<style scoped>
</style>
