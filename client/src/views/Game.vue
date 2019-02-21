<template>
  <div class="game-page">
    <div class="container page-header">
      <div class="row">
        <div class="col-sm-10">
          <h4>Game</h4>
        </div>
        <div class="col-sm-2 text-right">
          <a @click="$router.go(-1)">
            <button type="button" class="btn btn-secondary">Back</button>
          </a>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row">
        <div
          v-for="hero in heroes"
          :value="hero.name"
          :key="hero.id"
          class="card hero"
        >
          <img :src="hero.icon_url" class="card-img-top" :alt="hero.name" />
          <div class="card-body">
            <h6 class="card-title">{{ hero.name }}</h6>
          </div>
        </div>
      </div>
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
