import VueRouter from 'vue-router';
import Home from '@/views/Home.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/new-follower',
    name: 'New Follower',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "newFollower" */ '../views/NewFollower.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
