import { createRouter, createWebHistory } from 'vue-router'
import AlertRule from '../pages/AlertRule/AlertRule.vue'
import AlertType from '../pages/AlertType/AlertType.vue'
import PrometheusConfig from '../pages/PrometheusConfig/PrometheusConfig.vue'
import portal from '../pages/portal/portal.vue'


const routes = [
  { path: '/', redirect: '/alert/rules' },
  { path: '/alert/rules', name: 'AlertRule', component: AlertRule },
  { path: '/alert/type', name: 'AlertType', component: AlertType },
  { path: '/alert/prometheus', name: 'PrometheusConfig', component: PrometheusConfig },
  { path: '/portal', name: 'portal', component: portal },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
