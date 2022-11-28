<template>
   <div
      class="p-5 position-absolute top-50 start-50 translate-middle border border-2 shadow-lg p-3 mb-5 bg-body rounded">
      <div class="d-flex justify-content-center">
         <h2>Sign in</h2>
      </div>
      <form @submit.prevent="handleSubmit">
         <div>
            <div class="mb-2">
               <label for="exampleInputPassword1" class="form-label mt-4">Username</label>
            </div>
            <div class="input-group mb-3">
               <span class="input-group-text" id="basic-addon1">@</span>
               <input type="username" class="form-control" v-model="user_name" />
            </div>
         </div>
         <div class="mb-3">
            <label for="exampleInputPassword1" class="form-label">Password</label>
            <input type="password" class="form-control" v-model="password" />
         </div>
         <div class="mb-4 form-check">
            <input type="checkbox" class="form-check-input" id="exampleCheck1" />
            <label class="form-check-label" for="exampleCheck1">Check me out</label>
         </div>
         <div class="row">
            <button type="submit" class="btn btn-primary">Sign in</button>
         </div>
      </form>
      <div class="m-2">
         <p class="text-center">
            <a href="/register" class="">Create an account!!</a>
         </p>
      </div>
   </div>
</template>

<script>
import axios from 'axios'

export default {
   name: "login",
   data(){
      return {
         user_name:'',
         password: ''
      }
   },
   methods: {
      async handleSubmit() {
      const response = await axios.post('api/auth/login', {
         user_name: this.user_name,
         password: this.password
      }).then((response) => {
         if (response.data.status == "ok") {
            // Set cookie for JWT
            let d = new Date();
            d.setTime(d.getTime() + 14 * 24 * 60 * 60 * 1000);
            let expires = "expires=" + d.toUTCString()
            document.cookie = "token=" + response.data.data + ";" + expires + ";path=/";
            // Redirect to profile
            this.$router.push("/profile")
         } else if (response.data.status == "error") {
            // Invalid credentials.
            // Display user to try again.
            console.log("err")
         }
      })
      
    }
   }
};
</script>

<style>

</style>
