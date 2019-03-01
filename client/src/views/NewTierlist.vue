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
      <div v-for="(tier, key) in tiers" :key="key" class="row row-tier">
        <div class="col-sm-1">{{ tier.name }}</div>
        <draggable
          class="col-sm-11"
          v-model="tierHeroes"
        >
          <div v-for="(hero, key) in tier.heroes" :key="key">
            <div class="card hero">
              <img :src="hero.icon_url" class="card-img-top" :alt="hero.name" />
              <div class="card-body">
                <p class="card-title">{{ hero.name }}</p>
              </div>
            </div>
          </div>
        </draggable>
      </div>
    </div>

    <div class="container">
      <draggable
        class="row"
        :list="heroes"
      >
        <div v-for="hero in heroes" :key="hero.id">
          <div class="card hero">
            <img :src="hero.icon_url" class="card-img-top" :alt="hero.name" />
            <div class="card-body">
              <p class="card-title">{{ hero.name }}</p>
            </div>
          </div>
        </div>
      </draggable>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import draggable from "vuedraggable";
import { FETCH_HEROES, FETCH_TIERLIST_TIERS } from "@/store/actions.type";
import { ADD_HERO_TO_TIER } from "@/store/mutations.type";
export default {
  name: "game",
  components: {
    draggable
  },
  data: () => {
    return {
      targetHero: "",
      targetTier: ""
    };
  },
  mounted() {
    this.$store.dispatch(FETCH_HEROES, this.$route.params.gameId);
    this.$store.dispatch(FETCH_TIERLIST_TIERS, this.$route.params.gameId);
  },
  computed: {
    ...mapGetters(["tiers", "heroes"]),
    tierHeroes: {
      get() {
        console.log(this.$store.state.tierlist.tiers);
        return this.$store.state.tierlist.tiers.map((tier) => {
          return tier.heroes;
        })
      },
      set(value) {
        this.$store.commit('updateList', value);
      }
    }
  },
  methods: {
    onMove({ relatedContext, draggedContext }) {
      console.log(relatedContext);
      console.log(draggedContext);
    },
    chooseHero: function(event) {
      console.log(event);
      this.targetHero = event.from.id;
    },
    newHero: function(event) {
      console.log(event);
      this.targetTier = event.to.id;
      this.$store.dispatch(ADD_HERO_TO_TIER, this.targetTier, this.targetHero);
    },
    log: function(evt) {
      window.console.log(evt);
    }
  }
};
</script>
