import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";
import { API_URL } from "@/common/config";

const ApiService = {
  init() {
    Vue.use(VueAxios, axios);
    Vue.axios.defaults.baseURL = API_URL;
  },

  query(resource, params) {
    return Vue.axios.get(resource, params).catch(error => {
      throw new Error(`ApiService ${error}`);
    });
  },

  get(resource, slug = "") {
    return Vue.axios.get(`${resource}/${slug}`).catch(error => {
      throw new Error(`ApiService ${error}`);
    });
  },

  post(resource, params) {
    return Vue.axios.post(`${resource}`, params);
  }
};

export default ApiService;

export const GamesService = {
  all() {
    return ApiService.get("games", "all");
  },
  get(id) {
    return ApiService.get("game", id);
  }
};

export const HeroesService = {
  all(id) {
    return ApiService.get("heroes/all", id);
  }
};

export const TierlistService = {
  tiers() {
    return ApiService.get("tierlist/tiers/all")
  }
};
