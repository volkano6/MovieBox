<template>
  <section class="navColor border-bottom">
    <div class="container-sm border-bottom-dark">
      <nav class="navbar navbar-expand-lg bg-gray">
        <div class="container-fluid">
          <a class="navbar-brand" href="#">
            <img id="logo" src="https://a.ltrbxd.com/logos/letterboxd-mac-icon.png" alt="MDB Logo" draggable="false"
              height="40" />
          </a>
          <a class="navbar-brand text-light" href="#">MovieBoxd</a>
          <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
            data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false"
            aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <div class="col-4">
              <div class="col-md-5">
                <form class="d-flex input-group w-auto my-auto mb-3 mb-md-0">
                  <input autocomplete="off" type="search" class="form-control rounded" placeholder="Search" />
                </form>
              </div>
            </div>
            <div class="col-8">
              <div class="col-md-9">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                  <li class="nav-item">
                    <a class="nav-link active text-light text-opacity-75" aria-current="page" href="/home">Home</a>
                  </li>
                  <li class="nav-item">
                    <a class="nav-link text-light text-opacity-75" href="/films">Films</a>
                  </li>
                  <ul class="navbar-nav ms-auto" v-if="!logged">
                    <li class="nav-item">
                      <a class="nav-link d-flex align-items-center text-light text-opacity-75" href="/login">
                        <button type="button" class="btn btn-sm btn-outline-light me-2">Login</button>
                      </a>
                    </li>
                    <li class="nav-item">
                      <a class="nav-link d-flex align-items-center text-light text-opacity-75" href="/register">
                        <button type="button" class="btn btn-sm btn-warning">Sign-up</button>
                      </a>
                    </li>
                  </ul>
                  <ul class="navbar-nav ms-auto" v-else>
                    <div class="dropdown text-end">
                      <a href="#" class="d-block link-dark text-decoration-none dropdown-toggle" id="dropdownUser1"
                        data-bs-toggle="dropdown" aria-expanded="false">
                        
                        <div v-if="!avatar">
                          <img src="https://avatars.githubusercontent.com/u/1669855?v=4" alt="mdo" width="32" height="32" class="rounded-circle">
                        </div>
                        <div v-else>
                          <img :src="avatar" alt="mdo" width="32" height="32" class="rounded-circle">
                        </div>
                      </a>
                      <ul class="dropdown-menu text-small" aria-labelledby="dropdownUser1" style="">
                        <li><a class="dropdown-item" href="/profile">Profile</a></li>
                        <li>
                          <hr class="dropdown-divider">
                        </li>
                        <li><a class="dropdown-item" href="/login" @click="deleteCookie">Sign out</a></li>
                      </ul>
                    </div>
                  </ul>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </nav>
    </div>
  </section>
</template>


<script>
export default {
  name: "navbar",
  props: ['logged', 'avatar'],
  methods: {
    deleteCookie(){
      document.cookie = "token=;expires=Thu, 01 Jan 1970 00:00:00 GMT"
    }
  }
}
</script>

<style>
.navColor {
  background-image: #484848;
}
</style>