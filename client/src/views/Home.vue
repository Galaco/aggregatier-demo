<template>
  <div class="home-page">
    <div class="container page-header">
      <div class="row">
        <div class="col-sm-12">
          <h4>Home</h4>
        </div>
      </div>
    </div>
    <div class="container">
      <div class="row">
        <div v-for="game in games" :key="game.id" class="col-sm-4">
          <div class="card">
            <div class="card-body">
              <h6 class="card-title">
                {{ game.name }}
              </h6>
              <router-link :to="{ name: 'game', params: { gameId: game.id } }">
                Open
              </router-link>
            </div>
          </div>
        </div>
      </div>
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
