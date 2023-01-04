
<template>
  <div>
    <Nav :logged=this.data.logged :perm=data.perm ></Nav>
    <div class="container">
      <div class="row">
        <div class="col"></div>
        <div class="col"></div>
        <div class="col"></div>
        <div class="col">
          <form @submit.prevent="search" class="d-flex">
            <input v-model="inputSearch" class="form-control me-2" type="search" placeholder="Search"
              aria-label="Search">
            <button class="btn btn-outline-success" type="submit">Search</button>
          </form>
        </div>
      </div>
      <div class="text-center" style="padding-left:35%;padding-right:35%">
        <h2 class="text-center"
          style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Result
          of Movies</h2>
      </div>
      <div class="row">
        <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
          <div v-for="movie in this.data.movies" :key="movie.id">
            <MovieCard :id=movie.id :poster=movie.poster :title=movie.title :date="movie.release_date"
              :genre=movie.genres :rating="movie.rating" :length=movie.length :number_of_favorites=movie.favorite_count
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
          <div v-for="user in this.data.users" :key="user.id">
            <ProfileCard :id=user.id :username=user.username :avatar=user.avatar />
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
import ProfileCard from "../components/profile_card.vue"


export default {
  name: "search",
  data() {
    return {
      data: []
    }
  },
  async created() {
    const response = await axios.get(`api/search/?search=xxxxxxxxxxxxxxxxxxxxxx`);
      this.data = response.data
      console.log(this.data)
  },
  components: {
    Nav,
    Footer,
    MovieCard,
    ProfileCard,
    ReviewCard
  },
  methods: {
    async search() {
      
      const response = await axios.get(`api/search/?search=${this.inputSearch}`);
      this.data = response.data
      console.log(this.data)
    }
  }

}
</script>
  
