<template>

    <div>
        <Nav :logged=response.logged></Nav>
        <div class="container mt-3">
            <!--search row-->
            <div class="row pb-3 border-bottom">
        </div>
            <div class="row">
                <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
                    <div v-for="movie in movies" :key="movie.id">
                        <MovieCard :id=movie.id :poster=movie.poster :title=movie.title :date="movie.release_date" :genre=movie.genres
                            :rating="movie.rating" :length=movie.length :number_of_favorites=movie.favorite_count :number_of_watches=movie.watched_count />
                    </div>
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
            response: [],
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
            this.$router.push("/films")
        } else {
            this.movies = response.data.data
            this.response = response.data
            console.log(response.data.logged)
        }
    }
};
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