<template>
  <div>
    <Nav></Nav>
    <div class="container mt-3">
      <div class="row">
        <div class="col-3 border">
          <img src="https://picsum.photos/300/400" alt="" />
          <div class="row">
            <div class="col-6 border">
                <div class="row mt-1">
                    <div class="col">
                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-heart" viewBox="0 0 16 16">
                        <path d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z"/>
                        </svg>
                    </div>
                    <div class="col"><p>345</p></div>
                </div>
            </div>
            <div class="col-6 border">
                <div class="row mt-1">
                    <div class="col">
                        <svg xmlns="http://www.w3.org/2000/svg" width="31" height="31" fill="currentColor" class="bi bi-eye-fill" viewBox="0 0 16 16">
                        <path d="M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z"/>
                        <path d="M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8zm8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7z"/>
                        </svg>
                    </div>
                    <div class="col"><p>345</p></div>
                </div>
                
            </div>
          </div>
        </div>
        <div class="col-9">
          <div class="row justify-content-start">
            <div class="col-1"><h3 class="fw-bold">Title</h3></div>
            <div class="col-1"><h4 class="fw-light">date</h4></div>
            <div class="col-1"><h4 class="fw-light">director</h4></div>
          </div>
          <div class="row">
            <div class="col-9 ">
              <div class="row "><h3>Des</h3></div>
              <div class="row ">
                <h3>Genre</h3>
                <p>
                  Lorem ipsum dolor sit amet consectetur adipisicing elit.
                  Mollitia, incidunt consectetur! Alias perspiciatis repellendus
                  esse porro quae, voluptate sapiente accusantium perferendis,
                  fugit delectus reiciendis iusto quibusdam? Placeat consequatur
                  natus quaerat Lorem ipsum dolor sit amet consectetur,
                  adipisicing elit. Quas incidunt nulla sequi ipsum obcaecati,
                  aut quos id. Repellendus dolor, quia consequuntur, saepe magni
                  voluptatibus ea ad doloribus nesciunt aspernatur sit?
                </p>
              </div>
              <div class="row ">
                <h3>Cast</h3>
                <p>
                  Lorem ipsum dolor sit amet consectetur, adipisicing elit.
                  Nulla incidunt porro cum odit dolor tenetur beatae eos magnam
                  provident dolorem fugiat, odio quod quas aspernatur
                  reprehenderit, ipsa quidem, quisquam saepe. Nihil quo commodi
                  aliquam iste soluta earum est consectetur? Quia deleniti
                  tempora accusantium! Fugit provident, doloribus quidem omnis
                  et, distinctio tempore laudantium incidunt vero iusto ipsa
                  numquam nesciunt quis illo!
                </p>
              </div>
            </div>
            <div class="col bg-success">
              <div class="row">
                <div class="col-4">watch</div>
                <div class="col-4">list</div>
                <div class="col-4">watcchh</div>
              </div>
              <div class="row">
                <div class="col">Ratings</div>
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
import Nav from "../components/nav.vue";
import Footer from "../components/footer.vue";

export default {
  name: "films",
  components: {
    Nav,
    Footer,
  },
};
</script>

<style>
</style>