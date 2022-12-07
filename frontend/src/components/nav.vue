<template>
  <section class=nav>
    <div class="container-sm border-bottom-dark">
      <nav class="navbar navbar-expand-lg bg-gray">
        <div class="container-fluid">
          <a class="navbar-brand" href="#"><img id="logo" src="https://a.ltrbxd.com/logos/letterboxd-mac-icon.png"
              alt="MDB Logo" draggable="false" height="40" /></a>
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
                    <a class="nav-link text-light text-opacity-75" href="/profile">Profile</a>
                  </li>
                  <li class="nav-item">
                    <a class="nav-link text-light text-opacity-75" href="/#">Films</a>
                  </li>
                  <ul class="navbar-nav ms-auto">
                    <li class="nav-item">
                      <a class="nav-link d-flex align-items-center text-light text-opacity-75" href="/login">Sign In</a>
                    </li>
                    <li class="nav-item">
                      <a class="nav-link d-flex align-items-center text-light text-opacity-75" href="/login"> Log
                        Out</a>
                    </li>

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

}
</script>

<style>
.nav {

  background-image: radial-gradient(circle, #6e799e, #304b51, #273738, #1e2423)
}
</style>