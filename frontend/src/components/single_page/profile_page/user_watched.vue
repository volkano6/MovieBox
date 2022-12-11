<template>
    <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
        <MovieCard :poster="poster" :title="title" :date="date" :rating="rating" :genre="genre" :length="length" :number_of_likes="number_of_likes" :number_of_watches="number_of_watches" />
    </div>
</template>

<script>
import MovieCard from "../../movie_card.vue"

export default {
    data() {
        return {

            poster: "https://cdn.shopify.com/s/files/1/0057/3728/3618/products/avatar-the-way-of-water_sncuhzap_480x.progressive.jpg?v=1669382994",
            title: "Avatar",
            date: "2010",
            rating: "4.8",
            genre: "science fiction",
            length: "90" + "min.",
            number_of_watches: "110k",
            number_of_likes: "200k",
        }
    },
    components: {
        MovieCard,
    },
}
</script>

<style>

</style>