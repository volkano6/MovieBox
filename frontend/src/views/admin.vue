<template>
    <div>
        <Nav :logged=data.logged></Nav>


        <div class="container mt-3">

            <div class="row">
                <div class="col">
                    <h3 class="fw-bold" style="color:white;">My DashBoard</h3>
                    <h3></h3>
                    <!-- <p style="color:white;">My Dashboard</p> -->
                </div>
            </div>
            <div class="row">
                <h3 class="fw-bold" style="color:white;">Films</h3>
            </div>
            <div class="row mt-3">
                <div class="col-1 border ">
                    <h5 class="fw-bold" style="color:white;">Id</h5>
                </div>
                <div class="col-2 border">
                    <h5 class="fw-bold" style="color:white;">Title</h5>
                </div>
                <div class="col-3 border">
                    <h5 class="fw-bold" style="color:white;">Description</h5>
                </div>
                <div class="col-1 border">
                    <h5 class="fw-bold" style="color:white;">Release Date</h5>
                </div>
                <div class="col-2 border">
                    <!-- <p style="color:white;">Poster</p> -->
                    <h5 class="fw-bold" style="color:white;">Poster</h5>
                </div>
                <div class="col-1 border">
                    <h5 class="fw-bold" style="color:white;">Rating</h5>
                </div>
                <div class="col-1 border">
                    <h5 class="fw-bold" style="color:white;">Length</h5>

                </div>
                <div class="col-1">
                </div>
            </div>
            <div v-for="movie in this.data.movies" :key="movie.id">
                <div class="row">
                    <div class="col-1 border ">
                        <p class="mt-2" style="color:white;">{{ movie.id }}</p>
                    </div>
                    <div class="col-2 border">
                        <p class="mt-2" style="color:white;">{{ movie.title }}</p>
                    </div>
                    <div class="col-3 border">
                        <p class="mt-2" style="color:white;">{{ movie.description }}</p>
                    </div>
                    <div class="col-1 border">
                        <p class="mt-2" style="color:white;">{{ movie.release_date }}</p>
                    </div>
                    <div class="col-2 border">
                        <p class="mt-2" style="color:white;">{{ movie.poster }}</p>

                    </div>
                    <div class="col-1 border">
                        <p class="mt-2" style="color:white;">{{ movie.rating }}</p>

                    </div>
                    <div class="col-1 border">
                        <p class="mt-2" style="color:white;">{{ movie.length }}</p>

                    </div>
                    <div class="col-1 border">
                        <div class="row">
                            <div class="col-1">
                                <div style="position:relative; top:5px;">
                                    <button type="button" class="btn btn-primary" data-bs-toggle="modal"
                                        :data-bs-target="`#update${movie.id}`">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25"
                                            fill="currentColor" class="bi bi-pencil-square" viewBox="0 0 16 16">
                                            <path
                                                d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z">
                                            </path>
                                            <path fill-rule="evenodd"
                                                d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z">
                                            </path>
                                        </svg>
                                    </button>
                                    <div class="modal fade" :id="`update${movie.id}`" tabindex="-1"
                                        aria-labelledby="exampleModalLabel" aria-hidden="true">
                                        <div class="modal-dialog">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="exampleModalLabel">Film: {{ movie.id }}
                                                    </h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal"
                                                        aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">

                                                    <form @submit.prevent="updateFilm">
                                                        <div class="mb-3">
                                                            <label for="exampleFormControlInput1"
                                                                class="form-label">Title</label>
                                                            <input v-model=title type="displayname" class="form-control"
                                                                id="exampleFormControlInput1"
                                                                :placeholder="movie.title">
                                                        </div>
                                                        <div class="mb-3">
                                                            <label for="exampleFormControlTextarea1"
                                                                class="form-label">Description</label>
                                                            <textarea v-model="description" class="form-control"
                                                                id="exampleFormControlTextarea1" rows="3"
                                                                :placeholder="movie.description"></textarea>
                                                        </div>
                                                        <div class="mb-3">
                                                            <label for="exampleFormControlInput1"
                                                                class="form-label">Release Date</label>
                                                            <input v-model="relase_date" type="avatar"
                                                                class="form-control" id="exampleFormControlInput1"
                                                                :placeholder="movie.release_date">
                                                        </div>
                                                        <div class="mb-3">
                                                            <label for="exampleFormControlInput1"
                                                                class="form-label">Poster</label>
                                                            <input v-model="poster" type="location" class="form-control"
                                                                id="exampleFormControlInput1"
                                                                :placeholder="movie.poster">
                                                        </div>
                                                        <div class="mb-3">
                                                            <label for="exampleFormControlInput1"
                                                                class="form-label">Length</label>
                                                            <input v-model="length" type="location" class="form-control"
                                                                id="exampleFormControlInput1"
                                                                :placeholder="movie.length">
                                                        </div>
                                                        <div class="mb-3">
                                                            <label for="exampleFormControlInput1"
                                                                class="form-label">Genres</label>
                                                            <input v-model="genres" type="twitter" class="form-control"
                                                                id="exampleFormControlInput1"
                                                                :placeholder="movie.genres">
                                                        </div>

                                                        <div class="modal-footer">
                                                            <button type="button" class="btn btn-secondary"
                                                                data-bs-dismiss="modal">Close</button>
                                                            <button type="submit" class="btn btn-primary">Save
                                                                changes</button>
                                                        </div>
                                                    </form>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="col-1">
                                <div style="position:relative; left:30px; top: 5px;">
                                    <button type="button" class="btn btn-primary" data-bs-toggle="modal"
                                        :data-bs-target="`#deleteMovie${movie.id}`">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25"
                                            fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">
                                            <path
                                                d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5ZM11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H2.506a.58.58 0 0 0-.01 0H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1h-.995a.59.59 0 0 0-.01 0H11Zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5h9.916Zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47ZM8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z" />
                                        </svg>
                                    </button>
                                    <div class="modal fade" :id="`deleteMovie${movie.id}`" tabindex="-1"
                                        aria-labelledby="exampleModalLabel" aria-hidden="true">
                                        <div class="modal-dialog">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="exampleModalLabel">Delete Film</h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal"
                                                        aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">
                                                    Are you sure to delete "{{ movie.title }}" ?
                                                    <div class="modal-footer">
                                                        <button type="button" class="btn btn-secondary"
                                                            data-bs-dismiss="modal">No</button>
                                                        <button type="button" class="btn btn-primary">Yes</button>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="row mt-3">
                <h3 style="color:white;">Users</h3>
            </div>
            <div class="row">
                <div class="col-1 border ">
                    <h6 class="fw-bold mt-3" style="color:white;">User ID</h6>
                </div>
                <div class="col-1 border ">
                    <h6 class="fw-bold mt-3" style="color:white;">userName</h6>
                </div>
                <div class="col-2 border">
                    <h6 class="fw-bold mt-3" style="color:white;">disName</h6>
                </div>
                <div class="col-2 border">
                    <h6 class="fw-bold mt-3" style="color:white;">email</h6>
                </div>
                <div class="col-1 border">
                    <h6 class="fw-bold mt-3" style="color:white;">dateOfBirth</h6>
                </div>
                <div class="col-1 border">
                    <!-- <p style="color:white;">Poster</p> -->
                    <h6 class="fw-bold mt-3" style="color:white;">bio</h6>
                </div>
                <div class="col-1 border">
                    <h6 class="fw-bold mt-3" style="color:white;">location</h6>
                </div>
                <div class="col-1 border">
                    <h6 class="fw-bold mt-3" style="color:white;">twitter</h6>
                </div>
                <div class="col-1 border">
                    <h6 class="fw-bold mt-3" style="color:white;">instagram</h6>
                </div>

                <div class="col-1"></div>
            </div>
            <div v-for="user in this.data.users" :key="user.id">
                <div class="row">
                    <div class="col-1 border ">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.id }}</h6>
                    </div>
                    <div class="col-1 border ">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.username }}</h6>
                    </div>
                    <div class="col-2 border">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.displayname }}</h6>
                    </div>
                    <div class="col-2 border">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.email }}</h6>
                    </div>
                    <div class="col-1 border">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.date_of_birth }}</h6>
                    </div>
                    <div class="col-1 border">
                        <!-- <p style="color:white;">Poster</p> -->
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.bio }}</h6>
                    </div>
                    <div class="col-1 border">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.location }}</h6>
                    </div>
                    <div class="col-1 border">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.twitter }}</h6>
                    </div>
                    <div class="col-1 border">
                        <h6 class="fw-bold mt-3" style="color:white;">{{ user.instagram }}</h6>
                    </div>

                    <div class="col-1 border">
                        <div class="row">
                            <div class="col-1">
                                <div style="position:relative; top:5px;">
                                    <button type="button" class="btn btn-primary" data-bs-toggle="modal"
                                        :data-bs-target="`#updateUser${user.id}`">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25"
                                            fill="currentColor" class="bi bi-pencil-square" viewBox="0 0 16 16">
                                            <path
                                                d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z">
                                            </path>
                                            <path fill-rule="evenodd"
                                                d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z">
                                            </path>
                                        </svg>
                                    </button>
                                    <div class="modal fade" :id="`updateUser${user.id}`" tabindex="-1"
                                        aria-labelledby="exampleModalLabel" aria-hidden="true">
                                        <div class="modal-dialog">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="exampleModalLabel">User: {{ user.id }}
                                                    </h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal"
                                                        aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">

                                                    <form @submit.prevent="updateFilm(user.id)">

                                                        <div class="modal-footer">
                                                            <button type="button" class="btn btn-secondary"
                                                                data-bs-dismiss="modal">Close</button>
                                                            <button type="submit" class="btn btn-primary">Save
                                                                changes</button>
                                                        </div>
                                                    </form>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="col-1">
                                <div style="position:relative; left:30px; top: 5px;">
                                    <button type="button" class="btn btn-primary" data-bs-toggle="modal"
                                        :data-bs-target="`#deleteUser${user.id}`">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25"
                                            fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">
                                            <path
                                                d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5ZM11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H2.506a.58.58 0 0 0-.01 0H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1h-.995a.59.59 0 0 0-.01 0H11Zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5h9.916Zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47ZM8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z" />
                                        </svg>
                                    </button>
                                    <div class="modal fade" :id="`deleteUser${user.id}`" tabindex="-1"
                                        aria-labelledby="exampleModalLabel" aria-hidden="true">
                                        <div class="modal-dialog">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="exampleModalLabel">Delete User</h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal"
                                                        aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">
                                                    Are you sure to delete {{ user.username }} ?
                                                    <div class="modal-footer">
                                                        <button type="button" class="btn btn-secondary"
                                                            data-bs-dismiss="modal">No</button>
                                                        <button type="button" class="btn btn-primary">Yes</button>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
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

export default {
    name: "admin",
    data() {
        return {
            data: []
        }
    },
    components: {
        Nav,
        Footer,
    },
    async created() {
        const response = await axios.get("api/dashboard");
        if (response.data.status == "error") {
            this.$router.push("/home")
        } else {
            this.data = response.data
        }
    },
    methods: {
        deleteFilm(movie_id = 9) {
            alert("delete:", movie_id)
        },
        updateFilm: function (movie_id) {
            alert("update:", movie_id)
        },
        deleteUser(user_id) { },
        updateUser(user_id) { },
    }
}
</script>

<style>
svg {
    fill: white;
}

.col-1{
display: flex;
align-items: center;
overflow: scroll;
}
.col-2{
display: flex;
align-items: center;
overflow: scroll;
}
.col-3{
display: flex;
align-items: center;
overflow: scroll;
}
</style>