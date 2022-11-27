import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/login.vue"

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
      component: () => import("../views/login.vue"),
    },
    {
      path: "/register",
      name: "register",
      component: () => import("../views/login.vue"),
    },
    {
      path: "/profile",
      name: "profile",
      component: () => import("../views/profile.vue"),
    }
  ],
});

export default router;
