<template>
  <div class="game-page">
    <div class="container page-header">
      <div class="row">
        <div class="col-sm-10">
          <h4>Create Tierlist</h4>
        </div>
        <div class="col-sm-2 text-right">
          <a @click="$router.go(-1)">
            <button type="button" class="btn btn-secondary">Back</button>
          </a>
        </div>
      </div>
    </div>

    <div class="container">
      <div
        v-for="(tier, key) in tiers"
        :key="key"
        class="row row-eq-height row-tier"
      >
        <div class="col-sm-1">
          <h5>{{ tier.name }}</h5>
        </div>
        <div class="col-sm-11">
          <draggable class="row" v-model="tierHeroes" :options="{group:'TIER'}">
            <div v-for="(hero, key) in tier.heroes" :key="key">
              <div class="card hero">
                <img
                  :src="hero.icon_url"
                  class="card-img-top"
                  :alt="hero.name"
                />
                <div class="card-body">
                  <p class="card-title">{{ hero.name }}</p>
                </div>
              </div>
            </div>
          </draggable>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import draggable from "vuedraggable";
import { FETCH_HEROES, FETCH_TIERLIST_TIERS } from "@/store/actions.type";
export default {
  name: "game",
  components: {
    draggable
  },
  mounted() {
    Promise.all([
      this.$store.dispatch(FETCH_HEROES, this.$route.params.gameId),
      this.$store.dispatch(FETCH_TIERLIST_TIERS, this.$route.params.gameId)
    ]).then(() => {
      console.log(this.$store.state);
      let unassigned = this.$store.state.tierlist.tiers.filter(boundTier => {
        return boundTier.id === -1;
      })[0];
      this.$store.state.home.heroes.forEach(hero => {
        unassigned.heroes.push(hero);
      });
    });
  },
  computed: {
    ...mapGetters(["tiers", "heroes"]),
    tierHeroes: {
      get() {
        console.log(this.$store.state.tierlist.tiers);
        return this.$store.state.tierlist.tiers.map(tier => {
          return tier.heroes;
        });
      },
      set(value) {
        this.$store.commit("updateList", value);
      }
    }
  },
  methods: {
    log: function(evt) {
      window.console.log(evt);
    }
  }
};
</script>
