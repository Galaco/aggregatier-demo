<template>
  <div class="home-page">
    Home
    <div v-for="game in games" :value="game.name" :key="game.id">
      <router-link :to="{ name: 'game', params: { gameId: game.id } }">
        {{ game.name }}
      </router-link>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { FETCH_GAMES } from "@/store/actions.type";
export default {
  name: "home",
  components: {},
  mounted() {
    this.$store.dispatch(FETCH_GAMES);
  },
  computed: {
    ...mapGetters(["games"])
  }
};
</script>
