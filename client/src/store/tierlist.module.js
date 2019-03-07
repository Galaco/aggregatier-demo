import { FETCH_TIERLIST_TIERS } from "./actions.type";
import {
  FETCH_END,
  FETCH_START,
  ADD_HERO_TO_TIER,
  REMOVE_HERO_FROM_TIER
} from "./mutations.type";
import { TierlistService } from "@/common/api.service";

const state = {
  tiers: [],
  tiersCount: 0
};

const getters = {
  tiersCount(state) {
    return state.tiersCount;
  },
  tiers(state) {
    return state.tiers;
  }
};

const actions = {
  [FETCH_TIERLIST_TIERS]({ commit }, params) {
    commit(FETCH_START);
    return TierlistService.tiers(params)
      .then(({ data }) => {
        commit(FETCH_END, data.message);
      })
      .catch(error => {
        throw new Error(error);
      });
  }
};

/* eslint no-param-reassign: ["error", { "props": false }] */
const mutations = {
  [FETCH_START](state) {
    state.isLoading = true;
  },
  [FETCH_END](state, tiers) {
    tiers.forEach(tier => {
      tier.heroes = [];
    });
    state.tiers = tiers;
    state.tiersCount = tiers.length;
    state.tiers.push({
      id: -1,
      name: "N/A",
      heroes: []
    });
    state.isLoading = false;
  },
  [ADD_HERO_TO_TIER](state, props) {
    state.tiers.forEach((tier, tierIdx) => {
      if (tier.id !== props.tierId) {
        return;
      }
      state.tiers[tierIdx].heroes.push(props.hero);
    });
  },
  [REMOVE_HERO_FROM_TIER](state, props) {
    state.tiers.forEach((tier, tierIdx) => {
      if (tier.id !== props.tierId) {
        return;
      }
      let idx = -1;
      state.tiers[tierIdx].heroes.forEach((hero, key) => {
        if (hero.id === props.hero.id) {
          idx = key;
        }
      });
      if (idx !== -1) {
        state.tiers[tierIdx].heroes.splice(idx, 1);
      }
      console.log(idx);
    });
    console.log(state.tiers);
  }
};

export default {
  state,
  getters,
  actions,
  mutations
};
