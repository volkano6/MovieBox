<template>
  <div>
    <Nav></Nav>
      <div class="container">
        <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
      <div v-for="movie in movies" :key="movie.id">
        <MovieCard :poster=movie.Poster :title=movie.Title :date="movie.ReleaseDate" genre="null" :rating="movie.Rating" length="95" number_of_likes="null" number_of_watches="null" />
      </div>
    </div>
      </div>
    <Footer></Footer>
  </div>
</template>

<script>
import axios from 'axios'
import Nav from "../components/nav.vue";
import Footer from "../components/footer.vue";
import MovieCard from "../components/movie_card.vue"

export default {
  name: "films",
  data() {
    return {
      movies: []
    }
  },
  components: {
    Nav,
    Footer,
    MovieCard
  },
  async created() {
    const response = await axios.get("api/movies");
    if (response.data.status == "error") {
      this.$router.push("/login")
    } else {
      this.movies = response.data.data

    }
  }
};
</script>

<style>

</style>