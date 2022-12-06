<template>
  <div>
    <Nav></Nav>
    <div class="container">
      <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>
  
<script>
import axios from 'axios'
import Nav from "../components/nav.vue"
import Footer from "../components/footer.vue"
import Card from "../components/movie_card.vue"

export default {
  name: "profile",
  components: {
    Nav,
    Footer,
    Card
  },
  data() {
    return {
      user: null
    }
  },
  async created() {
    const response = await axios.get("api/profile");
    if (response.data.status == "error") {
      this.$router.push("/login")
    } else {
      this.user = response.data.data
    }
  }
}
</script>
  
<style>

</style>