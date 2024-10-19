import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import RulesPage from '@/components/RulesPage.vue'
import PlayersPage from '@/components/PlayersPage.vue'

const routes = [
  { path: '/', component: HomePage },
  { path: '/rules', component: RulesPage },
  { path: '/add_players', component: PlayersPage },
]
const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
