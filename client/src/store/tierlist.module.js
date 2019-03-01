import { FETCH_TIERLIST_TIERS } from "./actions.type";
import { FETCH_END, FETCH_START, ADD_HERO_TO_TIER } from "./mutations.type";
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
    state.tiers = tiers;
    state.tiersCount = tiers.length;
    state.tiers.forEach(tier => {
      tier.heroes = [];
    });
    state.isLoading = false;
  },
  [ADD_HERO_TO_TIER](state, tier, hero) {
    console.log(tier, hero);
  }
};

export default {
  state,
  getters,
  actions,
  mutations
};
