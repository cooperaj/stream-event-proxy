<template>
  <div id="home">
    <h1>Stream Alert Proxy</h1>
    <form id="alert-choice" class="pure-form">
      <button v-on:click="showAlertDisplay" :disabled="checkedAlerts.length == 0" class="pure-button">Open selected Alerts display</button>
      <label for="newFollower" class="pure-checkbox">
        <input id="newFollower" type="checkbox" value="newFollower" v-model="checkedAlerts" /> New Follower</label>
      <label for="newSubscriber" class="pure-checkbox">
        <input id="newSubscriber" type="checkbox" value="newSubscriber" v-model="checkedAlerts" disabled /> New Subscriber</label>
    </form>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        checkedAlerts: ['newFollower']
      }
    },

    methods: {
      showAlertDisplay: function() {
        this.$router.push({ name: 'alerts-display', params: { types: this.checkedAlerts.toString() }})
      }
    }
  }
</script>

<style lang="scss" scoped>

#home {
  font-family: "Raleway", sans-serif;
  max-width: 768px;
  margin: 3rem auto 0 auto;

  form#alert-choice {
    background-color: #444;
    color: white;
    padding: 1rem;
  }

  button {
    float: right;
  }
}

</style>
