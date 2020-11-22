import { createWebHistory, createRouter, Router } from "vue-router";
import Edit from "@/views/Edit.vue";
import Article from "@/views/Article.vue";
import Home from '@/views/Home.vue';

const routes: any[] = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/:name",
    name: "Article",
    component: Article,
    props: true
  },
  {
    path: "/edit/:name",
    name: "Edit",
    component: Edit,
    props: true
  }
];

const router: Router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
