<template>
  <div>

    <Nav :logged=this.data.logged :avatar=this.avatar></Nav>

    <div class="container-sm" v-if="Object.keys(this.data).length != 0">
      <div class="row">
        <div class="col-4">
          <div class="">
            <div class="user text-center">
              <div class="profile mt-3">
                <div v-if="this.avatar == ''">
                  <img src="https://picsum.photos/200" class="rounded-circle" width="90">
                </div>
                <div v-else>
                  <img :src="this.avatar" class="rounded-circle" width="90">
                </div>
              </div>
            </div>
            <div class="text-center">
              <h4 class="m-3 text-light text-opacity-75">
                @{{ this.displayname }}</h4>
            </div>
            <div class="row">
              <div class="text-center">
                <!-- Button trigger modal -->
                <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#exampleModal">
                  Update Profile
                </button>

                <!-- Modal -->
                <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel"
                  aria-hidden="true">
                  <div class="modal-dialog">
                    <div class="modal-content">
                      <div class="modal-header">
                        <h5 class="modal-title" id="exampleModalLabel">Update Profile</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                      </div>
                      <div class="modal-body">
                        <form @submit.prevent="userUpdate">
                          <div class="mb-3">
                            <label for="exampleFormControlInput1" class="form-label">Display Name</label>
                            <input v-model=displayname type="displayname" class="form-control"
                              id="exampleFormControlInput1" :placeholder="this.data.user.displayname">
                          </div>
                          <div class="mb-3">
                            <label for="exampleFormControlInput1" class="form-label">Email address</label>
                            <input v-model="email" type="email" class="form-control" id="exampleFormControlInput1"
                              :placeholder="this.email">
                          </div>
                          <div class="mb-3">
                            <label for="exampleFormControlInput1" class="form-label">Avatar</label>
                            <input v-model="avatar" type="avatar" class="form-control" id="exampleFormControlInput1"
                              :placeholder="this.avatar">
                          </div>
                          <div class="mb-3">
                            <label for="exampleFormControlInput1" class="form-label">Date of Birth</label>
                            <input v-model="date_of_birth" type="date_of_birth" class="form-control"
                              id="exampleFormControlInput1" :placeholder="this.date_of_birth">
                          </div>
                          <div class="mb-3">
                            <label for="exampleFormControlInput1" class="form-label">Location</label>
                            <input v-model="location" type="location" class="form-control" id="exampleFormControlInput1"
                              :placeholder="this.location">
                          </div>
                          <div class="mb-3">
                            <label for="exampleFormControlInput1" class="form-label">Twitter</label>
                            <input v-model="twitter" type="twitter" class="form-control" id="exampleFormControlInput1"
                              :placeholder="this.twitter">
                          </div>
                          <div class="mb-3">
                            <label for="exampleFormControlInput1" class="form-label">Instagram</label>
                            <input v-model="instagram" type="instagram" class="form-control"
                              id="exampleFormControlInput1" :placeholder="this.instagram">
                          </div>
                          <div class="mb-3">
                            <label for="exampleFormControlTextarea1" class="form-label">Bio</label>
                            <textarea v-model="bio" class="form-control" id="exampleFormControlTextarea1" rows="3"
                              :placeholder="this.data.user.bio"></textarea>
                          </div>
                          <div class="modal-footer">
                            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                            <button type="submit" class="btn btn-primary"  data-bs-dismiss="modal">Save changes</button>
                          </div>
                        </form>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col-8 mt-5">
          <div class="row justify-content-start" v-if="this.data.length != 0">
            <div class="col-2">
              <div class="row text-center">
                <h4 class="text-light text-opacity-75" v-if="data.user_watched != null">
                  {{ this.data.user_watched.length }}</h4>
                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                <p class="text-light text-opacity-75">Watched</p>
              </div>
            </div>
            <div class="col-2">
              <div class="row text-center">
                <h4 class="text-light text-opacity-75" v-if="data.user_watchlist != null">
                  {{ this.data.user_watchlist.length }}</h4>
                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                <p class="text-light text-opacity-75">Watchlist</p>
              </div>
            </div>
            <div class="col-2">
              <div class="row text-center">
                <h4 class="text-light text-opacity-75" v-if="data.user_comments != null">
                  {{ this.data.user_comments.length }}</h4>
                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                <p class="text-light text-opacity-75">Comments</p>
              </div>
            </div>
            <div class="col-2">
              <div class="row text-center">
                <h4 class="text-light text-opacity-75" v-if="data.user_favorites != null">
                  {{ this.data.user_favorites.length }}</h4>
                <h4 class="text-light text-opacity-75" v-else>{{ 0 }}</h4>
                <p class="text-light text-opacity-75">Favorites</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="row m-1 pt-2 rounded" style="background-color:#bdbdc6">
        <div class="row">
          <div class="col" v-if="this.data.user.username != ''">
            <p style="color:black">User Name: {{ this.data.user.username }}</p>
          </div>
          <div class="col" v-if="this.email != ''">
            <p style="color:black">Email: {{ this.email }}</p>
          </div>
          <div class="col" v-if="this.date_of_birth != ''">
            <p style="color:black">Date Of Birth: {{ this.date_of_birth }}</p>
          </div>
        </div>
        <div class="row">
          <div class="col" v-if="this.instagram != ''">
            <p style="color:black">Instagram: {{ this.instagram }}</p>
          </div>
          <div class="col" v-if="this.twitter != ''">
            <p style="color:black">Twitter: {{ this.twitter }}</p>
          </div>
          <div class="col" v-if="this.location != ''">
            <p style="color:black">Location: {{ this.location }}</p>
          </div>
        </div>
        <div class="row">
          <div class="col" v-if="this.bio != ''">
            <p style="color:black">Bio: {{ this.bio  }}</p>
          </div>
        </div>
      </div>
      <div class="row">

        <div class="btn-group" role="group" aria-label="Basic radio toggle button group">

          <input type="radio" @click="pageContent('UserWatched')" class="btn-check" name="btnradio" id="btnradio1"
            autocomplete="off">
          <label class="btn btn-light me-1" for="btnradio1">Watched</label>

          <input type="radio" @click="pageContent('UserWatchlit')" class="btn-check" name="btnradio" id="btnradio2"
            autocomplete="off">
          <label class="btn btn-light me-1" for="btnradio2">Watchlist</label>

          <input type="radio" @click="pageContent('UserComments')" class="btn-check" name="btnradio" id="btnradio3"
            autocomplete="off">
          <label class="btn btn-light me-1" for="btnradio3">Comments</label>

          <input type="radio" @click="pageContent('UserFavorites')" class="btn-check" name="btnradio" id="btnradio4"
            autocomplete="off">
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
  name: "profile",
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
      displayname: '',
      email: '',
      avatar: '',
      date_of_birth: '',
      location: '',
      twitter: '',
      instagram: '',
      bio: '',

    }
  },
  methods: {
    pageContent(a) {
      this.content = a
    },
    userUpdate() {
      axios.put('api/profile', {
        displayname: this.displayname,
        email: this.email,
        date_of_birth: this.date_of_birth,
        avatar: this.avatar,
        bio: this.bio,
        location: this.location,
        twitter: this.twitter,
        instagram: this.instagram
      })
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });
    }

  },
  async created() {
    const response = await axios.get("api/profile");
    if (response.data.status == "error") {
      this.$router.push("/login")
    } else {
      this.data = response.data
      this.displayname = this.data.user.displayname
      this.email = this.data.user.email
      this.avatar = this.data.user.avatar
      this.date_of_birth = this.data.user.date_of_birth
      this.location = this.data.user.location
      this.twitter = this.data.user.twitter
      this.instagram = this.data.user.instagram
      this.bio = this.data.user.bio
    }
  }
}
</script>
  
<style>
.row {
  color: #626988;
}
</style>