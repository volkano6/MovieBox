import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/login.vue"
import Login from "../views/login.vue"
import Register from "../views/register.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/login",
      name: "login",
      component: Login,
    },
    {
      path: "/register",
      name: "register",
      component: Register,
    },
    {
      path: "/profile",
      name: "profile",
      component: () => import("../views/profile.vue"),
    }
  ],
});

export default router;
