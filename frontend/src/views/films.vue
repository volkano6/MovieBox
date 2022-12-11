<template>
  <div>
    <Nav></Nav>
    
    <Footer></Footer>
  </div>
</template>

<script>
import Nav from "../components/nav.vue";
import Footer from "../components/footer.vue";

export default {
  name: "films",
  components: {
    Nav,
    Footer,
  },
};
</script>

<style>
</style>