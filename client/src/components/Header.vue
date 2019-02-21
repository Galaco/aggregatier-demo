<template>
  <nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
      <div class="navbar-header">
        <router-link class="navbar-brand" :to="{ name: 'home' }">
          Aggregatier
        </router-link>
      </div>
      <div>
        <ul class="navbar navbar-nav navbar-right">
          <li>
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
  </nav>
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
