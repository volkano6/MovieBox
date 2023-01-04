<template>
    <div>
        <Nav :logged=this.data.logged></Nav>
        <div class="container mt-3">
            <div class="row" v-if="Object.keys(this.data).length != 0">
                <div class="col-3">
                    <img :src="data.movie.poster" alt="" width="320" />
                    <div class="row">
                        <div class="col-6">
                            <div class="row mt-1">
                                <div class="col">
                                    <div style="position:relative; left:50px; top:2px; color:white;">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25"
                                            fill="currentColor" class="bi bi-heart" viewBox="0 0 16 16">
                                            <path
                                                d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z" />
                                        </svg>
                                    </div>
                                </div>
                                <div class="col" style="position:relative; top:4px;">
                                    <p style="color:white;">{{ data.movie.favorite_count }}</p>
                                </div>
                            </div>
                        </div>
                        <div class="col-6">
                            <div class="row mt-1">
                                <div class="col">
                                    <div style="position:relative; left:35px; color:white;">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="31" height="31"
                                            fill="currentColor" class="bi bi-eye-fill" viewBox="0 0 16 16">
                                            <path d="M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z" />
                                            <path
                                                d="M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8zm8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7z" />
                                        </svg>
                                    </div>
                                </div>
                                <div class="col" style="position:relative; right:10px; top:4px;">
                                    <p style="color:white;">{{ data.movie.watched_count }}</p>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-9">
                    <div class="row justify-content-start">
                        <div class="col">
                            <h3 class="fw-bold" style="color:white;">{{ data.movie.title }} </h3>
                        </div>
                        <div class="col">
                            <h4 class="fw-light" style="color:white;">{{ data.movie.release_date }}</h4>
                        </div>
                    </div>
                    <br>
                    <div class="row">
                        <div class="col-9 ">
                            <div class="row ">
                                <h3 style="color:white;">Description: </h3>
                                <h5 style="color:white;">{{ data.movie.description }}</h5>
                            </div>
                            <div class="row ">
                                <h3 style="color:white;">Genres: {{ data.movie.genres }}</h3>
                            </div>
                            <div class="row ">
                                <h3 style="color:white;">Cast</h3>

                            </div>
                        </div>
                        <div class="col">

                            <!-- watch or remove from watch -->
                            <div class="row">
                                <div class="d-grid gap-2" style="margin:5px; right:5px; color:white;">
                                    <button @click="watched" class="btn btn-primary active" data-bs-toggle="button"
                                        aria-pressed="true">
                                        <div style="position:relative; ">
                                            <svg xmlns="http://www.w3.org/2000/svg" width="31" height="31"
                                                fill="currentColor" class="bi bi-eye-fill" viewBox="0 0 16 16">
                                                <path d="M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z" />
                                                <path
                                                    d="M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8zm8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7z" />
                                            </svg>

                                        </div>
                                    </button>
                                </div>
                                <div class="row">
                                    <div v-if="in_watch">
                                        <div class="d-grid gap-2" style="margin-left:27px;">
                                            <p style="color:white;">Remove from Watched</p>
                                        </div>
                                    </div>
                                    <div v-else>
                                        <div class="d-grid gap-2" style="margin-left:53px;">
                                            <p style="left:240px; color:white;">Watched</p>
                                        </div>
                                    </div>

                                </div>
                            </div>
                            <!-- add or remove favorites -->
                            <div class="row">
                                <div class="d-grid gap-2" style="margin:5px; right:5px;">
                                    <button @click="favorite" class="btn btn-primary active" data-bs-toggle="button">
                                        <div style="position:relative; ">
                                            <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25"
                                                fill="currentColor" class="bi bi-heart" viewBox="0 0 16 16">
                                                <path
                                                    d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z" />
                                            </svg>
                                        </div>
                                    </button>
                                </div>
                                <div class="row">
                                    <div v-if="in_favorite">
                                        <div class="d-grid gap-2" style="margin-left:27px;">
                                            <p style="left:240px; color:white;">Favorites</p>
                                        </div>
                                    </div>
                                    <div v-else>
                                        <div class="d-grid gap-2" style="margin-left:53px;">
                                            <p style="left:240px; color:white;">Favorites</p>
                                        </div>
                                    </div>
                                </div>
                            </div>

                            <!-- watchlist row -->
                            <div class="row">
                                <div @click="watchlist" class="d-grid gap-2" style="margin:5px; right:5px;">
                                    <button class="btn btn-primary active" data-bs-toggle="button">
                                        <div style="position:relative; ">
                                            <div v-if="in_watchlist">
                                                <div class="d-grid gap-2" style="margin-left:15px;">
                                                    <p style="left:240px; color:white;">Remove from Watchlist</p>
                                                </div>
                                            </div>
                                            <div v-else>
                                                <div class="d-grid gap-2" style="margin-left:5px;">
                                                    <p style="left:240px; color:white;">Watchlist</p>
                                                </div>
                                            </div>
                                        </div>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <form @submit.prevent="comment_and_rating">
                            <h4 style="color:white;">Write comment:</h4>
                            <div class="mb-3 bg-light ">

                                <div class="row m-2">

                                    <div class="col">
                                        <h5> Rate: </h5>
                                        <fieldset class="rating">
                                            <input v-model="rating" type="radio" id="star5" name="rating"
                                                value="5" /><label class="full" for="star5"
                                                title="Awesome - 5 stars"></label>
                                            <input v-model="rating" type="radio" id="star4half" name="rating"
                                                value="4 and a half" /><label class="half" for="star4half"
                                                title="Pretty good - 4.5 stars"></label>
                                            <input v-model="rating" type="radio" id="star4" name="rating"
                                                value="4" /><label class="full" for="star4"
                                                title="Pretty good - 4 stars"></label>
                                            <input v-model="rating" type="radio" id="star3half" name="rating"
                                                value="3 and a half" /><label class="half" for="star3half"
                                                title="Meh - 3.5 stars"></label>
                                            <input v-model="rating" type="radio" id="star3" name="rating"
                                                value="3" /><label class="full" for="star3"
                                                title="Meh - 3 stars"></label>
                                            <input v-model="rating" type="radio" id="star2half" name="rating"
                                                value="2 and a half" /><label class="half" for="star2half"
                                                title="Kinda bad - 2.5 stars"></label>
                                            <input v-model="rating" type="radio" id="star2" name="rating"
                                                value="2" /><label class="full" for="star2"
                                                title="Kinda bad - 2 stars"></label>
                                            <input v-model="rating" type="radio" id="star1half" name="rating"
                                                value="1 and a half" /><label class="half" for="star1half"
                                                title="Meh - 1.5 stars"></label>
                                            <input v-model="rating" type="radio" id="star1" name="rating"
                                                value="1" /><label class="full" for="star1"
                                                title="Sucks big time - 1 star"></label>
                                            <input v-model="rating" type="radio" id="starhalf" name="rating"
                                                value="half" /><label class="half" for="starhalf"
                                                title="Sucks big time - 0.5 stars"></label>
                                        </fieldset>


                                    </div>
                                    <textarea v-model="comment" class="form-control mb-2"
                                        id="exampleFormControlTextarea1" rows="3"></textarea>

                                </div>
                                <button type="submit" class="btn btn-primary mb-2 ms-2">Submit</button>

                            </div>
                        </form>

                    </div>
                    <div v-if="data.comments == null">
                        no command
                    </div>
                    <div v-else>
                        <div class="row">
                            <div v-for="value in data.comments" :key=value.user_id>
                                <ReviewCard :UserID="value.user_id" :Username=value.user_name :MovieID=value.movie_id
                                    :MovieTitle=value.movie_title :Comment=value.comment
                                    :CommentDate=value.comment_date />
                            </div>
                        </div>
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
import ReviewCard from "../components/review_card.vue"

export default {
    name: "films",
    data() {
        return {
            data: [],
            movie_id: this.$route.params.id,
            in_watch: false,
            in_watchlist: false,
            in_favorite: false,
        }

    },
    components: {
        Nav,
        Footer,
        ReviewCard
    },
    async created() {
        if (Object.keys(this.data).length == 0) {
            console.log("boş")
        } else console.log("dolu")


        const response = await axios.get(`api/movies/${this.movie_id}`);
        this.data = response.data
        if (this.data.status == "error") {
            this.$router.push("/home")
        }
        if (Object.keys(this.data).length == 0) {
            console.log("boş")
        } else console.log("dolu")

    },
    methods: {
        async watched() {
            const id = this.$route.params.id
            await axios.post(`api/movies/${this.movie_id}/watched`)
                .then(async function (response) {

                    if (Object.keys(response.data).length == 2) {
                        
                        await axios.delete(`api/movies/${id}/watched`)
                            .then(function (response) {
                                alert(`This movie removed from your watched.`)
                            })
                            .catch(function (error) {
                                console.log(error);
                            });
                            
                    }else {
                        alert(`This movie added to your watched.`)
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });

        },
        async watchlist() {
            const id = this.$route.params.id
            await axios.post(`api/movies/${this.movie_id}/watchlist`)
                .then(async function (response) {

                    if (Object.keys(response.data).length == 2) {
                        
                        await axios.delete(`api/movies/${id}/watchlist`)
                            .then(function (response) {
                                alert(`This movie removed from your watchlist.`)
                            })
                            .catch(function (error) {
                                console.log(error);
                            });
                    }else {
                        alert(`This movie added to your watchlist.`)
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        async favorite() {
            const id = this.$route.params.id
            await axios.post(`api/movies/${this.movie_id}/favorite`)
                .then(async function (response) {

                    if (Object.keys(response.data).length == 2) {
                        
                        await axios.delete(`api/movies/${id}/favorite`)
                            .then(function (response) {
                                alert(`This movie removed from your favorites.`)
                            })
                            .catch(function (error) {
                                console.log(error);
                            });
                    }else {
                        alert(`This movie added to your favorites.`)
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        async comment_and_rating() {
            if (this.rating != undefined) {
                switch (this.rating) {
                    case "half":
                        this.rating = 0.5;
                        break;
                    case "1":
                        this.rating = 1;
                        break;
                    case "1 and a half":
                        this.rating = 1.5;
                        break;
                    case "2":
                        this.rating = 2;
                        break;
                    case "2 and a half":
                        this.rating = 2.5;
                        break;
                    case "3":
                        this.rating = 3;
                        break;
                    case "3 and a half":
                        this.rating = 3.5;
                        break;
                    case "4":
                        this.rating = 4;
                        break;
                    case "4 and a half":
                        this.rating = 4.5;
                        break;
                    case "5":
                        this.rating = "5";
                        break;
                }
                
                await axios.post(`api/movies/${this.movie_id}/rating?star=${this.rating}`)
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });

            }
            if (this.comment != undefined) {
                
                await axios.post(`api/movies/${this.movie_id}/comment`, {
                    comment: this.comment,
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            }
        }
    }
};
</script>
  
<style>
@import url(//netdna.bootstrapcdn.com/font-awesome/3.2.1/css/font-awesome.css);

fieldset,
label {
    margin: 0;
    padding: 0;
}

h1 {
    font-size: 1.5em;
    margin: 10px;
}

/****** Style Star Rating Widget *****/

.rating {
    border: none;
    float: left;
}

.rating>input {
    display: none;
}

.rating>label:before {
    margin: 5px;
    font-size: 1.25em;
    font-family: FontAwesome;
    display: inline-block;
    content: "\f005";
}

.rating>.half:before {
    content: "\f089";
    position: absolute;
}

.rating>label {
    color: #ddd;
    float: right;
}

.rating>input:checked~label,
/* show gold star when clicked */
.rating:not(:checked)>label:hover,
/* hover current star */
.rating:not(:checked)>label:hover~label {
    color: #FFD700;
}

/* hover previous stars in list */

.rating>input:checked+label:hover,
/* hover current star when changing rating */
.rating>input:checked~label:hover,
.rating>label:hover~input:checked~label,
/* lighten current selection */
.rating>input:checked~label:hover~label {
    color: #FFED85;
}
</style>