<template>
    <div>

        <Nav :logged=this.data.logged></Nav>

        <div class="container-sm">
            <div class="row">
                <div class="col-4">
                    <div>
                        <div class="">
                            <div class="user text-center">
                                <div class="profile mt-3">
                                    <img src="https://picsum.photos/200" class="rounded-circle" width="90">
                                </div>
                            </div>
                            <div class="text-center">
                                <h4 class="m-3 text-light text-opacity-75" v-if="Object.keys(this.data).length != 0">{{ this.data.user.username }}</h4>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-8 mt-5">
                    <div class="row justify-content-start" v-if="Object.keys(this.data).length != 0">
                        <div class="col-2" >
                            <div class="row text-center">
                                <h4 class="text-light text-opacity-75" v-if="data.user_watched != null">{{ this.data.user_watched.length }}</h4>
                                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                                <p class="text-light text-opacity-75">Watched</p>
                            </div>
                        </div>
                        <div class="col-2" >
                            <div class="row text-center">
                                <h4 class="text-light text-opacity-75" v-if="data.user_watchlist != null">{{ this.data.user_watchlist.length  }}</h4>
                                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                                <p class="text-light text-opacity-75">Watchlist</p>
                            </div>
                        </div>
                        <div class="col-2" >
                            <div class="row text-center">
                                <h4 class="text-light text-opacity-75" v-if="data.user_comments != null">{{ this.data.user_comments.length  }}</h4>
                                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                                <p class="text-light text-opacity-75">Comments</p>
                            </div>
                        </div>
                        <div class="col-2">
                            <div class="row text-center">
                                <h4 class="text-light text-opacity-75" v-if="data.user_favorites != null">{{ this.data.user_favorites.length  }}</h4>
                                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                                <p class="text-light text-opacity-75">Favorites</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="row m-1 pt-2 rounded" style="background-color:#bdbdc6">
        <div class="row">
          <div class="col" v-if="this.data.user.username != ''"><p style="color:black">User Name: {{this.data.user.username}}</p></div>
          <div class="col" v-if="this.data.user.email != ''"><p style="color:black">Email: {{this.data.user.email}}</p></div>
          <div class="col" v-if="this.data.user.date_of_birth != ''"><p style="color:black">Date Of Birth: {{this.data.user.date_of_birth}}</p></div>
        </div>
        <div class="row">
          <div class="col" v-if="this.data.user.instagram != ''"><p style="color:black">Instagram: {{this.data.user.instagram}}</p></div>
          <div class="col" v-if="this.data.user.twitter != ''"><p style="color:black">Twitter: {{this.data.user.twitter}}</p></div>
          <div class="col" v-if="this.data.user.location != ''"><p style="color:black">Location: {{this.data.user.location}}</p></div>
        </div>
        <div class="row">
          <div class="col" v-if="this.data.user.bio != ''"><p style="color:black">Bio: {{this.data.user.bio}}</p></div>
        </div>
      </div>
            <div class="row">

                <div class="btn-group" role="group" aria-label="Basic radio toggle button group">

                    <input type="radio" @click="pageContent('UserWatched')" class="btn-check" name="btnradio"
                        id="btnradio1" autocomplete="off">
                    <label class="btn btn-light me-1" for="btnradio1">Watched</label>

                    <input type="radio" @click="pageContent('UserWatchlit')" class="btn-check" name="btnradio"
                        id="btnradio2" autocomplete="off">
                    <label class="btn btn-light me-1" for="btnradio2">Watchlist</label>

                    <input type="radio" @click="pageContent('UserComments')" class="btn-check" name="btnradio"
                        id="btnradio3" autocomplete="off">
                    <label class="btn btn-light me-1" for="btnradio3">Comments</label>

                    <input type="radio" @click="pageContent('UserFavorites')" class="btn-check" name="btnradio"
                        id="btnradio4" autocomplete="off">
                    <label class="btn btn-light" for="btnradio4">Favorites</label>
                </div>
            </div>
            <div class="row">
                <div>
                </div>
                <div v-if="content == 'UserWatched'">
                    <UserWatched :watched=data.user_watched />
                </div>
                <div v-if="content == 'UserWatchlit'">
                    <UserWatchlit :watchlist=data.user_watchlist />
                </div>
                <div v-if="content == 'UserComments'">
                    <UserComments :comments=data.user_comments />
                </div>
                <div v-if="content == 'UserFavorites'">
                    <UserFavorites :favorites=data.user_favorites />
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
import UserWatched from "../components/single_page/profile_page/user_watched.vue"
import UserWatchlit from "../components/single_page/profile_page/user_watchlist.vue"
import UserComments from "../components/single_page/profile_page/user_comments.vue"
import UserFavorites from "../components/single_page/profile_page/user_favorites.vue"


export default {
    name: "user",
    components: {
        Nav,
        Footer,
        UserWatched,
        UserWatchlit,
        UserComments,
        UserFavorites
    },
    data() {
        return {
            data: [],
            content: "UserWatched",
        }
    },
    methods: {
        pageContent(a) {
            this.content = a
        },

    },
    async created() {
        const id = this.$route.params.id
        const response = await axios.get(`api/user/${id}`);
        this.data = response.data
        if (this.data.status == "error") {
            this.$router.push("/home")
        } else if (id == this.data.logged_id) {
            this.$router.push("/profile")
        }
    }
}
</script>
    
<style>
.row {
    color: #626988;
}
</style>