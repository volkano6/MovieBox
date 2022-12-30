
<template>
    <div class="profile1">
      <Nav :logged=this.movies.logged></Nav>
      <div class="container">
        <div class="text-center" style="padding-left:35%;padding-right:35%">
          <h2 class="text-center"
            style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Result
            of Movies {{ this.movies.month }} </h2>
        </div>
        <div class="row">
          <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
            <div v-for="movie in movies.monthly_popular" :key="movie.id">
              <MovieCard :id=movie.id :poster=movie.poster :title=movie.title :date="movie.release_date" :genre=movie.genres
                :rating="movie.rating" :length=movie.length :number_of_favorites=movie.favorite_count
                :number_of_watches=movie.watched_count />
            </div>
          </div>
        </div>
  
  
      </div>
  
      <div class="container">
        <div class="text-center" style="margin-top: 3%;padding-left:35%;padding-right:35%">
          <h2 class="text-center"
            style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Result
            of Profiles </h2>
        </div>
        <div class="row">
          <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
            <div v-for="movie in movies.all_time_favorites" :key="movie.id">
              <MovieCard :id=movie.id :poster=movie.poster :title=movie.title :date="movie.release_date" :genre=movie.genres
                :rating="movie.rating" :length=movie.length :number_of_favorites=movie.favorite_count
                :number_of_watches=movie.watched_count />
            </div>
          </div>
        </div>
  
  
      </div>
      <Footer></Footer>
    </div>
  </template>
    
  <script>
  import axios from 'axios'
  import Nav from "../components/nav.vue"
  import Footer from "../components/footer.vue"
  import MovieCard from "../components/movie_card.vue"
  import ReviewCard from "../components/review_card.vue"
  
  export default {
    name: "search",
    data() {
      return {
        movies: []
      }
    },
    components: {
      Nav,
      Footer,
      MovieCard,
      ReviewCard
    },
 
  }
  </script>
    
  