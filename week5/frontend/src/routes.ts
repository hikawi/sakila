import type { RouteRecordRaw } from "vue-router";
import HomePage from "./components/home/HomePage.vue";
import PostPage from "./components/post/PostPage.vue";

export const routes: RouteRecordRaw[] = [
  { path: "/", component: HomePage },
  { path: "/post/:id", component: PostPage },
];
