<template>
  <div class="col">
    <div class="movie-card">
      <div class="movie-header theDarkTower" :style="{ 'background-image': 'url(' + poster + ')' }"></div>
      <!--movie-header-->
      <div class="movie-content">
        <div class="row">
          <div class="col-8">
            <h4>{{ title }}</h4>
          </div>
          <div class="col-4 mt-1">
            <h5>{{ date }}</h5>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <div class="row mt-2">
              <p class="text-decoration-underline">{{ genre }}</p>
            </div>
          </div>
          <div class="col">
            <div class="row mt-2">
              <p class="text-decoration-underline">{{ length }}</p>
            </div>
          </div>
          <div class="col">
            <div class="row  mt-2">
              <p class="text-decoration-underline">{{ rating }}</p>
            </div>
          </div>

        </div>
        <div class="row">
          <div class="col">
            <div class="row">
              <div class="col ms-1">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-eye"
                  viewBox="0 0 16 16">
                  <path
                    d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8zM1.173 8a13.133 13.133 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5c2.12 0 3.879 1.168 5.168 2.457A13.133 13.133 0 0 1 14.828 8c-.058.087-.122.183-.195.288-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5c-2.12 0-3.879-1.168-5.168-2.457A13.134 13.134 0 0 1 1.172 8z" />
                  <path d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5zM4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0z" />
                </svg>
              </div>
              <div class="col me-2">{{ number_of_watches }}</div>
            </div>

          </div>
          <div class="col border-start border-2 border-dark">
            <div class="row">
              <div class="col ms-1">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-heart"
                  viewBox="0 0 16 16">
                  <path
                    d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z" />
                </svg>
              </div>
              <div class="col me-2">{{ number_of_likes }}</div>
            </div>
          </div>
        </div>
      </div>

      <!--movie-content-->
    </div>
    <!--movie-card-->
  </div>
</template>

<script>
export default {
  props:['poster','title','date','rating','genre','length','number_of_watches','number_of_likes'] ,
  data() {
    return {}
  }
}
</script>

<style scoped>
* {
  -webkit-transition: 300ms;
  transition: 300ms;
}

.movie-card {
  background: #ffffff;
  box-shadow: 0px 6px 18px rgba(130, 10, 10, 0.1);
  width: 100%;
  max-width: 290px;
  border-radius: 10px;
  display: inline-block;
}

.movie-header {
  padding: 0;
  margin: 0;
  height: 380px;
  width: 100%;
  display: block;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
}

.theDarkTower {
  background-size: cover;
}


.movie-card:hover {
  -webkit-transform: scale(1.03);
  transform: scale(1.03);
  box-shadow: 0px 10px 25px rgba(0, 0, 0, 0.08);
}

.movie-content {
  
  padding: 8px 12px 8px 12px;
  font-size: 10px;
}
.movie-content-header,
.movie-info {
  display: table;
  width: 100%;
}

@media screen and (max-width: 500px) {
  .movie-card {
    width: 95%;
    max-width: 95%;
    margin: 1em;
    display: block;
  }
}
</style>
