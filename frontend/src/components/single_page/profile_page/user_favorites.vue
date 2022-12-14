<template>
  <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
    <div v-for="movie in favorites" :key="movie.id">
      <MovieCard :poster=movie.poster :title=movie.title :date="movie.release_date" :genre=movie.genres
        :rating="movie.rating" :length=movie.length :number_of_favorites=movie.favorite_count
        :number_of_watches=movie.watched_count />
    </div>
  </div>
</template>

<script>
import MovieCard from "../../movie_card.vue"

export default {
  name: "userFavorites",
  props: ["favorites"],
  components: {
    MovieCard
  },
  data() {

  },
  created() {

  }

}
</script>

<style>

</style>