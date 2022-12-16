import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/login.vue"
import Login from "../views/login.vue"
import Register from "../views/register.vue"
import Films from "../views/films.vue"
import NotFound from "../views/404.vue"

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
    },
    {
      path: "/user/:id",
      name: "user",
      component: () => import("../views/user.vue"),
    },
    {
      path: "/home",
      name: "home",
      component: () => import("../views/home.vue"),
    },
    {
      path: "/movie",
      name: "movie",
      component: () => import("../views/movie.vue"),
    },
    {
      path: "/films",
      name: "films",
      component: () => import("../views/films.vue"),
    },
    {
      path: "/admin",
      name: "admin",
      component: () => import("../views/admin.vue"),
    },
    // If no path matches, return 404 component
    {
      path: "/:pathMatch(.*)*",
      name: "404",
      component: () => import("../views/home.vue"),
    }
  ],
});

export default router;
