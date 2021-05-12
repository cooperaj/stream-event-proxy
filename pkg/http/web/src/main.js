import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VueNativeSock from 'vue-native-websocket'
import PureCss from 'purecss'

import router from './router'
import store from './store'

import NewFollower from "./components/NewFollower.vue"

Vue.config.productionTip = false

Vue.use(VueRouter)

let socketAddress = '//' + document.location.hostname + ':' + document.location.port + '/ws'
Vue.use(VueNativeSock, socketAddress, { store: store, format: 'json', reconnection: true })

Vue.component('pure-css', PureCss)

Vue.component('new-follower', NewFollower)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
