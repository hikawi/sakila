import type { RouteRecordRaw } from "vue-router";
import HomePage from "./components/home/HomePage.vue";
import LoginPage from "./components/login/LoginPage.vue";

export const routes: RouteRecordRaw[] = [
  { path: "/", component: HomePage },
  { path: "/post/:id", component: import("./components/post/PostPage.vue") },
  { path: "/login", component: LoginPage },
];
