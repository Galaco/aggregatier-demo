<template>
  <div class="game-page">
    Game
    <div v-for="hero in heroes" :value="hero.name" :key="hero.id">
      <img :src="hero.icon_url" />
      <p>{{ hero.name }}</p>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { FETCH_HEROES } from "@/store/actions.type";
export default {
  name: "game",
  components: {},
  mounted() {
    this.$store.dispatch(FETCH_HEROES, this.$route.params.gameId);
  },
  computed: {
    ...mapGetters(["heroes"])
  }
};
</script>
