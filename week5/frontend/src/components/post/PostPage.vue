<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const route = useRoute();

type Post = {
  id: number;
  title: string;
  content: string;
  date_created: string;
  date_updated: string;
  author: {
    id: number;
    name: string;
  };
};

const loading = ref(true);
const post = ref<Post | undefined>();

onMounted(async () => {
  if (!route.params.id) {
    router.push({ path: "/" });
    return;
  }

  const res = await fetch(
    `${import.meta.env.VITE_DIRECTUS_API}/items/posts/${route.params.id}?fields=*.*`,
    {
      method: "GET",
      mode: "cors",
      credentials: "include",
    },
  );

  if (res.status != 200) {
    loading.value = false;
    return;
  }

  const json = await res.json();
  post.value = json.data;
  loading.value = false;
});
</script>

<template>
  <section class="prose mx-auto w-full p-6" v-if="post">
    <h1>{{ post.title }}</h1>
    <p>By {{ post.author.name }}</p>
    <p>Published on {{ new Date(post.date_created).toLocaleString() }}</p>
    <hr />
    <div v-html="post.content"></div>
  </section>
</template>
