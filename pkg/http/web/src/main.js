import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import router from './router'
import { MotionPlugin } from '@vueuse/motion'

Vue.config.productionTip = false

Vue.use(VueRouter)
Vue.use(MotionPlugin)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
