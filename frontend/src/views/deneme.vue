<template>
    <div style="height: 100%;background: rgb(0,4,47);
background: linear-gradient(90deg, rgba(0,4,47,1) 0%, rgba(0,7,45,1) 51%, rgba(0,2,37,1) 100%);">
        <div style="width: 100%;padding: 5% 15% 0px 15%;height: max-content;display: flex;
    background: rgb(0,4,47);
background: linear-gradient(90deg, rgba(0,4,47,1) 0%, rgba(0,17,112,1) 51%, rgba(0,2,37,1) 100%)">
            <div style="width: 14.6%;padding: 0px;margin: 0px;flex: 1;">
                <div style="border: 1px solid black;border-radius: 50%;height: 100%;width:65%">
                <img src="https://www.teknozum.com/wp-content/uploads/2019/12/whatsapp-profil-foto%C4%9Fraflar%C4%B1-10-1024x819.jpg" style="height: 100%;width:100%;border-radius: 50%;"/>
                </div>
            </div>
            <div style="width: 16.6%;margin-top: 1%;">
                <h2 class="text-left" style="color:white">Username</h2>
                <button class="text-left" style="padding: 1% 5% 2% 5%;font-weight:bold"> Edit Profile</button>
            </div>
            <div style="width: 18.8%">
            </div>
            <div style="width: 12.5%;border-right: 2px solid white;">
                <h2 class="text-center" style="padding-top: 7%;color:white">5</h2>
                <h5 class="text-center" style="color: white;">Films</h5>
            </div>
            <div style="width: 12.5%;border-right: 2px solid white;">
                <h2 class="text-center" style="padding-top: 7%;color:white">5</h2>
                <h5 class="text-center" style="color: white;">This Year</h5>
            </div>
            <div style="width: 12.5%;border-right: 2px solid white;">
                <h2 class="text-center" style="padding-top: 7%;color:white">5</h2>
                <h5 class="text-center" style="color: white;">Following</h5>
            </div>
            <div style="width: 12.5%">
                <h2 class="text-center" style="padding-top: 7%;color:white">5</h2>
                <h5 class="text-center" style="color: white;">Followers</h5>
            </div>
        </div>
        <div style="width: 100%;padding: 5% 15% 0px 15%;height: max-content;display: flex;
    background: rgb(0,4,47);
background: linear-gradient(90deg, rgba(0,4,47,1) 0%, rgba(0,17,112,1) 51%, rgba(0,2,37,1) 100%);padding-bottom: 3%;">
            <div style="width: 30%">
            </div>
            <div style="width:10%;flex:1;border:2px solid white;display:flex;justify-items:center;justify-content: center;">
                <a href="https://www.google.com" target="_blank" style="padding: 1% 40% 2% 40%;font-size: larger;font-weight: bold;">Profile</a>
            </div>
            <div style="width:10%;flex:1;border:2px solid white;display:flex;justify-items:center;justify-content: center;">
                <a href="https://www.google.com" target="_blank" style="padding: 1% 40% 2% 40%;font-size: larger;font-weight: bold;">Films</a>
            </div>
            <div style="width:10%;flex:1;border:2px solid white;display:flex;justify-items:center;justify-content: center;">
                <a href="https://www.google.com" target="_blank" style="padding: 1% 40% 2% 40%;font-size: larger;font-weight: bold;">Watchlist</a>
            </div>
            <div style="width:10%;flex:1;border:2px solid white;display:flex;justify-items:center;justify-content: center;">
                <a href="https://www.google.com" target="_blank" style="padding: 1% 40% 2% 40%;font-size: larger;font-weight: bold;">Likes</a>
            </div>
            <div style="width: 30%">
            </div>
        </div>
        <div style="width: 100%;padding: 5% 15% 0px 15%;height: max-content;display: flex">
            <div style="width:40%;flex:1;border-bottom:2px solid white;display:flex;justify-items:center;justify-content: left;">
                <h4 style="font-size: x-large;font-weight: bold;color: white;">Recent Films</h4>
            </div>
            <div style="width: 60%">
            </div>
        </div>
        <div style="width: 100%;padding: 5% 15% 5% 15%;height: max-content;display: flex">
            <div style="width:25%;flex:1;padding:0% 2% 0% 2%">
                <img style="width:100%" src="https://tr.web.img2.acsta.net/pictures/19/03/21/08/26/4285272.jpg"/>
            </div>
            <div style="width:25%;flex:1;padding:0% 2% 0% 2%">
                <img style="width:100%" src="https://tr.web.img2.acsta.net/pictures/19/03/21/08/26/4285272.jpg"/>
            </div>
            <div style="width:25%;flex:1;padding:0% 2% 0% 2%">
                <img style="width:100%" src="https://tr.web.img2.acsta.net/pictures/19/03/21/08/26/4285272.jpg"/>
            </div>
            <div style="width:25%;flex:1;padding:0% 2% 0% 2%">
                <img style="width:100%" src="https://tr.web.img2.acsta.net/pictures/19/03/21/08/26/4285272.jpg"/>
            </div>
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