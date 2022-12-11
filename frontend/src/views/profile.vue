<template>
  <div>

    <Nav></Nav>

      <div class="container-sm">
        <div class="row">
          <div class="col-4">
            <div>
              <div class="">
                <div class="user text-center">
                  <div class="profile mt-3">
                    <img src="https://i.imgur.com/JgYD2nQ.jpg" class="rounded-circle" width="90">
                  </div>
                </div>
                <div class="text-center">
                  <h4 class="m-3 text-light text-opacity-75" >{{user.Username}}</h4>
                </div>
              </div>
            </div>
          </div>
          <div class="col-8 mt-5">
            <div class="row justify-content-start">
              <div class="col-2">
                <div class="row text-center">
                  <h4 class="text-light text-opacity-75">70</h4>
                  <p class="text-light text-opacity-75">Watched</p>
                </div>
              </div>
              <div class="col-2">
                <div class="row text-center">
                  <h4 class="text-light text-opacity-75">45</h4>
                  <p class="text-light text-opacity-75">Watchlist</p>
                </div>
              </div>
              <div class="col-2">
                <div class="row text-center">
                  <h4 class="text-light text-opacity-75">360</h4>
                  <p class="text-light text-opacity-75">Comments</p>
                </div>
              </div>
              <div class="col-2">
                <div class="row text-center">
                   <h4 class="text-light text-opacity-75">450</h4>
                  <p class="text-light text-opacity-75">Likes</p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="row">
          
          <div class="btn-group" role="group" aria-label="Basic radio toggle button group">

            <input type="radio" @click="pageContent('UserWatched')" class="btn-check" name="btnradio" id="btnradio1" autocomplete="off">
            <label class="btn btn-light" for="btnradio1" >Watched</label>

            <input type="radio" @click="pageContent('UserWatchlit')" class="btn-check" name="btnradio" id="btnradio2" autocomplete="off">
            <label class="btn btn-light" for="btnradio2">Watchlist</label>

            <input type="radio" @click="pageContent('UserComments')" class="btn-check" name="btnradio" id="btnradio3" autocomplete="off">
            <label class="btn btn-light" for="btnradio3">Comments</label>

            <input type="radio" @click="pageContent('UserLikes')" class="btn-check" name="btnradio" id="btnradio4" autocomplete="off">
            <label class="btn btn-light" for="btnradio4">Likes</label>
          </div>

        </div>
        <div v-if="content == 'UserWatched'">
          <UserWatched />
        </div>
        <div v-if="content == 'UserWatchlit'">
          <UserWatchlit />
        </div>
        <div v-if="content == 'UserComments'">
          <UserComments />
        </div>
        <div v-if="content == 'UserLikes'">
          <UserLikes />
        </div>  
      </div>

      <Footer></Footer>

    </div>

</template>
  
<script>

import axios from 'axios'
import Nav from "../components/nav.vue"
import Footer from "../components/footer.vue"
import UserInfo from "../components/single_page/profile_page/user_information.vue"
import UserWatched from "../components/single_page/profile_page/user_watched.vue"
import UserWatchlit from "../components/single_page/profile_page/user_watchlist.vue"
import UserComments from "../components/single_page/profile_page/user_comments.vue"
import UserLikes from "../components/single_page/profile_page/user_likes.vue"


export default {
  name: "profile",
  components: {
    Nav,
    Footer,
    UserInfo,
    UserWatched,
    UserWatchlit,
    UserComments,
    UserLikes
    },
  data() {
    return {
      user: "User",
      content: "UserWatched",
    }
  },
  methods: {
    pageContent (a) {
      this.content = a
    }
  },
  async created() {
    const response = await axios.get("api/profile");
    if (response.data.status == "error") {
      this.$router.push("/login")
    } else {
      this.user = response.data.data
    }
  }
}
</script>
  
<style>
.row {
  color: #626988;
}
</style>