<template>
  <div class="welcome-container" ref="container">
    <div class="content">
      <h1 class="title">ようこそ</h1>
      <p class="subtitle">素敵なプロジェクトへ</p>
    </div>
  </div>
</template>

<script>
import confetti from 'canvas-confetti'

export default {
  name: 'WelcomePage',
  mounted() {
    this.triggerConfetti()
    this.startBackgroundAnimation()
  },
  methods: {
    triggerConfetti() {
      const duration = 3 * 1000
      const end = Date.now() + duration

      const colors = ['#ff0000', '#00ff00', '#0000ff', '#ffff00', '#ff00ff']

      ;(function frame() {
        confetti({
          particleCount: 5,
          angle: 60,
          spread: 55,
          origin: { x: 0 },
          colors: colors
        })
        confetti({
          particleCount: 5,
          angle: 120,
          spread: 55,
          origin: { x: 1 },
          colors: colors
        })

        if (Date.now() < end) {
          requestAnimationFrame(frame)
        }
      }())
    },
    startBackgroundAnimation() {
      const container = this.$refs.container
      let hue = 0

      const animate = () => {
        hue = (hue + 0.1) % 360
        container.style.background = `linear-gradient(45deg, 
          hsl(${hue}, 70%, 50%), 
          hsl(${(hue + 60) % 360}, 70%, 50%), 
          hsl(${(hue + 120) % 360}, 70%, 50%))`
        requestAnimationFrame(animate)
      }

      animate()
    }
  }
}
</script>

<style scoped>
.welcome-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: background 0.5s;
  overflow: hidden;
}

.content {
  text-align: center;
  color: white;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
  animation: fadeIn 2s ease-out;
}

.title {
  font-family: 'Tsukimi Rounded', sans-serif;
  font-size: 5.5rem;
  margin: 0;
  font-weight: 700;
  letter-spacing: 0.05em;
  animation: bounce 2s ease-in-out infinite;
}

.subtitle {
  font-family: 'Tsukimi Rounded', sans-serif;
  font-size: 2.2rem;
  margin-top: 1rem;
  opacity: 0.9;
  font-weight: 300;
  letter-spacing: 0.1em;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}
</style> 