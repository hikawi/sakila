<script setup lang="ts">
import { onMounted, ref } from "vue";

type Profile = {
  email: string;
  first_name: string;
  last_name: string;
  role: {
    name: string;
  };
};

const loading = ref(true);
const profile = ref<Profile | undefined>();

onMounted(async () => {
  const res = await fetch(
    `${import.meta.env.VITE_DIRECTUS_API}/users/me?fields=*.*`,
    {
      method: "GET",
      mode: "cors",
      credentials: "include",
    },
  );

  loading.value = false;
  if (res.status != 200) {
    return;
  }

  const json = await res.json();
  profile.value = json.data;
});
</script>

<template>
  <div class="w-full p-6">
    <div
      class="flex w-full flex-col gap-4 rounded-lg border border-zinc-300 p-4 shadow-xl"
    >
      <p class="text-lg italic" v-if="loading">Loading...</p>
      <p class="text-lg" v-else-if="!profile">Not logged in</p>
      <p class="text-lg" v-else>Logged in idk</p>
    </div>
  </div>
</template>
