<template>

    <div>
        <Nav></Nav>
        <div class="container mt-3">
            <div class="btn-group" role="group" aria-label="Basic mixed styles example">
                <div class="col-2" style="margin-top: 2px;">
                    <p>BROWSED BY</p>
                </div>
                <div class="col-2 " >

                    <div class="Year ms-2">
                        <button class="btn btn-secondary btn-lg dropdown-toggle" type="button" id="dropdownMenuButton"
                            data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            Year
                        </button>
                        <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                            <a class="dropdown-item" href="#">2012</a>
                            <a class="dropdown-item" href="#">2015</a>
                            <a class="dropdown-item" href="#">2018</a>
                        </div>
                    </div>
                </div>
                <div class="col-2">
                    <div class="Rating ms-2">

                        <button class="btn btn-secondary btn-lg dropdown-toggle" type="button" id="dropdownMenuButton"
                            data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            Rating
                        </button>
                        <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                            <a class="dropdown-item" href="#">Highest First</a>
                            <a class="dropdown-item" href="#">Lowest First</a>

                        </div>
                    </div>
                </div>
                <div class="col-2">
                    <div class="Genre ms-2">
                        <button class="btn btn-secondary btn-lg dropdown-toggle" type="button" id="dropdownMenuButton"
                            data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            Genre
                        </button>
                        <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                            <a class="dropdown-item" href="#">Action</a>
                            <a class="dropdown-item" href="#">Comedy</a>
                            <a class="dropdown-item" href="#">Horror</a>
                            <a class="dropdown-item" href="#">Drama</a>
                            <a class="dropdown-item" href="#">Science Fiction</a>
                            <a class="dropdown-item" href="#">Romance</a>
                        </div>
                    </div>
                </div>
                <form class="d-flex ms-5" role="Find a film">
                    <input class="form-control me-2" type="Film" placeholder="Find a Film" aria-label="Search">
                    <button class="btn btn-outline-success" type="submit">Search</button>
                </form>
            </div>
            <div class="text-center" style="padding-left:35%;padding-right:35%">
                <h2 class="text-center mt-5"
                    style="margin-top: 2%; color: aliceblue; border-bottom: 2px solid aliceblue; padding-bottom: 0.7%;">
                    Favorite of Weeks </h2>
            </div>
            <div class="row mt-2 row-cols-1 row-cols-xl-5 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-3">
                <Card />
                <Card />
                <Card />
                <Card />
                <Card />

            </div>

        </div>
        <Footer></Footer>
    </div>




</template>
  
<script>

import Nav from "../components/nav.vue";
import Footer from "../components/footer.vue";
import axios from 'axios'
import Card from "../components/movie_card.vue"

export default {
    name: "films",
    components: {
        Nav,
        Footer,
        Card
    },
};
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

.post-list .label {
    font-weight: normal;
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
    background-color: rgba(0, 0, 0, 0);
    font-weight: 400;
    opacity: 0.8;
    padding: 7px 16px;
}
</style>