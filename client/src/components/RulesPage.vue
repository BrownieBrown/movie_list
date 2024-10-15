<template>
  <div class="rules_page">
    <h1>The Rules of the movie list are as follows</h1>
    <p>
      1. During the first round each randomly selected player may add a chosen
      amount of movies.
    </p>
    <p>
      2. During the second round each player may strike out a selected amount of
      movies.
    </p>
    <p>
      3. During the third round the players watch trailers for all still
      available movies.
    </p>
    <p>
      4. During the fourth and final round the players each cast their vote for
      their movie of choice.
    </p>
    <p>5. The movie with the most votes wins.</p>
    <p>In case of a tie the players vote again!</p>
    <button @click="navigateToAddPlayers">ADD PLAYERS</button>
  </div>
</template>

<script lang="ts">
export default {
  methods: {
    navigateToAddPlayers() {
      this.$router.push('/add_players')
    },
  },
}
</script>

<style scoped>
.rules_page {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  text-align: center;
  background-image: url('../assets/img/rules_background.jpg');
  background-size: cover;
  background-position: center;
}

.rules_page p:last-of-type {
  color: grey; /* Change this to your desired color */
}

h1 {
  color: black;
  font-size: 2rem;
}

p {
  font-size: 1.2rem;
  margin: 20px 0;
  color: black;
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
