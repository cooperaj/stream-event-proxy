<template>
  <div id="alerts">
    <new-follower v-if="types.includes('newFollower')" :followerData=followerData />
  </div>
</template>

<script>

import { mapState } from 'vuex';

export default {
  props: {
    types: Array
  },

  data() {
    return {
      followerData: {},
    }
  },

  computed: mapState(['currentEvent']),

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