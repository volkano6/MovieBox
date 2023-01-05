<template>
  <div
    class="p-5 position-absolute top-50 start-50 translate-middle border border-2 shadow-lg p-3 mb-5 bg-body rounded">
    <div class="d-flex justify-content-center mb-2">
      <h2>Sign Up</h2>
    </div>
    <form @submit.prevent="handleSubmit">
      <div class="">
        <div class="mb-2">
          <label for="exampleInputPassword1" class="form-label">Username</label>
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon1">@</span>
          <input type="username" class="form-control" v-model="user_name" />
        </div>
      </div>
      <div class="mb-3">
        <label for="exampleInputEmail1" class="form-label">Email address</label>

        <input type="email" class="form-control" v-model="email">
      </div>
      <div class="mb-3">
        <label for="exampleInputPassword1" class="form-label">Password</label>

        <input type="password" class="form-control" v-model="password" />
      </div>
      <div class="row">
        <div class="mb-4">
          <label for="exampleInputPassword1" class="form-label">Date of birth</label>

          <input class="datepicker form-control" v-model="date_of_birth" data-date-format="yyyy/mm/dd"
            placeholder="yyyy-mm-dd">
        </div>
      </div>
      <div class="row">
        <button type="submit" class="btn btn-primary">Sign up</button>
      </div>
    </form>
    <div class="m-2">
      <p class="text-center">
        <a href="/login" class="text-dark text-decoration-underline">Already have an account!!</a>
      </p>
    </div>
    <div class="row" v-if="error">
      <div class="row">
        <div class="border border-danger text-danger ms-2">
          <p class="">{{ err_message }}</p>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios'

export default {
  name: "register",
  data() {
    return {
      user_name: '',
      email: '',
      password: '',
      date_of_birth: '',
      error: false,
      err_message: '',
    }
  },
  methods: {
    async handleSubmit() {
      await axios.post('api/auth/register', {
        user_name: this.user_name,
        email: this.email,
        password: this.password,
        date_of_birth: this.date_of_birth,
      }).then((response) => {

        if (response.data.status == "ok") {
          this.$router.push("/login")

        } else {
          this.error = true
          this.err_message = response.data.message
          
        }
      })

    }
  }
};
</script>
  
<style>

</style>