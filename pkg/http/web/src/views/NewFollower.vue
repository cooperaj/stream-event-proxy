<template>
  <transition name="bounce">
    <div 
        v-if="show" 
        v-on:animationend="show = false"
        class="follower"
    >
      <div class="follower__cayde">
        <img src="../assets/Torso.png">
      </div>
      <div class="follower__box">
        <h1>Hey, a new follower... that's cute!</h1>
        <h2>{{ follower }}</h2>
      </div>
      <div class="follower__caydeUpper">
        <img src="../assets/UpperBody.png">
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  data() {
    return {
      connection: null,
      connectionTimeout: 250,
      show: false,
      follower: 'testFollower'
    }
  },

  created() {
    console.log('Starting connection to WebSocket Server')
    this.connect()

    // loading images to stop FOC when receiving the first alert
    let torso = new Image()
    torso.src = require('../assets/Torso.png')
    let upperBody = new Image()
    upperBody.src = require('../assets/UpperBody.png')
    let box = new Image()
    box.src = require('../assets/Box.png')
  },

  methods: {
    connect() {
      var protocol = document.location.protocol == 'https:' ? 'wss' : 'ws'
      this.connection = new WebSocket(protocol + '://' + document.location.hostname + ':' + document.location.port + '/ws')

      this.connection.onmessage = (event) => {
        var follow = JSON.parse(event.data);

        if (follow.type == 'follow') {
          this.follower = follow.data.user_name
          this.show = true
        }
      }

      this.connection.onopen = () => {
        console.log('Successfully connected to the websocket server...')
        this.connectionTimeout = 250
      }

      this.connection.onclose = (event) => {
        console.log('Socket is closed. Reconnect will be attempted in ' + this.connectionTimeout/1000 + ' seconds.', event.reason);
        setTimeout(this.connect, Math.min(10000, this.connectionTimeout += this.connectionTimeout))
      }
    },
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Russo+One');

.follower {
  font-family: 'Russo One';
  width: 782px;
  margin: 0 auto;
}

.follower__cayde, .follower__caydeUpper, .follower__box {
  position: absolute;
}

.follower__cayde, .follower__caydeUpper {
  transform: translateX(-90px)
}

.follower__box {
  background: url('../assets/Box.png');
  margin-top: 388px;
  width: 782px;
  height: 202px;
  color: white;
  text-align: center;
}

.follower__box h1 {
  font-size: 2.5rem;
  margin: 2.5rem 0.5rem 0rem;
}

.follower__box h2 {
  font-size: 4rem;
  margin: 0.5rem 0;
}

.bounce-enter-active {
  animation: slide-through 6s;
}
.bounce-enter-active .follower__caydeUpper, .bounce-enter-active .follower__cayde {
  animation: cayde-slide-right 6s linear;
}
.bounce-leave-active {
  opacity: 0;
}

@keyframes slide-through {
  0% {
    transform: translateX(-1200px);
    opacity: 0;
    animation-timing-function: cubic-bezier( 0, 1, 0.32, 0.97 );
  }
  20% {
    transform: translateX(-30px);
    opacity: 1;
    animation-timing-function: linear;
  }
  90% {
    transform: translateX(30px);
    opacity: 1;
    animation-timing-function: cubic-bezier( 1, 0.03, 1, 0.68 );
  }
  100% {
    opacity: 0;
    transform: translateX(1200px);
  }
}
@keyframes cayde-slide-right {
  0% {
    transform: translateX(-90px);
  }
  100% {
    transform: translateX(0px);
  }
}
</style>