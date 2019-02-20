<template>
  <div class="md-tabs">
    <div class="mobile-nav-bar title-bar">
      <div class="title-bar-center">
        <router-link class="title-bar-text" :to="{ name: 'home' }">
          Aggregatier
        </router-link>
        <ul class="nav navbar-nav pull-xs-right">
          <li class="nav-item">
            <router-link
              class="nav-link"
              active-class="active"
              exact
              :to="{ name: 'home' }"
            >
              Home
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "Header",
  computed: {
    ...mapGetters([])
  }
};
</script>
