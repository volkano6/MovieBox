<template>
  <div>
    <h1>
      {{user.Username}}
      {{user.Email}}
    </h1>
  </div>
</template>
  
<script>
import axios from 'axios'

export default {
  name: "profile",
  data() {
    return {
      user: null
    }
  },
  async created() {
    const response = await axios.get("api/profile");
    if (response.data.status == "error") {
      this.$router.push("/login")
    } else {
      this.user = response.data.data
    }
  }
}
</script>
  
<style>

</style>