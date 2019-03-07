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
        v-for="tier in tiers"
        :key="tier.id"
        class="row row-eq-height row-tier"
      >
        <div class="col-sm-1">
          <h5>{{ tier.name }}</h5>
        </div>
        <div class="col-sm-11" :key="tier.id">
          <draggable
            class="row"
            v-model="tierHeroes"
            :options="{ group: 'TIER' }"
            @add="add($event, tier)"
            @remove="remove($event, tier)"
            :key="tier.id"
          >
            <div v-for="hero in tier.heroes" :key="hero.id" class="card hero">
              <img :src="hero.icon_url" class="card-img-top" :alt="hero.name" />
              <div class="card-body">
                <p class="card-title">{{ hero.name }}</p>
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
import {
  ADD_HERO_TO_TIER,
  REMOVE_HERO_FROM_TIER
} from "@/store/mutations.type";
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
        return this.$store.state.tierlist.tiers.map(tier => {
          return tier.heroes;
        });
      },
      set() {}
    }
  },
  methods: {
    add: function(evt, tier) {
      console.log(evt);
      console.log(evt.from);
      // @TODO this does not produce the right index
      // It needs the index of the hero in the current tier, not the complete heroes list
      console.log(this.$store.state.tierlist.tiers[evt.from.__vue__]);
      console.log(this.$store.state.tierlist.tiers[evt.from.__vue__.context.index].heroes);
      const hero = this.$store.state.tierlist.tiers[evt.from.__vue__.context.index].heroes[evt.oldIndex];
      console.log(hero);
      this.$store.commit(ADD_HERO_TO_TIER, {
        tierId: tier.id,
        hero: hero
      });
    },
    remove: function(evt, tier) {
      this.$store.commit(REMOVE_HERO_FROM_TIER, {
        tierId: tier.id,
        hero: tier.heroes[evt.oldIndex]
      });
    }
  }
};
</script>
