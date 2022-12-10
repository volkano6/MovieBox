<template>
  <div class="profile1">
    <Nav></Nav>
    <div class="container" >
      <div class="text-center" style="padding-left:35%;padding-right:35%">
      <h2 class="text-center" style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Favorite of Weeks </h2> 
    </div>
      <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />

      </div>

    </div>

    <div class="container" >
      <div class="text-center" style="margin-top: 3%;padding-left:35%;padding-right:35%">
      <h2 class="text-center" style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> All Time Favorites </h2> 
    </div>
      <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
        <Card />
        <Card />
        <Card />
        <Card />
        <Card />

      </div>

    </div>
      <div class="text-left" style="padding-left:7.7%;padding-right:7%">
      <h3 class="text-left" style="margin-top: 5%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;"> Last Reviews </h3> 
      </div>
    <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css" rel="stylesheet">
<div class="container bootstrap snippets bootdey" style="margin-top: 1.5%; padding-left: 1.7%;">
    <div class="row">
      <ReviewCard />
      <ReviewCard />
      <ReviewCard />
      <ReviewCard />
    </div>
</div>
    <Footer></Footer>
  </div>
</template>
  
<script>
import axios from 'axios'
import Nav from "../components/nav.vue"
import Footer from "../components/footer.vue"
import Card from "../components/movie_card.vue"
import ReviewCard from "../components/review_card.vue"

export default {
  name: "profile",
  components: {
    Nav,
    Footer,
    Card,
    ReviewCard
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

.post-list .label{
    font-weight:normal;    
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
  background-color: rgba(0,0,0,0);
  font-weight: 400;
  opacity: 0.8;
  padding: 7px 16px;
}
</style>

