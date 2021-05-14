import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    socket: {
      isConnected: false,
      reconnectError: false,
    },
    currentEvent: null,
    pendingEvents: []
  },
  mutations: {
    SOCKET_ONOPEN (state, event)  {
      Vue.prototype.$socket = event.currentTarget
      state.socket.isConnected = true
    },
    SOCKET_ONCLOSE (state)  {
      state.socket.isConnected = false
    },
    SOCKET_ONERROR (state, event)  {
      console.error(state, event)
    },
    // default handler called for all methods
    SOCKET_ONMESSAGE (state, message)  {
      state.pendingEvents.push(message)
      if (state.currentEvent == null) {
        state.currentEvent = state.pendingEvents.shift()
      }
    },
    // mutations for reconnect methods
    SOCKET_RECONNECT(state, count) {
      console.info(state, count)
    },
    SOCKET_RECONNECT_ERROR(state) {
      state.socket.reconnectError = true
    },
    EVENT_RESOLVED(state) {
      state.currentEvent = null
      if (state.pendingEvents.length > 0) {
        state.currentEvent = state.pendingEvents.shift()
      }
    }
  },
  actions: {
    eventResolved({commit}) {
      setTimeout(() => {
        commit('EVENT_RESOLVED')
      }, 1000) // wait a second to before (possibly) queuing the next event
    }
  },
})
