<template>
  <div class="home-page">
    Egghat
    <select name="Game list">
      <option v-for="game in games" :value="game.name" :key="game.id">
        {{ game.name }}
      </option>
    </select>
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
