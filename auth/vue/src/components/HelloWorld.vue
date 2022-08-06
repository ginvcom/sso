<template>
  <div class="center-on-page">
    <div class="logo">
      <div class="logo-inner">
        <div class="round"></div>
        <div class="line">
        </div>
        <div class="line line2">
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import WebLogo from './WebLogo.vue'
import { ref } from 'vue'

defineProps<{ msg: string }>()

const count = ref(0)
</script>
<style scoped>
*{
  box-sizing: border-box;
}

body {
  background: #ecf0f1;
}

.center-on-page {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.ball {
  position: relative;
  animation: fall .25s ease-in-out, shake 1.25s cubic-bezier(.36,.07,.19,.97) 3;
}

@keyframes shake {
  0% {
    transform: translate(0, 0) rotate(0);
  }

  20% {
    transform: translate(-10px, 0) rotate(-20deg);
  }

  30% {
    transform: translate(10px, 0) rotate(20deg);
  }

  50% {
    transform: translate(-10px, 0) rotate(-10deg);
  }

  60% {
    transform: translate(10px, 0) rotate(10deg);
  }

  100% {
    transform: translate(0, 0) rotate(0);
  }
}

@keyframes fall {
  0% {
    top: -200px
  }

  60% {
    top: 0
  }

  80% {
    top: -20px
  }

  100% {
    top: 0
  }
}
</style>
