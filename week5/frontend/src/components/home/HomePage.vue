<script setup lang="ts">
import { onMounted, ref } from "vue";

type Post = {
  id: number;
  title: string;
  date_created: string;
  date_updated?: string;
  author: {
    id: number;
    name: string;
  };
};

const posts = ref<Post[]>([]);

onMounted(async () => {
  const res = await fetch(
    `${import.meta.env.VITE_DIRECTUS_API}/items/posts?fields=id,title,date_created,date_updated,author.*`,
    {
      method: "GET",
      mode: "cors",
      credentials: "include",
    },
  );

  if (res.status != 200) {
    return;
  }

  const json = await res.json();
  posts.value = json.data;
});
</script>

<template>
  <div class="p-6 text-lg font-bold" v-if="posts.length == 0">
    Oopsie, no posts!
  </div>
  <div
    v-else
    class="flex w-full flex-col items-center justify-center gap-8 p-6"
  >
    <h1 class="text-xl font-bold italic underline">Posts!</h1>

    <div class="flex w-full flex-col gap-4">
      <RouterLink v-for="post in posts" :to="`post/${post.id}`">
        <div class="rounded-lg border border-zinc-300 p-6 shadow-xl">
          <h2 class="text-lg font-semibold">{{ post.title }}</h2>
          <p>By {{ post.author.name }}</p>
        </div>
      </RouterLink>
    </div>
  </div>
</template>
