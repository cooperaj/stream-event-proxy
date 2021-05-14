<template>
  <transition name="fly-by">
    <div 
        v-if="show" 
        v-on:animationend="finished"
        class="follower"
    >
      <div class="follower__cayde">
        <img src="../assets/Torso.png">
      </div>
      <div class="follower__box">
        <h1>Hey, a new follower... that's cute!</h1>
        <h2>{{ followerData.name }}</h2>
      </div>
      <div class="follower__caydeUpper">
        <img src="../assets/UpperBody.png">
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  props: {
    followerData: Object
  },

  data() {
    return {
      show: false,
    }
  },

  watch: {
    followerData: function() {
      this.show = true
      this.playAudio()
    }
  },

  created() {
    // loading images to stop FOC when receiving the first alert
    let torso = new Image()
    torso.src = require('../assets/Torso.png')
    let upperBody = new Image()
    upperBody.src = require('../assets/UpperBody.png')
    let box = new Image()
    box.src = require('../assets/Box.png')
  },

  methods: {
    playAudio() {
      var audio = new Audio('/audio/youre_my_favourite.ogg')
      setTimeout(() => {
        audio.play()
      }, 1000);
    },

    finished(animationEvent) {
      // we animate 3 elements so make sure to just listen to the finished event of one.
      if (animationEvent.srcElement.classList.contains('follower')) {
        this.show = false
        this.$store.dispatch('eventResolved')
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import url('https://fonts.googleapis.com/css?family=Russo+One');

.follower {
  font-family: 'Russo One';
  width: 782px;
  margin: 0 auto;

  &__cayde, &__caydeUpper {
    position: absolute;
    transform: translateX(-90px);
  }

  &__box {
    position: absolute;
    background: url('../assets/Box.png');
    margin-top: 388px;
    width: 782px;
    height: 202px;
    color: white;
    text-align: center;

    h1 {
      font-size: 2.5rem;
      margin: 2.5rem 0.5rem 0rem;
    }

    h2 {
      font-size: 4rem;
      margin: 0.5rem 0;
    }
  }
}

.fly-by-enter-active {
  animation: slide-through 6s;

  .follower__caydeUpper, .follower__cayde {
    animation: cayde-slide-right 6s linear;
  }
}

.fly-by-leave-active {
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