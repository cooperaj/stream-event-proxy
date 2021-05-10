<template>
  <div id="alerts">
    <new-follower v-if="types.includes('newFollower')" :followerData=followerData />
  </div>
</template>

<script>
export default {
  props: {
    types: Array
  },

  data() {
    return {
      followerData: {},
      currentEvent: null
    }
  },

  created() {
    this.unsubscribe = this.$store.subscribe((mutation, state) => {
      if (mutation.type === 'SOCKET_ONMESSAGE') {
        this.currentEvent = state.currentEvent
        this.currentEvent.created_at = Date.now()
      }
    });
  },

  beforeDestroy() {
    this.unsubscribe();
  },

  watch: {
    currentEvent() {
      if (this.currentEvent == null) {
        return
      }

      switch (this.currentEvent.type) {
        case 'follow' :
          this.followerData = {
            name: this.currentEvent.data.user_name,
            updated_at: this.currentEvent.created_at
          }
          break
      }
    }
  }
}
</script>