<template>
  <footer class="footer">
    <div class="container">
      <router-link class="logo-font" :to="{ name: 'home', params: {} }">
        Aggregatier
      </router-link>
      <span class="attribution">
        A Go x Vue application by
        <a rel="noopener noreferrer" target="blank" href="https://galaco.me">
          Galaco</a
        >.
      </span>
    </div>
  </footer>
</template>

<script>
export default {
  name: "Footer"
};
</script>
