
<template>
  <div class="profile1">
    <Nav></Nav>
    <div class="container">
      <div class="text-center" style="padding-left:35%;padding-right:35%">
        <h2 class="text-center"
          style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Most
          Watched Movies of {{ this.movies.month }} </h2>
      </div>
      <div class="row">
        <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
          <div v-for="movie in movies.monthly_popular" :key="movie.id">
            <MovieCard :poster=movie.poster :title=movie.title :date="movie.release_date" :genre=movie.genres
              :rating="movie.rating" :length=movie.length :number_of_likes=movie.favorite_count
              :number_of_watches=movie.watched_count />
          </div>
        </div>
      </div>


    </div>

    <div class="container">
      <div class="text-center" style="margin-top: 3%;padding-left:35%;padding-right:35%">
        <h2 class="text-center"
          style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Most
          Favorited Movies of All Time </h2>
      </div>
      <div class="row">
        <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
          <div v-for="movie in movies.all_time_favorites" :key="movie.id">
            <MovieCard :poster=movie.poster :title=movie.title :date="movie.release_date" :genre=movie.genres
              :rating="movie.rating" :length=movie.length :number_of_likes=movie.favorite_count
              :number_of_watches=movie.watched_count />
          </div>
        </div>
      </div>


    </div>
    <div class="text-left" style="padding-left:7.7%;padding-right:7%">
      <h3 class="text-left"
        style="margin-top: 5%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Latest
        Movie Reviews </h3>
    </div>
    <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css" rel="stylesheet">
    <div class="container bootstrap snippets bootdey" style="margin-top: 1.5%; padding-left: 1.7%;">
      <div class="row">
        <div v-for="value in movies.latest_reviews" :key=value.user_id>
          <ReviewCard :Username=value.user_name :MovieTitle=value.movie_title :Comment=value.comment
            :CommentDate=value.comment_date />
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
  name: "home",
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
  async created() {
    const response = await axios.get("api/homepage");
    if (response.data.status == "error") {
      this.$router.push("/home")
    } else {
      this.movies = response.data
    }
    console.log(this.movies.monthly_popular)

  }
}
</script>
  
<style>
.post-list {
  position: relative;
  padding: 5px 0;
}

.post-list .picture {
  max-width: 400px;
  overflow: hidden;
  height: auto;
  border-radius: 6px;
}

.post-list .label {
  font-weight: normal;
}

.post-list .picture {
  max-width: 210px;
}

.post-list .picture img {
  width: 100%;
}

.post-list h4 {
  font-size: 20px;
}

.post-list h5 {
  color: #888;
}

.post-list p {
  float: left;
}

.post-list:after {
  background: #EEEEEE;
  width: 83.3333%;
  bottom: 0;
  right: 0;
  content: "";
  display: block;
  position: absolute;
}

.post-list .btn-download {
  margin-top: 45px;
}

.btn-info {
  color: #1084FF;
  border-color: #269abc;
}

.btn-round {
  border-width: 1px;
  border-radius: 30px !important;
  opacity: 0.79;
  padding: 9px 18px;
}

.btn {
  border-width: 2px;
  background-color: rgba(0, 0, 0, 0);
  font-weight: 400;
  opacity: 0.8;
  padding: 7px 16px;
}
</style>
