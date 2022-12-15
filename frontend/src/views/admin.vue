<template>
    <div>
        <Nav></Nav>


        <div class="container mt-3">
            <div class="row">
                <div class="col">
                    <h3 class="fw-bold" style="color:white;">My DashBoard</h3>
                    <!-- <p style="color:white;">My Dashboard</p> -->
                </div>
            </div>

            <div class="row mt-3">
                <div class="col-1 border ">
                    <h5 class="fw-bold" style="color:white;">Id</h5>
                </div>
                <div class="col-2 border">
                    <h5 class="fw-bold" style="color:white;">Title</h5>
                </div>
                <div class="col-3 border">
                    <h5 class="fw-bold" style="color:white;">Description</h5>
                </div>
                <div class="col-2 border">
                    <h5 class="fw-bold" style="color:white;">Release Date</h5>
                </div>
                <div class="col-1 border">
                    <!-- <p style="color:white;">Poster</p> -->
                    <h5 class="fw-bold" style="color:white;">Poster</h5>
                </div>
                <div class="col-1 border">
                    <h5 class="fw-bold" style="color:white;">Rating</h5>

                </div>
                <div class="col-1 border">
                    <h5 class="fw-bold" style="color:white;">Length</h5>

                </div>
                <div class="col-1">
                </div>
            </div>
            <div class="row">
                <div class="col-1 border ">
                    <p class="mt-2" style="color:white;">1</p>
                </div>
                <div class="col-2 border">
                    <p class="mt-2" style="color:white;">Lord Of the Rings</p>               
                </div>
                <div class="col-3 border">
                    <p class="mt-2" style="color:white;">Dunyanin en iyi filmlerinden</p>
                </div>
                <div class="col-2 border">
                    <p class="mt-2" style="color:white;">1999</p>                
                </div>
                <div class="col-1 border">
                    <!-- <p style="color:white;">Poster</p> -->
                    <p class="mt-2" style="color:white;">posterling</p>
                    
                </div>
                <div class="col-1 border">
                    <p class="mt-2" style="color:white;">8.2</p>

                </div>
                <div class="col-1 border">
                    <p class="mt-2" style="color:white;">168</p>

                </div>
                <div class="col-1 border">

                    <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-pencil-square mt-2" viewBox="0 0 16 16">
                    <path d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z"></path>
                    <path fill-rule="evenodd" d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z"></path>
                    </svg>
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
    name: "admin",
    components: {
        Nav,
        Footer,
    },
}
</script>

<style>

</style>