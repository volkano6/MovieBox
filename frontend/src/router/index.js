import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/login.vue"
import Login from "../views/login.vue"
import Register from "../views/register.vue"
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
      path: "/deneme",
      name: "deneme",
      component: () => import("../views/deneme.vue"),
    },
    {
      path: "/home",
      name: "home",
      component: () => import("../views/home.vue"),
    },
    // If no path matches, return 404 component
    {
      path: "/:pathMatch(.*)*",
      name: "404",
      component: NotFound,
    }
  ],
});

export default router;
