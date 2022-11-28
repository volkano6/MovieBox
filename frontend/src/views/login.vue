<template>
   <div
      class="p-5 position-absolute top-50 start-50 translate-middle border border-2 shadow-lg p-3 mb-5 bg-body rounded">
      <div class="d-flex justify-content-center">
         <h2>Sign in</h2>
      </div>
      <br />
      <form @submit.prevent="handleSubmit">
         <div>
            <div class="mb-2">
               <label for="exampleInputPassword1" class="form-label">Username</label>
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
            <a href="/register" class="">Already have an account!!</a>
         </p>
      </div>
   </div>
</template>

<script>
export default {
   name: "login",
   data(){
      return {
         user_name:'',
         password: ''
      }
   },
   methods: {
      handleSubmit(){
         const data = {
            user_name: this.user_name,
            password: this.password
         }
         console.log(data)
      }
   }
};
</script>

<style>

</style>
