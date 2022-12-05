<template>
  <div>
    <Nav></Nav>
    <div class="container-sm mt-3">
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
                <h4 class="m-3">{{user.Username}}</h4>
              </div>
            </div>
          </div>
        </div>
        <div class="col-8 mt-5">
          <div class="row justify-content-start">
            <div class="col-2">
              <div class="row text-center">
                <h4>456</h4>
                <p class="text-secondary text-opacity-75">Watched</p>
              </div> 
            </div>
            <div class="col-2">
              <div class="row text-center">
                <h4>456</h4>
                <p class="text-secondary text-opacity-75">Watchlist</p>
              </div>
            </div>
            <div class="col-2">
              <div class="row text-center">
                <h4>456</h4>
                <p class="text-secondary text-opacity-75">Comments</p>
              </div>
            </div>
            <div class="col-2">
              <div class="row text-center">
                <h4>569</h4>
                <p class="text-secondary text-opacity-75">Likes</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="btn-group" role="group" aria-label="Basic radio toggle button group">
          <input type="radio" class="btn-check" name="btnradio" id="btnradio1" autocomplete="off" checked>
          <label class="btn btn-outline-dark" for="btnradio1">Profile</label>

          <input type="radio" class="btn-check" name="btnradio" id="btnradio2" autocomplete="off">
          <label class="btn btn-outline-dark" for="btnradio2">Watched</label>

          <input type="radio" class="btn-check" name="btnradio" id="btnradio3" autocomplete="off">
          <label class="btn btn-outline-dark" for="btnradio3">Watchlist</label>

          <input type="radio" class="btn-check" name="btnradio" id="btnradio4" autocomplete="off">
          <label class="btn btn-outline-dark" for="btnradio4">Comments</label>

          <input type="radio" class="btn-check" name="btnradio" id="btnradio5" autocomplete="off">
          <label class="btn btn-outline-dark" for="btnradio5">Likes</label>
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

export default {
  name: "profile",
  components: {
    Nav,
    Footer
  },
  data() {
    return {
      user: null
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

</style>