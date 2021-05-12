import VueRouter from 'vue-router';
import Home from '@/views/Home.vue'
import Alerts from '@/views/Alerts.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/alerts/:types?',
    name: 'alerts-display',
    component: Alerts,
    props: (route) => { 
      if (route.params.types != undefined ) {
        return { types: route.params.types.split(',') }
      }

      return { types: [] }
    }
  }
]

const router = new VueRouter({
  routes
})

export default router
