<template>
  <div class="col" style="margin-right: 3.5%;">
    <div class="profile-card" style="width: 18rem;">
  <img src="https://lh3.googleusercontent.com/-XdUIqdMkCWA/AAAAAAAAAAI/AAAAAAAAAAA/4252rscbv5M/photo.jpg
" class="card-img-top">
  <div class="profile-content">
    <p class="text-center">{{username}}</p>
  </div>
</div>
    
  </div>
</template>

<script>
export default {
  name: "user",
  props: ['id', 'username', 'displayname', 'email', 'date_of_birth', 'avatar', 'bio', 'location', 'twitter', 'instagram'],

}

</script>


<style scoped>
* {
  -webkit-transition: 300ms;
  transition: 300ms;
}

.profile-card {
  background: #ffffff;
  box-shadow: 0px 6px 18px rgba(130, 10, 10, 0.1);
  width: 100%;
  max-width: 290px;
  border-radius: 10px;
  display: inline-block;
}



.profile-card:hover {
  -webkit-transform: scale(1.03);
  transform: scale(1.03);
  box-shadow: 0px 10px 25px rgba(0, 0, 0, 0.08);
}

.profile-content {
  padding: 8px 12px 8px 12px;
  font-size: 18px;
  background: transparent;
}


@media screen and (max-width: 500px) {
  .profile-card {
    width: 95%;
    max-width: 95%;
    margin: 1em;
    display: block;
  }
}
</style>

