<template>
  <div class="home-page">
    Home
    <div class="md-tabs">
      <ul
        class="tabs"
        data-responsive-accordion-tabs="tabs small-accordion medium-tabs large-tabs"
        id="collapsing-tabs"
        data-allow-all-closed="true"
        data-multi-expand="true"
      >
        <li
          v-for="game in games"
          :key="game.id"
        >
          <router-link :to="{ name: 'game', params: { gameId: game.id } }">
            {{ game.name }}
          </router-link>
        </li>
        <div class="slide"></div>
      </ul>
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
