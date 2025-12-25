import { createRouter, createWebHistory } from 'vue-router'
import AlertRule from '../pages/AlertRule/AlertRule.vue'
import AlertType from '../pages/AlertType/AlertType.vue'
import PrometheusConfig from '../pages/PrometheusConfig/PrometheusConfig.vue'
import portal from '../pages/portal/portal.vue'
import SqlTask from '../pages/collector/SqlTask/index.vue'
import JobTask from '../pages/job/jobtask/index.vue'

const routes = [
  { path: '/', redirect: '/alert/rules' },
  { path: '/portal', name: 'portal', component: portal, meta: { title: '首页' } },
  { path: '/jobtask', name: 'JobTask', component: JobTask, meta: { title: '定时任务管理' } },
  { path: '/job', name: 'Job', meta: { title: '定时任务' } ,
    children:[
      { path: 'jobtask', name: 'JobTask', component: JobTask, meta: { title: '任务列表' } },
    ]
  },
  { path: '/collector', name: 'Collector', meta: { title: '收集器' } ,
    children:[
      { path: 'sqltask', name: 'SqlTask', component: SqlTask, meta: { title: 'SQL采集' } },
    ]
  },
  {
    path: '/alert',
    name: 'Alert',
    meta: { title: '告警中心' },
    children: [
      { path: 'rules', name: 'AlertRule', component: AlertRule, meta: { title: '告警规则' } },
      { path: 'type', name: 'AlertType', component: AlertType, meta: { title: '告警类型' } },
      { path: 'prometheus', name: 'PrometheusConfig', component: PrometheusConfig, meta: { title: '普米管理' } },
    ]
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
