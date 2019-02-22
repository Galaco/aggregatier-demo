<template>
  <div class="game-page">
    <div class="container page-header">
      <div class="row">
        <div class="col-sm-10">
          <h4>{{ game.name }}</h4>
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
        <router-link
          :to="{
            name: 'tierlist/create',
            params: { gameId: this.$route.params.gameId }
          }"
        >
          <button class="btn btn-primary">Create Tierlist</button>
        </router-link>
      </div>
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
import { FETCH_HEROES, FETCH_GAME } from "@/store/actions.type";
export default {
  name: "game",
  components: {},
  mounted() {
    this.$store.dispatch(FETCH_HEROES, this.$route.params.gameId);
    this.$store.dispatch(FETCH_GAME, this.$route.params.gameId);
  },
  computed: {
    ...mapGetters(["heroes"]),
    ...mapGetters(["game"])
  }
};
</script>
