<template>
    <div>
      <Nav></Nav>
    </div>
  </template>
    
  <script>
  import axios from 'axios'
  
  export default {
    name: "profile",
    components: {
      Nav
     },
    data() {
      return {
        user: null
      }
    },
  }
  </script>
    
  <style>
  
  </style>