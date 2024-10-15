<template>
  <div class="homepage">
    <h1>Welcome to the Movie List</h1>
    <button @click="viewRules">VIEW THE RULES</button>
  </div>
</template>

<script lang="ts">
export default {
  methods: {
    viewRules() {
      this.$router.push('/rules')
    },
  },
}
</script>

<style scoped>
html,
body,
#app {
  height: 100%;
  margin: 0;
}

.homepage {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  text-align: center;
  background-image: url('../assets/img/background_homepage.jpg');
  background-size: cover;
  background-position: center;
}

h1 {
  color: black;
  font-size: 2rem;
}

button {
  padding: 10px 20px;
  margin-top: 20px;
  background-color: black;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>
