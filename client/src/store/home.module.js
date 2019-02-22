import { GamesService, HeroesService } from "@/common/api.service";
import { FETCH_GAME, FETCH_GAMES, FETCH_HEROES } from "./actions.type";
import {
  FETCH_END,
  FETCH_START,
  SET_HEROES,
  SET_GAME,
  UPDATE_GAMES_IN_LIST
} from "./mutations.type";

const state = {
  game: {},
  games: [],
  heroes: [],
  isLoading: true,
  gamesCount: 0,
  heroesCount: 0
};

const getters = {
  game(state) {
    return state.game;
  },
  gamesCount(state) {
    return state.gamesCount;
  },
  heroesCount(state) {
    return state.heroesCount;
  },
  games(state) {
    return state.games;
  },
  isLoading(state) {
    return state.isLoading;
  },
  heroes(state) {
    return state.heroes;
  }
};

const actions = {
  [FETCH_GAMES]({ commit }) {
    commit(FETCH_START);
    return GamesService.all()
      .then(({ data }) => {
        commit(FETCH_END, data.message);
      })
      .catch(error => {
        throw new Error(error);
      });
  },
  [FETCH_HEROES]({ commit }, params) {
    return HeroesService.all(params)
      .then(({ data }) => {
        commit(SET_HEROES, data.message);
      })
      .catch(error => {
        throw new Error(error);
      });
  },
  [FETCH_GAME]({ commit }, params) {
    return GamesService.get(params)
      .then(({ data }) => {
        commit(SET_GAME, data.message);
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
  [FETCH_END](state, games) {
    state.games = games;
    state.gamesCount = games.length;
    state.isLoading = false;
  },
  [UPDATE_GAMES_IN_LIST](state, games) {
    state.games = games;
  },
  [SET_GAME](state, data) {
    state.game = data;
  },
  [SET_HEROES](state, data) {
    state.heroes = data;
    state.heroesCount = data.length;
    state.isLoading = false;
  }
};

export default {
  state,
  getters,
  actions,
  mutations
};
