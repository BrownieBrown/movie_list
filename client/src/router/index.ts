import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import RulesPage from '@/components/RulesPage.vue'

const routes = [
  { path: '/', component: HomePage },
  { path: '/rules', component: RulesPage },
]
const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
